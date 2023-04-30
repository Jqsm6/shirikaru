package repository

const (
	uploadAnime = `INSERT INTO anime_list (title, alternativeTitle, description, productionStatus, picture, episode) 
                      VALUES ($1, $2, $3, $4, $5, $6) 
                      RETURNING anime_id`
	getAnimeByID = `SELECT * FROM anime_list 
                        WHERE anime_id = $1`
	getAnimeByTitle = `SELECT anime_id AS id, title, alternativeTitle, description, productionStatus, picture, episode
					       FROM anime_list
                           WHERE title LIKE '%' || $1 || '%' OR alternativeTitle LIKE '%' || $1 || '%'`
	deleteAnime = `DELETE FROM anime_list 
                       WHERE anime_id = $1`
)
