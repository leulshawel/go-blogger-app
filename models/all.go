package models

// user object
type User struct {
	Name      string `json:"name" bson:"name" validator:"required"`
	Sub_field string `json:"subfield" bson:"name" validator:"required"`
	Age       uint8  `json:"age" bson:"name" validator:"required"`
	Field     string `json:"field" bson:"name" validator:"required"`
}

// the content of a blog post
type Data struct {
	Text  string `json:"text" bson:"text"`
	Image []byte `json:"image" bson:"image"`
	Video []byte `json:"video" bson:"video"`
}

// a post
type Post struct {
	Content_type string `json:"contenttype" bson:"name" validator:"required"`
	User         *User  `json:"user" bson:"name" validator:"required"`
	Data         *Data  `json:"data" bson:"name" validator:"required"`
}
