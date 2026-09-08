-- Copyright (C) 2026 Asyraf Mubarak
-- This file is part of gopos-api and licensed under GNU AGPLv3.

CREATE TABLE cash_payments (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL,
    sale_id uuid NOT NULL,
    tendered_rupiah bigint NOT NULL CHECK (tendered_rupiah > 0),
    applied_rupiah bigint NOT NULL CHECK (applied_rupiah > 0),
    change_rupiah bigint NOT NULL CHECK (change_rupiah >= 0),
    paid_at timestamptz NOT NULL,
    CONSTRAINT cash_payments_sale_unique UNIQUE (root_id, sale_id),
    CONSTRAINT cash_payments_root_id_id_unique UNIQUE (root_id, id),
    CONSTRAINT cash_payments_amounts_check CHECK (
        tendered_rupiah = applied_rupiah + change_rupiah
    ),
    FOREIGN KEY (root_id, sale_id) REFERENCES sales(root_id, id) ON DELETE RESTRICT
);

CREATE TABLE cash_refunds (
    id uuid PRIMARY KEY,
    root_id uuid NOT NULL,
    reversal_id uuid NOT NULL,
    cash_payment_id uuid NOT NULL,
    amount_rupiah bigint NOT NULL CHECK (amount_rupiah > 0),
    refunded_at timestamptz NOT NULL,
    CONSTRAINT cash_refunds_reversal_unique UNIQUE (root_id, reversal_id),
    CONSTRAINT cash_refunds_payment_unique UNIQUE (root_id, cash_payment_id),
    FOREIGN KEY (root_id, reversal_id) REFERENCES sale_reversals(root_id, id) ON DELETE RESTRICT,
    FOREIGN KEY (root_id, cash_payment_id) REFERENCES cash_payments(root_id, id) ON DELETE RESTRICT
);
