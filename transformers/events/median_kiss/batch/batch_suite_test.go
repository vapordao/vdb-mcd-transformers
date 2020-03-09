package batch_test

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestMedianKissBatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MedianKissBatch Suite")
}

var _ = BeforeSuite(func() {
	logrus.SetOutput(ioutil.Discard)
})