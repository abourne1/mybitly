CREATE TABLE short_link (
    uuid serial,
    slug varchar(100),
    url varchar(500),
    is_custom bool,
    created_at int,
    PRIMARY KEY (uuid)
);

CREATE INDEX slug_index ON short_link (slug);