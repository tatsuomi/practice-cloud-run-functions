# ===== ステージ1: ビルド用 =====
FROM golang:1.23-alpine AS builder

# 作業ディレクトリ設定
WORKDIR /app

# Goのソースコードをコピー
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# バイナリをビルド
RUN go build -o api-read reference-api/main.go

# ===== ステージ2: 実行用（軽量） =====
# ステージ1で作成したバイナリのみを使用する
FROM alpine:latest

# 作業ディレクトリ
WORKDIR /root/

# ビルド済みバイナリをコピー
COPY --from=builder /app/api-read .

# コンテナ起動時のコマンド
CMD ["./api-read"]

# ポートを公開
EXPOSE 8080
