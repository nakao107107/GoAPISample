## プロジェクトの開始
``` bash
# infra ディレクトリへ移動
cd infra

# docker-compose build
docker-compose build

# databaseコンテナに接続
docker-compose run db bash 

# データベースを作成
mysql -uroot -p #password入力
create database sample;
exit

## 環境変数のコピー
cp ../app/.env.example ../app/.env

#docker-compose up
docker-compose up
```

localhost:8080/posts で空配列が帰ってくればOK!