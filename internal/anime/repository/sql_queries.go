package repository

const (
	uploadAnime = `INSERT INTO anime_list (animeID, title, alternativeTitle, description, productionStatus, picture, episode) 
                      VALUES ($1, $2, $3, $4, $5, $6, $7) 
                      RETURNING animeID`
	getAnimeAll = `SELECT animeID, title, alternativeTitle, description, productionStatus, picture, episode
	                    FROM anime_list`
	getAnimeByID = `SELECT * FROM anime_list 
                        WHERE animeID = $1`
	searchAnimeByTitle = `SELECT *
					       FROM anime_list
                           WHERE title LIKE '%' || $1 || '%' OR alternativeTitle LIKE '%' || $1 || '%'`
	deleteAnime = `DELETE FROM anime_list 
                       WHERE animeID = $1`
)
