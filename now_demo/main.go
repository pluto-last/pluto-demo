package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

func main() {

}

func loUtils() {

	// 数组去重
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	log.Println(names) // [Samuel John]

	// 数组去重，按照指定函数进行去重
	uniqValues := lo.UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	log.Println(uniqValues) // [0 1 2]

	// 按照指定的长度拆分数据
	respV1 := lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	log.Println(respV1) // [[0 1] [2 3] [4 5]]

	// 打乱元素的先后顺序
	randomOrder := lo.Shuffle([]int{0, 1, 2, 3, 4, 5})
	log.Println(randomOrder) // [5 0 1 2 4 3]

	// 反转数组
	reverseOrder := lo.Reverse([]int{0, 1, 2, 3, 4, 5})
	log.Println(reverseOrder) // []int{5, 4, 3, 2, 1, 0}
}

func nowUtils() {
	log.Println(time.Now())

	// 一些工具类函数
	log.Println(now.BeginningOfMinute())
	log.Println(now.BeginningOfWeek())

	// 根据另外一个时间计算时间
	t := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	log.Println(now.With(t).EndOfMonth())

	// 字符串转时间
	t, _ = now.Parse("2023-10-13")
	log.Println(t)

}

func decimalUtils() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
