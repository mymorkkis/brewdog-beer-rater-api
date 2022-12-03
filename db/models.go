package db

type Model interface {
	BeerRating
}

type BeerRating struct {
	BeerID int  `json:"beerID"`
	Rating int8 `json:"rating"`
}
