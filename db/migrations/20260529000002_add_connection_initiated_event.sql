-- +goose Up
ALTER TABLE qbo_connection_events DROP CONSTRAINT qbo_connection_events_type_check;

ALTER TABLE qbo_connection_events ADD CONSTRAINT qbo_connection_events_type_check CHECK (
    event_type IN (
        'connection_created',
        'connected',
        'connection_initiated',
        'connection_failed',
        'reconnect_required',
        'disconnected',
        'reconnected',
        'import_succeeded',
        'import_failed'
    )
);

-- +goose Down
ALTER TABLE qbo_connection_events DROP CONSTRAINT qbo_connection_events_type_check;

ALTER TABLE qbo_connection_events ADD CONSTRAINT qbo_connection_events_type_check CHECK (
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
);
