package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccess(t *testing.T) {
	tests := map[string]struct {
		input string
		app   string
		res   string
		verb  string
	}{
		"simple":            {input: "a:b:c", app: "a", res: "b", verb: "c"},
		"empty":             {input: "", app: "", res: "", verb: ""},
		"bad sep":           {input: "a/b/c", app: "", res: "", verb: ""},
		"too few segments":  {input: "a:b", app: "", res: "", verb: ""},
		"too many segments": {input: "a:b:c:d", app: "", res: "", verb: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			acc := Access{Permission: tc.input}
			assert.Equal(t, tc.app, acc.Application())
			assert.Equal(t, tc.res, acc.Resource())
			assert.Equal(t, tc.verb, acc.Verb())
		})
	}
}

func TestIsAllowed(t *testing.T) {
	tests := map[string]struct {
		input   AccessList
		app     string
		res     string
		verb    string
		allowed bool
	}{
		"single": {input: AccessList{
			Access{Permission: "a:b:c"},
		}, app: "a", res: "b", verb: "c", allowed: true},
		"multiple": {input: AccessList{
			Access{Permission: "a:b:c"},
			Access{Permission: "d:e:f"},
		}, app: "d", res: "e", verb: "f", allowed: true},
		"wildcard resource": {input: AccessList{
			Access{Permission: "a:*:c"},
		}, app: "a", res: "b", verb: "c", allowed: true},
		"wildcard verb": {input: AccessList{
			Access{Permission: "a:b:*"},
		}, app: "a", res: "b", verb: "c", allowed: true},
		"wildcard both": {input: AccessList{
			Access{Permission: "a:*:*"},
		}, app: "a", res: "b", verb: "c", allowed: true},
		"empty": {input: AccessList{}, app: "a", res: "b", verb: "c", allowed: false},
		"not allowed": {input: AccessList{
			Access{Permission: "a:b:c"},
		}, app: "d", res: "e", verb: "f", allowed: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.IsAllowed(tc.app, tc.res, tc.verb)
			assert.Equal(t, tc.allowed, got)
		})
	}
}
