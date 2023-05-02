package service

import (
	"github.com/rs/zerolog"
	"github.com/sirjager/gomicro/config"
	"github.com/sirjager/gomicro/internal/db/sqlc"
	"github.com/sirjager/gomicro/pkg/email"

	rpc "github.com/sirjager/rpcs/gomicro/go"
)

// CoreService represents the core service of your application.
type CoreService struct {
	rpc.UnimplementedGoMicroServer
	Logr   zerolog.Logger   // Logger is used for logging.
	Config config.Config    // Config holds the configuration settings.
	mailer email.MailSender // Mailer is used for sending emails.
	store  sqlc.Store
}

// NewService creates a new instance of the CoreService.
func NewService(logr zerolog.Logger, config config.Config, mailer email.MailSender, store sqlc.Store) (*CoreService, error) {

	return &CoreService{
		Logr:   logr,
		Config: config,
		mailer: mailer,
		store:  store,
	}, nil
}
