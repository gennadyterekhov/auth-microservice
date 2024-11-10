package repositories

import (
	"context"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock is used in tests in place of a real-DB repo. naive implementation using maps.
type RepositoryMock struct {
	mock.Mock
	users                            map[int64]*models2.User
	lastUsedUserID                   int64
	lastUsedOrderID                  int64
	lastUsedCategoryID               int64
	lastUsedOrderCategoryID          int64
	lastUsedCategoryFollowedByUserID int64
	orders                           map[int64]*models2.Order
	categories                       map[int64]*models2.Category
	categoriesFollowedByUsers        map[string]*models2.CategoryFollowedByUser
	ordersCategories                 map[string]*models2.OrderCategory
}

var _ repositories.RepositoryInterface = &RepositoryMock{}

func NewMock() *RepositoryMock {
	return &RepositoryMock{
		users:                     make(map[int64]*models2.User),
		orders:                    make(map[int64]*models2.Order),
		categories:                make(map[int64]*models2.Category),
		categoriesFollowedByUsers: make(map[string]*models2.CategoryFollowedByUser),
		ordersCategories:          make(map[string]*models2.OrderCategory),
	}
}

func (repo *RepositoryMock) Clear() {
	repo.users = make(map[int64]*models2.User)
	repo.orders = make(map[int64]*models2.Order)

	repo.categories = make(map[int64]*models2.Category)
	repo.categoriesFollowedByUsers = make(map[string]*models2.CategoryFollowedByUser)
	repo.ordersCategories = make(map[string]*models2.OrderCategory)
}

// users

func (repo *RepositoryMock) InsertUser(ctx context.Context, login string, password string, bio string) (*models2.User, error) {
	alreadyExisting, err := repo.SelectUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	if alreadyExisting != nil {
		return nil, fmt.Errorf("ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)")
	}

	repo.lastUsedUserID += 1
	newID := repo.lastUsedUserID
	user := &models2.User{
		ID:       newID,
		Login:    login,
		Password: password,
		Bio:      bio,
	}

	repo.users[newID] = user

	return user, nil
}

func (repo *RepositoryMock) SelectUserByID(_ context.Context, id int64) (*models2.User, error) {
	user, ok := repo.users[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}

func (repo *RepositoryMock) SelectUserByLogin(_ context.Context, login string) (*models2.User, error) {
	for _, v := range repo.users {
		if v.Login == login {
			return v, nil
		}
	}

	return nil, nil
}

func (repo *RepositoryMock) UpdateUser(_ context.Context, id int64, bio string) error {
	_, ok := repo.users[id]
	if ok {
		repo.users[id].Bio = bio
		return nil
	}

	return fmt.Errorf("user with id %d not found", id)
}
