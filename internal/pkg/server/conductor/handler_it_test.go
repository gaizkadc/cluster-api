/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
RUN_INTEGRATION_TEST=true
*/

package conductor

import (
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/cluster-api/internal/pkg/server/ithelpers"
	"github.com/nalej/cluster-api/internal/pkg/utils"
	"github.com/nalej/grpc-utils/pkg/test"
	"github.com/onsi/ginkgo"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var _ = ginkgo.Describe("Roles", func() {

	if !utils.RunIntegrationTests() {
		log.Warn().Msg("Integration tests are skipped")
		return
	}

	/*
		var (
			systemModelAddress = os.Getenv("IT_SM_ADDRESS")
			userManagerAddress = os.Getenv("IT_USER_MANAGER_ADDRESS")
		)

		if systemModelAddress == "" || userManagerAddress == "" {
			ginkgo.Fail("missing environment variables")
		}
	*/

	// gRPC server
	var server *grpc.Server
	// grpc test listener
	var listener *bufconn.Listener
	// client
	/*
		var smConn *grpc.ClientConn
		var umConn *grpc.ClientConn
		var client grpc_public_api_go.RolesClient
	*/

	// Target organization.
	//var token string

	ginkgo.BeforeSuite(func() {
		listener = test.GetDefaultListener()
		authConfig := ithelpers.GetAuthConfig("/cluster_api.Roles/List")
		server = grpc.NewServer(interceptor.WithServerAuthxInterceptor(
			interceptor.NewConfig(authConfig, "secret", ithelpers.AuthHeader)))

		/*
			conn, err := test.GetConn(*listener)
			gomega.Expect(err).To(gomega.Succeed())

			manager := NewManager()
			handler := NewHandler(manager)
		*/
		//grpc_public_api_go.RegisterRolesServer(server, handler)
		test.LaunchServer(server, listener)

		/*
			client = grpc_public_api_go.NewRolesClient(conn)
			targetOrganization = ithelpers.CreateOrganization(fmt.Sprintf("testOrg-%d", ginkgo.GinkgoRandomSeed()), orgClient)
			targetRole = ithelpers.CreateRole(targetOrganization.OrganizationId, umClient)
			token = ithelpers.GenerateToken("email@nalej.com",
				targetOrganization.OrganizationId, "Owner", "secret",
				[]grpc_authx_go.AccessPrimitive{grpc_authx_go.AccessPrimitive_ORG})
		*/
	})

	ginkgo.AfterSuite(func() {
		server.Stop()
		listener.Close()
		/*
			smConn.Close()
			umConn.Close()
		*/
	})

	ginkgo.It("should be able to ...", func() {
	})

})
