/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package conductor

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestConductorPackage(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Conductor package suite")
}
