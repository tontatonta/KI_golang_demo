# 使用するGoのバージョンを指定
FROM golang:1.23.2

# 作業ディレクトリを設定
WORKDIR /app

# Goのモジュールをコピー
COPY go.mod ./
RUN go mod download

# アプリのソースコードをコピー
COPY . .

# アプリをビルド
RUN go build -o main .

# アプリを実行
CMD ["./main"]
