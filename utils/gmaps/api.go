package gmaps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"

	"github.com/umahmood/haversine"
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

func GetDistance(attr *domain.Attraction, query *domain.LocQuery) (domain.DistanceMatrix, error) {
	start := haversine.Coord{Lat: query.UserLat, Lon: query.UserLng}
	dest := haversine.Coord{Lat: attr.MapsDetail.GeometryInfo.Loc.Lat, Lon: attr.MapsDetail.GeometryInfo.Loc.Lng}

	_, distanceKM := haversine.Distance(start, dest)

	dst := domain.DistanceMatrix{}

	fmt.Println(distanceKM)

	dst.DistanceValue = int64(distanceKM)

	return dst, nil
}
