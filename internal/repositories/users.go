package repositories

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/models"
)

func (repo *Repository) InsertUser(ctx context.Context, login, password, bio string) (*models.User, error) {
	const query = `INSERT INTO users ( login, password, bio) VALUES ( $1, $2, $3) RETURNING id;`

	row := repo.DB.QueryRowContext(ctx, query, login, password, bio)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       id,
		Login:    login,
		Password: password,
		Bio:      bio,
	}

	return &user, nil
}

func (repo *Repository) SelectUserByID(ctx context.Context, id int64) (*models.User, error) {
	const query = `SELECT id, login, password, bio from users WHERE  id = $1`
	row := repo.DB.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user := models.User{}
	err := user.ScanRow(row)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) SelectUserByLogin(ctx context.Context, login string) (*models.User, error) {
	const query = `SELECT id, login, password, bio from users WHERE  login = $1`
	row := repo.DB.QueryRowContext(ctx, query, login)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user := models.User{}
	err := user.ScanRow(row)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) UpdateUser(ctx context.Context, id int64, bio string) error {
	const query = `UPDATE users SET bio = $2 WHERE id = $1;`

	_, err := repo.DB.ExecContext(ctx, query, id, bio)
	if err != nil {
		return err
	}

	return nil
}
