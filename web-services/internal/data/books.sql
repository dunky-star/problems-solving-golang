CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title STRING NOT NULL,
    published INT NOT NULL,
    pages INT NOT NULL,
    genres STRING[] NOT NULL,
    rating FLOAT NOT NULL,
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp()
);
