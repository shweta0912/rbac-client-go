package rbac

const identityHeader = `X-RH-Identity`

// PaginatedBody represents the response body format from the RBAC service
type PaginatedBody struct {
	Meta  PaginationMeta  `json:"meta"`
	Links PaginationLinks `json:"links"`
	Data  interface{}     `json:"data"`
}

// PaginationMeta contains metadata for pagination
type PaginationMeta struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// PaginationLinks provides links to additional pages of response data
type PaginationLinks struct {
	First    string `json:"first"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Last     string `json:"last"`
}
