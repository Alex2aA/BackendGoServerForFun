-- Bookings
CREATE TABLE IF NOT EXISTS booked (
                                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_house UUID REFERENCES free_houses(id) ON DELETE CASCADE,
    id_hostel UUID REFERENCES hostels(id) ON DELETE CASCADE,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );
