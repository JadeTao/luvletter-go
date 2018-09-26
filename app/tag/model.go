package tag

// Tag ...
type Tag struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Account      string `json:"account"`
	Count        int64  `json:"count"`
	CreateTime   string `json:"createTime"`
	LastUsedTime string `json:"lastUsedTime"`
}
