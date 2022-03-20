CREATE table presets (
    id varchar(32) not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name varchar(128) not null,
    display_order integer not null,
    loop_count integer not null,
    waits_confirm_each boolean not null,
    waits_confirm_last boolean not null
);


CREATE table timer_units (
    "order" integer not null,
    duration integer not null,
    preset_id varchar(32) not null
);
