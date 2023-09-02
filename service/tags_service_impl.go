package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (*TagsServiceImpl) Delete(tagsId int) {
	panic("unimplemented")
}

// FindAll implements TagsService.
func (*TagsServiceImpl) FindAll() []response.TagsResponse {
	panic("unimplemented")
}

// FindById implements TagsService.
func (*TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	panic("unimplemented")
}

// Update implements TagsService.
func (*TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	panic("unimplemented")
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}
