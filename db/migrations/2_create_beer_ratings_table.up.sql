CREATE TABLE beer_ratings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    beer_id INTEGER NOT NULL CHECK (beer_id >= 0),
    rating SMALLINT NOT NULL CHECK (rating >= 0 AND rating <= 10)
);
