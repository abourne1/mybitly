CREATE TABLE short_link (
    id serial,
    slug varchar(100),
    url varchar(500),
    datestr varchar(25),
    created_at int,
    PRIMARY KEY (id)
);

-- I should create an index on slug to improve lookup performance