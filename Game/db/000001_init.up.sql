CREATE TABLE IF NOT EXISTS catalog(
                                      id TEXT PRIMARY KEY,
                                      name TEXT UNIQUE NOT NULL,
                                      points INT NOT NULL
);