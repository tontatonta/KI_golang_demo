# 使用するGoのイメージを指定
FROM golang:1.23.2

# アプリケーションの作業ディレクトリを作成
WORKDIR /app

# モジュールの設定ファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係のインストール
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリをビルド
RUN go build -o server .

# ポート8080でアプリを提供
CMD ["/app/server"]
