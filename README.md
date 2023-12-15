# URL Shortener

## Description
This is a Web Application made in [Golang](https://go.dev/) as a URL shortener.
This application used [PostgreSQL](https://www.postgresql.org/) running on [Docker](https://www.docker.com/) in order to have a database.

To build this it was also used the [GIN Framework](https://gin-gonic.com/) as Web Framework and [GORM Framework](https://gorm.io) to handle the data. 

## How to run it:

1. Since the models will be automated created as tables in the database, to start the application you can simply make the following command, and then you can start to make requests to `localhost:8080`.

```shell
make run
```

## How to test it:

### Create a new short version of a URL:
 * POST request to `localhost:8080/short-url`

```json
{
  "longUrl": "https://github.com/renatompf"
}
```

### Get (and redirects) a longer URL:
 * DELETE request to `localhost:8080/short-url/:shortURL`
