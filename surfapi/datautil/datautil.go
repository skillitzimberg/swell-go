package datautil

import (
	"errors"
	"strconv"
	"strings"
)

func parseDataInRows(bouyData string) (bouyDataArray []string) {
	arrayOfRows := strings.Split(bouyData, "\n")
	bouyDataArray = arrayOfRows[2:]
	return
}

func PackageStructsForJson(bouyDataArray []string) (packagedBouyDataStructs []BouyData) {
	for _, bouyData := range bouyDataArray {
		rowArray := rowIntoArray(bouyData)
		hourlyData, _ := RowArrayToStruct(rowArray)
		packagedBouyDataStructs = append(packagedBouyDataStructs, hourlyData)
	}
	return
}

func getCurrentData(dataRows []string) (row string) {
	row = dataRows[0]
	return row
}

func rowIntoArray(row string) (rowArray []string) {
	rowArray = strings.Fields(row)
	return
}

func RowArrayToStruct(rowArray []string) (hourlyData BouyData, err error) {

	if len(rowArray) != 15 {
		err = errors.New("Cannot work with an array less or greater than length of 15")
		return BouyData{}, err
	}

	for j := 0; j < len(rowArray); j++ {
		Year, err := strconv.Atoi(rowArray[0])
		if err != nil {
			return BouyData{}, err
		}
		Month, err := strconv.Atoi(rowArray[1])
		if err != nil {
			return BouyData{}, err
		}
		Day, err := strconv.Atoi(rowArray[2])
		if err != nil {
			return BouyData{}, err
		}
		Hour, err := strconv.Atoi(rowArray[3])
		if err != nil {
			return BouyData{}, err
		}
		Minute, err := strconv.Atoi(rowArray[4])
		if err != nil {
			return BouyData{}, err
		}
		WVHT, err := strconv.ParseFloat(rowArray[5], 64)
		if err != nil {
			return BouyData{}, err
		}
		SwH, err := strconv.ParseFloat(rowArray[6], 64)
		if err != nil {
			return BouyData{}, err
		}
		SwP, err := strconv.ParseFloat(rowArray[7], 64)
		if err != nil {
			return BouyData{}, err
		}
		WWH, err := strconv.ParseFloat(rowArray[8], 64)
		if err != nil {
			return BouyData{}, err
		}
		WWP, err := strconv.ParseFloat(rowArray[9], 64)
		if err != nil {
			return BouyData{}, err
		}
		SwD := rowArray[10]
		WWD := rowArray[11]
		Steepness := rowArray[12]
		APD, err := strconv.ParseFloat(rowArray[13], 64)
		if err != nil {
			return BouyData{}, err
		}
		MWD, err := strconv.Atoi(rowArray[14])
		if err != nil {
			return BouyData{}, err
		}

		hourlyData = BouyData{
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
	return
}
