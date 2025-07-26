CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE player_stats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_espn_id VARCHAR(10) NOT NULL,
    player_espn_id VARCHAR(10) NOT NULL,
    team_espn_id VARCHAR(10) NOT NULL,

    min INT,
    fg_made INT,
    fg_att INT,
    threept_made INT,
    threept_att INT,
    ft_made INT,
    ft_att INT,
    oreb INT,
    dreb INT,
    reb INT,
    ast INT,
    stl INT,
    blk INT,
    turnover INT,
    pf INT,
    pts INT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
