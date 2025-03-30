package scraper

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/shimabukuromeg/ageage-search/ent"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"golang.org/x/time/rate"
)

// Article はスクレイピング対象の記事情報を表します
type Article struct {
	ArticleID      string
	Title          string
	ImageURL       string
	StoreName      string
	Address        string
	SiteURL        string
	PublishedDate  string
	MunicipalityID int
}

// ProcessResult は記事処理の結果を表します
type ProcessResult struct {
	Article *Article
	Meshi   *ent.Meshi
	Error   error
}

// Worker はスクレイピング処理を行うワーカーを表します
type Worker struct {
	ID          int
	JobChan     chan *Article
	ResultChan  chan *ProcessResult
	RateLimiter *rate.Limiter
	Client      *ent.Client
	Processor   func(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error)
}

// NewWorker は新しいワーカーを作成します
func NewWorker(id int, jobChan chan *Article, resultChan chan *ProcessResult, 
    limiter *rate.Limiter, client *ent.Client, 
    processor func(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error)) *Worker {
	
	return &Worker{
		ID:          id,
		JobChan:     jobChan,
		ResultChan:  resultChan,
		RateLimiter: limiter,
		Client:      client,
		Processor:   processor,
	}
}

// Start はワーカーを起動します
func (w *Worker) Start() {
	log.Printf("Starting worker %d", w.ID)
	go func() {
		for article := range w.JobChan {
			log.Printf("Worker %d processing article %s", w.ID, article.ArticleID)
			// レート制限に従って待機
			if err := w.RateLimiter.Wait(context.Background()); err != nil {
				result := &ProcessResult{
					Article: article,
					Error:   fmt.Errorf("rate limiter error: %w", err),
				}
				w.ResultChan <- result
				continue
			}

			result := w.Process(article)
			w.ResultChan <- result
		}
		log.Printf("Worker %d finished", w.ID)
	}()
}

// Process は記事を処理します
func (w *Worker) Process(article *Article) *ProcessResult {
	ctx := context.Background()
	
	// 既存データのチェック
	exists, err := w.Client.Meshi.
		Query().
		Where(meshi.ArticleIDEQ(article.ArticleID)).
		Exist(ctx)
	
	if err != nil {
		return &ProcessResult{
			Article: article,
			Error:   fmt.Errorf("error checking existence: %w", err),
		}
	}
	
	if exists {
		// 既存データがある場合はスキップ
		return &ProcessResult{
			Article: article,
			Error:   nil,
		}
	}
	
	// 新規データ作成
	result := &ProcessResult{
		Article: article,
	}
	meshi, err := w.Processor(ctx, w.Client, article)
	result.Meshi = meshi
	result.Error = err
	
	return result
}

// WorkerPool はワーカーのプールを表します
type WorkerPool struct {
	Workers    []*Worker
	JobQueue   chan *Article
	ResultChan chan *ProcessResult
	Client     *ent.Client
	wg         sync.WaitGroup
}

// NewWorkerPool は新しいワーカープールを作成します
func NewWorkerPool(size int, client *ent.Client, rps float64, 
    processor func(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error)) *WorkerPool {
	
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
		pool.Workers[i] = NewWorker(i, pool.JobQueue, pool.ResultChan, limiter, client, processor)
	}
	
	return pool
}

// Start はワーカープールを起動します
func (p *WorkerPool) Start() {
	log.Printf("Starting worker pool with %d workers", len(p.Workers))
	for _, worker := range p.Workers {
		p.wg.Add(1)
		go func(w *Worker) {
			defer p.wg.Done()
			w.Start()
		}(worker)
	}
}

// IsRetryableError はエラーが再試行可能かどうかを判定します
func IsRetryableError(err error) bool {
	// ネットワークエラーなど再試行可能なエラーの判定ロジック
	// 実際の実装では具体的なエラータイプに応じて判定
	return true
}

// ProcessWithRetry は再試行機能付きで記事を処理します
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