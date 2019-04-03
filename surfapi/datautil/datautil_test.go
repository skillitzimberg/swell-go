package datautil

import (
	"strings"
	"testing"
)

var sampleRow = `2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286`

var validSampleRowArray = []string{"2019", "04", "02", "16", "00", "1.8", "1.8", "14.8", "0.2", "3.4", "WNW", "NNW", "SWELL", "9.3", "286"}

var invalidSampleRowArray = []string{"2019", "04", "02", "16", "00", "1.8", "1.8", "14.8", "0.2", "3.4", "WNW", "NNW", "SWELL", "9.3"}

var sampleData = `#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286
2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288
2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285
2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288
2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289`

var sampleDataStruct = SurfData{
	Year:      2019,
	Month:     04,
	Day:       02,
	Hour:      16,
	Minute:    00,
	WVHT:      1.8,
	SwH:       1.8,
	SwP:       14.8,
	WWH:       0.2,
	WWP:       3.4,
	SwD:       "WNW",
	WWD:       "NNW",
	Steepness: "SWELL",
	APD:       9.3,
	MWD:       286,
}

var url = "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"

func TestGetLatestData(t *testing.T) {
	bouyData := getLatestData(sampleData)
	testArray := strings.Split(sampleRow, " ")

	if len(bouyData) != len(testArray) {
		t.Errorf("Expected %v, got %v", len(testArray), len(bouyData))
	}
}
func TestRemoveEmptySpace(t *testing.T) {
	arrayWithoutEmptySpaces := removeEmptySpace(sampleRow)

	for i := 0; i < len(arrayWithoutEmptySpaces); i++ {
		if arrayWithoutEmptySpaces[i] == "" {
			t.Errorf("Expected %v, got %v", "empty string", arrayWithoutEmptySpaces[i])
		}
	}
}

func TestValidRowDataToStruct(t *testing.T) {
	surfData, err := rowDataToStruct(validSampleRowArray)

	if err != nil {
		t.Error(err.Error())
	}

	if surfData.Year != 2019 {
		t.Errorf("Expected %v, got %v", 2019, surfData.Year)
	}

	if surfData.Month != 04 {
		t.Errorf("Expected %v, got %v", 04, surfData.Month)
	}

	if surfData.Day != 02 {
		t.Errorf("Expected %v, go %v", 02, surfData.Day)
	}

	if surfData.Hour != 16 {
		t.Errorf("Expected %v, got %v", 16, surfData.Hour)
	}

	if surfData.Minute != 00 {
		t.Errorf("Expected %v, got %v", 00, surfData.Minute)
	}

	if surfData.WVHT != 1.8 {
		t.Errorf("Expected %v, got %v", 1.8, surfData.WVHT)
	}

	if surfData.SwH != 1.8 {
		t.Errorf("Expected %v, got %v", 1.8, surfData.SwH)
	}

	if surfData.SwP != 14.8 {
		t.Errorf("Expected %v, got %v", 14.8, surfData.SwP)
	}

	if surfData.WWH != 0.2 {
		t.Errorf("Expected %v, got %v", 0.2, surfData.WWH)
	}

	if surfData.WWP != 3.4 {
		t.Errorf("Expected %v, got %v", 0.2, surfData.WWP)
	}

	if surfData.SwD != "WNW" {
		t.Errorf("Expected %v, got %v", "WNW", surfData.SwD)
	}

	if surfData.WWD != "NNW" {
		t.Errorf("Expected %v, got %v", "NNW", surfData.WWD)
	}

	if surfData.Steepness != "SWELL" {
		t.Errorf("Expected %v, got %v", "SWELL", surfData.Steepness)
	}

	if surfData.APD != 9.3 {
		t.Errorf("Expected %v, got %v", 9.3, surfData.APD)
	}

	if surfData.MWD != 286 {
		t.Errorf("Expected %v, got %v", 286, surfData.MWD)
	}
}

func TestInvalidRowDataToStruct(t *testing.T) {
	_, err := rowDataToStruct(invalidSampleRowArray)

	if err == nil {
		t.Error(err.Error())
	}
}

func TestGetSwellPeriodScore(t *testing.T) {

	tables := []struct {
		got  float64
		want float64
	}{
		{17, 5},
		{13.4, 4},
		{16, 5},
		{5, 1},
		{10.4, 3},
		{13, 4},
		{11.3, 3},
	}

	for _, table := range tables {
		score := getSwellPeriodScore(table.got)
		if score != table.want {
			t.Errorf(" Expected %v , got %v", score, table.want)
		}
	}
}

func TestGetWindDirectionScore(t *testing.T) {

	tables := []struct {
		got  string
		want float64
	}{
		{"E", 5},
		{"NE", 4},
		{"SE", 4},
		{"S", 3},
		{"SSW", 1},
		{"NNW", 1},
		{"N", 1},
		{"W", 1},
	}

	for _, table := range tables {
		score := getWindDirectionScore(table.got)
		if score != table.want {
			t.Errorf(" Expected %v , got %v", score, table.want)
		}
	}
}

func TestGetWaveSizeScore(t *testing.T) {
	tables := []struct {
		swellHeight float64
		swellPeriod float64
		want        float64
	}{
		{2.0, 16.8, 5},
		{1.8, 14.8, 4},
		{1.5, 13.6, 3},
		{1.3, 9.8, 2},
		{0.9, 11.2, 1},
	}

	for _, table := range tables {
		score := getWaveSizeScore(table.swellHeight, table.swellPeriod)
		if score != table.want {
			t.Errorf(" Expected %v , got %v", score, table.want)
		}
	}
}

func TestCalculateSurRating(t *testing.T) {
	testScore := 9 / 3.0
	score := CalculateSurfRating(sampleDataStruct)
	if score != testScore {
		t.Errorf(" Expected %v , got %v", testScore, score)
	}
}

// func TestGetBouyData(t *testing.T) {
// 	rawBouyData := GetBouyData(url)
// 	var testData []byte

// 	if reflect.TypeOf(rawBouyData) != reflect.TypeOf(testData) {
// 		t.Errorf("Expected %T, got %T", testData, rawBouyData)
// 	}
// }

// func TestGetLatestData(t *testing.T) {
// 	rawBouyData := GetBouyData(url)
// 	bouyData := HandleRawData(rawBouyData)
// 	latestData := getLatestData(bouyData)
// 	var testData []string

// 	if reflect.TypeOf(latestData) != reflect.TypeOf(testData) {
// 		t.Errorf("Expected %T, got %T", testData, latestData)
// 	}
// }

// func TestGetSwellHeight(t *testing.T) {
// 	rawBouyData := GetBouyData(url)
// 	bouyData := HandleRawData(rawBouyData)
// 	swellHeight := getSwellHeight(bouyData)
// 	var testData float64

// 	if reflect.TypeOf(swellHeight) != reflect.TypeOf(testData) {
// 		t.Errorf("Expected %T, got %T", testData, swellHeight)
// 	}
// }

// func TestGetSwellPeriod(t *testing.T) {
// 	rawBouyData := GetBouyData(url)
// 	bouyData := HandleRawData(rawBouyData)
// 	swellPeriod := getSwellPeriod(bouyData)
// 	var testData float64

// 	if reflect.TypeOf(swellPeriod) != reflect.TypeOf(testData) {
// 		t.Errorf("Expected %T, got %T", testData, swellPeriod)
// 	}
// }
