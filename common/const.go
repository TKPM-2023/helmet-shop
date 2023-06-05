package common

type DbType int

const (
	DbTypeCategory DbType = 1
	DbTypeUser     DbType = 2
	DbTypeProduct  DbType = 3
	DbTypeOrder	   DbType = 4
	DbTypeOrder_Detail DbType=5
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
