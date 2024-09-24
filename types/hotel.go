package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   int                  `json:"rating" bson:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaViewRoomType
	SuiteRoomType
)

type Room struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Size    string             `json:"size" bson:"size"`
	SeaView bool               `json:"sea_view" bson:"sea_view"`
	Price   float64            `json:"price" bson:"price"`
	HotelID primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
}
