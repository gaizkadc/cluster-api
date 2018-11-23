/*
 * Copyright (C)  2018 Nalej - All Rights Reserved
 */

package ithelpers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/authx/pkg/token"
	"github.com/nalej/grpc-authx-go"
	"github.com/onsi/gomega"
	"google.golang.org/grpc/metadata"
	"time"
)

const AuthHeader = "authorization"

func GenerateToken(email string, organizationID string, roleName string, secret string, primitives []grpc_authx_go.AccessPrimitive) string {
	p := make([]string, 0)
	for _, prim := range primitives {
		p = append(p, prim.String())
	}

	pClaim := token.PersonalClaim{
		UserID:         email,
		Primitives:     p,
		RoleName:       roleName,
		OrganizationID: organizationID,
	}

	claim := token.NewClaim(pClaim, "it", time.Now(), time.Minute)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := t.SignedString([]byte(secret))
	gomega.Expect(err).To(gomega.Succeed())

	return tokenString
}

func GetContext(token string) (context.Context, context.CancelFunc) {
	md := metadata.New(map[string]string{AuthHeader: token})
	baseContext, cancel := context.WithTimeout(context.Background(), time.Minute)
	return metadata.NewOutgoingContext(baseContext, md), cancel
}

func GetAuthConfig(endpoints ...string) *interceptor.AuthorizationConfig {
	permissions := make(map[string]interceptor.Permission, 0)
	for _, e := range endpoints {
		permissions[e] = interceptor.Permission{
			Must: []string{grpc_authx_go.AccessPrimitive_ORG.String()},
		}
	}
	return &interceptor.AuthorizationConfig{
		AllowsAll:   false,
		Permissions: permissions,
	}
}
