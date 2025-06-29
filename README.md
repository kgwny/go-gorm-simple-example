# go-gorm-simple-example

## セットアップ

### golang のバージョンアップ（Mac/Homebrew）
```
brew upgrade go
```

### golang のバージョン確認
```
go version
```

### golang のバージョン確認結果
```
❯ go version
go version go1.24.4 darwin/arm64
```

### go mod 初期化
```
go mod init my-app
```

### echo のインストール
```
go get -u github.com/labstack/echo/v4
```

### gorm のインストール
```
go get -u gorm.io/gorm
```

### mysql のインストール
```
go get -u gorm.io/driver/mysql 
```

## 動作確認

### air + Dockerfile + docker-compose.yml によるライブリロード
```
docker compose up --build
```

起動後に以下のログが出力される
```
sample-api  |
sample-api  |   __    _   ___
sample-api  |  / /\  | | | |_)
sample-api  | /_/--\ |_| |_| \_ v1.62.0, built with Go go1.24.4
sample-api  |
sample-api  | mkdir /app/tmp
sample-api  | watching .
sample-api  | watching db
sample-api  | watching models
sample-api  | watching routes
sample-api  | !exclude tmp
sample-api  | building...
sample-api  | running...
sample-api  | start.
sample-api  | dbUser = appuser
sample-api  | dbPassword = apppass
sample-api  | dbName = sample
sample-api  | dsn = appuser:apppass@tcp(db:3306)/sample?charset=utf8mb4&parseTime=true&loc=Local
sample-api  |
sample-api  |    ____    __
sample-api  |   / __/___/ /  ___
sample-api  |  / _// __/ _ \/ _ \
sample-api  | /___/\__/_//_/\___/ v4.13.4
sample-api  | High performance, minimalist Go web framework
sample-api  | https://echo.labstack.com
sample-api  | ____________________________________O/_______
sample-api  |                                     O\
sample-api  | ⇨ http server started on [::]:8080
```

ブラウザを起動して、以下のURLへアクセスする  
http://localhost:8080/user  
API実行によりレスポンスが返却されたら確認OK  

### dockerコンテナ内の環境変数の確認
```
docker compose exec api env
```

### docker-compose.yml で定義された設定を検証する
```
docker compose config
```

### dockerコンテナ内のdbへ接続する
```
docker-compose exec db bash
```

mysql-client のインストール
```
apk add mysql-client
```

mysqlにログイン
```
mysql -h db -u appuser -papppass sample
```

### dockerコンテナ停止＆イメージ削除
```
docker-compose down -v
```
