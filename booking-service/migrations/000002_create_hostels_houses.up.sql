-- Hostels and Houses
CREATE TABLE IF NOT EXISTS hostels (
                                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    rate INT CHECK (rate BETWEEN 1 AND 5),
    created_at TIMESTAMPTZ DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS free_houses (
                                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_user UUID REFERENCES users(id) ON DELETE SET NULL,
    id_hostel UUID REFERENCES hostels(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    number_of_rooms INT NOT NULL,
    price_for_one_day INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );