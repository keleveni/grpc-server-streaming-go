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
```
go generate ./...
go run ./cmd/serve
```
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
