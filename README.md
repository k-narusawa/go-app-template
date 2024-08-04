# go-rest-template

Golang で簡易な RestAPI を作成用のテンプレート

## 起動

```shell
make dev # ログを確認したい場合
make up
```

## テスト

```shell
make test
```

## 補足

### Docker

```shell
docker build -t go-rest-template .
docker run -p 8080:8080 go-rest-template
```
