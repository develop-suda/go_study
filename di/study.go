package di

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// TotalValueImpl TotalValueの実態
type TotalValueImpl struct{}

// ProductPriceImpl ProductPriceの実態
type ProductPriceImpl struct{}

// SumValue TotalValueが持つメソッドの実装
func (t TotalValueImpl) SumValue() int {
	pp := ProductPriceImpl{} // ProductPriceImplのインスタンスを直接生成
	record, err := pp.AllData()
	if err != nil {
		return 0
	}
	// 合計を計算してその値を返す
	var sumValue int
	for _, row := range record {
		for _, value := range row {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			sumValue += intValue
		}
	}
	return sumValue
}

// AllData ProductPriceが持つメソッドの実装
func (p ProductPriceImpl) AllData() ([][]string, error) {
	// ファイルを読み込んで、データを返す
	file, err := os.Open("sample_data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func main() {
	tv := TotalValueImpl{} // TotalValueImplのインスタンスを直接生成

	// 合計値を計算
	sum := tv.SumValue()
	fmt.Println("Total Value:", sum)
}
