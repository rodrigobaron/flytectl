package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterExamplesFunc(t *testing.T) {
	setup()
	registerFilesSetup()
	args = []string{""}
	err := registerExamplesFunc(ctx, args, cmdCtx)
	assert.NotNil(t, err)
}
func TestRegisterExamplesFuncErr(t *testing.T) {
	setup()
	registerFilesSetup()
	flytesnacks = "testingsnacks"
	args = []string{""}

	err := registerExamplesFunc(ctx, args, cmdCtx)
	// TODO (Yuvraj) make test to success after fixing flytesnacks bug
	assert.NotNil(t, err)
	flytesnacks = "flytesnacks"
}
