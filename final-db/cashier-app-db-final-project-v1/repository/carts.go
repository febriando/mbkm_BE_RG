package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	return []model.JoinCart{}, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	totalPrice := product.Price - ((product.Discount / 100) * product.Price)

	cart := model.Cart{
		ProductID:  product.ID,
		Quantity:   1,
		TotalPrice: totalPrice,
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		cartData := model.Cart{}

		if err := tx.Where("product_id = ?", cart.ProductID).First(&cartData).Error; err != nil {
			return err
		}

		if cartData.Quantity != 0 || cartData.ProductID != 0 {
			cartData.Quantity += cart.Quantity
			cartData.TotalPrice += cart.TotalPrice

			if err := tx.Table("carts").Where("product_id = ?", cart.ProductID).Updates(cartData).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Create(&cart).Error; err != nil {
				return err
			}
		}

		if err := tx.Exec("UPDATE TABLE products SET stock = stock-1 WHERE id = ?", cart.ProductID).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	return nil // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	return nil // TODO: replace this
}
