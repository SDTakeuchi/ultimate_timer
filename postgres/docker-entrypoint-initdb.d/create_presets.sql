CREATE TABLE IF NOT EXISTS presets (
    id varchar(36) PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    name varchar(128) NOT NULL,
    display_order integer NOT NULL,
    loop_count integer NOT NULL,
    waits_confirm_each boolean NOT NULL,
    waits_confirm_last boolean NOT NULL
);
