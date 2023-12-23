package validate_file_creator

import (
	"log"
	"project/contract/validate_file_api"
	"project/dto"
	"project/enum"
)

type FileTypeValida struct {
	next validate_file_api.DepartmentFileCreator
}

func (typeFile FileTypeValida) Execute(dto *dto.TemplateGenerator) {
	folder, err := enum.FolderNames[dto.GetFileType()]
	if !err || len(folder) == 0 {
		log.Print("Incorrect type folder")
		dto.SetCanCreated(false)

		return
	}

	dto.SetOutputDir(folder)
}

func (typeFile *FileTypeValida) SetNext(next validate_file_api.DepartmentFileCreator) {
	typeFile.next = next
}
