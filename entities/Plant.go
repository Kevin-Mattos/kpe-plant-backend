package entities

type Plants []*Plant

type Plant struct {
	ID    int    `form:"id" json:"id" xml:"id"`
	Nome  string `form:"nome" json:"nome" xml:"nome" binding:"required"`
	Idade int    `form:"idade" json:"idade" xml:"idade" binding:"required"`
}
