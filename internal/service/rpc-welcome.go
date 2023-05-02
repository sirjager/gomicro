package service

import (
	"context"
	"fmt"

	rpc "github.com/sirjager/rpcs/gomicro/go"
)

// welcomeMessage returns a formatted welcome message string.
func welcomeMessage(name string) string {
	return fmt.Sprintf("Welcome to %s Api", name)
}

// Welcome is a service method that returns a welcome message.
func (s *CoreService) Welcome(ctx context.Context, req *rpc.WelcomeRequest) (*rpc.WelcomeResponse, error) {
	return &rpc.WelcomeResponse{Message: welcomeMessage(s.Config.ServiceName)}, nil
}
