CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS presets (
    id uuid DEFAULT uuid_generate_v4 (),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name varchar(128) NOT NULL,
    display_order integer NOT NULL,
    loop_count integer NOT NULL,
    waits_confirm_each boolean NOT NULL,
    waits_confirm_last boolean NOT NULL,
    timer_units jsonb,
    PRIMARY KEY (id)
);
