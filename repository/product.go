package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mini_shop/global"
	"mini_shop/model"
)

type ProductDAO struct {
	db *gorm.DB
}

func NewProductDAO() *ProductDAO {
	return &ProductDAO{
		db: global.GetDB(),
	}
}

func (d *ProductDAO) ListProducts(offset, pageSize, categoryID int, keyword, sort string) ([]model.Product, int64, error) {
	var products []model.Product
	var count int64

	currentDB := d.db.Model(&model.Product{}).Where("status = ?", "on_sale")

	if keyword != "" {
		currentDB = currentDB.Where("`name` LIKE ?", "%"+keyword+"%")
	}
	if categoryID != 0 {
		currentDB = currentDB.Where("`category_id` = ?", categoryID)
	}

	switch sort {
	case "price_desc":
		currentDB = currentDB.Order("price DESC")
	case "price_asc":
		currentDB = currentDB.Order("price ASC")
	case "sold_desc":
		currentDB = currentDB.Order("sold DESC")
	default:
		currentDB = currentDB.Order("created_at DESC")
	}

	currentDB.Count(&count)
	err := currentDB.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}

func (d *ProductDAO) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := d.db.Where("id = ? AND status = ?", id, "on_sale").First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
func (d *ProductDAO) GetProductByIDWithLock(tx *gorm.DB, id uint) (*model.Product, error) {
	var product model.Product
	// 用 tx 执行，加悲观锁（FOR UPDATE），防止并发修改
	return &product, tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", id).First(&product).Error
}

func (d *ProductDAO) DecreaseStockInTx(tx *gorm.DB, id uint, quantity int) error {
	// 用 tx 执行，确保扣库存操作在事务内
	return tx.Model(&model.Product{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (d *ProductDAO) CreateProduct(product *model.Product) error {
	return d.db.Create(product).Error
}

func (d *ProductDAO) UpdateProduct(product *model.Product) error {
	return d.db.Save(product).Error //不用updates是为里跳过0值更新（例如stock为0和默认值一样时，updates就会默认跳过更新）
}

func (d *ProductDAO) DeleteProduct(id uint) error {
	return d.db.Delete(&model.Product{}, id).Error
}
