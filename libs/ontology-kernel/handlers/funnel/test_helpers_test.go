package funnel

import (
	"github.com/google/uuid"

	authmw "github.com/Izaiaspertrelly/openfoundry/libs/auth-middleware"
)

// mustClaims fabricates a Claims struct with the requested subject
// and roles for table-driven access-control tests.
func mustClaims(sub uuid.UUID, roles []string) *authmw.Claims {
	return &authmw.Claims{Sub: sub, Roles: roles}
}
