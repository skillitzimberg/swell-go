package datautil

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetBouyData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func HandleRawData(responseData []byte) [][]string {
	text := string(responseData)
	stringsOnReturn := strings.Split(text, "\n")
	var stringsOnTabs []string
	cleanedData := make([]string, 0, len(stringsOnTabs))
	dataRows := make([][]string, 0, len(stringsOnReturn))

	for i := 2; i < 24; i++ {
		stringsOnTabs = strings.Split(stringsOnReturn[i], " ")
		cleanedData = removeEmptySpace(stringsOnTabs)
		dataRows = append(dataRows, cleanedData)
	}

	return dataRows
}

func removeEmptySpace(arrayWithSpaces []string) []string {
	cleanedData := make([]string, 0, len(arrayWithSpaces))
	for j := 0; j < len(arrayWithSpaces); j++ {
		if arrayWithSpaces[j] != "" {
			cleanedData = append(cleanedData, arrayWithSpaces[j])
		}
	}
	return cleanedData
}

func ParseRow(row string) (SurfData, error) {
	errFunc := func(err error) {
		return SurfData{}, err
	}
	hourlyData := strings.Split(row, " ")
	year, err := strconv.Atoi(hourlyData[0])
	if err != nil {
		return SurfData{}, err
	}
	month, err := strconv.Atoi(hourlyData[1])
	day, _ := strconv.Atoi(hourlyData[2])
	hour, _ := strconv.Atoi(hourlyData[3])
	minute, _ := strconv.Atoi(hourlyData[4])
	WVHT, _ := strconv.ParseFloat(hourlyData[5], 64)
	SwH, _ := strconv.ParseFloat(hourlyData[6], 64)
	SwP, _ := strconv.ParseFloat(hourlyData[7], 64)
	WWH, _ := strconv.ParseFloat(hourlyData[8], 64)
	WWP, _ := strconv.ParseFloat(hourlyData[9], 64)
	SwD := hourlyData[10]
	WWD := hourlyData[11]
	steepness := hourlyData[12]
	APD, _ := strconv.ParseFloat(hourlyData[13], 64)
	MWD, _ := strconv.Atoi(hourlyData[14])

	surfData := SurfData{
		Year:      0,
		Month:     0,
		Day:       0,
		Hour:      0,
		Minute:    0,
		WVHT:      0.0,
		SwH:       0.0,
		SwP:       0.0,
		WWH:       0.0,
		WWP:       0.0,
		SwD:       "",
		WWD:       "",
		Steepness: "",
		APD:       0.0,
		MWD:       0,
	}
}

func DataToStructs(dataRows [][]string) []SurfData {
	var surfDataStructs []SurfData

	for i := 0; i < len(dataRows); i++ {
		hourlyData := dataRows[i]
		for j := 0; j < len(hourlyData); j++ {
		}

		year, _ := strconv.Atoi(hourlyData[0])
		month, _ := strconv.Atoi(hourlyData[1])
		day, _ := strconv.Atoi(hourlyData[2])
		hour, _ := strconv.Atoi(hourlyData[3])
		minute, _ := strconv.Atoi(hourlyData[4])
		WVHT, _ := strconv.ParseFloat(hourlyData[5], 64)
		SwH, _ := strconv.ParseFloat(hourlyData[6], 64)
		SwP, _ := strconv.ParseFloat(hourlyData[7], 64)
		WWH, _ := strconv.ParseFloat(hourlyData[8], 64)
		WWP, _ := strconv.ParseFloat(hourlyData[9], 64)
		SwD := hourlyData[10]
		WWD := hourlyData[11]
		steepness := hourlyData[12]
		APD, _ := strconv.ParseFloat(hourlyData[13], 64)
		MWD, _ := strconv.Atoi(hourlyData[14])

		var surfData = SurfData{
			year,
			month,
			day,
			hour,
			minute,
			WVHT,
			SwH,
			SwP,
			WWH,
			WWP,
			SwD,
			WWD,
			steepness,
			APD,
			MWD,
		}

		surfDataStructs = append(surfDataStructs, surfData)
	}
	return surfDataStructs
}

func getLatestData(cleanedData [][]string) []string {
	return cleanedData[0]
}

func getSwellHeight(cleanedData [][]string) float64 {
	latestData := getLatestData(cleanedData)
	convertSwHToInt, err := strconv.ParseFloat(latestData[6], 64)
	if err != nil {
		log.Fatal(err)
	}
	return convertSwHToInt
}

func getSwellPeriod(cleanedData [][]string) float64 {
	latestData := getLatestData(cleanedData)
	convertSwPToInt, err := strconv.ParseFloat(latestData[7], 64)
	if err != nil {
		log.Fatal(err)
	}
	return convertSwPToInt
}

func getSwellPeriodScore(swellPeriod float64) float64 {
	if swellPeriod >= 16 {
		return 5
	} else if swellPeriod >= 13 {
		return 4
	} else if swellPeriod >= 10 {
		return 3
	} else {
		return 1
	}
}

func getWindDirection(cleanedData [][]string) string {
	latestData := getLatestData(cleanedData)
	return latestData[11]
}

func getWindDirectionScore(windDirection string) float64 {
	if windDirection == "E" {
		return 5
	} else if windDirection == "NE" || windDirection == "SE" {
		return 4
	} else if windDirection == "S" {
		return 3
	} else {
		return 1
	}
}

func getWaveSizeScore(cleanedData [][]string) float64 {
	swellHeight := getSwellHeight(cleanedData)
	swellPeriod := getSwellPeriod(cleanedData)
	waveSize := swellHeight * swellPeriod

	if waveSize >= 30 {
		return 5
	} else if waveSize >= 25 {
		return 4
	} else if waveSize >= 20 {
		return 3
	} else if waveSize >= 11 {
		return 2
	} else {
		return 1
	}
}

func CalculateSurfRating(cleanedData [][]string) float64 {
	swellPeriod := getSwellPeriod(cleanedData)
	swellPeriodScore := getSwellPeriodScore(swellPeriod)
	windDirection := getWindDirection(cleanedData)
	windDirectionScore := getWindDirectionScore(windDirection)
	waveSizeScore := getWaveSizeScore(cleanedData)
	surfRating := ((swellPeriodScore + waveSizeScore + windDirectionScore) / 3.0)
	return surfRating
}
