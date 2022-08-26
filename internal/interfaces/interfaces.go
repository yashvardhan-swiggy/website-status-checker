package interfaces

type PostRequestURLs struct {
	URLs []string `json:"websites"`
}

var WebsitesMap = make(map[string]interface{})
