package redis

import (
	"context"
	"fmt"
	"github.com/AaronGulman/chat-gin/internal/chat"
	"github.com/redis/go-redis/v9"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
	Protocol: 2,
})

var ctx = context.Background()

func Connect() {

	err := client.Set(ctx, "foo", "GoLang value", 0).Err()
	if err != nil {
		fmt.Printf("There has been an error: %s", err)
	}

	val, err := client.Get(ctx, "foo").Result()

	if err != nil {
		fmt.Printf("There has been an error: %s", err)
	}

	fmt.Println("The received value: ", val)

}

func SubscribeToMessages(hub *chat.Hub) {
	var sub = client.Subscribe(ctx, "messages")
	ch := sub.Channel()

	defer sub.Close()
	go func() {
		for msg := range ch {
			fmt.Println("Message received: ", msg.Channel, msg.Payload)
		}
	}()
}

func PubMsg(msg string) {

	message := fmt.Sprintf("Text %s", msg)
	err := client.Publish(ctx, "messages", message).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Messages sent: ", message)
	time.Sleep(time.Duration(1000))
}
