package cartstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartmodel"
)

func (s *sqlStore) RemoveProductsFromCart(context context.Context, cartID int, data cartmodel.RemoveCartProducts) error {
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
