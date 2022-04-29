CREATE TABLE IF NOT EXISTS timer_units (
    id varchar(36) PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    "order" integer NOT NULL,
    duration integer NOT NULL,
    preset_id varchar(36) REFERENCES presets
);
