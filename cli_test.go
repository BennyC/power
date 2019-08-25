package power

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CliTestSuite struct {
	suite.Suite
}

func (suite *CliTestSuite) TestNewCliGeneratesCli() {
	c := NewCli()
	assert.NotNilf(suite.T(), c, "NewCli should return pointer to Cli struct: got %v", c)
}

func (suite *CliTestSuite) TestCliReturnsAZeroCode() {
	c := NewCli()
	exit := c.Run()

	assert.Equalf(suite.T(), 0, exit, "Cli run should return 0 exit code: got %v", exit)
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, &CliTestSuite{})
}
