# GOTH

## (Go, Templ, htmx)

This simple todo app uses connect to a postgresql database which need a table like this

```sql
CREATE TABLE todos (
    id   SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    done INTEGER DEFAULT 0
);
```

and a .env file that looks like this

```
DATABASE_URL=postgres://*username*:*password*@localhost:*port*/*database_name*
```
