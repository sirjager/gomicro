package tokens

import (
	"os"
	"testing"

	"github.com/sirjager/gomicro/pkg/utils"
)

var small_secret_key string
var valid_secret_key string

func TestMain(m *testing.M) {
	small_secret_key = utils.RandomString(30)
	valid_secret_key = utils.RandomString(32)
	os.Exit(m.Run())
}
