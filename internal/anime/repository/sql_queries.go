package repository

const (
	uploadAnime = `INSERT INTO anime_list (title, alternative_title, description, production_status, picture, episode) 
                      VALUES ($1, $2, $3, $4, $5, $6) 
                      RETURNING anime_id`
	getAnimeByID = `SELECT * FROM anime_list 
                        WHERE anime_id = $1`
	deleteAnime = `DELETE FROM anime_list 
                       WHERE anime_id = $1`
)
