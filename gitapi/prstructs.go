package gitapi

type PRSeachResult struct {
	TotalCount	int
	Items		[]PRStruct
}

type PRStruct struct {
	Id 		int
	HTMLUrl string `json:html_url`
	State 	string
	Title 	string
	User	*User
}

type User struct{
	Login	string
	HTMLUrl	string `json:html_url`
}