-- Active: 1680439612142@@127.0.0.1@5432@buildbus
CREATE TABLE if not exists me_fiat_conversion_rates (
    base VARCHAR,
    us_dollar FLOAT,    
    turkish_lira FLOAT, 
    ts TIMESTAMPTZ
);


