package entities

type Plants []*Plant

type Plant struct {
	ID    int
	Nome  string
	Idade int
}
