package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Endereco struct {
	DeptId primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Rua    string             `json:"rua"`
	Bairro string             `json:"bairro"`
	Cidade string             `json:"cidade"`
	Estado string             `json:"estado"`
}
