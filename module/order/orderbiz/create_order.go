package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/product/productbiz"
	"context"
)

type CreateOrderStore interface {
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) error
	FindOrderWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*ordermodel.Order, error)
}

type createOrderBusiness struct {
	createorder_store CreateOrderStore
	product_store	productbiz.GetProductStore
	orderdetail_store orderdetailbiz.CreateOrderDetailStore
}

func NewCreateOrderBusiness(createorder_store CreateOrderStore,
							product_store	productbiz.GetProductStore,
							orderdetail_store orderdetailbiz.CreateOrderDetailStore) *createOrderBusiness {
	return &createOrderBusiness{createorder_store: createorder_store,
								product_store: product_store,
								orderdetail_store: orderdetail_store,
	}
}

func (business *createOrderBusiness) CreateOrder(context context.Context, data *ordermodel.OrderCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := business.createorder_store.CreateOrder(context, data); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	for i := range data.Products {
		data.Products[i].Order_ID = data.Id
		//get product info from model products
		//product, err := business.product_store.GetProduct(context.Request.Context(), int(data.Products[i].Product_Origin.UID.GetLocalID()))
		product, err := business.product_store.FindProductWithCondition(context, map[string]interface{}{"id": data.Products[i].Product_Origin.UID.GetLocalID()})


		if err != nil {
			panic(err)
		}

		//assign to prodcut_origin
		data.Products[i].Product_Origin.Description = product.Description
		data.Products[i].Product_Origin.Name = product.Name
		data.Products[i].Price = (float64(product.Price) * float64(data.Products[i].Quantity)) - (float64(product.Price) * float64(data.Products[i].Discount))

		if err := business.orderdetail_store.CreateOrderDetail(context, &data.Products[i]); err != nil {
			panic(err)
		}
	}
	return nil
}
