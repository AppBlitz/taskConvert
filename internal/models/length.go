// Package models
package models

type LengthUnit struct {
	UnitToConvertFrom string  `json:"unitToConvertFrom"`
	UnitToConvertTo   string  `json:"unitToConvertTo"`
	ResponseUnit      float64 `json:"responseUnit"`
}

type RequestLength struct {
	UnitToConvertFrom string  `json:"unitToConvertFrom"`
	UnitToConvertTo   string  `json:"unitToConvertTo"`
	RequestLength     float64 `json:"requestLength"`
}

type StrucCalculatorLength struct {
	UnitToConvertFrom string
	UnitToConvertTo   string
	Value             float64
}
