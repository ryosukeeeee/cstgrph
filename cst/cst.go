package cst

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func CostAndUsage() (result *costexplorer.GetCostAndUsageOutput, err error) {
	var opt = session.Options{
		//　デフォルト以外のプロファイルを使用する場合はこの行のコメントを外して書き換える
		// Profile: "default",
		Config: aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	}
	sess := session.Must(
		session.NewSessionWithOptions(opt),
	)

	svc := costexplorer.New(sess)
	day := time.Now()
	const layout = "2006-01-02"

	// 8日前から1日前までのコストと内訳を取得する
	timeperiod := costexplorer.DateInterval{
		Start: aws.String(day.AddDate(0, 0, -8).Format(layout)),
		End:   aws.String(day.AddDate(0, 0, -1).Format(layout)),
	}
	response, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("DAILY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		GroupBy: []*costexplorer.GroupDefinition{&costexplorer.GroupDefinition{
			Key:  aws.String("SERVICE"),
			Type: aws.String("DIMENSION"),
		}},
		TimePeriod: &timeperiod,
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return response, nil
}
