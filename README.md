# wantedly01

## 課題1

### 実行方法

以下のコマンドを``task01``以下で実行してください

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