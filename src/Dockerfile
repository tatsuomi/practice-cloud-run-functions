# ベースイメージをGo 1.23.10に設定
FROM golang:1.23.10

# 作業用ディレクトリの設定
WORKDIR /app

# ソースをコンテナ内に配置
COPY . .

# ソースのビルド
RUN go build -o app .
