package rbac

import "strings"

const permissionDelimiter = ":"
const wildcard = "*"

// AccessList is a slice of Accesses and is generally used to represent a principal's
// full set of permissions for an application
type AccessList []Access

// Access represents a permission and an optional resource definition
type Access struct {
	ResourceDefinitions []ResourceDefinition `json:"resourceDefinitions,omitempty"`
	Permission          string               `json:"permission"`
}

// ResourceDefinition limits an Access to specific resources
type ResourceDefinition struct {
	Filter ResourceDefinitionFilter `json:"attributeFilter"`
}

// ResourceDefinitionFilter represents the key/values used for filtering
type ResourceDefinitionFilter struct {
	Key       string `json:"key"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}

// IsAllowed returns whether an action against a resource is allowed by an AccessList
// taking wildcards into consideration
// TODO: Take resource definitions into account
func (l AccessList) IsAllowed(app, res, verb string) bool {
	for _, a := range l {
		if a.Application() == app && matchWildcard(a.Resource(), res) && matchWildcard(a.Verb(), verb) {
			return true
		}
	}
	return false
}

// Application returns the name of the application in the permission
func (a Access) Application() string {
	return permIndex(a.Permission, 0)
}

// Resource returns the name of the resource in the permission
func (a Access) Resource() string {
	return permIndex(a.Permission, 1)
}

// Verb returns the verb in the permission
func (a Access) Verb() string {
	return permIndex(a.Permission, 2)
}

func permIndex(p string, i int) string {
	s := strings.Split(p, permissionDelimiter)
	if len(s) == 3 {
		return s[i]
	}
	return ""
}

func matchWildcard(s1, s2 string) bool {
	return s1 == s2 || s1 == wildcard
}
