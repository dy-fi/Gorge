package message

import (
	"github.com/globalsign/mgo/bson"
)

// Message model
type Message struct {
	ID []bson.ObjectId 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Content []byte		`json:"content,omitempty" bson:"content,omitempty"`
}