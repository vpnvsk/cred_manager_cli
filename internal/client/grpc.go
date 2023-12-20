package client

import (
	"context"
	"errors"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/pterm/pterm"
	authv1 "github.com/vpnvsk/protos_go/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"p_s_cli/internal/models"
	"time"
)

type Client struct {
	api   authv1.AuthClient
	appId int32
}

func NewClient(
	ctx context.Context,
	addr string,
	timeout time.Duration,
	retriesCount int,
	appId int32,
) *Client {
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		return nil
	}
	return &Client{
		api:   authv1.NewAuthClient(cc),
		appId: appId,
	}
}

func getAndSaveToken(resp *authv1.LoginResponse) error {
	responseBody := resp.Token

	config := models.Token{
		JwtToken: responseBody,
	}
	filePath := "token.json"
	go config.WriteToken(filePath)
	return nil
}
func (c *Client) LogIn(ctx context.Context, user models.User) error {
	resp, err := c.api.Login(ctx, &authv1.LoginRequest{Login: user.UserName, Password: user.Password, AppId: c.appId})
	if err != nil {
		return errors.New("bad credentials")
	}
	if err = getAndSaveToken(resp); err != nil {
		return err
	}
	pterm.Success.Printfln("LogIn successfully")

	return err
}
func (c *Client) SignUp(ctx context.Context, user models.User) error {
	_, err := c.api.Register(ctx, &authv1.RegisterRequest{Login: user.UserName, Password: user.Password})
	if err != nil {
		return err
	}
	loginResult := make(chan error)

	// Start a goroutine to perform the login.
	go func() {
		loginErr := c.LogIn(ctx, user)
		if loginErr != nil {
			loginResult <- loginErr
		} else {
			loginResult <- nil // Login succeeded
		}
	}()

	// Wait for the login result from the goroutine.
	loginErr := <-loginResult

	return loginErr

}
