package model

type Anime struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	AlternativeTitle string `json:"alternativeTitle"`
	Description      string `json:"description"`
	ProductionStatus string `json:"productionStatus"`
	Picture          string `json:"picture"`
	Episode          int    `json:"episode"`
}

type DBAnime struct {
	Title            string `pg:"title"`
	AlternativeTitle string `pg:"alternativeTitle"`
	Description      string `pg:"description"`
	ProductionStatus string `pg:"productionStatus"`
	Picture          string `pg:"picture"`
	Episode          int    `pg:"episode"`
}

func (a *Anime) ToDB() *DBAnime {
	return &DBAnime{
		Title:            a.Title,
		AlternativeTitle: a.AlternativeTitle,
		Description:      a.Description,
		ProductionStatus: a.ProductionStatus,
		Picture:          a.Picture,
		Episode:          a.Episode,
	}
}
