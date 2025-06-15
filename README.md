# go-gorm-simple-example

## このリポジトリで最初にやること

### go mod の導入 -> 自身のGitHubリポジトリに合わせて初期化する
```
go mod init github.com/kgwny/go-gorm-simple-example
```

### godotenv のインストール -> .env による環境変数の取り扱いを可能にする
```
go get github.com/joho/godotenv
```

### golang-migrate のインストール -> DBのマイグレーションをできるようにする
```
go get github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/mysql
go get -u github.com/golang-migrate/migrate/v4/source/file
```

### go mod の更新
```
go mod tidy
```


## その他

### Dockerfileで使用するdockerイメージは下記リンク先から確認できる
https://hub.docker.com/_/golang/tags
