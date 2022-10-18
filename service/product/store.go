package product

import (
	"dizzy1021.dev/interview-impact/model"
	"dizzy1021.dev/interview-impact/util"
	"github.com/gin-gonic/gin"

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
	result := store.db.Create(&product)

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
	result := store.db.Where("id = ?", id).Delete(&model.Product{})

	if result.Error != nil {
		return result.Error
	}
	return nil	
}

func (store *ProductStore) FindProduct(ctx *gin.Context) ([]*model.Product, *util.Pagination, error) {
	var products []*model.Product
	pagination := util.NewPagination(ctx)
	result := store.db.Scopes(util.Paginate(products, &pagination, store.db)).Find(&products)

	if result.Error != nil {
		return nil, nil, result.Error
	}
	return products, &pagination, nil	
}

func (store *ProductStore) FindOneProduct(id string) (*model.Product, error) {
	var product = model.Product{ID: id}
	result := store.db.First(&product)

	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil	
}

func (store *ProductStore) FindOneProductByCode(code string) (*model.Product) {
	var product *model.Product
	_ = store.db.Where("code = ?", code).First(&product)	
	
	return product	
}