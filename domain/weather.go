package domain

type Weather struct {
	Main string `json:"main"`
	Desc string `json:"description"`
}

type WeatherInfo struct {
	Weather []Weather `json:"weather"`
}

type WeatherList struct {
	List []WeatherInfo `json:"list"`
}
