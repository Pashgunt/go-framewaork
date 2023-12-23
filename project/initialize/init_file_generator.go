package initialize

import (
	"flag"
	"fmt"
	"os"
	"project/chain_responsibility/validate_file_creator"
	"project/dto"
	"sync"
)

type file struct {
}

func (file file) Init() {
	fileType := flag.String("type", "", "")
	fileName := flag.String("name", "", "")
	flag.Parse()
	dtoTemplateGenerator := dto.NewTemplateGenerator(
		*fileType,
		*fileName,
		"",
		1,
	)
	fileNameValidator := validate_file_creator.FileNameValid{}
	fileNameValidator.SetNext(&validate_file_creator.FileTypeValida{})
	fileNameValidator.Execute(dtoTemplateGenerator)
	if !dtoTemplateGenerator.GetCanCreated() {
		return
	}
	err := file.generateFiles(*dtoTemplateGenerator)
	if err != nil {
		return
	}
}

func (file *file) generateFiles(dto dto.TemplateGenerator) error {
	var wg sync.WaitGroup
	wg.Add(dto.GetFileCount())

	semaphore := make(chan struct{}, 1)

	for i := 0; i < dto.GetFileCount(); i++ {
		go func(index int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			filename := fmt.Sprintf("%s/%s.go", dto.GetOutputDir(), dto.GetFileName())

			err := os.WriteFile(filename, []byte(""), 0644)
			if err != nil {
				fmt.Printf("Failed to write file %s: %v\n", filename, err)
			} else {
				fmt.Printf("File %s generated successfully\n", filename)
			}
		}(i)
	}

	wg.Wait()

	return nil
}
