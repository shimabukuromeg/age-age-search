# ageage-collector 技術ドキュメント

## アーキテクチャ概要

`ageage-collector`はGoで実装されたウェブスクレイパーで、沖縄テレビの「あげあげ」ページから飲食店情報を抽出し、構造化されたデータとしてSQLiteまたはPostgreSQLデータベースに保存します。

## 主要コンポーネント

### データモデル

```
┌────────────────┐       ┌─────────────────┐
│   Meshi        │       │   Municipality   │
├────────────────┤       ├─────────────────┤
│ ID             │       │ ID               │
│ ArticleID      │       │ Name             │
│ Title          │       │ Zipcode          │
│ ImageURL       │       └─────────────────┘
│ StoreName      │                ▲
│ Address        │                │
│ SiteURL        │                │
│ PublishedDate  │                │
│ Latitude       │                │
│ Longitude      ├────────────────┘
│ MunicipalityID │
└────────────────┘
```

### 外部依存性

- **goquery**: HTMLパース用ライブラリ
- **ent**: データベースのORM
- **lib/pq**: PostgreSQLドライバー
- **mattn/go-sqlite3**: SQLiteドライバー

### 外部API連携

1. **国土地理院API**:
   - エンドポイント: `https://msearch.gsi.go.jp/address-search/AddressSearch`
   - 用途: 住所から緯度・経度情報を取得
   - パラメータ: `q=<住所文字列>`
   - レスポンス形式: GeoJSON

2. **Zipcloud API**:
   - エンドポイント: `https://zipcloud.ibsnet.co.jp/api/search`
   - 用途: 郵便番号から市町村情報を取得
   - パラメータ: `zipcode=<郵便番号>`
   - レスポンス形式: JSON

## 主要アルゴリズム

### 記事抽出フロー

1. ページURLからHTML取得
2. `goquery`で記事リンク要素を選択
3. 各記事ページへ遷移し詳細データ抽出
4. 店舗名と住所を抽出

### 住所処理

```go
// 郵便番号と住所の抽出（正規表現使用）
func GetZipcodeAndAddress(fullAddress string) (string, string, error) {
    r := regexp.MustCompile(`〒([0-9]{3})-([0-9]{4})\s?(.*)`)
    match := r.FindStringSubmatch(fullAddress)
    if len(match) > 3 {
        zipCode := match[1] + match[2]         // 郵便番号
        address := strings.TrimSpace(match[3]) // 住所
        return zipCode, address, nil
    }
    return "", "", fmt.Errorf("unable to find postal code and address in: %s", fullAddress)
}

// 市町村名の抽出（正規表現使用）
func GetMunicipalityByAddress(address string) (string, error) {
    r := regexp.MustCompile(`(沖縄県)?([^市町村]*郡)?([^市町村]*?[市町村])`)
    match := r.FindStringSubmatch(address)
    if len(match) > 3 {
        return match[3], nil // 市町村名を返す
    }
    return "", fmt.Errorf("unable to find municipality in: %s", address)
}
```

### データ保存の最適化

- **UPSERT処理**: 同一ArticleIDの重複登録を防止
- **トランザクション管理**: エンティティ間の整合性確保

```go
// Entを使用したUPSERT処理例
id, err := client.Meshi.
    Create().
    SetArticleID(article.ArticleID).
    // ... その他のフィールド ...
    OnConflictColumns("article_id").
    UpdateNewValues().
    ID(ctx)
```

## コマンドラインオプション

| オプション | デフォルト値 | 説明 |
|-----------|------------|------|
| `-t` | `sqlite3` | データベースタイプ（`sqlite3`または`postgres`） |
| `-d` | `file:database.sqlite?_fk=1` | データベース接続文字列 |
| `-target` | `first` | スクレイピング対象（`first`または`all`） |
| `-isCreateSchema` | `false` | スキーマ作成を実行するかどうか |

## エラーハンドリング

- **カスタムエラー型**: 特定のエラーケースに対応
- **エラーラッピング**: コンテキスト情報を保持
- **グレースフルリカバリー**: 一部データの取得失敗時も処理継続

## パフォーマンス最適化

- **リクエスト間隔制御**: `time.Sleep(time.Second * 1)` でレート制限対応
- **段階的なデータベースアクセス**: 存在確認後の挿入処理

## 今後の改善点

- コネクションプーリングの実装
- エラーリトライメカニズムの強化
- より堅牢な地理情報正規化処理
- 並列スクレイピングの実装（適切なレート制限付き） 