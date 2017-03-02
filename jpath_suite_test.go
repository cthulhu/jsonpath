package jpath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJpath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jpath Suite")
}
