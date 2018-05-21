CREATE TABLE short_link_visit (
    uuid serial,
    slug varchar(100),
    datestr varchar(25),
    created_At int,
    PRIMARY KEY (uuid)
);