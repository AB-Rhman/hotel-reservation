package main

import (
	"context"
	"flag"
	"log"

	"github.com/AB-Rhman/hotel-reservation/db"

	"github.com/AB-Rhman/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userStore = db.NewMongoUserStore(client)
		hotelStore = db.NewMongoHotelStore(client)
		roomStore = db.NewMongoRoomStore(client, hotelStore)
		store = &db.Store{
			User: userStore,
			Hotel: hotelStore,
			Room: roomStore,
		}
		userHandler = api.NewUserHandler(userStore)
		hotelHandeler = api.NewHotelHandler(store)
		app = fiber.New(config)
		apiv1 = app.Group("/api/v1")
	)
	
	// user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	
	// hotel handlers
	apiv1.Get("/hotel", hotelHandeler.HandleGetHotels)
	apiv1.Get("/hotel/:id/rooms", hotelHandeler.HandleGetHotelRooms)
	apiv1.Get("/hotel/:id", hotelHandeler.HandleGetHotel)
	
	app.Listen(*listenAddr)
}
