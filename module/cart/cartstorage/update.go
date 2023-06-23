package cartstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) AddProductsToCart(context context.Context, cartID int, data cartmodel.CartProductDetails) error {
	db := s.db.Begin()

	for _, product := range data {
		cartProduct := &cartmodel.CartProduct{
			CartID:    cartID,
			ProductID: int(product.ProductUID.GetLocalID()),
			Quantity:  product.Quantity,
		}

		if err := db.Table(cartmodel.CartProduct{}.TableName()).Create(cartProduct).Error; err != nil {
			db.Rollback()
			return common.ErrDB(err)
		}
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateCartItemQuantity(context context.Context, cartID int, data *cartmodel.CartProductDetail) error {
	db := s.db.Table(cartmodel.CartProduct{}.TableName())

	if err := db.Where("cart_id = ? AND product_id = (?)", cartID, data.ProductId).Update("quantity", data.Quantity).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseTotalProduct(context context.Context, id, quantity int) error {
	if err := s.db.Table(cartmodel.Cart{}.TableName()).Where("id = ?", id).
		Update("total_product", gorm.Expr("total_product + ?", quantity)).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DecreaseTotalProduct(context context.Context, id, quantity int) error {
	if err := s.db.Table(cartmodel.Cart{}.TableName()).Where("id = ?", id).
		Update("total_product", gorm.Expr("total_product - ?", quantity)).Error; err != nil {
		return err
	}
	return nil
}
