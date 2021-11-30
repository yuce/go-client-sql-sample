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

SINK INTO employees(__key, age, name) VALUES (1, 30, 'Mike McGregor');
SINK INTO employees(__key, age, name) VALUES (2, 41, 'Jane Brown');
SINK INTO employees(__key, age, name) VALUES (3, 22, 'Joe Taylor');
SINK INTO employees(__key, age, name) VALUES (4, 33, 'Mandy Bronson');
