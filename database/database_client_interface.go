package database

type IDatabaseClient interface {
	Ping() error
	CloseDatabase()
}
