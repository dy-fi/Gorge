package message

import (
	"github.com/globalsign/mgo/bson"
)

// Repository defines the model behavior for data being saved to the backend
type Repository interface {
	FindbyId(id interface{}) ([]*Message, error)
	Query(query interface{}) ([]*Message, error)
	FindAll() ([]*Message, error)
	Update(message *Message) (error)
	Save(message *Message) (bson.ObjectId, error)
	Delete(message *Message) (bson.ObjectId, error)
}

