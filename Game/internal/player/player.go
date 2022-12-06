package player

type Player struct {
	Id     string `bson,json:"id"`
	Name   string `bson,json:"name"`
	Points int32  `bson,json:"points"`
}
