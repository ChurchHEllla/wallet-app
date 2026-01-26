CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

