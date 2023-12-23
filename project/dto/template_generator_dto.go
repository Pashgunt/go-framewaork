package dto

type TemplateGenerator struct {
	fileType   string
	fileName   string
	outputDir  string
	fileCount  int
	canCreated bool
}

func (template *TemplateGenerator) GetCanCreated() bool {
	return template.canCreated
}

func (template *TemplateGenerator) SetCanCreated(canCreated bool) {
	template.canCreated = canCreated
}

func (template *TemplateGenerator) SetOutputDir(outputDir string) {
	template.outputDir = outputDir
}

func (template TemplateGenerator) GetOutputDir() string {
	return template.outputDir
}

func (template TemplateGenerator) GetFileCount() int {
	return template.fileCount
}

func (template TemplateGenerator) GetFileType() string {
	return template.fileType
}

func (template TemplateGenerator) GetFileName() string {
	return template.fileName
}

func NewTemplateGenerator(fileType string, fileName string, outputDir string, fileCount int) *TemplateGenerator {
	return &TemplateGenerator{fileType: fileType, fileName: fileName, outputDir: outputDir, fileCount: fileCount}
}
