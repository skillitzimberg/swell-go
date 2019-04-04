package datautil

import (
	"reflect"
	"testing"
)

var sampleRow = `2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286`
var sampleRowArray = []string{"2019", "04", "02", "16", "00", "", "1.8", "", "1.8", "14.8", "", "0.2", "", "3.4", "WNW", "NNW", "", "", "", "", "", "SWELL", "", "9.3", "286"}

var validSampleRowArray = []string{"2019", "04", "02", "16", "00", "1.8", "1.8", "14.8", "0.2", "3.4", "WNW", "NNW", "SWELL", "9.3", "286"}

var invalidSampleRowArray = []string{"2019", "04", "02", "16", "00", "1.8", "1.8", "14.8", "0.2", "3.4", "WNW", "NNW", "SWELL", "9.3"}

var sampleData = `#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286
2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288
2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285
2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288
2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289`

var sampleDataArray = []string{"2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286",
	"2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288",
	"2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285",
	"2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288",
	"2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289"}

var sampleDataStruct = BouyData{
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

var packagedBouyData = []BouyData{
	{
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
	},
	{
		Year:      2019,
		Month:     04,
		Day:       02,
		Hour:      15,
		Minute:    00,
		WVHT:      1.7,
		SwH:       1.7,
		SwP:       16,
		WWH:       0.2,
		WWP:       3.4,
		SwD:       "WNW",
		WWD:       "NNE",
		Steepness: "SWELL",
		APD:       9.4,
		MWD:       288,
	},
	{
		Year:      2019,
		Month:     04,
		Day:       02,
		Hour:      14,
		Minute:    00,
		WVHT:      1.5,
		SwH:       1.5,
		SwP:       13.8,
		WWH:       0.2,
		WWP:       3.6,
		SwD:       "WNW",
		WWD:       "NNW",
		Steepness: "SWELL",
		APD:       9.1,
		MWD:       285,
	},
	{
		Year:      2019,
		Month:     04,
		Day:       02,
		Hour:      13,
		Minute:    00,
		WVHT:      1.7,
		SwH:       1.7,
		SwP:       17.4,
		WWH:       0.2,
		WWP:       3.8,
		SwD:       "WNW",
		WWD:       "NNE",
		Steepness: "SWELL",
		APD:       10.1,
		MWD:       288,
	},
	{
		Year:      2019,
		Month:     04,
		Day:       02,
		Hour:      12,
		Minute:    00,
		WVHT:      1.9,
		SwH:       1.8,
		SwP:       16.0,
		WWH:       0.2,
		WWP:       3.7,
		SwD:       "WNW",
		WWD:       "NNW",
		Steepness: "SWELL",
		APD:       10.5,
		MWD:       289,
	},
}

var url = "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"

func TestParseDataInRows(t *testing.T) {
	got := parseDataInRows(sampleData)
	want := sampleDataArray

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%v'; want '%v'", got, want)
	}
}

func TestGetCurrentData(t *testing.T) {
	got := getCurrentData(sampleDataArray)
	want := sampleRow

	if got != want {
		t.Errorf("Got '%s'; want '%s'", got, want)
	}
}

func TestRowIntoArray(t *testing.T) {
	got := rowIntoArray(sampleRow)
	want := validSampleRowArray

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%v'; want '%v'", got, want)
	}
}

func TestRowArrayToStruct(t *testing.T) {
	got, _ := RowArrayToStruct(validSampleRowArray)
	want := sampleDataStruct

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%v'; want '%v'", got, want)
	}
}

func TestPackageStructsForJson(t *testing.T) {
	got := PackageStructsForJson(sampleDataArray)
	want := packagedBouyData

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%v'; want '%v'", got, want)
	}
}

func TestBouyDataMethods(t *testing.T) {
	t.Run("WaveSize()", func(t *testing.T) {
		bouyData, _ := RowArrayToStruct(validSampleRowArray)
		got := bouyData.WaveSize()
		want := 26.64

		if got != want {
			t.Errorf("Got '%f'; want '%f'", got, want)
		}
	})

	t.Run("GetWaveScore", func(t *testing.T) {
		bouyData, _ := RowArrayToStruct(validSampleRowArray)
		got := bouyData.getWaveScore()
		want := 4

		if got != want {
			t.Errorf("Got '%d'; want '%d'", got, want)
		}
	})

	t.Run("GetSwellPeriodScore", func(t *testing.T) {
		bouyData, _ := RowArrayToStruct(validSampleRowArray)
		got := bouyData.getSwellPeriodScore()
		want := 4

		if got != want {
			t.Errorf("Got '%d'; want '%d'", got, want)
		}
	})

	t.Run("getWindDirectionScore", func(t *testing.T) {
		bouyData, _ := RowArrayToStruct(validSampleRowArray)
		got := bouyData.getWindDirectionScore()
		want := 1

		if got != want {
			t.Errorf("Got '%d'; want '%d'", got, want)
		}
	})

	t.Run("CalculateSurfRating", func(t *testing.T) {
		bouyData, _ := RowArrayToStruct(validSampleRowArray)
		got := bouyData.CalculateSurfRating()
		want := 3.0

		if got != want {
			t.Errorf("Got '%f'; want '%f'", got, want)
		}
	})
}

func TestFetchData(t *testing.T) {
	t.Run("GetBouyData", func(t *testing.T) {
		var fetchBouyData = FetchData{"https://www.ndbc.noaa.gov/data/realtime2/46029.spec"}
		got := fetchBouyData.GetBouyData()
		want := 

	})
}
