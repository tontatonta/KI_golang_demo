# 公式のGoイメージをベースにする
FROM golang:1.23.2

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main .

# 8080番ポートを公開
EXPOSE 8080

# アプリを起動
CMD ["./main"]
