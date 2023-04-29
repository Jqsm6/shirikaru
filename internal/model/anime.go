package model

type Anime struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	AlternativeTitle string `json:"alternative_title"`
	Description      string `json:"description"`
	ProductionStatus string `json:"production_status"`
	Picture          string `json:"picture"`
	Episode          int    `json:"episode"`
}

type DBAnime struct {
	Title            string `pg:"title,notnull"`
	AlternativeTitle string `pg:"alternative_title,notnull"`
	Description      string `pg:"description,notnull"`
	ProductionStatus string `pg:"production_status,notnull"`
	Picture          string `pg:"picture,notnull"`
	Episode          int    `pg:"episode,notnull"`
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
