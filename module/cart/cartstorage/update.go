package cartstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"context"
)

func (s *sqlStore) AddProductsToCart(ctx context.Context, cartID int, data cartmodel.CartProductDetails) error {
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

func (s *sqlStore) RemoveProductsFromCart(ctx context.Context, cartID int, data cartmodel.RemoveCartProducts) error {
	db := s.db.Begin()

	for _, product := range data {
		if err := db.Table(cartmodel.CartProduct{}.TableName()).Where("cart_id = ? AND product_id = (?)", cartID, int(product.ProductUID.GetLocalID())).Delete(nil).Error; err != nil {
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

func (s *sqlStore) UpdateCartItemQuantity(ctx context.Context, cartID int, data *cartmodel.CartProductDetail) error {
	db := s.db.Table(cartmodel.CartProduct{}.TableName())

	if err := db.Where("cart_id = ? AND product_id = (?)", cartID, data.ProductId).Update("quantity", data.Quantity).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
