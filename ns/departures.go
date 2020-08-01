type NsStation []struct {
	Direction             string `json:"direction"`
	Name                  string `json:"name"`
	PlannedDateTime       string `json:"plannedDateTime"`
	PlannedTimeZoneOffset int    `json:"plannedTimeZoneOffset"`
	ActualDateTime        string `json:"actualDateTime"`
	ActualTimeZoneOffset  int    `json:"actualTimeZoneOffset"`
	PlannedTrack          string `json:"plannedTrack"`
	Product               struct {
		Number            string `json:"number"`
		CategoryCode      string `json:"categoryCode"`
		ShortCategoryName string `json:"shortCategoryName"`
		LongCategoryName  string `json:"longCategoryName"`
		OperatorCode      string `json:"operatorCode"`
		OperatorName      string `json:"operatorName"`
		Type              string `json:"type"`
	} `json:"product"`
	TrainCategory string `json:"trainCategory"`
	Cancelled     bool   `json:"cancelled"`
	RouteStations []struct {
		UicCode    string `json:"uicCode"`
		MediumName string `json:"mediumName"`
	} `json:"routeStations"`
	DepartureStatus string `json:"departureStatus"`
	Messages        []struct {
		Message string `json:"message"`
		Style   string `json:"style"`
	} `json:"messages,omitempty"`
}
