package repository

import (
	"context"
	"encoding/json"

	"github.com/brunoOchoa.com/api-REST-FULL/domain"
	"github.com/brunoOchoa.com/api-REST-FULL/queue"
	"github.com/brunoOchoa.com/api-REST-FULL/requests"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClienteRepository interface {
	GetAllClientes() ([]domain.Cliente, error)
	GetCliente(id string) (domain.Cliente, error)
	CreateCliente(requests.ClienteCreateRequest) (domain.Cliente, error)
	UpdateCliente(string, requests.ClienteUpdateRequest) error
	DeleteCliente(string) error
}

type repository struct {
	collection *mongo.Collection
	ctx        context.Context
	ch         *amqp.Channel
}

const (
	GetQueue    = "publisher.get"
	CreateQueue = "publisher.create"
	UpdateQueue = "publisher.update"
	DeleteQueue = "publisher.delete"
)

func NewMongoRepository(collection *mongo.Collection, ctx context.Context,
	ch *amqp.Channel) ClienteRepository {

	return &repository{
		collection: collection,
		ctx:        ctx,
		ch:         ch,
	}
}

func (r *repository) GetAllClientes() ([]domain.Cliente, error) {

	q := createQueues(GetQueue, r.ch)

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Getting todos clientes"),
	}

	publishMessage(r.ch, q.Name, msg)

	cursor, err := r.collection.Find(r.ctx, bson.D{})
	defer cursor.Close(r.ctx)

	if err != nil {
		return []domain.Cliente{}, err
	}

	var accounts []domain.Cliente

	if cursor.All(r.ctx, &accounts); err != nil {
		return []domain.Cliente{}, err
	}

	return accounts, nil
}

func (r *repository) GetCliente(id string) (domain.Cliente, error) {
	q := createQueues(GetQueue, r.ch)

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Requesting cliente for ID " + id),
	}

	publishMessage(r.ch, q.Name, msg)

	account := domain.Cliente{}
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, err
	}
	err = r.collection.FindOne(r.ctx, bson.M{
		"_id": uid,
	}).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (r *repository) CreateCliente(request requests.ClienteCreateRequest) (domain.Cliente, error) {

	request.Id = primitive.NewObjectID()
	request.Endereco.DeptId = primitive.NewObjectID()

	q := createQueues(CreateQueue, r.ch)
	body, err := json.Marshal(request)

	if err != nil {
		return domain.Cliente{}, err
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	publishMessage(r.ch, q.Name, msg)

	_, err = r.collection.InsertOne(r.ctx, request)

	if err != nil {
		return domain.Cliente{}, err
	}

	return domain.Cliente{
		Id:         request.Id,
		Name:       request.Name,
		Nascimento: request.Nascimento,
		CPF:        request.CPF,
	}, nil

	return domain.Cliente{}, err
}

func (r *repository) UpdateCliente(id string, request requests.ClienteUpdateRequest) error {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(r.ctx, bson.M{
		"_id": uid,
	}, bson.M{
		"$set": request,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteCliente(id string) error {
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	q := createQueues(DeleteQueue, r.ch)
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Request to delete account with id " + id),
	}
	publishMessage(r.ch, q.Name, msg)

	_, err = r.collection.DeleteOne(r.ctx, bson.M{
		"_id": uid,
	})

	if err != nil {
		return err
	}

	return nil
}

func createQueues(name string, ch *amqp.Channel) amqp.Queue {
	return queue.NewQueue(name, ch).CreateQueue()
}

// Default exchange
func publishMessage(ch *amqp.Channel, name string, msg amqp.Publishing) {
	ch.Publish("", name, false, false, msg)
}
