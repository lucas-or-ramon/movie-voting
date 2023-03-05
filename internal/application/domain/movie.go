package domain

type MovieRepository interface {
	CreateMovie(movie *Movie) error
	FindMovieByTitle(title string) (*[]Movie, error)
}

type Movie struct {
	ID               string
	Image            string
	Title            string
	Description      string
	Genres           string
	ImDbRating       float64
	ImDbRatingVotes  int64
	MetacriticRating float64
	Plot             string
	Stars            string
}
