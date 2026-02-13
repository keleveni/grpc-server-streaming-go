# grpc-golang

gRPCとProtocol Buffersを使用したレシピ共有アプリケーションのサンプルプロジェクトです。
Server Streamingによるリアルタイム通知機能を実装しています。

## 概要

このプロジェクトは、以下の機能を持つgRPCベースのバックエンドアプリケーションです。

### 主な機能

- **レシピ管理**: レシピの作成、取得、更新、削除（CRUD操作）
- **ユーザーフォロー**: ユーザー間のフォロー機能
- **リアルタイム通知**: フォローしたユーザーが新しいレシピを作成した際に、Server Streamingでリアルタイムに通知

### 使用技術

- **Go** 1.25.5
- **gRPC** / **Protocol Buffers**: API定義とRPC通信
- **PostgreSQL**: データベース
- **Bun**: Go用のSQL-firstなORM
- **Wire**: 依存性注入（DI）ツール
- **Atlas**: データベースマイグレーション管理
- **OpenTelemetry**: 分散トレーシング・可観測性
- **Jaeger**: トレース可視化UI
- **zerolog**: 構造化ログ

### プロジェクト構成

```
.
├── cmd/              # アプリケーションエントリーポイント
├── internal/
│   ├── app/         # アプリケーションロジック
│   ├── db/          # データベース接続
│   ├── grpc/        # gRPCサーバー実装
│   ├── orm/         # ORMモデル
│   └── pb/          # Protocol Buffers生成コード
├── proto/           # Protocol Buffers定義ファイル
│   ├── recipe/      # レシピサービス
│   ├── follow/      # フォローサービス
│   └── notification/ # 通知サービス（Server Streaming）
└── database/
    └── migrations/  # データベースマイグレーション
```

## Getting Started

### 1. インフラストラクチャの起動
```bash
docker compose up -d
```

これにより以下のサービスが起動します：
- PostgreSQL (ポート: 25432)
- Jaeger UI (ポート: 16686)
- OpenTelemetry Collector (ポート: 4317, 4318)

### 2. gRPCサーバーの起動
```bash
go generate ./...
go run ./cmd/serve
```

### 3. トレースの確認
ブラウザで http://localhost:16686 を開き、Jaeger UIでトレースを確認できます。

gRPCリクエストを送信すると、自動的にトレースが記録され、Jaeger UIで可視化されます。
## Testing Server Streaming

Server streamingの動作を確認するには、2つのターミナルを使用します。

### Terminal 1: SubscribeRecipeFeed (ストリーミングを受信)
```bash
# ユーザーをフォロー
grpcurl --plaintext \
  -d '{"follower_id":"01HZ0000000000000000000001", "followee_id":"01HZ0000000000000000000002"}' \
  localhost:50051 \
  follow.FollowService/FollowUser

# レシピフィードをストリーミング受信
grpcurl -plaintext \
  -d '{"user_id":"01HZ0000000000000000000001"}' \
  localhost:50051 \
  notification.NotificationService.SubscribeRecipeFeed
```

### Terminal 2: CreateRecipe (レシピを作成)
```bash
# 新しいレシピを作成
grpcurl -plaintext \
  -d '{"user_id":"01HZ0000000000000000000002", "title": "Pasta", "description": "Delicious pasta"}' \
  localhost:50051 \
  recipe.RecipeService/CreateRecipe
```

Terminal 1でストリーミングが開始された状態で、Terminal 2でレシピを作成すると、Terminal 1にリアルタイムでレシピ情報が配信されます。

## Observability (可観測性)

このプロジェクトはOpenTelemetryを使用した分散トレーシングを実装しています。

### アーキテクチャ

```
gRPCサーバー
  └─> OpenTelemetry SDK (OTLP Exporter)
        └─> OpenTelemetry Collector (localhost:4317)
              └─> Jaeger (localhost:14250)
                    └─> Jaeger UI (http://localhost:16686)
```

### トレース情報

各gRPCリクエストについて、以下の情報が自動的に記録されます：

- **スパン名**: gRPCメソッド名（例: `/recipe.RecipeService/CreateRecipe`）
- **属性**:
  - `rpc.service`: サービス名
  - `rpc.method`: メソッド名
  - `rpc.grpc.status_code`: gRPCステータスコード
  - その他のメタデータ
- **タイミング**: リクエストの開始時刻、終了時刻、処理時間
- **エラー情報**: エラーが発生した場合のスタックトレース

### Jaeger UIの使い方

1. http://localhost:16686 を開く
2. 左側の「Service」ドロップダウンから `grpc-golang` を選択
3. 「Find Traces」をクリック
4. トレース一覧から任意のトレースをクリックして詳細を確認

### カスタムトレーシング

独自のスパンを追加する場合：

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

func MyFunction(ctx context.Context) error {
    tracer := otel.Tracer("my-component")
    ctx, span := tracer.Start(ctx, "MyFunction")
    defer span.End()

    span.SetAttributes(attribute.String("key", "value"))

    // 処理...

    return nil
}
```
