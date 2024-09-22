package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
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
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type      RoomType           `json:"type" bsom:"type"`
	BasePrice float64            `json:"base_price" bson:"base_price"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
}
