package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var CategoryTester = struct {
	category     *Category
	categoryRepo *CategoryRepository
}{
	category: &Category{
		Name: "Tablets",
	},
}

func TestCategoryModel(t *testing.T) {
	// Setup dependency can be replaced with mocked db.DBClient
	CategoryTester.categoryRepo = NewCategoryRepository(Tester.dbClient)

	t.Run("Test list categories", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		db, close := CategoryTester.categoryRepo.Connect()
		defer close()

		assert.True(t, db.NewRecord(CategoryTester.category))
		categories, err := CategoryTester.categoryRepo.List()
		assert.True(t, len(categories) > 0)
		assert.Empty(t, err)
	})
}
