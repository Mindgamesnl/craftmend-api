type EftelingApi struct {
	OpeningHours struct {
		Date           string `json:"Date"`
		BusyIndication string `json:"BusyIndication"`
		HourFrom       string `json:"HourFrom"`
		HourTo         string `json:"HourTo"`
	} `json:"OpeningHours"`
	AttractionInfo []struct {
		ID              string        `json:"Id"`
		Type            string        `json:"Type"`
		MapLocation     string        `json:"MapLocation"`
		ShowDuration    int           `json:"ShowDuration,omitempty"`
		State           string        `json:"State"`
		StateColor      string        `json:"StateColor,omitempty"`
		IsTheaterShow   bool          `json:"IsTheaterShow,omitempty"`
		PastShowTimes   []interface{} `json:"PastShowTimes,omitempty"`
		WaitingTime     string        `json:"WaitingTime,omitempty"`
		StatePercentage int           `json:"StatePercentage,omitempty"`
		ShowTimes       []interface{} `json:"ShowTimes,omitempty"`
		Timeslot        struct {
			State string `json:"State"`
		} `json:"Timeslot,omitempty"`
	} `json:"AttractionInfo"`
}
