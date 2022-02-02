package usecase

import (
	"github.com/codeedu/esquenta-imersao-go-course/entity"
	"github.com/google/uuid"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {

	course := entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)
	if err != nil {
		return CreateCourseOutputDto{}, err
	}

	output := CreateCourseOutputDto{}
	output.ID = course.ID
	output.Name = course.Name
	output.Description = course.Description
	output.Status = course.Status

	return output, nil

}
