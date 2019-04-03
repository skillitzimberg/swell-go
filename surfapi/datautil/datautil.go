package datautil

import (
	"errors"
	"strconv"
	"strings"
)

// func GetBouyData(url string) []byte {

// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return responseData
// }

// func HandleRawData(responseData []byte) [][]string {
// 	text := string(responseData)
// 	stringsOnReturn := strings.Split(text, "\n")
// 	var stringsOnTabs []string
// 	cleanedData := make([]string, 0, len(stringsOnTabs))
// 	dataRows := make([][]string, 0, len(stringsOnReturn))

// 	for i := 2; i < 24; i++ {
// 		stringsOnTabs = strings.Split(stringsOnReturn[i], " ")
// 		cleanedData = removeEmptySpace(stringsOnTabs)
// 		dataRows = append(dataRows, cleanedData)
// 	}

// 	return dataRows
// }

func getLatestData(rawData string) []string {
	text := string(rawData)
	stringsOnReturn := strings.Split(text, "\n")
	latestData := strings.Split(stringsOnReturn[2], " ")
	return latestData
}

func removeEmptySpace(dataRow string) []string {
	dataArray := strings.Split(dataRow, " ")
	var cleanedData = make([]string, 0, len(dataArray))
	for j := 0; j < len(dataArray); j++ {
		if dataArray[j] != "" {
			cleanedData = append(cleanedData, dataArray[j])
		}
	}
	return cleanedData
}

func rowDataToStruct(hourlyData []string) (SurfData, error) {
	var surfData SurfData
	var err error

	if len(hourlyData) != 15 {
		err = errors.New("Cannot work with an array less or greater than length of 15")
		return SurfData{}, err
	}

	for j := 0; j < len(hourlyData); j++ {
		Year, err := strconv.Atoi(hourlyData[0])
		if err != nil {
			return SurfData{}, err
		}
		Month, err := strconv.Atoi(hourlyData[1])
		if err != nil {
			return SurfData{}, err
		}
		Day, err := strconv.Atoi(hourlyData[2])
		if err != nil {
			return SurfData{}, err
		}
		Hour, err := strconv.Atoi(hourlyData[3])
		if err != nil {
			return SurfData{}, err
		}
		Minute, err := strconv.Atoi(hourlyData[4])
		if err != nil {
			return SurfData{}, err
		}
		WVHT, err := strconv.ParseFloat(hourlyData[5], 64)
		if err != nil {
			return SurfData{}, err
		}
		SwH, err := strconv.ParseFloat(hourlyData[6], 64)
		if err != nil {
			return SurfData{}, err
		}
		SwP, err := strconv.ParseFloat(hourlyData[7], 64)
		if err != nil {
			return SurfData{}, err
		}
		WWH, err := strconv.ParseFloat(hourlyData[8], 64)
		if err != nil {
			return SurfData{}, err
		}
		WWP, err := strconv.ParseFloat(hourlyData[9], 64)
		if err != nil {
			return SurfData{}, err
		}
		SwD := hourlyData[10]
		WWD := hourlyData[11]
		Steepness := hourlyData[12]
		APD, err := strconv.ParseFloat(hourlyData[13], 64)
		if err != nil {
			return SurfData{}, err
		}
		MWD, err := strconv.Atoi(hourlyData[14])
		if err != nil {
			return SurfData{}, err
		}

		surfData = SurfData{
			Year,
			Month,
			Day,
			Hour,
			Minute,
			WVHT,
			SwH,
			SwP,
			WWH,
			WWP,
			SwD,
			WWD,
			Steepness,
			APD,
			MWD,
		}
	}
	return surfData, err
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

func getWaveSizeScore(swellHeight, swellPeriod float64) float64 {
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

func CalculateSurfRating(surfData SurfData) float64 {
	swellPeriodScore := float64(getSwellPeriodScore(surfData.SwP))
	windDirectionScore := float64(getWindDirectionScore(surfData.WWD))
	waveSizeScore := float64(getWaveSizeScore(surfData.SwP, surfData.SwH))
	numerator := (swellPeriodScore + waveSizeScore + windDirectionScore)
	surfRating := (numerator / 3.0)
	return surfRating
}

// func DataToStructs(dataRows [][]string) []SurfData {
// 	var surfDataStructs []SurfData

// 	for i := 0; i < len(dataRows); i++ {
// 		hourlyData := dataRows[i]
// 		for j := 0; j < len(hourlyData); j++ {
// 		}

// 		year, _ := strconv.Atoi(hourlyData[0])
// 		month, _ := strconv.Atoi(hourlyData[1])
// 		day, _ := strconv.Atoi(hourlyData[2])
// 		hour, _ := strconv.Atoi(hourlyData[3])
// 		minute, _ := strconv.Atoi(hourlyData[4])
// 		WVHT, _ := strconv.ParseFloat(hourlyData[5], 64)
// 		SwH, _ := strconv.ParseFloat(hourlyData[6], 64)
// 		SwP, _ := strconv.ParseFloat(hourlyData[7], 64)
// 		WWH, _ := strconv.ParseFloat(hourlyData[8], 64)
// 		WWP, _ := strconv.ParseFloat(hourlyData[9], 64)
// 		SwD := hourlyData[10]
// 		WWD := hourlyData[11]
// 		steepness := hourlyData[12]
// 		APD, _ := strconv.ParseFloat(hourlyData[13], 64)
// 		MWD, _ := strconv.Atoi(hourlyData[14])

// 		var surfData = SurfData{
// 			year,
// 			month,
// 			day,
// 			hour,
// 			minute,
// 			WVHT,
// 			SwH,
// 			SwP,
// 			WWH,
// 			WWP,
// 			SwD,
// 			WWD,
// 			steepness,
// 			APD,
// 			MWD,
// 		}

// 		surfDataStructs = append(surfDataStructs, surfData)
// 	}
// 	return surfDataStructs
// }

// func getSwellHeight(cleanedData [][]string) float64 {
// 	latestData := getLatestData(cleanedData)
// 	convertSwHToInt, err := strconv.ParseFloat(latestData[6], 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return convertSwHToInt
// }

// func getSwellPeriod(cleanedData [][]string) float64 {
// 	latestData := getLatestData(cleanedData)
// 	convertSwPToInt, err := strconv.ParseFloat(latestData[7], 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return convertSwPToInt
// }

// func getWindDirection(cleanedData [][]string) string {
// 	latestData := getLatestData(cleanedData)
// 	return latestData[11]
// }
