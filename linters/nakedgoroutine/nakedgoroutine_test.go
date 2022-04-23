package nakedroutine

import (
	"github.com/stretchr/testify/suite"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

type linterSuite struct {
	suite.Suite
}

func (suite *linterSuite) TestContextLinter() {
	analysistest.Run(
		suite.T(), "nakedgoroutine",
		NakedRoutineCodeAnalyzer, "testlintdata/nakedroutine")
}

func TestLinterSuite(t *testing.T) {
	suite.Run(t, new(linterSuite))
}
