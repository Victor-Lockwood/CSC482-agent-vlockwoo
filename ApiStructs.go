package main

/*
 * This file will hold any structs used for APIs - requests or responses.
 * This is to keep things clean in our other files.
 */

// Used this: https://transform.tools/json-to-go
type SatelliteInfo struct {
	Info struct {
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

type SatelliteInfoDynamo struct {
	Timestamp    int     `json:"timestamp"`
	Satname      string  `json:"satname"`
	Satid        int     `json:"satid"`
	Satlatitude  float64 `json:"satlatitude"`
	Satlongitude float64 `json:"satlongitude"`
	Sataltitude  float64 `json:"sataltitude"`
	Azimuth      float64 `json:"azimuth"`
	Elevation    float64 `json:"elevation"`
	Ra           float64 `json:"ra"`
	Dec          float64 `json:"dec"`
	Eclipsed     bool    `json:"eclipsed"`
}

// We only really need one of the positions; the second one is where it is the next second.
// We'ree be polling hourly/half hourly so we don't need that level of granularity
func transformApiResponse(info SatelliteInfo) SatelliteInfoDynamo {
	return SatelliteInfoDynamo{
		Timestamp:    info.Positions[0].Timestamp,
		Satname:      info.Info.Satname,
		Satid:        info.Info.Satid,
		Satlatitude:  info.Positions[0].Satlatitude,
		Satlongitude: info.Positions[0].Satlongitude,
		Sataltitude:  info.Positions[0].Sataltitude,
		Azimuth:      info.Positions[0].Azimuth,
		Elevation:    info.Positions[0].Elevation,
		Ra:           info.Positions[0].Ra,
		Dec:          info.Positions[0].Dec,
		Eclipsed:     info.Positions[0].Eclipsed,
	}
}
