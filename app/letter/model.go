package letter

// Letter ...
type Letter struct {
	ID         int64    `json:"id"`
	Account    string   `json:"account"`
	Nickname   string   `json:"nickname"`
	CreateTime string   `json:"createTime"`
	Content    string   `json:"content"`
	Mood       string   `json:"mood"`
	Tags        []string `json:"tags"`
}
