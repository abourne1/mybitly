# MyBitly

A URL shortenting API

## Getting Started

- go build
- go run main.go

## Design decisions

- Postgresql
    -- I started this project assuming that I was going to use a NoSQL style db to store short links. This would decrease read time, and improve redirection speed.
    -- However, I also wanted to use auto-incrementing primary keys as the base10 representation of my randomly generated short links
    -- So, I decided that I would use a relational database to store my short links, and a NoSQL cache (probably redis) to store recently accessed links
    -- Unfortunately, I ran out of time to implement the cache, so for now, this solution just relies on Postgres

- Data Model
    -- I decided to store randomly generated short links and custom short links in the same table b/c I thought it would save me some time
    -- Obviously there are some drawbacks to this approach. For one, it reduces the total number of short links I can make. Assuming that I want a fixed-width url shortener (like bit.ly and goo.gl), then there's at most 62^7 short links I can make. Placing custom links in the same table as my random links means that I have less unique keys to generate slugs from.
