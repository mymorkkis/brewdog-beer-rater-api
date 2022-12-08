CREATE FUNCTION set_updated_at() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE PROCEDURE  set_updated_at();

CREATE TRIGGER set_beer_ratings_updated_at
    BEFORE UPDATE ON beer_ratings
    FOR EACH ROW EXECUTE PROCEDURE  set_updated_at();
