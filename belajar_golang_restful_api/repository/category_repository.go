package repository

import (
	"context"
	"database/sql"
	"zalllrizalll/belajar_golang_restful_api/model/domain"
)

type CategoryRepository interface {
	// Restfull API
	// Func Create
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Func Update
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Func Delete
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	// Func Get by ID
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	// Func Get All
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
