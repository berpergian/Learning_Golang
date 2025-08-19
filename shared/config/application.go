package config

import "github.com/berpergian/chi_learning/shared/database"

type Application struct {
	Env      *Env
	Database database.IDatabaseClient
}
