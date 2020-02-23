package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type SlackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// channelにmessageを投稿する
func PostMessage(apiToken, channel, message string) error {
	values := url.Values{}
	values.Set("token", apiToken)
	values.Add("channel", channel)
	values.Add("text", message)

	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		strings.NewReader(values.Encode()),
	)
	handleError(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleError(err)
	defer resp.Body.Close()

	handleSlackReponse(resp)

	return err
}

func UploadFile(apiToken, channel string) error {
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	// Slackのトークンをセット
	err := mw.WriteField("token", apiToken)
	handleError(err)

	// 投稿先のチャンネルのIDをセット
	err = mw.WriteField("channels", channel)
	handleError(err)

	err = mw.WriteField("initial_comment", "this week AWS costs and usage")
	// チャンネルに投稿する画像を開く
	imgfile, err := os.Open(os.Getenv("IMG_PATH"))
	handleError(err)
	defer imgfile.Close()

	// 画像を書き込むパートを作成する
	fw, err := mw.CreateFormFile("file", "content")
	handleError(err)

	// 画像を書き込む
	_, err = io.Copy(fw, imgfile)
	handleError(err)

	// 終端のバウンダリを書き込む
	err = mw.Close()
	handleError(err)

	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/files.upload",
		io.Reader(body),
	)
	handleError(err)

	// バウンダリを含むcontent-typeをリクエストにセットする
	contentType := mw.FormDataContentType()
	req.Header.Set("Content-Type", contentType)

	/// [For debug]
	///	作成したリクエストをダンプする
	//
	// dump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// fmt.Println(string(dump))

	client := &http.Client{}
	resp, err := client.Do(req)
	handleError(err)
	defer resp.Body.Close()

	handleSlackReponse(resp)

	return err
}

func handleSlackReponse(resp *http.Response) {
	if resp.StatusCode == http.StatusOK {
		slackRespArray, _ := ioutil.ReadAll(resp.Body)
		slackRespJsonBytes := ([]byte)(slackRespArray)
		slackRespData := new(SlackResponse)

		// Slack APIレスポンスのokがfalseの場合、メッセージ投稿エラー
		if err := json.Unmarshal(slackRespJsonBytes, slackRespData); err != nil {
			fmt.Println("JSON Unmarshal error: ", err)
		}
		if slackRespData.Ok == false {
			fmt.Println("メッセージ投稿に失敗")
			fmt.Println("メッセージ投稿結果=[", slackRespData.Ok, "] メッセージ投稿エラーメッセージ=[", slackRespData.Error, "]")
		} else {
			fmt.Println("メッセージ投稿に成功")
		}
	} else {
		fmt.Println("メッセージ投稿に失敗")
	}
	fmt.Println("メッセージ投稿ステータスコード=[", resp.StatusCode, "] レスポンス内容=[", resp.Status, "]")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
