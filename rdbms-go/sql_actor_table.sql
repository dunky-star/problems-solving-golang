CREATE TABLE actor (
    actor_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name STRING NOT NULL,
    last_name  STRING NOT NULL,
    email STRING UNIQUE NOT NULL,
    country STRING,
    created_at TIMESTAMP DEFAULT current_timestamp(),
    updated_at TIMESTAMP DEFAULT current_timestamp() ON UPDATE current_timestamp()
);
