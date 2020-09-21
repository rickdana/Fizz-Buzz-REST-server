package Repository

import (
	"encoding/json"
	"github.com/rickdana/fizzbuzzApi/Model"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

type testFileRepoConfig struct {
	filePath string
}

func (trc *testFileRepoConfig) GetConfig() string {
	return trc.filePath
}

var requests1 = Model.FizzBuzz{
	FirstMultiple:  3,
	SecondMultiple: 5,
	Limit:          90,
	FizzWord:       "Fizz",
	BuzzWord:       "Buzz",
}

var requests2 = Model.FizzBuzz{
	FirstMultiple:  7,
	SecondMultiple: 9,
	Limit:          100,
	FizzWord:       "Hello",
	BuzzWord:       "World",
}

func TestFileRepository_FindAll(t *testing.T) {
	file := setup()

	tests := []struct {
		name    string
		config  Config
		want    interface{}
		wantErr bool
	}{
		{"FindAll should read all data from file", &testFileRepoConfig{filePath: file}, []Model.FizzBuzz{requests1, requests2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FileRepository{
				config: tt.config,
			}
			got, err := fr.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileRepository_Save(t *testing.T) {
	type fields struct {
	}
	type args struct {
		entity interface{}
	}
	tests := []struct {
		name    string
		config  Config
		args    args
		wantErr bool
	}{
		{"Save should write data to file", &testFileRepoConfig{filePath: createTempFile(os.TempDir())}, args{entity: "My data"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FileRepository{
				config: tt.config,
			}
			if err := fr.Save(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			b, err := ioutil.ReadFile(tt.config.GetConfig())

			if err != nil {
				t.Fatal(err)
			}

			var str string
			if err := json.Unmarshal(b, &str); err != nil {
				t.Fatal(err)
			}

			str = strings.TrimSuffix(str, "\n")
			if str != tt.args.entity {
				t.Errorf("Save() written data: %v, want written data: %v", str, tt.args.entity)

			}
			os.Remove(tt.config.GetConfig())
		})
	}
}

func Test_createFile(t *testing.T) {
	filePath := createTempFile(os.TempDir())
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Create file should create new file at the expect path", args{filePath: filePath}, filePath, false},
		{"Create file should fail when path does not exist", args{filePath: "dir/fake_file_path"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("createFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createFile() got = %v, want %v", got, tt.want)
			}
		})
		os.Remove(tt.args.filePath)
	}
}

func createTempFile(dir string) (filePath string) {
	file, err := ioutil.TempFile(dir, "file-")
	if err != nil {
		log.Fatal(err)
	}
	return file.Name()
}

func setup() string {
	file := createTempFile(os.TempDir())
	repository := FileRepository{config: &testFileRepoConfig{filePath: file}}
	for _, fz := range []Model.FizzBuzz{requests1, requests2} {
		if err := repository.Save(fz); err != nil {
			panic(err)
		}
	}
	return file
}
