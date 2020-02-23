## AWS CostAndUsage Monitor

### 設定
Makefileの2行目、３行目を書き換える

```
PROGRAM=cost_explorer
SLACK_TOKEN=(slack-token)
SLACK_CHANNEL_ID=(slack-channel)
```

### ローカルで実行

```
$ make debug
```

### Deploy

```
$ make deploy
```

### Reference

#### Go language
[Goプログラミング言語仕様 - golang.jp](http://golang.jp/go_spec)

[Go言語でSlack,Chatwork,LINEへチャットメッセージを投稿するコマンド作成 (Windows,Linux,Macに対応したコマンドを作成する) - Qiita](https://qiita.com/na0AaooQ/items/7f937ff1fd938eb89b6a#_reference-a1d0bb5cd8f60a21af76)

[Go言語のpackageの作り方: 長くなったコードを別ファイルに切り出す方法 - Qiita](https://qiita.com/suin/items/f4ad02e4123c8a3c75eb)

[Goの構造体にメタ情報を付与するタグの基本 - Qiita](https://qiita.com/itkr/items/9b4e8d8c6d574137443c)

[Golang の defer 文と panic/recover 機構について - CUBE SUGAR CONTAINER](https://blog.amedama.jp/entry/2015/10/11/123535)

[Goのarrayとsliceを理解するときがきた - Qiita](https://qiita.com/seihmd/items/d9bc98a4f4f606ecaef7)

[【Go言語】assignment to entry in nil map - DRYな備忘録](https://otiai10.hatenablog.com/entry/2014/08/09/154256)

[golang チートシート - Qiita](https://qiita.com/jca02266/items/56a4fb7b07b692a6bf34#%E9%85%8D%E5%88%97%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9)

[Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](http://matope.hatenablog.com/entry/2014/04/22/101127)

[Goで、mapをrangeでイテレーションすると、取り出す順番は実行ごとに異なる罠 - Qiita](https://qiita.com/yuki2006/items/5a43644e278c0777ca52)

[[Go]小数点第Nで四捨五入する - Qiita](https://qiita.com/k-kurikuri/items/bba031bbe8ff17b21890)

[Goで学ぶポインタとアドレス - Qiita](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)

[Go言語のmime/multipartパッケージでファイルをアップロードしましょう - Eureka Engineering - Medium](https://medium.com/eureka-engineering/multipart-file-upload-in-golang-c4a8eb15a3ee)

[Go net/httpパッケージの概要とHTTPクライアント実装例 - Qiita](https://qiita.com/jpshadowapps/items/463b2623209479adcd88)

#### AWS
[costexplorer - Amazon Web Services - Go SDK](https://docs.aws.amazon.com/sdk-for-go/api/service/costexplorer/#CostExplorer.GetCostAndUsage)

[AWS Lambda 環境変数の使用 - AWS Lambda](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime)

[aws-sdk-go を使用するときの注意点 | TECHSCORE BLOG](https://www.techscore.com/blog/2017/12/25/aws-sdk-go-を使用するときの注意点/)

[Serverless Framework(Go) でHello worldしてみる - Qiita](https://qiita.com/seike460/items/b54a61ec8d07b2be8c87)

#### Slack
[files.upload method | Slack](https://api.slack.com/methods/files.upload)

[chat.postMessage method | Slack](https://api.slack.com/methods/chat.postMessage)


#### Makefile

[Go言語開発を便利にするMakefileの書き方 - Qiita](https://qiita.com/yoskeoka/items/317a3afab370155b3ae8)

[2016年だけどMakefileを使ってみる - Qiita](https://qiita.com/petitviolet/items/a1da23221968ee86193b)
