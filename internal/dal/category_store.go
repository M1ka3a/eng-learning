package dal

import (
	"backend-lab-go/internal/model"
	"database/sql"
	"strings"
)

type CategoryStore struct {
	db *sql.DB
}

func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

func (s *CategoryStore) GetAll() ([]model.Category, error) {
	rows, err := s.db.Query(
		"SELECT id, title FROM categories ORDER BY id DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}

	for rows.Next() {
		var category model.Category

		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryStore) Add(title string) (model.Category, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return model.Category{}, model.ErrTitleInvalid
	}

	category := model.Category{}

	err := s.db.QueryRow(
		`INSERT INTO categories (title)
         VALUES ($1)
         RETURNING id, title`,
		title,
	).Scan(&category.ID, &category.Title)

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}
