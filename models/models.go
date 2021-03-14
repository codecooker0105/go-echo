package models

type Model interface{}

var Models = []Model{
	User{},
	Blog{},
	Comment{},
}
