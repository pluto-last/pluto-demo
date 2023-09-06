package main

import (
	"encoding/json"
)

func main() {

	// 围栏数据， 长沙主城区
	var coordinates []Coordinates
	json.Unmarshal([]byte(coordinatesJson), &coordinates)

	// 围栏内每600m 打一个点
	points, _ := ExtractCircleEdgePoints(coordinates)

	var coordinateList [][]float64
	for _, item := range points {
		var temp []float64
		temp = append(temp, item.Lng)
		temp = append(temp, item.Lat)

		coordinateList = append(coordinateList, temp)
	}

	// 打好的点位，生成geojson文件，在网站上可视化查看
	// https://geojsonviewer.nsspot.net/
	CreateGeoJson(coordinateList)
}
