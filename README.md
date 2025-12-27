Next(Frontend)＋Go(Backend)＋PostgreSQL＋nginxで作る、Docker Composeの雛形
---
使い方は以下のコマンドを打つ

### 開発環境
1回目
```
./dev_build.sh
```

2回目以降
```
./dev_run.sh
```

停止
```
./dev_stop.sh
```


### 確認環境
1回目
```
./prod_build.sh
```

2回目以降
```
./prod_run.sh
```

停止
```
./prod_stop.sh
```

### VSCode (Dev Containers) での起動
frontend / api のどちらにもアタッチ可能

1. VSCode で `Dev Containers: Open Folder in Container...` を実行
2. `.devcontainer/frontend` または `.devcontainer/api` を選択
3. コンテナが起動し、デフォルトの作業ディレクトリで開く
