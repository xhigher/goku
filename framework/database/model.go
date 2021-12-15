package database

type Model interface {
	DatabaseName() string
	TableName() string
}

type BaseModel interface {
	DatabaseName() string
	TableName() string
}

type Transaction struct {
}
