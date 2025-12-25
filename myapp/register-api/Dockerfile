# ===== ステージ1: ビルド用 =====
FROM golang:1.23-alpine AS builder

# 作業ディレクトリ設定
WORKDIR /app

# Goのソースコードをコピー
COPY . .

# バイナリをビルド
RUN go build -o api-write

# ===== ステージ2: 実行用（軽量） =====
# ステージ1で作成したバイナリのみを使用する
FROM alpine:latest

# 作業ディレクトリ
WORKDIR /root/

# ビルド済みバイナリをコピー
COPY --from=builder /app/api-write .

# コンテナ起動時のコマンド
CMD ["./api-write"]

# ポートを公開
EXPOSE 8080
