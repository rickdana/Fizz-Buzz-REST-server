package Config

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	} `yaml:"server"`
	DataSource struct {
		FileStore string `yaml:"fileStore"`
	} `yaml:"dataSource"`
	Logger struct {
		FilePath string `yaml:"filePath"`
	} `yaml:"logger"`
	Auth struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"auth"`
}
