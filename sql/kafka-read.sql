CREATE MAPPING IF NOT EXISTS trades (
    id BIGINT,
    ticker VARCHAR,
    price DOUBLE,
    amount BIGINT
)
TYPE Kafka
OPTIONS (
    'valueFormat' = 'json-flat',
    'bootstrap.servers' = '127.0.0.1:9092'
);

SELECT ticker, ROUND(price * 100) AS price_cents, amount
FROM trades
WHERE price * amount > 100;
