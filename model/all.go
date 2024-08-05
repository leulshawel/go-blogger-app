package model

// user object
type User struct {
	name      string `json:"name"`
	sub_field string `json:"subfield"`
	age       uint8  `json:"age"`
	field     string `json:"field"`
}

// create new user
func newUser(name string, sub_field string, age uint8, field string) *User {
	return &User{name: name, sub_field: sub_field, age: age, field: field}
}

// the content of a blog post
type Data struct {
	text  string `json:"text"`
	image []byte `json:"image"`
	video []byte `json:"video"`
}

// create new data
func newData(text string, image []byte, video []byte) *Data {
	return &Data{text: text, image: image, video: video}
}

// a post
type Post struct {
	content_type byte  `json:"contenttype`
	user         *User `json:"user`
	data         *Data `json:"data"`
}

// create new post
func newPost(content_type byte, user *User, data *Data) *Post {
	return &Post{content_type: content_type, user: user, data: data}
}
