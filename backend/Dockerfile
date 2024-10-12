# ベースイメージを指定
FROM golang:1.23.2 AS builder

# ワーキングディレクトリを設定
WORKDIR /app

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Goのビルドを実行
RUN go build -o server ./main.go

# 実行時に実行するコマンドを指定
FROM alpine:latest

# ビルドしたバイナリを新しいイメージにコピー
WORKDIR /root/
COPY --from=builder /app/server .

# 実行コマンドを指定
CMD ["./server"]

# コンテナがリッスンするポートを指定
EXPOSE 8000
