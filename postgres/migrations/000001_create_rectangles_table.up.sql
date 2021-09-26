CREATE TABLE IF NOT EXISTS triangles (
                                           id INTEGER PRIMARY KEY,
                                           first_name VARCHAR(255) NOT NULL,
                                           last_name VARCHAR(255) NOT NULL,
                                           email VARCHAR(320) NOT NULL,
                                           created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at TIMESTAMPTZ
);