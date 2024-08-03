package ds

//user object
type User struct{
	name string
	sub_field string
	age uint8
	field string
}

//create new user
func newUser (name string, sub_field string, age uint8, field string) *User { 
	return &User{name: name, sub_field: sub_field, age: age, field: field} 
}

//the content of a blog post
type Data struct{
	text string
	image []byte 
	video []byte 
}

//create new data
func newData (text string, image []byte, video []byte) *Data { 
	return &Data{text: text, image: image, video: video} out
}

//a post 
type Post struct {
	content_type byte
	user *User
	data *Data
}

//create new post
func newPost (content_type byte, user *User, data *Data) *Post { 
	return &Post{ content_type: content_type, user: user, data: data } 
}




