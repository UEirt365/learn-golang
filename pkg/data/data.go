package data

import "github.com/UEirt365/learn-golang/pkg/dto"

var Todos []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Content: "Maths"},
		{ID: 2, Content: "Literature"},
		{ID: 3, Content: "Physics"},
		{ID: 4, Content: "Chemistry"},
		{ID: 5, Content: "Foreign language"},
	}
}
