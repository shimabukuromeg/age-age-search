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
$ ./ageage-collector -target single

# Scrape all pages
$ ./ageage-collector -target all

# Create database schema if it doesn't exist
$ ./ageage-collector -isCreateSchema true

# Run with 4 parallel workers (default is 8)
$ ./ageage-collector -workers 4

# Set rate limit to 2 requests per second (default is 2.0)
$ ./ageage-collector -rps 2.0

# Limit to scrape only the latest 10 articles
$ ./ageage-collector -limit 10

# Combined options example
$ ./ageage-collector -target single -workers 4 -rps 1.0 -isCreateSchema true -limit 20
```

### Example Usage Scenarios

Here are some common use cases and the appropriate command options:

```bash
# 初回実行時 - スキーマ作成して最初のページのみスクレイピング
$ ./ageage-collector -isCreateSchema true -target single

# 高速処理 - 8並列で4 req/secのレート制限を設定
$ ./ageage-collector -workers 8 -rps 4.0

# サイト負荷に配慮 - 2並列で0.5 req/secの低レート設定
$ ./ageage-collector -workers 2 -rps 0.5 -target all

# 最初のページから最大20件を取得（ページに20件未満の場合はその分のみ）
$ ./ageage-collector -limit 20

# 最新の20件を確実に取得（複数ページにまたがる場合あり）
$ ./ageage-collector -target all -limit 20

# 全ページから最大100件を取得 - 制限に達したら終了
$ ./ageage-collector -target all -limit 100 -t postgres -d "postgresql://postgres@localhost:5555/ageagedb?sslmode=disable"

# 定期実行用 - 最新の50件を効率的に取得（必要なページ数だけ処理）
$ ./ageage-collector -target all -limit 50 -workers 4 -rps 2.0
```

#### Parallel Scraping

The tool now supports parallel scraping with multiple workers, which can significantly speed up the scraping process. Use the following options to control parallelism:

- `-workers`: Set the number of concurrent workers (default: 8)
- `-rps`: Set the maximum rate limit in requests per second (default: 2.0)

For example, to scrape with 4 workers at a rate of 2 requests per second:

```bash
$ ./ageage-collector -workers 4 -rps 2.0
```

Note: Be mindful of the target website's load capacity when increasing these values.

#### Target Pages

The `-target` option controls how many pages to scrape:

- `single`: Scrape only the first page (default)
- `all`: Scrape all available pages

For example, to scrape all pages:

```bash
$ ./ageage-collector -target all
```

#### Limiting Articles

If you only want to scrape a specific number of the most recent articles:

- `-limit`: Set the maximum number of articles to scrape (default: 0, meaning no limit)

For example, to scrape only the latest 50 articles:

```bash
$ ./ageage-collector -limit 50
```

Note: When using `-limit` alone, the default target (`single`) will apply, meaning only the first page will be scraped. If the first page contains fewer articles than the limit, you won't get the full number of articles. To ensure you get the exact number of articles specified by limit, combine it with `-target all`:

```bash
$ ./ageage-collector -target all -limit 20
```

When combined with `-target all`, the scraper will continue fetching pages until it reaches the specified limit, then stop. This is useful for getting a specific number of the most recent articles across multiple pages:

```bash
$ ./ageage-collector -target all -limit 20
```

This will fetch pages sequentially until it has collected 20 articles, then stop processing additional pages.

#### Creating Schema

The `-isCreateSchema` option ensures the database tables exist before scraping:

- `true`: Create database schema if it doesn't exist
- `false`: Don't create schema (default)

This is useful for first-time setup:

```bash
$ ./ageage-collector -isCreateSchema true
```

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
