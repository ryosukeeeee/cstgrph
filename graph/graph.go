package graph

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type GraphSource struct {
	date    []string
	dataSet map[string]*ServiceData
}

type ServiceData struct {
	serviceName string
	costs       []float64
}

func BarPlot(c *costexplorer.GetCostAndUsageOutput) error {
	gs, err := parser(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	p, err := plot.New()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	// グラフにタイトルを追加する
	p.Title.Text = "Bar chart"

	// 棒グラフの横幅を指定する
	w := vg.Points(20)

	var bars = []*plotter.BarChart{}
	var cnt = 0
	for key, sd := range gs.dataSet {
		fmt.Println(key, ":", sd.serviceName, sd.costs)
		groupA := plotter.Values(sd.costs)
		bar, err := plotter.NewBarChart(groupA, w)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		// 枠線の太さは0にする
		bar.LineStyle.Width = vg.Length(0)
		// 塗りつぶし色を指定する
		bar.Color = plotutil.Color(cnt)
		// 凡例を追加する
		p.Legend.Add(sd.serviceName, bar)

		bars = append(bars, bar)
		if cnt != 0 {
			// 棒グラフを積み上げる
			bar.StackOn(bars[cnt-1])
		}
		p.Add(bar)
		cnt = cnt + 1
	}

	// 凡例を上部に表示する
	p.Legend.Top = true

	// x軸の目盛に日付を設定する
	p.NominalX(gs.date...)

	if err := p.Save(600, 300, os.Getenv("IMG_PATH")); err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return nil
}

func parser(c *costexplorer.GetCostAndUsageOutput) (GraphSource, error) {
	gs := GraphSource{}
	gs.dataSet = map[string]*ServiceData{}

	for _, results := range c.ResultsByTime {
		// xlabel用に日付を配列に追加する
		gs.date = append(gs.date, aws.StringValue(results.TimePeriod.Start))
		if len(results.Groups) != 0 {
			for _, group := range results.Groups {
				serviceName := aws.StringValue(group.Keys[0])
				if _, isExist := gs.dataSet[serviceName]; isExist != true {
					sd := ServiceData{
						serviceName: serviceName,
					}
					gs.dataSet[serviceName] = &sd
				}
			}
		}
	}

	for _, results := range c.ResultsByTime {
		one_day_result := map[string]float64{}

		for _, group := range results.Groups {
			var key = aws.StringValue(group.Keys[0])
			// 小数点第3桁を四捨五入
			f, _ := strconv.ParseFloat(
				aws.StringValue(group.Metrics["UnblendedCost"].Amount),
				32,
			)
			f = math.Round(f*100) / 100
			one_day_result[key] = f
		}

		for key, _ := range gs.dataSet {
			// このgroupにサービスが含まれていればamountをプッシュ
			if isInclude(key, results.Groups) {
				gs.dataSet[key].costs = append(
					gs.dataSet[key].costs,
					one_day_result[key],
				)
			} else {
				gs.dataSet[key].costs = append(
					gs.dataSet[key].costs,
					0,
				)
			}
		}
	}

	deleteNoUseEntry(&gs)
	return gs, nil
}

// sはAWSサービス名
// costexplorere.Groupの配列にsが含まれているかチェックする
func isInclude(s string, groups []*costexplorer.Group) bool {
	for _, group := range groups {
		if aws.StringValue(group.Keys[0]) == s {
			return true
		}
	}

	return false
}

// cost explorerのAPIのレスポンスには利用料が$0のサービスも含まれているので
// 期間中すべてが$0のサービスはdataSetから除く
func deleteNoUseEntry(gs *GraphSource) {
	for key, sd := range gs.dataSet {
		deleteFlg := true
		for _, amount := range sd.costs {
			if amount != 0 {
				deleteFlg = false
			}
		}
		if deleteFlg {
			delete(gs.dataSet, key)
		}
	}
}
