# goのイメージをDockerHubから流用する(Alpine Linux)
FROM golang:1.24.3-alpine

# コンテナ内の作業ディレクトリを指定
WORKDIR /app

# Airのバイナリをインストール
RUN go install github.com/air-verse/air@latest

# 起動時のコマンド(airを使用する)
CMD ["air"]
