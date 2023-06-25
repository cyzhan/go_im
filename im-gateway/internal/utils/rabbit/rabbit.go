package rabbit

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

func exitIfError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func InitRabbit() {
	var err error
	conn, err = amqp.Dial(os.Getenv("AMQP_DIAL"))
	exitIfError(err)

	channel, err = conn.Channel()
	exitIfError(err)
	log.Printf("package init ok: rabbit")
}

// func SendWithPB(m protoreflect.ProtoMessage) {
// 	r, err := proto.Marshal(m)
// 	if err != nil {
// 		log.Printf("proto.Marshal fail")
// 	}

// 	channel.Publish(
// 		"amq.topic",  // exchange
// 		"ws.generic", // routing key
// 		false,        // mandatory
// 		false,        // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(r),
// 		})
// }

func PbConvertAndSendMQ(m protoreflect.ProtoMessage, exchange string, routingKey string) {
	r, err := proto.Marshal(m)
	if err != nil {
		log.Printf("proto.Marshal fail")
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
