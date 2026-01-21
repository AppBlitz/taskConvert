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
		http.Error(w, "Error read body  ", http.StatusConflict)
	}
	if r.Method == http.MethodPost {
		_, errrorWrite := w.Write(convertDataToByte(dataRequest))
		if errrorWrite != nil {
			http.Error(w, "Error in write response", http.StatusConflict)
		}
	} else {
		http.Error(w, "Method invalid", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-type", "Application/json")
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + valueZeroCelsiusInKelvin
}

func createSolutionTemperature(dataRequest models.RequestLength) float64 {
	var solutionTemperature float64
	switch dataRequest.UnitToConvertFrom {
	case "c":
		if dataRequest.UnitToConvertTo == "k" {
			solutionTemperature = celsiusToKelvin(dataRequest.RequestLength)
		}
	case "k":
		switch dataRequest.UnitToConvertTo {
		case "c":
			solutionTemperature = kelvinToCelsius(dataRequest.RequestLength)
		case "f":
			solutionTemperature = kelvinToFarenheit(dataRequest.RequestLength)
		}

	}
	return solutionTemperature
}

func convertDataToByte(data models.RequestLength) []byte {
	var modelResponse models.LengthUnit
	var solution []byte
	modelResponse = models.LengthUnit{UnitToConvertFrom: data.UnitToConvertFrom, UnitToConvertTo: data.UnitToConvertTo, ResponseUnit: createSolutionTemperature(data)}
	solution, _ = json.Marshal(&modelResponse)
	return solution
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - valueZeroCelsiusInKelvin
}

func kelvinToFarenheit(degreeKelvin float64) float64 {
	var solutionDegree float64
	solutionDegree = degreeKelvin - float64(273.15)
	solutionDegree = solutionDegree * float64(1.8)
	return solutionDegree + 32
}
