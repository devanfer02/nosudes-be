package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
)

func GetWeatherInfo(lat, lng float64) ([]domain.Weather, error) {
	mapsEndpoint := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s",
		lat, lng,
		env.ProcEnv.WeatherAPIKey,
	)

	apiResp, err := http.Get(mapsEndpoint)

	if err != nil {
		logger.ErrLog(layers.Service, "failed to fetch api", err)
		return nil, domain.ErrFailedFetchOtherAPI
	}

	body, err := io.ReadAll(apiResp.Body)

	if err != nil {
		logger.ErrLog(layers.Service, "failed to read response body", err)
		return nil, domain.ErrFailedFetchOtherAPI
	}

	var infos domain.WeatherList

	if err = json.Unmarshal(body, &infos); err != nil {
		logger.ErrLog(layers.Service, "failed to unmarshal json", err)
		return nil, domain.ErrFailedFetchOtherAPI

	}

	if len(infos.List) == 0 {
		return nil, nil
	}

	weathers := make([]domain.Weather, 0)

	for _, info := range infos.List {
		weathers = append(weathers, info.Weather[0])
	}

	return weathers[:7], nil
}
