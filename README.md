# age-age-search

This tool scrapes information from a designated [ageage](https://www.otv.co.jp/okitive/collaborator/ageage/page/1/) website and saves it into a SQLite or PostgreSQL database.

## Setup and Running

This project is written in Go and requires a Go development environment. Please make sure you have Go installed and your PATH is set properly.

```bash
$ go build -o age-age-search cmd/ageage-collector/main.go
```

This will start the tool with default settings. By default, it uses a SQLite database (database.sqlite).

### Using PostgreSQL

You can switch to a PostgreSQL database by using the -t option. You can specify the PostgreSQL connection information using the -d option or the DSN environment variable.

```bash
$ docker-compose up -d
$ export DSN="postgresql://postgres@localhost:5555/ageagedb?sslmode=disable"
$ ./age-age-search
```

If you specify -t postgres but do not specify -d or DSN, the tool will prompt you to provide the PostgreSQL connection information.
