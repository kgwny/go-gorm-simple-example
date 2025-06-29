# goのイメージをDockerHubから流用する（Alpine Linux）
# https://hub.docker.com/_/golang
FROM golang:1.24.4-alpine3.22

# Linuxパッケージ情報のアップデートと git と bash のインストール
RUN apk update && apk add --no-cache git bash

# ログのタイムゾーンを指定する
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# コンテナ内の作業ディレクトリを指定する
WORKDIR /app

# ソースコードをコンテナ内にコピー
COPY ./app ./

# 追加: wait-for-it をコピー
COPY ./app/wait-for-it.sh ./wait-for-it.sh
RUN chmod +x ./wait-for-it.sh

# /app/go.modに記載された依存関係の解決＋必要なパッケージのダウンロードを実行
RUN go mod download

# Airのバイナリをインストール
RUN go install github.com/air-verse/air@latest

# コンテナの公開するポートを指定
EXPOSE 8080

# 起動時のコマンド（airを使用する）
CMD ["air", "-c", ".air.toml"]
