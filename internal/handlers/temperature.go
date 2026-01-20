package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AppBlitz/taskConvert/internal/models"
)

var valueZeroCelsiusInKelvin float64 = 273.15

func HandlerTemperature(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dataRequest models.RequestLength
	data, _ := io.ReadAll(r.Body)
	erro := json.Unmarshal(data, &dataRequest)
	if erro != nil {
	}
	if r.Method == http.MethodPost {
		w.Write(convertDataToByte(dataRequest))
	} else {
		http.Error(w, "Method invalid", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "Application/json")
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + valueZeroCelsiusInKelvin
}

func createSolutionTemperature() float64 {
	var solutionTemperature float64

	return solutionTemperature
}

func convertDataToByte(data models.RequestLength) []byte {
	var modelResponse models.LengthUnit
	var solution []byte
	modelResponse = models.LengthUnit{UnitToConvertFrom: data.UnitToConvertFrom, UnitToConvertTo: data.UnitToConvertTo, ResponseUnit: createSolutionTemperature()}
	solution, _ = json.Marshal(&modelResponse)
	return solution
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - valueZeroCelsiusInKelvin
}
