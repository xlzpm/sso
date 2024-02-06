package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xlzpm/sso/internal/domain/models"
	myjwt "github.com/xlzpm/sso/internal/lib/jwt"
)

func TestCreateNewToken(t *testing.T) {
	str, err := myjwt.NewToken(models.User{
		ID:       1,
		Email:    "grpc",
		PassHash: []byte("AYE"),
	}, models.App{
		ID:     1,
		Name:   "grpc",
		Secret: "AYE",
	},
		time.Second)
	assert.NotEqual(t,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjEsImVtYWlsIjoiY3VycyIsImV4cCI6MTcwNTU5MzEwOSwidWlkIjoxfQ.iKZkg1HTmbOoXJTC_ZKWHAVPb45wXsFCRWPJWY5orTA",
		str)
	require.NoError(t, err)
}
