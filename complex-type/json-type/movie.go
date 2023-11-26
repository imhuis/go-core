package json_type

type Movie struct {
	Title  string `json:"title"`
	Year   int    `json:"release"`
	Color  bool   `json:"color,omitempty"`
	Actors []string
}
