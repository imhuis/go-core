package json1

type Movie struct {
	Title  string `json1:"title"`
	Year   int    `json1:"release"`
	Color  bool   `json1:"color,omitempty"`
	Actors []string
}
