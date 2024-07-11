package queue

import (
	"github.com/streadway/amqp"
)

// Rabbitmq struct
type Rabbitmq struct{}

// Find message from rabbitmq queue with uuid
// func (rabbitmq *Rabbitmq) Consume(queueName string, uuid string) (<-chan amqp.Delivery, error) {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Declare a queue that will be created if not exists with some args
// 	q, err := ch.QueueDeclare(
// 		queueName, // name
// 		false,     // durable
// 		true,      // delete when unused
// 		false,     // exclusive
// 		false,     // no-wait
// 		nil,       // arguments
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get message from the queue
// 	msg, _, err := ch.Get(q.Name, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if msg.Headers["correlation_id"] != uuid {
// 		// Requeue the message
// 		err = ch.Nack(msg.DeliveryTag, false, false)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return rabbitmq.Consume(queueName, uuid)
// 	}

// 	return msg, nil
// }

// Func publish message to rabbitmq queue
func (rabbitmq *Rabbitmq) Publish(queueName string, message []byte, correlationId string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue that will be created if not exists with some args
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Publish a message to the queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Headers:     amqp.Table{"correlation_id": correlationId},
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}

	return nil
}
