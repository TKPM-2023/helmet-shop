package orderstorage

import (
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
	"strconv"
)

type NResult struct {
	N float64 //or int ,or some else
}

func (s *sqlStore) GetRevenue(year int) [12]float64 {
	db := s.db.Table(ordermodel.EntityName)

	var sum [12]float64
	for i := 0; i < 12; i++ {
		var n NResult
		date_from := strconv.Itoa(year) + "-" + strconv.Itoa(i+1) + "-01"
		date_to := strconv.Itoa(year) + "-" + strconv.Itoa(i+2) + "-01"
		if i == 11 {
			date_to = strconv.Itoa(year+1) + "-01-01"
		}
		//db.Select("sum(total_price) as n").Where("created_at between ? and ?", date_from, date_to).Scan(&n)
		db.Raw("select sum(total_price) as n from orders where created_at between ? and ?", date_from, date_to).Scan(&n)
		sum[i] = n.N
	}

	return sum
}
