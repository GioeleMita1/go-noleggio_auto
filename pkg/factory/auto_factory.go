package factory

import (
	"noleggio_auto/pkg/dto"
	"noleggio_auto/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModel(req dto.CreateAutoRequest) *model.Auto {
	return &model.Auto{
		ID:           primitive.NewObjectID(), 
		Marca:        req.Marca,
		Modello:      req.Modello,
		Targa:        req.Targa,
		Disponibile:  true, 
	}
}