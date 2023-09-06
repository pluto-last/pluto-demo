package goset

import (
	mapset "github.com/deckarep/golang-set/v2"
	"log"
)

func GoSetTest() {
	// Create a string-based set of required classes.
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")

	// Create a string-based set of science classes.
	sciences := mapset.NewSet[string]()
	sciences.Add("biology")
	sciences.Add("chemistry")

	// Create a string-based set of electives.
	electives := mapset.NewSet[string]()
	electives.Add("welding")
	electives.Add("music")
	electives.Add("automotive")

	// Create a string-based set of bonus programming classes.
	bonus := mapset.NewSet[string]()
	bonus.Add("beginner go")
	bonus.Add("python for dummies")

	// 多个set合并，会自动去重 (合并)
	all := required.Union(sciences).Union(electives).Union(bonus)
	log.Println(all)

	// 判断元素是否存在
	result := sciences.Contains("cooking")
	log.Println(result)

	//  去除部分数据
	notScience := all.Difference(sciences)
	log.Println(notScience)

	// 取交集
	reqScience := sciences.Intersect(required)
	log.Println(reqScience)
}
