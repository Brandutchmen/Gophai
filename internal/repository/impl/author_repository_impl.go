package impl

import (
	"app/internal/database"
	"app/internal/entity"
	"app/internal/repository"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthorRepositoryImpl struct {
	Conn *pgx.Conn
}

func NewAuthorRepositoryImpl(Conn *pgx.Conn) repository.AuthorRepository {
	return &AuthorRepositoryImpl{Conn: Conn}
}

func (a *AuthorRepositoryImpl) Insert(ctx context.Context, author entity.Author) (entity.Author, error) {
	q := database.New(a.Conn)

	params := database.CreateAuthorParams{
		Name: author.Name,
		Bio:  pgtype.Text{String: author.Bio},
	}

	createdAuthor, err := q.CreateAuthor(ctx, params)
	if err != nil {
		return entity.Author{}, err
	}

	return entity.Author{
		ID:   createdAuthor.ID,
		Name: createdAuthor.Name,
		Bio:  createdAuthor.Bio.String,
	}, nil

}

func (a *AuthorRepositoryImpl) Update(ctx context.Context, author entity.Author) (entity.Author, error) {
	return entity.Author{}, nil
}

func (a *AuthorRepositoryImpl) Delete(ctx context.Context, id int64) (entity.Author, error) {
	return entity.Author{}, nil
}

func (a *AuthorRepositoryImpl) List(ctx context.Context) ([]entity.Author, error) {
	return []entity.Author{}, nil
}

func (a *AuthorRepositoryImpl) Find(ctx context.Context, id int64) (entity.Author, error) {
	return entity.Author{}, nil
}
