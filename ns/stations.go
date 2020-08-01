type Stations []struct {
	Sporen               []interface{} `json:"sporen"`
	Synoniemen           []interface{} `json:"synoniemen"`
	HeeftFaciliteiten    bool          `json:"heeftFaciliteiten"`
	HeeftVertrektijden   bool          `json:"heeftVertrektijden"`
	HeeftReisassistentie bool          `json:"heeftReisassistentie"`
	Code                 string        `json:"code"`
	Namen                struct {
		Lang   string `json:"lang"`
		Kort   string `json:"kort"`
		Middel string `json:"middel"`
	} `json:"namen"`
	StationType   string  `json:"stationType"`
	Land          string  `json:"land"`
	UICCode       string  `json:"UICCode"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	Radius        int     `json:"radius"`
	NaderenRadius int     `json:"naderenRadius"`
	EVACode       string  `json:"EVACode"`
	IngangsDatum  string  `json:"ingangsDatum"`
}