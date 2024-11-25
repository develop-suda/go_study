package dii

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// TotalValue 商品の合計値を計算するinterface
type TotalValue interface {
	SumValue() int
}

// ProductPrice 商品の全てのデータを取得するinterface
type ProductPrice interface {
	AllData() ([][]string, error)
}

// TotalValueImpl TotalValueの実態(interfaceに書かれているメソッド(sumValueByName)を実装する必要がある)
type TotalValueImpl struct {
	pp ProductPrice
}

// ProductPriceImpl ProductPriceの実態(interfaceに書かれているメソッド(getAllData)を実装する必要がある)
type ProductPriceImpl struct {
}

// NewTotalValue TotalValueImplをインスタンス化するためのコンストラクタ
func NewTotalValue(pp ProductPrice) TotalValue {
	return &TotalValueImpl{pp: pp}
}

// NewProductPrice ProductPriceImplをインスタンス化するためのコンストラクタ
func NewProductPrice() ProductPrice {
	return ProductPriceImpl{}
}

// SumValue TotalValueが持つsumValueByNameの実装
func (t TotalValueImpl) SumValue() int {
	record, err := t.pp.AllData()
	if err != nil {
		return 0
	}
	// 合計をしてその値を返す
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

// AllData ProductPriceが持つAllDataの実装
func (p ProductPriceImpl) AllData() ([][]string, error) {
	// ファイルを読み込んで,返すだけの処理
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
	fmt.Println(records)

	return records, nil
}

func main() {
	// ProductPriceのインスタンスを生成
	pp := NewProductPrice()

	// TotalValueのインスタンスを生成し、ProductPriceを注入
	tv := NewTotalValue(pp)

	// 合計値を計算
	sum := tv.SumValue()
	fmt.Println("Total Value:", sum)
}
