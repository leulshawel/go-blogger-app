package models

// user object
type User struct {
	Name      string `json:"name"`
	Sub_field string `json:"subfield"`
	Age       uint8  `json:"age"`
	Field     string `json:"field"`
}

// the content of a blog post
type Data struct {
	Text  string `json:"text"`
	Image []byte `json:"image"`
	Video []byte `json:"video"`
}

// a post
type Post struct {
	Content_type string `json:"contenttype`
	User         *User  `json:"user`
	Data         *Data  `json:"data"`
}
