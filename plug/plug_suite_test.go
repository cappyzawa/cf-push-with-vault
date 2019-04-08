package plug_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plug Suite")
}
