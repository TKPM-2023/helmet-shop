package common

type DbType int

const (
	DbTypeCategory DbType = 1
	DbTypeUser     DbType = 2
	DbTypeProduct  DbType = 3
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
