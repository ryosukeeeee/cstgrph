package main

import (
	"fmt"
	"os"

	"./cst"
	"./graph"
	"./slack"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() {
	// AWS Cost Explorerでコストと使用状況を取得する
	cau, err := cst.CostAndUsage()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// コストと使用状況を積み上げ棒グラフに可視化する
	graph.BarPlot(cau)

	// Slackのチャンネルに画像を投稿する
	var apiToken = os.Getenv("SLACK_TOKEN")
	var channel = os.Getenv("SLACK_CHANNEL_ID")
	slack.UploadFile(apiToken, channel)
}

func main() {
	// 実行環境がlambdaかローカルかで処理を分ける
	_, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	if ok {
		os.Setenv("IMG_PATH", "/tmp/barchart.png")
		lambda.Start(Handler)
	} else {
		os.Setenv("IMG_PATH", "barchart.png")
		Handler()
	}
}
