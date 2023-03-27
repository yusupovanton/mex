CREATE TABLE if not exists me_active_ads ( 
    id BIGINT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY, 
    adv_no VARCHAR UNIQUE,    
    asset VARCHAR, 
    price VARCHAR,
    updated_at TIMESTAMPTZ,
    is_outdated BOOLEAN
);


