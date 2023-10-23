package main

/*
 * This file will hold any structs used for APIs - requests or responses.
 * This is to keep things clean in our other files.
 */

// Used this: https://transform.tools/json-to-go
type SatelliteInfo struct {
	Timestamp string `json:"Timestamp"`
	Info      struct {
		Satname           string `json:"satname"`
		Satid             int    `json:"satid"`
		Transactionscount int    `json:"transactionscount"`
	} `json:"info"`
	Positions []struct {
		Satlatitude  float64 `json:"satlatitude"`
		Satlongitude float64 `json:"satlongitude"`
		Sataltitude  float64 `json:"sataltitude"`
		Azimuth      float64 `json:"azimuth"`
		Elevation    float64 `json:"elevation"`
		Ra           float64 `json:"ra"`
		Dec          float64 `json:"dec"`
		Timestamp    int     `json:"timestamp"`
		Eclipsed     bool    `json:"eclipsed"`
	} `json:"positions"`
}
