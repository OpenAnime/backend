package constani

type Anime struct {
	// Anime ıd
	Id string `json:"id"`

	// Anime name
	Name string `json:"name"`

	// Anime İnfo
	Info string `json:"info"`

	//Anime Pictures
	Pictures Picture `json:"pictures"`

	//Episodes
	Episodes []Episodes `json:"episodes"`
}

type AnimeHome struct {
	Id string `json:"id"`

	// Anime name
	Name string `json:"name"`

	// Anime İnfo
	Info string `json:"info"`

	//Anime Pictures
	Pictures Picture `json:"pictures"`
}

type Episodes struct {
	// Episode Id
	Id string `json:"id"`

	// Episode Name
	Name string `json:"name"`

	// Episode Duration
	Duration string `json:"duration"`

	// Episode Description
	Description string `json:"description"`

	// Episode Avatar
	Avatar string `json:"avatar"`

	// Viedo Url
	VidoURL string `json:"viedoURL"`
}

type Picture struct {
	// Anime Avatar
	Avatar string `json:"avatar"`

	// Anime Banner
	Banner string `json:"banner"`
}
