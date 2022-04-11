package entities

import "time"

type Detail struct {
	ID               int       `json:"id" db:"id_detail"`
	PlantId          int64     `json:"plantId" db:"id_plant"`
	Time             time.Time `json:"time" db:"time"`
	InternalHumidity float32   `json:"internalHumidity" db:"internal_humidity" binding:"required"`
	ExternalHumidity float32   `json:"externalHumidity" db:"external_humidity" binding:"required"`
	Temperature      float32   `json:"temperature" db:"temp" binding:"required"`
	Luminosity       float32   `json:"luminosity" db:"luminosity" binding:"required"`
}
