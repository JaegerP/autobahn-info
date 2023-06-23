package lib

type Warning struct {
	Warning []struct {
		Extent              string        `json:"extent"`
		Identifier          string        `json:"identifier"`
		RouteRecommendation []interface{} `json:"routeRecommendation"`
		Coordinate          struct {
			Lat  string `json:"lat"`
			Long string `json:"long"`
		} `json:"coordinate"`
		Footer                   []interface{} `json:"footer"`
		Icon                     string        `json:"icon"`
		IsBlocked                string        `json:"isBlocked"`
		Description              []string      `json:"description"`
		Title                    string        `json:"title"`
		Point                    string        `json:"point"`
		DisplayType              string        `json:"display_type"`
		LorryParkingFeatureIcons []interface{} `json:"lorryParkingFeatureIcons"`
		Future                   bool          `json:"future"`
		Subtitle                 string        `json:"subtitle"`
		StartTimestamp           string        `json:"startTimestamp"`
	} `json:"warning"`
}
