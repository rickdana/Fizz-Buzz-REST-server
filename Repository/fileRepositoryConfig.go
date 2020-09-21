package Repository

type FileRepositoryConfig struct {
	filePath string
}

func NewFileRepositoryConfig(filePath string) *FileRepositoryConfig {
	return &FileRepositoryConfig{filePath: filePath}
}

func (frc *FileRepositoryConfig) GetConfig() string {
	return frc.filePath
}
