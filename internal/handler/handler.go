package handler

import db "github.com/maadiab/tasty-click/internal/db/sqlc"

type ApiConfig struct {
	DBQueries db.Queries
}
