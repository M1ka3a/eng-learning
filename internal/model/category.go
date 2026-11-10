package model

import "errors"

type Category struct {
	ID    int
	Title string
}

func AddCategory(categories []Category, title string) ([]Category, error) {
	if len(title) == 0 {
		return categories, errors.New("invalid title")
	}

	newCategory := Category{
		ID:    len(categories) + 1,
		Title: title,
	}

	categories = append(categories, newCategory)
	return categories, nil
}
