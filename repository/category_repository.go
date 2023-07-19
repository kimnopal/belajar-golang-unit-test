package repository

import "belajar-golang-unit-test/entity"

type CategoryRespository interface {
	FindById(id string) *entity.Category
}
