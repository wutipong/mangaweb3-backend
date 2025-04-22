# MangaWeb3-Backend

Backend server application for MangaWeb3.

## Configurations

Configurations are all based on environment variable value. If the `.env` file exists in the path that the application runs, it will be also used.

### Default values

| Variable              | Default value                                       | Description                                                                                                        |
| --------------------- |---------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| `MANGAWEB_ADDRESS`    | `:8972`                                             | The address the server runs at.                                                                                    |
| `MANGAWEB_DATA_PATH`  | `./data`                                            | The path where data files is kept. Can be either relative or absolute.                                             |
| `MANGAWEB_CACHE_PATH` | `./cache`                                           | The path where cache data is created. Can be either relative or absolute.                                          |
| `MANGAWEB_DB`         | `postgres://postgres:password@localhost:5432/manga` | Database connection string.                                                                                        |
| `MANGAWEB_DB_TYPE`    | `postgres`                                          | The type of database. Can be only either `postgres` or `sqlite3`.                                                  |
| `MANGAWEB_ENVIRONMENT`|                                                     | The environment type the server run as. Can be set to `development` for more readable logs. Ignored otherwise      |
