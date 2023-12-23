package validate_file_api

import "project/dto"

type DepartmentFileCreator interface {
	Execute(dto *dto.TemplateGenerator)
	SetNext(next DepartmentFileCreator)
}
