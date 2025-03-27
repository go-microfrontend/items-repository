-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    type TEXT NOT NULL,
    weight_in_grams int NOT NULL CHECK (weight_in_grams > 0),
    amount int DEFAULT 0
);

-- +goose Down
DROP TABLE items;

DROP EXTENSION IF EXISTS "uuid-ossp";
