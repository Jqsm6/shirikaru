package models

type Anime struct {
	AnimeID          int    `json:"animeID"`
	Title            string `json:"title"`
	AlternativeTitle string `json:"alternativeTitle"`
	Description      string `json:"description"`
	ProductionStatus string `json:"productionStatus"`
	Picture          string `json:"picture"`
	Episode          int    `json:"episode"`
}

type DBAnime struct {
	AnimeID          int    `pg:"animeID"`
	Title            string `pg:"title"`
	AlternativeTitle string `pg:"alternativeTitle"`
	Description      string `pg:"description"`
	ProductionStatus string `pg:"productionStatus"`
	Picture          string `pg:"picture"`
	Episode          int    `pg:"episode"`
}

func (a *Anime) ToDB() *DBAnime {
	return &DBAnime{
		AnimeID:          a.AnimeID,
		Title:            a.Title,
		AlternativeTitle: a.AlternativeTitle,
		Description:      a.Description,
		ProductionStatus: a.ProductionStatus,
		Picture:          a.Picture,
		Episode:          a.Episode,
	}
}
