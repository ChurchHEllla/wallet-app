package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const migrationsDir = "migrations/changelog/master"

type migration struct {
	version int
	path    string
}

// Run применяет все новые миграции
func Run(db *sql.DB) error {
	if err := ensureSchemaMigrations(db); err != nil {
		return fmt.Errorf("ensure schema_migrations: %w", err)
	}

	currentVersion, err := getCurrentVersion(db)
	if err != nil {
		return fmt.Errorf("get current version: %w", err)
	}

	migrations, err := collectMigrations(currentVersion)
	if err != nil {
		return err
	}

	if len(migrations) == 0 {
		log.Println("no new migrations")
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, m := range migrations {
		log.Printf("applying migration %d", m.version)

		sqlBytes, err := os.ReadFile(m.path)
		if err != nil {
			return err
		}

		if _, err := tx.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("migration %d failed: %w", m.version, err)
		}

		if _, err := tx.Exec(
			`INSERT INTO schema_migrations(version) VALUES ($1)`,
			m.version,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func collectMigrations(currentVersion int) ([]migration, error) {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	var migrations []migration

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		parts := strings.SplitN(f.Name(), "-", 2)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid migration name: %s", f.Name())
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid migration version: %s", f.Name())
		}

		if v > currentVersion {
			migrations = append(migrations, migration{
				version: v,
				path:    filepath.Join(migrationsDir, f.Name()),
			})
		}
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].version < migrations[j].version
	})

	return migrations, nil
}

func ensureSchemaMigrations(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INT PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT now()
		)
	`)
	return err
}

func getCurrentVersion(db *sql.DB) (int, error) {
	var v sql.NullInt64
	err := db.QueryRow(
		`SELECT max(version) FROM schema_migrations`,
	).Scan(&v)

	if err != nil {
		return 0, err
	}
	if !v.Valid {
		return 0, nil
	}
	return int(v.Int64), nil
}
