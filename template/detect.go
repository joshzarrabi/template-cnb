package {{ .buildpack }}

import (
	"github.com/cloudfoundry/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{}, nil
	}
}
