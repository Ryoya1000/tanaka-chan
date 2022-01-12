## 環境変数の設定
```shell
heroku config:set LINE_BOT_CHANNEL_SECRET=<シークレットトークン>
heroku config:set LINE_BOT_CHANNEL_TOKEN=<チャネルアクセストークン>
```

## heroku push
```shell
git push heroku main
```

## heroku実行
```shell
heroku run ./bin/tanaka-chan
```

