package cli

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

func (suite *CliTestSuite) TestCliReturnsAnExitCodeWhenNoFilePath() {
	c := NewCli()
	exit, err := c.Run()

	assert.Equalf(suite.T(), 2, exit, "cli Run() should return 2 exit code: got %v", exit)
	assert.Equalf(suite.T(), err, ErrNoFilePath, "cli Run() should return error, expected: %v, got: %v", ErrNoFilePath, err)
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, &CliTestSuite{})
}
