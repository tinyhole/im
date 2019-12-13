package entity

type AutoIncrement struct {
	Key       string `bson:"_id"`
	CurrentID int64  `bson:"current_id"`
}
