CREATE TABLE IF NOT EXISTS triangles (
                                           id           SERIAL PRIMARY KEY,
                                           side_a       FLOAT,
                                           side_b       FLOAT,
                                           side_c       FLOAT,
                                           injection_a  FLOAT,
                                           injection_b  FLOAT,
                                           injection_c  FLOAT,
                                           area         FLOAT,
                                           created_at   TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at   TIMESTAMPTZ
);