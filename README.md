# wantedly01

## 課題1

### 実行方法

以下のコマンドを``task01``ディレクトリ以下で実行してください

```shell
docker build -t echo .
docker run --rm -p 127.0.0.1:8080:8080 --name echo echo
```

### 実行結果

リクエスト

```shell
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```

レスポンス

```json
{
  "message": "Hello World!!"
}
```

## 課題２

リクエスト

```shell
curl -XGET -H 'Content-Type:application/json' https://task02.herokuapp.com/
```

レスポンス

```json
{
  "message": "Hello World!!"
}
```

### デプロイの方法

#### 1. main.go, Dockerfileの作成

#### 2. Heroku へログイン

```shell
heroku login
heroku container:login
```

#### 3. ビルド

```shell
docker build -t herokuexp .
docker run -e "PORT=3000" -p 3000:3000 -t herokuexp
```

#### 4. デプロイ

Herokuアプリの作成

```shell
heroku create {アプリ名}
```

push

```shell
heroku container:push web
```

release

```shell
heroku container:release web
```

ブラウザで確認

```shell
heroku open
```