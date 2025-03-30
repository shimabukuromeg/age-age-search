# age-age-search

This tool scrapes information from a designated [ageage](https://www.otv.co.jp/okitive/collaborator/ageage/page/1/) website and saves it into a SQLite or PostgreSQL database.


## Setup and Running

This project is written in Go and requires a Go development environment. Please make sure you have Go installed and your PATH is set properly.

```bash
$ go build -o ageage-collector cmd/ageage-collector/main.go
```

This will start the tool with default settings. By default, it uses a SQLite database (database.sqlite).

### Using SQLite

```bash
$ ./ageage-collector
```

### Using PostgreSQL

You can switch to a PostgreSQL database by using the -t option. You can specify the PostgreSQL connection information using the -d option or the DSN environment variable.

```bash
$ docker-compose up -d
$ export DSN="postgresql://postgres@localhost:5555/ageagedb?sslmode=disable"
$ ./ageage-collector
```

If you specify -t postgres but do not specify -d or DSN, the tool will prompt you to provide the PostgreSQL connection information.

### Command Options

Various options are available for customizing the scraping process:

```bash
# Scrape only the first page
$ ./ageage-collector -target first

# Scrape all pages
$ ./ageage-collector -target all

# Create database schema if it doesn't exist
$ ./ageage-collector -isCreateSchema true

# Run with 4 parallel workers (default is 1)
$ ./ageage-collector -workers 4

# Set rate limit to 2 requests per second (default is 1.0)
$ ./ageage-collector -rps 2.0

# Combined options example
$ ./ageage-collector -target first -workers 4 -rps 1.0 -isCreateSchema true
```

#### Parallel Scraping

The tool now supports parallel scraping with multiple workers, which can significantly speed up the scraping process. Use the following options to control parallelism:

- `-workers`: Set the number of concurrent workers (default: 1)
- `-rps`: Set the maximum rate limit in requests per second (default: 1.0)

For example, to scrape with 4 workers at a rate of 2 requests per second:

```bash
$ ./ageage-collector -workers 4 -rps 2.0
```

Note: Be mindful of the target website's load capacity when increasing these values.

### GraphQL Access

The scraped data can be accessed via a GraphQL interface. The server can be started with the following command:

```bash
$ go build -tags netgo -ldflags '-s -w' -o app ./cmd/ageage-server
$ ./app
```

After running the command, you can access GraphiQL by navigating to http://localhost:8081 on your web browser.

For PostgreSQL, you need to set the DSN environment variable:

```bash
$ export DSN="postgresql://postgres@localhost:5555/ageagedb?sslmode=disable"
$ go build -tags netgo -ldflags '-s -w' -o app ./cmd/ageage-server
$ ./app
```

The server assumes that data has already been scraped and exists in the database.

### Sample Queries

After the server is started, you can issue queries to the scraped data through the GraphQL interface. Here are some examples:

```graphql
query {
  municipalities(where: { or: [{ nameContains: "那覇" }] }) {
    edges {
      node {
        id
        name
        meshis(first: 100) {
          totalCount
          edges {
            node {
              __typename
              id
              title
              address
              storeName
              siteURL
              imageURL
            }
          }
        }
      }
      cursor
    }
  }
}
```
