package requests

import (
	"github.com/brunoOchoa.com/api-REST-FULL/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var EmptyCreateCliente = ClienteCreateRequest{}

type ClienteCreateRequest struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name       string             `json:"name" binding:"required,min=2,max=100"`
	CPF        string             `json:"cpf"`
	Nascimento string             `json:"nascimento"`
	Endereco   domain.Endereco    `json:"endereco"`
}

type ClienteUpdateRequest struct {
	Name       string          `json:"name" binding:"required,min=2,max=100"`
	CPF        string          `json:"cpf"`
	Nascimento string          `json:"nascimento"`
	Endereco   domain.Endereco `json:"endereco"`
}
