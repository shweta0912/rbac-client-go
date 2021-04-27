package rbac

type PaginatedBody struct {
	Meta  PaginationMeta  `json:"meta"`
	Links PaginationLinks `json:"links"`
	Data  interface{}     `json:"data"`
}

type PaginationMeta struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PaginationLinks struct {
	First    string `json:"first"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Last     string `json:"last"`
}
