CREATE TABLE connects (
    id UUID PRIMARY KEY,
    active BOOLEAN NOT NULL,
    userid UUID NOT NULL
);

CREATE TABLE messages(
    id UUID  PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TEXT NOT NULL
);

CREATE TABLE users(
    id UUID PRIMARY KEY,
    username TEXT NOT NULL
);

