-- Create users and blacklist
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    hash_password TEXT NOT NULL,
    alias VARCHAR(255) DEFAULT 'none',
    refresh_token TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS blacklist_tokens (
                                                token TEXT PRIMARY KEY,
                                                created_at TIMESTAMPTZ DEFAULT NOW()
    );