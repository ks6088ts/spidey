# gRPC

## grpcurl による疎通確認

[grpcurl](https://github.com/fullstorydev/grpcurl) をインストールする。

```bash
go get github.com/fullstorydev/grpcurl/...
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
grpcurl -help
```

[docker-compose.yaml](../docker-compose.yaml) の account サービスにポートフォワーディング設定を追加する。

```yaml
services:
  account:
+    ports:
+      - 8080:8080
```

サービスを起動

```bash
docker-compose up -d --build
```

[gRPC サーバーの動作確認を grpcurl でやってみた](https://qiita.com/yukina-ge/items/a84693f01f3f0edba482) を参考に grpcurl による動作確認

```bash
> grpcurl -plaintext localhost:8080 list
grpc.reflection.v1alpha.ServerReflection

> grpcurl -plaintext localhost:8080 list pb.AccountService
pb.AccountService.GetAccount
pb.AccountService.GetAccounts
pb.AccountService.PostAccount

> grpcurl -plaintext -d @ localhost:8080 pb.AccountService/PostAccount
{
"name": "hello"
}
EOM # WindowsではCtrl+Z、Unix系ではCtrl+D
{
  "account": {
    "id": "1hJCwfq0SuSWNrTA1VuNHxLhUkL",
    "name": "hello"
  }
}

> grpcurl -plaintext localhost:8080 pb.AccountService/GetAccounts
{
  "accounts": [
    {
      "id": "1hJCwfq0SuSWNrTA1VuNHxLhUkL",
      "name": "hello"
    }
  ]
}
```

# GraphQL

## gqlgen

API サーバの実装入門として [開発ライブ実況 #2 GoLand 編 (メルペイ Solutions チーム デフォルト好きエンジニア)](https://www.youtube.com/watch?v=8MdxqDb07eQ) を参照した。

```
query {
  todos {
    id
    name
    createdAt
  }
}

mutation Create {
  createTodo(input: {name: "text A"}) {
    id
    name
    createdAt
  }
}

mutation Update {
  updateTodo(
    input: {
      id: "49337d7b-efaf-4821-99dc-299e56ceba4a"
      name: "text A!"
    }
  ){
    id
    name
  }
}
```
