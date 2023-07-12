package types

type Route struct {
	Url    string
	Method string
}

type Service struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Root string `json:"root"`
	Icon string `json:"icon"`
}
