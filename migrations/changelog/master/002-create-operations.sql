CREATE TABLE wallet_operations (
    id UUID PRIMARY KEY,
    wallet_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    operation_type TEXT NOT NULL CHECK (operation_type IN ('DEPOSIT', 'WITHDRAW')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_wallet
       FOREIGN KEY (wallet_id)
           REFERENCES wallets(id)
);
