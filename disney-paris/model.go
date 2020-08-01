type DisneyEntity []struct {
	EntityID          string `json:"entityId"`
	Type              string `json:"type"`
	ParkID            string `json:"parkId"`
	Status            string `json:"status"`
	ResponseTimestamp string `json:"responseTimestamp"`
	SingleRider       struct {
		IsAvailable            bool        `json:"isAvailable"`
		SingleRiderWaitMinutes interface{} `json:"singleRiderWaitMinutes"`
	} `json:"singleRider"`
	PostedWaitMinutes string `json:"postedWaitMinutes"`
	RideDetails       struct {
		ContentType string `json:"contentType"`
		EntityType  string `json:"entityType"`
		ContentID   string `json:"contentId"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Location    struct {
			ID            string `json:"id"`
			Value         string `json:"value"`
			URLFriendlyID string `json:"urlFriendlyId"`
			IconFont      string `json:"iconFont"`
		} `json:"location"`
		Coordinates []struct {
			Lat  float64 `json:"lat"`
			Lng  float64 `json:"lng"`
			Type string  `json:"type"`
		} `json:"coordinates"`
		Closed    bool `json:"closed"`
		Schedules []struct {
			Language  string `json:"language"`
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
			Date      string `json:"date"`
			Status    string `json:"status"`
			Closed    bool   `json:"closed"`
		} `json:"schedules"`
		Age []struct {
			ID            string `json:"id"`
			Value         string `json:"value"`
			URLFriendlyID string `json:"urlFriendlyId"`
			IconFont      string `json:"iconFont"`
		} `json:"age"`
		Height []struct {
			ID            string `json:"id"`
			Value         string `json:"value"`
			URLFriendlyID string `json:"urlFriendlyId"`
			IconFont      string `json:"iconFont"`
		} `json:"height"`
		Interests []struct {
			ID            string `json:"id"`
			Value         string `json:"value"`
			URLFriendlyID string `json:"urlFriendlyId"`
			IconFont      string `json:"iconFont"`
		} `json:"interests"`
		Photopass   bool `json:"photopass"`
		FastPass    bool `json:"fastPass"`
		SingleRider bool `json:"singleRider"`
	} `json:"rideDetails,omitempty"`
}