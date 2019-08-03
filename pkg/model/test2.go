package model

type Test2Model struct {
	//ID    objectid.ObjectID "_id,omitempty"
	Name string `bson:"name",json:"name"`
}
