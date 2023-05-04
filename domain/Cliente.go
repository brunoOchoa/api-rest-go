package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cliente struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name       string             `json:"name" binding:"required,min=2,max=100"`
	CPF        string             `json:"cpf"`
	Nascimento string             `json:"nascimento"`
	Endereco   Endereco           `json:"endereco"`
}
