package entities

type Plant struct {
	ID      int    `json:"id" db:"id_plant"`
	Name    string `json:"name" db:"name" binding:"required"`
	Species string `json:"species" db:"species" binding:"required"`
}
