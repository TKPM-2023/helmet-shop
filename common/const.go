package common

type DbType int

const (
	DbTypeCategory DbType = 1
	DbTypeUser     DbType = 2
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
