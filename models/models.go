package models

type Response struct {
	IPAddress   string `json:"ip_address"`
	CountryCode string `json:"country_code"`
	Language    string `json:"language"`
	IsoString   string `json:"iso"`
}

//Language contains a Counyty and Language as strings
type Language struct {
	Language string `json:"language"`
	Country  string `json:"country"`
}

//DBlocation PRAMS
//Type: @{STRING}
//Location: @{STRING}
type DBLocation struct {
	Type     string
	Location string
}

//Languages
var Languages []Language

//ConfigFile is a string
var ConfigFile string

//CurrentHash is the current hash
var CurrentHash string
