package abstract

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
)

const (
	ErrorNotFound                 = "not found"
	ErrorExpectedAtLeastOneResult = "sql: no rows in result set"
)

type QueryMaker interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Ping() error
	Close() error
}

// Scannable interface for generic scan behavior
type Scannable interface {
	ScanRow(row *sql.Row) error
	ScanRows(rows *sql.Rows) error
}

type Repository struct {
	DB QueryMaker
}

var (
	_ Scannable = &models.User{}
	_ Scannable = &models.Order{}
	_ Scannable = &models.Category{}
	_ Scannable = &models.CategoryFollowedByUser{}
	_ Scannable = &models.OrderCategory{}
)

func New(db QueryMaker) *Repository {
	return &Repository{
		DB: db,
	}
}

// Clear is used only in tests
func (repo *Repository) Clear() {
	queries := []string{
		"delete from categories_followed_by_user;",
		"delete from orders_categories;",
		"delete from categories;",
		"delete from orders;",
		"delete from users;",
	}

	for _, v := range queries {
		_, err := repo.DB.Exec(v)
		if err != nil {
			panic("error when clearing " + err.Error())
		}
	}
}

// category

// SelectOneBy gets exactly one or returns error
func (repo *Repository) SelectOneBy(ctx context.Context, args ...interface{}) (*models.Category, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("expected at least 2 args (key and value)")
	}

	conditions, onlyArgVals := prepareConditionsAndArgValues(args)
	fullQuery := `SELECT * FROM categories  ` + " WHERE " + strings.Join(conditions, " AND ")

	return repo.querySingleRowAndScan(ctx, fullQuery, onlyArgVals)
}

// FindBy gets zero, one or several entities. no error if not found
func (repo *Repository) FindBy(ctx context.Context, args ...interface{}) ([]models.Category, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("expected at least 2 args (key and value)")
	}

	conditions, onlyArgVals := prepareConditionsAndArgValues(args)
	fullQuery := `SELECT * FROM categories  ` + " WHERE " + strings.Join(conditions, " AND ")

	return repo.queryMultipleRowsAndScan(ctx, fullQuery, onlyArgVals)
}

// querySingleRow Helper to execute a single row query and scan into a provided struct
func (repo *Repository) querySingleRow(ctx context.Context, query string, args []interface{}, scanFunc func(row *sql.Row) error) error {
	row := repo.DB.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return row.Err()
	}
	return scanFunc(row)
}

// queryMultipleRows Helper to execute a query and collect multiple rows
func (repo *Repository) queryMultipleRows(ctx context.Context, query string, args []interface{}, scanFunc func(rows *sql.Rows) error) error {
	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err = scanFunc(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

func (repo *Repository) querySingleRowAndScan(ctx context.Context, query string, args []interface{}) (*models.Category, error) {
	cat := models.Category{}
	err := repo.querySingleRow(ctx, query, args, func(row *sql.Row) error {
		scannedCategory, err := scanCategory(row)
		if err != nil {
			return err
		}
		cat = scannedCategory
		return nil
	})
	return &cat, err
}

func (repo *Repository) queryMultipleRowsAndScan(ctx context.Context, query string, args []interface{}) ([]models.Category, error) {
	cats := make([]models.Category, 0)
	err := repo.queryMultipleRows(ctx, query, args, func(rows *sql.Rows) error {
		scannedCategory, err := scanCategories(rows)
		if err != nil {
			return err
		}
		cats = append(cats, scannedCategory)
		return nil
	})

	return cats, err
}

// scanCategory Helper function to scan a Category from a row
func scanCategory(row *sql.Row) (models.Category, error) {
	cat := models.Category{}
	err := row.Scan(&(cat.ID), &(cat.Name))
	return cat, err
}

// scanCategories Helper function to scan multiple Categories from rows
func scanCategories(rows *sql.Rows) (models.Category, error) {
	cat := models.Category{}
	err := rows.Scan(&(cat.ID), &(cat.Name))
	return cat, err
}

func prepareConditionsAndArgValues(args []interface{}) ([]string, []interface{}) {
	parts := make([]string, 0)
	onlyArgVals := make([]interface{}, 0)
	argName := ""
	for i, v := range args {
		if i%2 == 0 {
			argName = v.(string)
		} else {
			onlyArgVals = append(onlyArgVals, v)
			parts = append(parts, fmt.Sprintf("%v = $%d", argName, len(onlyArgVals)))
		}
	}
	return parts, onlyArgVals
}

// GenericSelectOneBy Generic method for getting one item from the DB by key-value condition
func GenericSelectOneBy[T Scannable](ctx context.Context, repo *Repository, args ...interface{}) (T, error) {
	var obj T
	if len(args) == 0 {
		return obj, fmt.Errorf("expected at least 2 args (key and value)")
	}

	fullQuery, onlyArgVals := getFullQueryAndArgs(args, obj)
	logger.Debugln("fullQuery", fullQuery)
	logger.Debugln("onlyArgVals", onlyArgVals)

	err := repo.querySingleRow(ctx, fullQuery, onlyArgVals, func(row *sql.Row) error {
		return obj.ScanRow(row)
	})

	logger.Debugln("after repo.querySingleRow")
	logger.Debugln("obj", obj)

	return obj, err
}

// FindByGeneric Generic method for finding items from the DB by key-value condition
func FindByGeneric[T Scannable](ctx context.Context, repo *Repository, args ...interface{}) ([]T, error) {
	objs := make([]T, 0)
	if len(args) == 0 {
		return objs, fmt.Errorf("expected at least 2 args (key and value)")
	}
	var obj T

	fullQuery, onlyArgVals := getFullQueryAndArgs(args, obj)

	err := repo.queryMultipleRows(ctx, fullQuery, onlyArgVals, func(rows *sql.Rows) error {
		err := obj.ScanRows(rows)
		if err != nil {
			return err
		}
		objs = append(objs, obj)
		return nil
	})
	return objs, err
}

func FindAllGeneric[T Scannable](ctx context.Context, repo *Repository) ([]T, error) {
	objs := make([]T, 0)
	var obj T

	fullQuery := "SELECT * " + " FROM " + getTableName(obj)

	err := repo.queryMultipleRows(ctx, fullQuery, []any{}, func(rows *sql.Rows) error {
		err := obj.ScanRows(rows)
		if err != nil {
			return err
		}
		objs = append(objs, obj)
		return nil
	})
	return objs, err
}

func getFullQueryAndArgs(args []interface{}, obj any) (string, []interface{}) {
	conditions, onlyArgVals := prepareConditionsAndArgValues(args)

	// divided so IDE does not say it's incorrect sql
	fullQuery := "SELECT * " + " FROM " + getTableName(obj) + " WHERE " + strings.Join(conditions, " AND ")
	return fullQuery, onlyArgVals
}

func getTableName(obj any) string {
	switch any(obj).(type) {
	case *models.Category:
		return "categories"
	case models.Category:
		return "categories"
	case *models.User:
		return "users"
	case models.User:
		return "users"
	case *models.Order:
		return "orders"
	case models.Order:
		return "orders"
	case *models.CategoryFollowedByUser:
		return "categories_followed_by_users"
	case models.CategoryFollowedByUser:
		return "categories_followed_by_users"
	case *models.OrderCategory:
		return "orders_categories"
	case models.OrderCategory:
		return "orders_categories"
	default:
		return ""
	}
}

//
//func initializeObject[T Scannable](obj T) T {
//	switch any(obj).(type) {
//	case *models.Category:
//		return &models.Category{}
//	case models.Category:
//		return models.Category{}
//	case *models.User:
//		return &models.User{}
//	case models.User:
//		return models.User{}
//	case *models.Order:
//		return &models.Order{}
//	case models.Order:
//		return models.Order{}
//	case *models.CategoryFollowedByUser:
//		return &models.CategoryFollowedByUser{}
//	case models.CategoryFollowedByUser:
//		return models.CategoryFollowedByUser{}
//	case *models.OrderCategory:
//		return &models.OrderCategory{}
//	case models.OrderCategory:
//		return models.OrderCategory{}
//	default:
//		return nil
//	}
//}
