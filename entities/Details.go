package entities

type Details []*Detail

type Detail struct {
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}
