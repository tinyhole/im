package gateway

type AuthClient interface {
	Auth(uid int64, token string) (bool, error)
}
