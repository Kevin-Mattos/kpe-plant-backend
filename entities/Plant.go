package entities

type Plants []*Plant

type Plant struct {
	ID    int    `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome" binding:"required"`
	Idade int    `json:"idade" db:"idade" binding:"required"`
}
