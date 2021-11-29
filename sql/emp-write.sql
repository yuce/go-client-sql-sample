CREATE MAPPING IF NOT EXISTS employees (
    __key BIGINT,
    age BIGINT,
    name VARCHAR
)
TYPE IMAP
OPTIONS (
    'keyFormat' = 'bigint',
    'valueFormat' = 'json-flat'
);

SINK INTO employees(__key, age, name) VALUES (1, 10, 'Yalçın Görkem Tekol');
SINK INTO employees(__key, age, name) VALUES (2, 1, 'Ezgi Tekol');
SINK INTO employees(__key, age, name) VALUES (3, 39, 'Gülhan Tekol');
SINK INTO employees(__key, age, name) VALUES (4, 43, 'Yüce Tekol');
SINK INTO employees(__key, age, name) VALUES (5, 12, 'Ozan İltutmuş Tekol');
SINK INTO employees(__key, age, name) VALUES (6, 24, 'Ali Yankı Tekol');
