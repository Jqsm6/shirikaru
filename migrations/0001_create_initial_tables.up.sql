DROP TABLE IF EXISTS anime_list CASCADE;

CREATE TABLE anime_list
(
    animeID SERIAL,
    title TEXT NOT NULL UNIQUE,
    alternativeTitle TEXT,
    description TEXT,
    productionStatus TEXT NOT NULL,
    picture TEXT,
    episode INT NOT NULL
);

CREATE INDEX IF NOT EXISTS anime_list_idx
    ON anime_list (animeID);