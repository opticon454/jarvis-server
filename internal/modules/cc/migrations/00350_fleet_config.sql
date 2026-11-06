-- Fleet config (docs/cc/14-fleet-config-and-recovery.md): one household-sealed blob per config
-- type, readable by every node of the household. jarvisd stores ciphertext only; the household
-- key lives on the phone and the nodes. revision rises on every write; a node's ack row says
-- which revision it has applied, so a per-node override pushed later is never overwritten.
-- +goose Up
CREATE TABLE cc_fleet_config (
    household_id TEXT    NOT NULL,
    config_type  TEXT    NOT NULL,
    ciphertext   TEXT    NOT NULL,
    nonce        TEXT    NOT NULL,
    tag          TEXT    NOT NULL,
    revision     INTEGER NOT NULL,
    updated_at   TEXT    NOT NULL,
    PRIMARY KEY (household_id, config_type)
);

CREATE TABLE cc_fleet_config_acks (
    node_id      TEXT    NOT NULL REFERENCES cc_nodes (node_id) ON DELETE CASCADE,
    household_id TEXT    NOT NULL,
    config_type  TEXT    NOT NULL,
    revision     INTEGER NOT NULL,
    acked_at     TEXT    NOT NULL,
    PRIMARY KEY (node_id, household_id, config_type)
);

-- +goose Down
DROP TABLE cc_fleet_config_acks;
DROP TABLE cc_fleet_config;
