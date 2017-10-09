package ginkgo_jenkins_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/fabric8-jenkins/ginkgo-jenkins/utils"
)

var _ = Describe("connection to Jenkins", func() {

	Context("initially", func() {

		It("has 0 jobs", func() {
			c, err := utils.GetJenkinsClient()
			Expect(err).NotTo(HaveOccurred())

			numOfJobs, err := c.GetJobs()
			Expect(err).NotTo(HaveOccurred())

			Expect(len(numOfJobs)).Should(BeZero())
		})
	})
})