-- Active: 1680439612142@@127.0.0.1@5432@buildbus
CREATE TABLE if not exists me_active_ads (
    adv_no VARCHAR UNIQUE,    
    asset VARCHAR, 
    price VARCHAR,
    updated_at TIMESTAMPTZ,
    is_outdated BOOLEAN
);


