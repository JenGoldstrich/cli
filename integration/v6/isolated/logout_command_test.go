package isolated

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/util/configv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("logout command", func() {
	Context("help", func() {
		It("displays help", func() {
			session := helpers.CF("logout", "--help")
			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Say("   logout - Log user out"))
			Eventually(session).Should(Say("USAGE:"))
			Eventually(session).Should(Say("   cf logout"))
			Eventually(session).Should(Say("ALIAS:"))
			Eventually(session).Should(Say("   lo"))
			Eventually(session).Should(Exit(0))
		})
	})

	When("there's user information set in the config", func() {
		BeforeEach(func() {
			helpers.SetupCF(ReadOnlyOrg, ReadOnlySpace)
		})

		It("clears out user information in the config", func() {
			username, _ := helpers.GetCredentials()
			session := helpers.CF("logout")

			Eventually(session).Should(Say(`Logging out %s\.\.\.`, username))
			Eventually(session).Should(Say("OK"))
			Eventually(session).Should(Exit(0))

			config, err := configv3.LoadConfig()
			Expect(err).ToNot(HaveOccurred())

			Expect(config.ConfigFile.AccessToken).To(BeEmpty())
			Expect(config.ConfigFile.RefreshToken).To(BeEmpty())
			Expect(config.ConfigFile.TargetedOrganization.GUID).To(BeEmpty())
			Expect(config.ConfigFile.TargetedOrganization.Name).To(BeEmpty())
			Expect(config.ConfigFile.TargetedSpace.AllowSSH).To(BeFalse())
			Expect(config.ConfigFile.TargetedSpace.GUID).To(BeEmpty())
			Expect(config.ConfigFile.TargetedSpace.Name).To(BeEmpty())
			Expect(config.ConfigFile.UAAGrantType).To(BeEmpty())
			Expect(config.ConfigFile.UAAOAuthClient).To(Equal("cf"))
			Expect(config.ConfigFile.UAAOAuthClientSecret).To(BeEmpty())
		})

		When("a valid refresh token is present", func() {

			var introspectRequest *http.Request

			BeforeEach(func() {
				helpers.SkipIfClientCredentialsTestMode()

				config, err := configv3.LoadConfig()
				Expect(err).ToNot(HaveOccurred())
				jwt := helpers.ParseTokenString(config.ConfigFile.RefreshToken)
				tokenID, ok := jwt.Claims().JWTID()
				Expect(ok).To(BeTrue())
				uaaURL := path.Join(config.ConfigFile.UAAEndpoint, "introspect")
				body := strings.NewReader(fmt.Sprintf("token=%s", tokenID))

				// TODO: This requires the uaa.resource authority
				introspectRequest, err = http.NewRequest(http.MethodGet, uaaURL, body)

			})

			It("invalidates the Refresh token", func() {

				
			})
		})
	})
})
