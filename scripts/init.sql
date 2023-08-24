-- SCHEMA: manga

-- DROP SCHEMA IF EXISTS manga ;

CREATE SCHEMA IF NOT EXISTS manga
    AUTHORIZATION postgres;

-- Table: manga.items

-- DROP TABLE IF EXISTS manga.items;

CREATE TABLE IF NOT EXISTS manga.items
(
    name text COLLATE pg_catalog."default" NOT NULL,
    create_time timestamp with time zone,
    favorite boolean,
    file_indices integer[],
    thumbnail bytea,
    is_read boolean,
    tags text[] COLLATE pg_catalog."default",
    CONSTRAINT items_name UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS manga.items
    OWNER to postgres;

-- Table: manga.tags

-- DROP TABLE IF EXISTS manga.tags;

CREATE TABLE IF NOT EXISTS manga.tags
(
    name text COLLATE pg_catalog."default" NOT NULL,
    favorite boolean,
    hidden boolean,
    thumbnail bytea,
    CONSTRAINT tags_name UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS manga.tags
    OWNER to postgres;