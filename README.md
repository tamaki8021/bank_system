# 機能
- 顧客が口座を作成できるようにする
- 顧客がお金を引き出すことができるようにする
- 顧客が別の口座に送金できるようにする
- 顧客データと最終残高を含む取引明細を提供する
- 取引明細を印刷するためのWeb APIを、エンドポイントを通して公開する

# curl

```
$ curl http://localhost:8000/statement\?number=1001
1001 - John - 0

$ curl http://localhost:8000/deposit\?number=1001&amount=100
1001 - John - 100

```

# コマンド

```
// test
$ to test -v
```