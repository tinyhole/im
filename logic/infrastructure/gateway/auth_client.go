package gateway

import (
	"github.com/tinyhole/im/logic/domain/gateway"
)

type AuthClient struct{}

func NewAuthClient() gateway.AuthClient {
	return &AuthClient{}
}

func (a *AuthClient) Auth(uid int64, token string) (bool, error) {
	return true, nil
}
