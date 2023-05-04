package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brunoOchoa.com/api-REST-FULL/requests"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("publisher.create", false, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			newCliente := &requests.ClienteCreateRequest{}

			err := json.Unmarshal(d.Body, newCliente)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf(newCliente.Name)

			// conectar no banco
			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				panic(err)
			}

			if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
				panic(err)
			}

			collection := client.Database("clientes").Collection("clientes")
			result, err := collection.InsertOne(context.TODO(), newCliente)
			// check for errors in the insertion
			if err != nil {
				panic(err)
			}

			fmt.Println(result.InsertedID)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
