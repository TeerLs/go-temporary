package product

import "gorm.io/gorm"

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(product *Product) (*Product, error) {
	err := r.db.Table("products").Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepo) GetByID(id uint) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) GetAll() ([]*Product, error) {
	var products []*Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepo) Update(product *Product, id uint) error {
	return r.db.Model(&Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *ProductRepo) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}