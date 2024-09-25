package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AB-Rhman/hotel-reservation/db"
	"github.com/AB-Rhman/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	ctx        = context.Background()
)

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rating:   rating,
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []types.Room{
		{
			Size:    "small",
			SeaView: false,
			Price:   99.9,
		},
		{
			Size:    "normal",
			SeaView: false,
			Price:   199.9,
		},
		{
			Size:    "big",
			SeaView: true,
			Price:   299.9,
		},
		{
			Size:    "suite",
			SeaView: true,
			Price:   399.9,
		},
	}
	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	seedHotel("Hilton", "New York", 3)
	seedHotel("Marriott", "San Francisco", 4)
	seedHotel("Sheraton", "Los Angeles", 1)

	fmt.Println("Seed data inserted successfully")

}

func init() {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
