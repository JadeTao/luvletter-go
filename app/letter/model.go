package letter

// Letter ...
type Letter struct {
	ID         int64
	Account    string
	Nickname   string
	CreateTime string `json:"createTime"`
	Content    string
	Mood       string
	Tag        []string
}
