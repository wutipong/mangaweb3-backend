# MangaWeb3-Backend

Backend server application for MangaWeb3.

## Configurations

Configurations are all based on environment variable value. If the `.env` file exists in the path that the application runs, it will be also used.

### Default values

| Variable              | Default value                                       | Description                                                                                                        |
| --------------------- |------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `MANGAWEB_ADDRESS`    | `:8972`                                   | The address the server runs at.                                                                                    |
| `MANGAWEB_DATA_PATH`  | `./data`                                  | The path where data files is kept. Can be either relative or absolute.                                             |
| `MANGAWEB_CACHE_PATH` | `./cache`                                 | The path where cache data is created. Can be either relative or absolute.                                          |
| `MANGAWEB_DB`         | `file:db.sqlite3?_pragma=foreign_keys(1)` | Database connection string.                                                                                        |
| `MANGAWEB_DB_TYPE`    | `sqlite3`                                 | The type of database. Can be only either `postgres` or `sqlite3`.                                                  |
| `MANGAWEB_ENVIRONMENT`|                                           | The environment type the server run as. Can be set to `development` for more readable logs. Ignored otherwise      |

### Database Type

**MangaWeb3 backend** supports 2 different types of database, **SQLite** for development and **Postgres** for production.

In order to use **Postgres** as the database server, change the value of the `MANGAWEB_DB_TYPE` to `postgres` and `MANGAWEB_DB` to a connection string, for example `postgres://postgres:password@localhost:5432/manga`.

For **SQLite**, set the value of `MANGAWEB_DB_TYPE` to `sqlite` and the `MANGAWEB_DB` to a connection string, for example `file:db.sqlite3?_pragma=foreign_keys(1)`. Keep in mind that the foreign key support is required. For more information about foreign key support please visit [SQLite Foreign Key Support](https://sqlite.org/foreignkeys.html)