package boshinit_test

import (
	"github.com/pivotal-cf-experimental/bosh-bootloader/boshinit"
	"github.com/pivotal-cf-experimental/bosh-bootloader/ssl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manifest", func() {
	Describe("DirectorSSLKeyPair", func() {
		It("returns the director ssl keypair from a built manifest", func() {
			manifest := boshinit.Manifest{
				Jobs: []boshinit.Job{{
					Properties: boshinit.JobProperties{
						Director: boshinit.DirectorJobProperties{
							SSL: boshinit.SSLProperties{
								Cert: "some-cert",
								Key:  "some-key",
							},
						},
					},
				}},
			}

			keyPair := manifest.DirectorSSLKeyPair()
			Expect(keyPair).To(Equal(ssl.KeyPair{
				Certificate: []byte("some-cert"),
				PrivateKey:  []byte("some-key"),
			}))
		})

		It("returns an empty keypair if there are no jobs", func() {
			manifest := boshinit.Manifest{}

			keyPair := manifest.DirectorSSLKeyPair()
			Expect(keyPair).To(Equal(ssl.KeyPair{}))
		})
	})
})
