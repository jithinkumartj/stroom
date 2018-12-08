package auth

import(
	"fmt"
)

type AuthenticateClient struct {}

func (AuthenticateClient) Authenticate(token string) error{
	fmt.Println(token)
	return nil
}