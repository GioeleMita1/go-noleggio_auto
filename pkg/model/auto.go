package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auto struct{
	ID primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	Marca string 			`bson:"marca" json:"marca"`
	Modello string			`bson:"modello" json:"modello"`
	Targa string			`bson:"targa" json:"targa"`
	Disponibile bool		`bson:"disponibile" json:"disponibile"`
}