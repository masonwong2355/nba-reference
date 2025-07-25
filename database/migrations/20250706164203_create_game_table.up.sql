CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    espn_id VARCHAR(20) NOT NULL,  
    start_time TIMESTAMP NOT NULL,
    season_year VARCHAR(10),               
    type VARCHAR(20),         
    home_team_id VARCHAR(10),
    away_team_id VARCHAR(10),
    home_score INT,
    away_score INT,
    home_q1_score INT,
    home_q2_score INT,
    home_q3_score INT,
    home_q4_score INT,
    away_q1_score INT,
    away_q2_score INT,
    away_q3_score INT,
    away_q4_score INT,
    arena VARCHAR(100),
    referees TEXT,
    winner_team_id VARCHAR(10),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_games_espn_id ON games(espn_id);

