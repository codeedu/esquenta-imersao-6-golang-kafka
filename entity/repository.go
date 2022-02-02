package entity

type CourseRepository interface {
	Insert(course Course) error
}
