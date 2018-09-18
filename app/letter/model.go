package letter

// Letter ...
type Letter struct {
	ID         int
	Account    string
	Nickname   string
	CreateTime int `json:"createTime"`
	Content    string
	Mood       string
	Tag        []string
}
