package dotfileignorematcher

import (
	"testing"

	"github.com/deadcheat/goblet/generator"
)

func TestPrepareSuccess(t *testing.T) {
	e := generator.OptionFlagEntity{
		IgnoreDotFiles: false,
	}

	testee := New()

	err := testee.Prepare(e)
	if err != nil {
		t.Error("Prepare() should not return any errors")
	}
}
