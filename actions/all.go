package actions

import "github.com/leulshawel/go-blogger-app/models"

type actionsCollection struct {
	OnDelete func(post *models.Post)
	OnPost   func(Post *models.Post)
	OnEdit   func(Post *models.Post)
	OnRead   func(Post *models.Post)
}

var placeHolderFunc = func(post *models.Post) { return }

var Actions_ = new(actionsCollection)

func OnPost(f func(post *models.Post)) {
	Actions_.OnPost = f
}

func OnDelete(f func(post *models.Post)) {
	Actions_.OnDelete = f
}

func OnEdit(f func(post *models.Post)) {
	Actions_.OnEdit = f
}

func OnRead(f func(post *models.Post)) {
	Actions_.OnRead = f

}
