// Package handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AppBlitz/taskConvert/internal/models"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodPost {
		_, err := w.Write(structToBytes(models.RequestLength{UnitToConvertFrom: "Km", UnitToConvertTo: "m", RequestLength: 1.3}))
		if err != nil {
			http.Error(w, "Error with response", http.StatusConflict)
		}
	} else {
		http.Error(w, "Method no valid", http.StatusMethodNotAllowed)
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
	if structCalculatorLength.UnitToConvertFrom == "Km" {
		if structCalculatorLength.UnitToConvertTo == "m" {
			responseLength = structCalculatorLength.Value
		}
	}
	return responseLength
}
