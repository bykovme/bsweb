package main

// Config - structure of config file
type Config struct {
	Port          string `json:"port"`
	TemplatesPath string `json:"templates"`
	AssetsPath    string `json:"assets"`
	DBPath        string `json:"db"`
}

// ErrorPageData - data for error information on the page
type ErrorPageData struct {
	ErrorHeader  string
	ErrorMessage string
}
