package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   Point                  `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

func CreateGeoJson(coordinates [][]float64) {

	features := make([]Feature, len(coordinates))

	for i, coord := range coordinates {
		point := Point{
			Type:        "Point",
			Coordinates: coord,
		}

		feature := Feature{
			Type:       "Feature",
			Geometry:   point,
			Properties: make(map[string]interface{}),
		}

		features[i] = feature
	}

	featureCollection := FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}

	jsonData, err := json.Marshal(featureCollection)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fileName := "output.geojson"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.Write(jsonData)

	fmt.Println("GeoJSON file created:", fileName)
}
