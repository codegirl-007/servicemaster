-- +goose Up
CREATE TABLE qbo_connections (
    id UUID PRIMARY KEY,
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    realm_id TEXT NOT NULL,
    company_name TEXT,
    state TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT qbo_connections_state_check CHECK (
        state IN (
            'pending',
            'connected',
            'reconnect_required',
            'disconnected'
        )
    )
);

CREATE UNIQUE INDEX qbo_connections_tenant_unique
    ON qbo_connections (tenant_id);

CREATE UNIQUE INDEX qbo_connections_realm_unique
    ON qbo_connections (realm_id);

CREATE TABLE qbo_connection_events (
    id UUID PRIMARY KEY,
    qbo_connection_id UUID NOT NULL REFERENCES qbo_connections(id) ON DELETE CASCADE,
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    event_type TEXT NOT NULL,
    message TEXT,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT qbo_connection_events_type_check CHECK (
        event_type IN (
            'connection_created',
            'connected',
            'connection_failed',
            'reconnect_required',
            'disconnected',
            'reconnected',
            'import_succeeded',
            'import_failed'
        )
    )
);

CREATE INDEX qbo_connection_events_connection_occurred_at_idx
    ON qbo_connection_events (qbo_connection_id, occurred_at DESC);

CREATE INDEX qbo_connection_events_tenant_occurred_at_idx
    ON qbo_connection_events (tenant_id, occurred_at DESC);

-- +goose Down
DROP TABLE qbo_connection_events;
DROP TABLE qbo_connections;
