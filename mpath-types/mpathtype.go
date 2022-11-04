package mpathtypes

type FetchJSON struct {
	SampleCount int `json:"sample-count"`
	Results     map[string]Result `json:"results"`
	Disclaimer  string `json:"disclaimer"`
}

type Result struct {
	Meta_data               MetaData                `json:"meta-data"`
}

type MetaData struct {
	Dmp_sample_id         string  `json:"dmp_sample_id"`
}
