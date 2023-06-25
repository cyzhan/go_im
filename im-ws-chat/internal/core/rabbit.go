package core

import (
	"encoding/json"
	"log"
	"os"

	"imws/internal/constant"
	"imws/internal/constant/topic"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	conn             *amqp.Connection
	channel          *amqp.Channel
	randomQueueNames map[string]string
)

func init() {
	randomQueueNames = make(map[string]string)
	randomQueueNames[topic.CHAT_MESSAGE] = uuid.NewString()
	randomQueueNames[topic.WS_GENERIC] = uuid.NewString()
	for routingKey, name := range randomQueueNames {
		log.Printf("%s: %s", routingKey, name)
	}
}

func startMQ() {
	var err error
	conn, err = amqp.Dial(os.Getenv("AMQP_DIAL"))
	shutDownIfError(err)

	channel, err = conn.Channel()
	shutDownIfError(err)

	chatMsgQueue, err := channel.QueueDeclare(
		randomQueueNames[topic.CHAT_MESSAGE], // name
		false,                                // durable
		true,                                 // delete when unused
		false,                                // exclusive
		false,                                // no-wait
		nil,                                  // arguments
	)
	shutDownIfError(err)

	err = channel.QueueBind(
		chatMsgQueue.Name,  // queue name
		topic.CHAT_MESSAGE, // routing key
		constant.AMQ_TOPIC, // exchange
		false,
		nil)
	shutDownIfError(err)

	wsGenericQueue, err := channel.QueueDeclare(
		randomQueueNames[topic.WS_GENERIC], // name
		false,                              // durable
		true,                               // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)

	err = channel.QueueBind(
		wsGenericQueue.Name, // queue name
		topic.WS_GENERIC,    // routing key
		constant.AMQ_TOPIC,  // exchange
		false,
		nil)

	shutDownIfError(err)
}

func ConvertAndSendMQ(v any, exchange string, routingKey string) {
	r, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal fail")
	}

	channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(r),
		})
}

func PbConvertAndSendMQ(m protoreflect.ProtoMessage, exchange string, routingKey string) {
	r, err := proto.Marshal(m)
	if err != nil {
		log.Printf("proto.Marshal fail")
	}

	if err := channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(r),
		}); err != nil {
		log.Printf("channel.Publish fail")
	}
}

func shutDownIfError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
