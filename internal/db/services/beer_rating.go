package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
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

// TODO update the parameter types in these method to accept the proper types

func (s *BeerRatingService) InsertRating(userID, beerID, rating string) (*BeerRating, error) {
	var br BeerRating

	row := s.DBPool.QueryRow(
		context.Background(),
		`INSERT INTO beer_ratings (user_id, beer_id, rating) VALUES ($1, $2, $3)
			RETURNING id, user_id, beer_id, rating`,
		userID,
		beerID,
		rating,
	)

	if err := row.Scan(&br.ID, &br.UserID, &br.BeerID, &br.Rating); err != nil {
		return nil, err
	}

	return &br, nil
}

func (s *BeerRatingService) GetRating(ratingID string) (*BeerRating, error) {
	var br BeerRating

	row := s.DBPool.QueryRow(
		context.Background(),
		"SELECT id, user_id, beer_id, rating FROM beer_ratings WHERE id = $1",
		ratingID,
	)

	if err := row.Scan(&br.ID, &br.UserID, &br.BeerID, &br.Rating); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return &br, nil
}

func (s *BeerRatingService) GetRatingsByUser(userID string) ([]*BeerRating, error) {
	rows, err := s.DBPool.Query(
		context.Background(),
		"SELECT id, user_id, beer_id, rating FROM beer_ratings WHERE user_id = $1;",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ratings := []*BeerRating{}

	for rows.Next() {
		var br BeerRating

		err := rows.Scan(&br.ID, &br.UserID, &br.BeerID, &br.Rating)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, &br)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return ratings, nil
}
