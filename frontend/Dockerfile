# 公式のNodeイメージをベースにする
FROM node:18

# 作業ディレクトリを設定
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピー
COPY package*.json ./

# 依存関係をインストール
RUN npm install

# アプリケーションのコードをコピー
COPY . .

# アプリをビルド
RUN npm run build

# 3000番ポートを公開
EXPOSE 3000

# アプリを起動
CMD ["npm", "start"]
