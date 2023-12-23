package validate_file_creator

import (
	"log"
	"project/contract/validate_file_api"
	"project/dto"
)

type FileNameValid struct {
	next validate_file_api.DepartmentFileCreator
}

func (name FileNameValid) Execute(dto *dto.TemplateGenerator) {
	if len(dto.GetFileName()) == 0 {
		log.Print("File name doesnt pass")
		dto.SetCanCreated(false)

		return
	}

	dto.SetCanCreated(true)
	name.next.Execute(dto)
}

func (name *FileNameValid) SetNext(next validate_file_api.DepartmentFileCreator) {
	name.next = next
}
