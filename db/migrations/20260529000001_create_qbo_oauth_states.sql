-- +goose Up
CREATE TABLE qbo_oauth_states (
    id UUID PRIMARY KEY,
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    state_checksum TEXT NOT NULL,
    encrypted_state BYTEA NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    consumed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX qbo_oauth_states_checksum_idx
    ON qbo_oauth_states (state_checksum);

CREATE INDEX qbo_oauth_states_expires_idx
    ON qbo_oauth_states (expires_at)
    WHERE consumed_at IS NULL;

-- +goose Down
DROP TABLE qbo_oauth_states;
