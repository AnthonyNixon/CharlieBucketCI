package types

type Repo struct {
	Id           int    `json:"id"`
	Org          string `json:"org"`
	Name         string `json:"repo"`
	Loadbalancer string `json:"loadbalancer"`
	Domain       string `json:"domain"`
}
