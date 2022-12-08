package player

type Player struct {
	Id     string `bson,json:"id"`
	Name   string `bson,json:"name"`
	Points int32  `bson,json:"points"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL" envDefault:"postgresql://postgres:catalog@localhost:5432/catalog?sslmode=disable"`
	JwtKey        []byte `env:"JWT-KEY" `
}
