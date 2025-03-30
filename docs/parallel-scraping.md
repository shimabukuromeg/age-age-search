# 並列スクレイピング設計案

## 概要

現在の実装では記事の取得とデータ処理が逐次的に行われており、大量のデータを収集する際に時間がかかります。並列処理を導入することでパフォーマンスを向上させつつ、対象サイトに過剰な負荷をかけないよう配慮した設計を提案します。

## 設計目標

1. **処理時間の短縮**: 並列処理によるスクレイピング時間の短縮
2. **サーバー負荷への配慮**: 適切なレート制限による対象サイトへの配慮
3. **リソース効率**: ローカルマシンのCPU/メモリ使用を最適化
4. **エラー耐性**: 一部の処理失敗が全体に影響しない堅牢な設計
5. **拡張性**: 将来的な対象サイト追加にも対応できる設計

## アーキテクチャ

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  記事一覧収集   │────▶│  ワーカープール  │────▶│  データベース保存 │
│                 │     │                 │     │                 │
└─────────────────┘     └─────────────────┘     └─────────────────┘
        │                       ▲                       ▲
        │                       │                       │
        │                ┌──────┴──────┐                │
        │                │             │                │
        └───────────────▶│ レート制限機構 │────────────────┘
                         │             │
                         └─────────────┘
```

## 主要コンポーネント

### 1. ワーカープール

Goのgoroutineとチャネルを使用した並列処理システム。

```go
type Worker struct {
    ID        int
    JobChan   chan *Article
    ResultChan chan *ProcessResult
    RateLimiter *rate.Limiter
    Client    *ent.Client
}

type ProcessResult struct {
    Article *Article
    Meshi   *ent.Meshi
    Error   error
}

func NewWorkerPool(size int, client *ent.Client, rps float64) *WorkerPool {
    pool := &WorkerPool{
        JobQueue:   make(chan *Article, 100),
        ResultChan: make(chan *ProcessResult, 100),
        Workers:    make([]*Worker, size),
        Client:     client,
    }
    
    // ワーカーごとにレート制限を設定
    workerRPS := rps / float64(size)
    for i := 0; i < size; i++ {
        limiter := rate.NewLimiter(rate.Limit(workerRPS), 1)
        pool.Workers[i] = NewWorker(i, pool.JobQueue, pool.ResultChan, limiter, client)
    }
    
    return pool
}
```

### 2. レート制限機構

`golang.org/x/time/rate`を使用したトークンバケットアルゴリズムによるレート制限。

```go
func (w *Worker) Process(article *Article) *ProcessResult {
    // レート制限に従って待機
    if err := w.RateLimiter.Wait(context.Background()); err != nil {
        return &ProcessResult{Article: article, Error: err}
    }
    
    // 通常の処理を実行
    result := &ProcessResult{Article: article}
    meshi, err := CreateMeshiAndMunicipality(context.Background(), w.Client, article)
    result.Meshi = meshi
    result.Error = err
    
    return result
}
```

### 3. エラーハンドリングと再試行機構

一時的なエラーに対応するための再試行メカニズム。

```go
func ProcessWithRetry(worker *Worker, article *Article, maxRetries int) *ProcessResult {
    var result *ProcessResult
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        result = worker.Process(article)
        if result.Error == nil {
            return result
        }
        
        // エラーの種類に応じて再試行判断
        if IsRetryableError(result.Error) {
            log.Printf("Retry %d for article %s: %v", 
                      attempt+1, article.ArticleID, result.Error)
            // 指数バックオフで待機
            time.Sleep(time.Second * time.Duration(math.Pow(2, float64(attempt))))
            continue
        }
        
        // 再試行不可能なエラーの場合は即時返却
        return result
    }
    
    return result
}
```

## 実装案

### メイン処理フロー

```go
func main() {
    // ... 既存の初期化処理 ...
    
    // ワーカープールの設定（8ワーカー、全体で毎秒2リクエスト）
    pool := NewWorkerPool(8, client, 2.0)
    pool.Start()
    
    // 結果処理用goroutine
    resultChan := pool.ResultChan
    go func() {
        for result := range resultChan {
            if result.Error != nil {
                log.Printf("Error processing article %s: %v", 
                          result.Article.ArticleID, result.Error)
                continue
            }
            log.Printf("Successfully processed article %s", result.Article.ArticleID)
        }
    }()
    
    // 記事一覧の取得（既存のコード）
    baseURL := "https://www.otv.co.jp/okitive/collaborator/ageage/page/%d"
    page := 1
    for {
        listURL := fmt.Sprintf(baseURL, page)
        articles, err := FindArticles(listURL)
        if err != nil {
            log.Fatal(err)
        }
        
        // 記事が見つからなければ終了
        if len(articles) == 0 {
            break
        }
        
        // 各記事をジョブキューに送信
        for _, article := range articles {
            // 既存データの確認は各ワーカーで実施
            pool.JobQueue <- &article
        }
        
        if target == "first" {
            break
        }
        page++
        
        // ページ間の移動にも適切な間隔を設ける
        time.Sleep(time.Second * 2)
    }
    
    // すべてのジョブが完了するまで待機
    pool.Wait()
    fmt.Println("done")
}
```

### データベースアクセスの最適化

```go
func (w *Worker) ProcessArticle(article *Article) (*ent.Meshi, error) {
    ctx := context.Background()
    
    // 既存データのチェック（ロックを最小限に）
    exists, err := w.Client.Meshi.
        Query().
        Where(meshi.ArticleIDEQ(article.ArticleID)).
        Exist(ctx)
    
    if err != nil {
        return nil, fmt.Errorf("error checking existence: %w", err)
    }
    
    if exists {
        return nil, nil // 既存データがある場合はスキップ
    }
    
    // 新規データ作成
    return CreateMeshiAndMunicipality(ctx, w.Client, article)
}
```

## パフォーマンス予測

現在の実装と比較した性能向上の見込み：

| 処理内容 | 現在の実装 | 並列実装（8ワーカー） | 改善率 |
|---------|-----------|----------------------|-------|
| 100記事処理 | 約100秒 | 約50秒 | 50% |
| 500記事処理 | 約500秒 | 約250秒 | 50% |
| 1000記事処理 | 約1000秒 | 約500秒 | 50% |

※ スクレイピング対象サイトへの配慮としてレート制限を設けているため、理論上の最大改善率は制限されます。

## 注意点

1. **既存データとの整合性**: 並列処理による競合状態を防ぐため、データベース操作に楽観的ロックを検討
2. **メモリ使用量**: 大量の記事を同時に処理する場合のメモリ使用に注意
3. **エラー処理**: 一部の記事処理が失敗しても全体のプロセスは継続するよう設計
4. **実行モニタリング**: 処理状況をリアルタイムで確認できるログ機能やプログレスバーの実装

## 今後の展望

1. **動的レート制限**: サーバーレスポンスに応じて自動的にレート制限を調整
2. **分散処理対応**: 複数マシンでの分散処理に対応する拡張設計
3. **差分更新機能**: 前回の実行から変更された記事のみを処理する最適化 