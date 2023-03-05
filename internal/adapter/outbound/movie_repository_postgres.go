package outbound

import (
	"database/sql"
	"movie-voting/internal/application/domain"
)

type MovieRepositoryPostgres struct {
	DB *sql.DB
}

func (m *MovieRepositoryPostgres) CreateMovie(movie *domain.Movie) error {
	return nil
}
