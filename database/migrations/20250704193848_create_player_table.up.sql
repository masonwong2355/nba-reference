CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    espn_id VARCHAR(30) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    team_id VARCHAR(10),
    jersey_number INTEGER,
    position VARCHAR(40),
    height VARCHAR(20),          -- Keep as string unless you want to split feet/inches and store as int
    weight INTEGER,              -- lbs
    birthdate DATE,
    experience VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);