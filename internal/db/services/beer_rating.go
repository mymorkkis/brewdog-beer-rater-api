package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BeerRating struct {
	ID     uint  `json:"id"`
	UserID uint  `json:"userID"`
	BeerID uint  `json:"beerID"`
	Rating uint8 `json:"rating"`
}

type BeerRatingService struct {
	DBPool *pgxpool.Pool
}

func (m *BeerRatingService) InsertRating(userID, beerID, rating string) (*BeerRating, error) {
	var br BeerRating

	row := m.DBPool.QueryRow(
		context.Background(),
		"INSERT INTO beer_ratings (user_id, beer_id, rating) VALUES ($1, $2, $3) RETURNING id, user_id, beer_id, rating",
		userID,
		beerID,
		rating,
	)

	err := row.Scan(&br.ID, &br.UserID, &br.BeerID, &br.Rating)
	if err != nil {
		return nil, err
	}

	return &br, nil
}

func (m *BeerRatingService) GetRatingsByUser(userID string) ([]*BeerRating, error) {
	rows, err := m.DBPool.Query(
		context.Background(),
		"SELECT id, user_id, beer_id, rating FROM beer_ratings WHERE user_id = $1;",
		userID,
	)
	if err != nil {
		return nil, err
	}

	ratings := []*BeerRating{}

	for rows.Next() {
		var br BeerRating

		err := rows.Scan(&br.ID, &br.UserID, &br.BeerID, &br.Rating)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &br)
	}

	return ratings, nil
}
