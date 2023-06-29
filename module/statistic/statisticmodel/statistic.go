package statisticmodel

type Statistic struct {
	UserCount     int64 `json:"user_count"`
	OrderCount    int64 `json:"order_count"`
	ProductCount  int64 `json:"product_count"`
	CategoryCount int64 `json:"category_count"`
	Revenue       [12]float64	`json:"revenue"`
}
