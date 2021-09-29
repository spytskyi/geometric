CREATE TABLE IF NOT EXISTS triangles (
                                           id           SERIAL PRIMARY KEY,
                                           side_a       FLOAT,
                                           side_b       FLOAT,
                                           side_c       FLOAT,
                                           injection_ab FLOAT,
                                           injection_bc FLOAT,
                                           injection_ac FLOAT,
                                           area         FLOAT,
                                           created_at   TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at   TIMESTAMPTZ
);