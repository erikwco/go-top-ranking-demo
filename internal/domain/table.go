package domain

type Table struct {
	Rank        int16  `csv:"rank" json:"rank"`
	Item        string `csv:"item" json:"item"`
	RepoName    string `csv:"repo_name" json:"repo_name"`
	Stars       int32  `csv:"stars" json:"stars"`
	Forks       int32  `csv:"forks" json:"forks"`
	Language    string `csv:"language" json:"language"`
	RepoUrl     string `csv:"repo_url" json:"repo_url"`
	Username    string `csv:"username" json:"username"`
	Issues      int32  `csv:"issues" json:"issues"`
	LastCommit  string `csv:"last_commit" json:"last_commit"`
	Description string `csv:"description" json:"description"`
}
