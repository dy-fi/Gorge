package message

type Repository struct {
	FindbyId(id interface{}) (*Message, error)
	Query(query {}interace) (*)
	FindAll() ([]*Message, error)
	Update(message *Message) (error)
	Save(message *Message) (entity.ID, error)
	Delete(message *Message) (entity.ID, error)
}