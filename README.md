# GOTH

## (Go, Templ, htmx)

This simple todo app connects to a postgresql database which need a table like this

```sql
CREATE TABLE todos (
    id   SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    done INTEGER DEFAULT 0
);
```

Create a .env file in root directory with a url to connect to the postgresql database.

```
DATABASE_URL=postgres://*username*:*password*@localhost:*port*/*database_name*
```
