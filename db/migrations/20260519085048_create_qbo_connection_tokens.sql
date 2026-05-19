-- +goose Up
CREATE TABLE qbo_connection_tokens (
    qbo_connection_id UUID PRIMARY KEY REFERENCES qbo_connections(id) ON DELETE CASCADE,
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    encrypted_access_token BYTEA NOT NULL,
    encrypted_refresh_token BYTEA NOT NULL,
    access_token_expires_at TIMESTAMPTZ NOT NULL,
    refresh_token_expires_at TIMESTAMPTZ,
    token_type TEXT NOT NULL DEFAULT 'bearer',
    scope TEXT,
    version BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX qbo_connection_tokens_tenant_idx
    ON qbo_connection_tokens (tenant_id);

-- +goose Down
DROP TABLE qbo_connection_tokens;
