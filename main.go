package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}

	// ランダムに送りたいメッセージ一覧
	text := []string{
		"僕は、犬を買ってて名前は、「ひげだるま」っていうんだ。ひげだるまは、基本的に人なつこくて優しい心の持ち主なんだけど、家族の中で唯一僕とだけベッドで一緒に寝てくれないんだ。最近、お風呂に入ってないからかな？？",
		"なにか失敗した時に、友人から「どんまい」って言われるのムカつかない？",
		"中学の時、雨が降るとテンション上がる〜って言ってた友達の気持ちが今だによくわからない。。",
		"ぴ〜か〜ぶ〜〜〜〜！！！！",
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(text))
	fmt.Println(text[num])
	result := text[num]

	// // テキストメッセージを生成する
	message := linebot.NewTextMessage(result)
	// // テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
