package product

import (
	"dizzy1021.dev/interview-impact/model"
	"dizzy1021.dev/interview-impact/util"

	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (store *ProductStore) InsertProduct(product model.Product) error {
	result := store.db.Create(product)

	if result.Error != nil {
		return result.Error
	}
	return nil	
}

func (store *ProductStore) UpdateProduct(product model.Product) error {
	result := store.db.Save(product)

	if result.Error != nil {
		return result.Error
	}
	return nil	
}

func (store *ProductStore) DeleteProduct(id string) error {
	result := store.db.Delete(&model.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	return nil	
}

func (store *ProductStore) FindProduct(page, pageSize int) ([]*model.Product, error) {
	var products []*model.Product
	result := store.db.Scopes(util.Paginate(page, pageSize)).Find(products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil	
}

func (store *ProductStore) FindOneProduct(id string) (*model.Product, error) {
	var product = model.Product{ID: id}
	result := store.db.First(&product)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil	
}