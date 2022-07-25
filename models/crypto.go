package models

type Crypto struct {
	Id    int64  `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Votes int64  `json:"votes" bson:"votes"`
}
