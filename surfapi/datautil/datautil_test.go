package datautil

import (
	"reflect"
	"testing"
)

var url string = "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"

func TestGetBouyData(t *testing.T) {
	rawBouyData := GetBouyData(url)
	var testData []byte

	if reflect.TypeOf(rawBouyData) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, rawBouyData)
	}
}

func TestHandleRawData(t *testing.T) {
	rawBouyData := GetBouyData(url)
	bouyData := HandleRawData(rawBouyData)
	var testData [][]string

	if reflect.TypeOf(bouyData) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, bouyData)
	}
}

func TestGetLatestData(t *testing.T) {
	rawBouyData := GetBouyData(url)
	bouyData := HandleRawData(rawBouyData)
	latestData := getLatestData(bouyData)
	var testData []string

	if reflect.TypeOf(latestData) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, latestData)
	}
}

func TestGetSwellHeight(t *testing.T) {
	rawBouyData := GetBouyData(url)
	bouyData := HandleRawData(rawBouyData)
	swellHeight := getSwellHeight(bouyData)
	var testData float64

	if reflect.TypeOf(swellHeight) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, swellHeight)
	}
}

func TestGetSwellPeriod(t *testing.T) {
	rawBouyData := GetBouyData(url)
	bouyData := HandleRawData(rawBouyData)
	swellPeriod := getSwellPeriod(bouyData)
	var testData float64

	if reflect.TypeOf(swellPeriod) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, swellPeriod)
	}
}

func TestGetSwellPeriodScore(t *testing.T) {

	tables := []struct {
		x float64
		n float64
	}{
		{17, 5},
		{13.4, 4},
		{16, 5},
		{5, 1},
		{10.4, 3},
		{13, 4},
		{12, 3},
	}

	for _, table := range tables {
		score := getSwellPeriodScore(table.x)
		if score != table.n {
			t.Errorf(" Expected %v , got %v", score, table.n)
		}
	}
}

func TestGetWindDirectionScore(t *testing.T) {

	tables := []struct {
		x string
		n float64
	}{
		{"E", 5},
		{"NE", 4},
		{"SE", 4},
		{"S", 3},
		{"SSW", 1},
		{"NNW", 1},
	}

	for _, table := range tables {
		score := getWindDirectionScore(table.x)
		if score != table.n {
			t.Errorf(" Expected %v , got %v", score, table.n)
		}
	}
}

func TestGetWaveSizeScore(t *testing.T) {
	rawBouyData := GetBouyData(url)
	bouyData := HandleRawData(rawBouyData)
	waveSizeScore := getWaveSizeScore(bouyData)
	var testData float64

	if reflect.TypeOf(waveSizeScore) != reflect.TypeOf(testData) {
		t.Errorf("Expected %T, got %T", testData, waveSizeScore)
	}
}

// func TestCalculateSurfRating(t *testing.T) {
// 	// rawBouyData := GetBouyData(url)
// 	// bouyData := HandleRawData(rawBouyData)
// 	// surfRating := CalculateSurfRating(bouyData)
// 	tables := []struct {
// 		testArray [][]string
// 	}{
// 		{[][]string{["2019", "4", "1", "16", "0", "1.7", "1.7", "17.4", "0.3", "4", "W", "NNW", "SWELL", "11.4", "280"]}},
// 	}

// 	for _, table := range tables {
// 		score := CalculateSurfRating(table.testArray)
// 		if score != table.n {
// 			t.Errorf(" Expected %v , got %v", score, table.n)
// 		}
// 	}

// 	if reflect.TypeOf(surfRating) != reflect.TypeOf(testData) {
// 		t.Errorf("Expected %T, got %T", table.testType, surfRating)
// 	}
// }

func TestCalculateSurRatingf(t *testing.T) {
	tables := []struct {
		testArray [][]string
		testScore float64
	}{
		{[][]string{{"2019", "4", "1", "16", "0", "1.7", "1.7", "17.4", "0.3", "4", "W", "NNW", "SWELL", "11.4", "280"}, {"2019", "4", "1", "16", "0", "1.7", "2.3", "17.4", "0.3", "4", "W", "NNW", "SWELL", "11.4", "280"}}, (10 / 3.0)},
	}

	for _, table := range tables {
		score := CalculateSurfRating(table.testArray)
		if score != table.testScore {
			t.Errorf(" Expected %v , got %v", table.testScore, score)
		}
	}
}
