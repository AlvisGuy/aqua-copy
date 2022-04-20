package constraint_test

import (
	"testing"

	constraint "github.com/aquaproj/aqua/pkg/version-constraint"
)

func TestVersionConstraints_Check(t *testing.T) {
	t.Parallel()
	data := []struct {
		title       string
		constraints string
		version     string
		exp         bool
		isErr       bool
	}{
		{
			title:       "true",
			constraints: `semver(">= 0.4.0")`,
			version:     "v0.4.0",
			exp:         true,
		},
		{
			title:       "false",
			constraints: `semver(">= 0.4.0")`,
			version:     "v0.3.0",
			exp:         false,
		},
		{
			title:       "invalid expression",
			constraints: `>= 0.4.0`,
			version:     "v0.3.0",
			isErr:       true,
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			constraints := constraint.NewVersionConstraints(d.constraints)
			b, err := constraints.Check(d.version)
			if d.isErr {
				if err == nil {
					t.Fatal("err should be returned")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if b != d.exp {
				t.Fatalf("wanted %v, got %v", d.exp, b)
			}
		})
	}
}