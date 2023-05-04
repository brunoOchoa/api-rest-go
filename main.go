package main

import (
	"context"

	"github.com/brunoOchoa.com/api-REST-FULL/controllers"
	"github.com/brunoOchoa.com/api-REST-FULL/repository"
	"github.com/brunoOchoa.com/api-REST-FULL/routes"
	"github.com/brunoOchoa.com/api-REST-FULL/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database   = "goMongo"
	Collection = "account"
	MongoDbUrl = "mongodb://localhost:27017"
	AMQP_Url   = "amqp://guest@localhost:5672"
)

var (
	collection *mongo.Collection
	repo       repository.ClienteRepository
	serv       service.ClienteService
	controller controllers.Controller
	ctx        context.Context
	conn       *amqp.Connection
	ch         *amqp.Channel
)

func main() {
	log.Info("publishing")
	configMessagingQueue()
	initDatabase()
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()

	routes.RegisterHandlers(server, controller).RegisterHandlers()

	err := server.Run()

	if err != nil {
		failOnError(err, "Server nao conectou")
	}
}

func initializeLayers() {
	repo = repository.NewMongoRepository(collection, ctx, ch)
	serv = service.NewClienteService(repo)
	controller = controllers.NewController(serv)
}

func configMessagingQueue() {
	var err error
	conn, err = amqp.Dial(AMQP_Url)
	failOnError(err, "Failed to connect to rabbit mq")
	ch, err = conn.Channel()
	failOnError(err, "Failed to create channel")
}

func initDatabase() {
	log.Info("Connecting to MongoDb...")
	clientOptions := options.Client().ApplyURI(MongoDbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		failOnError(err, "Unable to connect to MongoDb")
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		failOnError(err, "MongoDb Connection is not responding")
	}

	log.Info("Connected to MongoDB!")

	collection = client.Database(Database).Collection(Collection)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}
