package data

import "github.com/jackc/pgx/v5/pgxpool"

type Models struct {
	Book BookModel
}

func NewModels(pool *pgxpool.Pool) Models {
	return Models{
		Book: BookModel{DB: pool},
	}
}
