package gmaps

import (
	"fmt"
	"io"
	"strings"
	"encoding/json"
	"net/http"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
	"github.com/devanfer02/nosudes-be/domain"


)


func GetRatings(attractionName string) (domain.MapsDetail, error) {
	mapsEndpoint := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/textsearch/json?query=%s&sensor=true&key=%s",
		strings.ReplaceAll(attractionName, " ", "+"),
		env.ProcEnv.MapsAPIKey,
	)

	apiResp, err := http.Get(mapsEndpoint)

	if err != nil {
		logger.ErrLog(layers.Service, "failed to fetch api", err)
		return domain.MapsDetail{}, domain.ErrFailedFetchOtherAPI
	}

	body, err := io.ReadAll(apiResp.Body)

	if err != nil {
		logger.ErrLog(layers.Service, "failed to read response body", err)
		return domain.MapsDetail{}, domain.ErrFailedFetchOtherAPI
	}

	var gmapsRef domain.GmapsRef

	err = json.Unmarshal(body, &gmapsRef)

	if err != nil {
		logger.ErrLog(layers.Service, "failed to unmarshal json", err)
		return domain.MapsDetail{}, domain.ErrFailedFetchOtherAPI
	}

	return gmapsRef.Results[0], nil
}
