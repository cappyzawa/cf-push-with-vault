package plug_test

import (
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cappyzawa/cf-push-with-vault/plug"
	"github.com/cappyzawa/cf-push-with-vault/vault/vaultfakes"
)

var _ = Describe("Command", func() {
	var (
		fakeCliConnection *pluginfakes.FakeCliConnection
		fakeVariables     *vaultfakes.FakeVariables
		command           *Command
	)

	BeforeEach(func() {
		fakeCliConnection = new(pluginfakes.FakeCliConnection)
		fakeVariables = new(vaultfakes.FakeVariables)
		command = &Command{
			CliConnection: fakeCliConnection,
			Variables:     fakeVariables,
		}
	})

	Describe("Push()", func() {
		Context("manifest file is missing", func() {
			It("an error is occurred", func() {
				err := command.Push("../testdata/missing.yml")
				Expect(err).To(HaveOccurred())
			})
		})
		Context("manifest file does not contain parameters", func() {
			It("access to the vault does not occur", func() {
				err := command.Push("../testdata/no_contains_params.yml")
				Expect(fakeVariables.GetCallCount()).To(BeZero())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("manifest file contains parameters", func() {
			JustBeforeEach(func() {
				fakeVariables.GetReturns(nil, true, nil)
				fakeVariables.GetReturns(nil, true, nil)
			})
			It("access to the vault occurs multiple times", func() {
				err := command.Push("../testdata/multi_params.yml")
				Expect(fakeVariables.GetCallCount()).To(Equal(2))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
