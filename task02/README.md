# デプロイの方法

## 1. main.go, Dockerfileの作成

## 2. Heroku へログイン

```shell
heroku login
heroku container:login
```

## 3. ビルド

```shell
docker build -t herokuexp .
docker run -e "PORT=3000" -p 3000:3000 -t herokuexp
```

## 4. デプロイ

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