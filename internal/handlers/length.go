// Package handlers
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AppBlitz/taskConvert/internal/models"
)

var (
	oneMetre      float64
	thousandMetre float64
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dataRequest models.RequestLength
	data, err := io.ReadAll(r.Body)
	erro := json.Unmarshal(data, &dataRequest)
	if erro != nil {
		http.Error(w, "read data error", http.StatusConflict)
	}
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == http.MethodPost {
		_, err := w.Write(structToBytes(dataRequest))
		if err != nil {
			http.Error(w, "Error with response", http.StatusConflict)
		}
	} else {
		http.Error(w, dataRequest.UnitToConvertFrom, http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "Application/json")
}

func createStructOfResponse(requestLength models.RequestLength) *models.LengthUnit {
	calculatorLength := models.StrucCalculatorLength{UnitToConvertFrom: requestLength.UnitToConvertFrom, UnitToConvertTo: requestLength.UnitToConvertTo, Value: requestLength.RequestLength}
	return &models.LengthUnit{UnitToConvertFrom: requestLength.UnitToConvertFrom, UnitToConvertTo: requestLength.UnitToConvertTo, ResponseUnit: calculatorLengthh(calculatorLength)}
}

func structToBytes(requestLength models.RequestLength) []byte {
	var response []byte
	response, _ = json.Marshal(createStructOfResponse(requestLength))
	return response
}

func calculatorLengthh(structCalculatorLength models.StrucCalculatorLength) float64 {
	var responseLength float64

	switch structCalculatorLength.UnitToConvertFrom {
	case "km":
		if structCalculatorLength.UnitToConvertTo == "m" {
			responseLength = validationKm(structCalculatorLength.Value)
		}

	case "m":
		if structCalculatorLength.UnitToConvertTo == "km" {
			responseLength = convertOfMetreToKm(structCalculatorLength.Value)
		}
	}
	return responseLength
}

func validationKm(km float64) float64 {
	// 1km=0.001m
	oneMetre = float64(1)
	// 1000m= 1 km
	thousandMetre = float64(1000)
	return (multiplicationtwoNumbers(thousandMetre, km) / oneMetre)
}

func convertOfMetreToKm(metre float64) float64 {
	// 1km=0.001m
	oneMetre = float64(1)
	// 1000m=1km
	thousandMetre = float64(1000)
	return (multiplicationtwoNumbers(oneMetre, metre) / thousandMetre)
}

func multiplicationtwoNumbers(numberOne, numberTwo float64) float64 {
	return numberOne * numberTwo
}
