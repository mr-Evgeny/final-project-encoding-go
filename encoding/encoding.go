package encoding

import (
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"fmt"
    "os"
    "gopkg.in/yaml.v3"
	"encoding/json"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonFile, err := os.ReadFile(j.FileInput)
    if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %s", err.Error())
    }

	// создаём yaml-файл
    yamlFile, err := os.Create(j.FileOutput)
    if err != nil {
        return fmt.Errorf("ошибка при создании файла: %s", err.Error())
    }
    defer yamlFile.Close()

    err = json.Unmarshal(jsonFile, &j.DockerCompose)
    if err != nil {
        return fmt.Errorf("ошибка при десериализации: %s", err.Error())
    }

	yamlData, err := yaml.Marshal(&j.DockerCompose)
    if err != nil {
        return fmt.Errorf("ошибка при сериализации в yaml: %s", err.Error())
    }

	// записываем слайс байт в файл
    _, err = yamlFile.Write(yamlData)
    if err != nil {
        return fmt.Errorf("ошибка при записи данных в файл: %s", err.Error())
    }

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
    if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %s", err.Error())
    }

	// создаём json-файл
    jsonFile, err := os.Create(y.FileOutput)
    if err != nil {
        return fmt.Errorf("ошибка при создании файла: %s", err.Error())
    }
    defer jsonFile.Close()

    err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
    if err != nil {
        return fmt.Errorf("ошибка при десериализации: %s", err.Error())
    }

	jsonData, err := json.Marshal(&y.DockerCompose)
    if err != nil {
        return fmt.Errorf("ошибка при сериализации в json: %s", err.Error())
    }

	// записываем слайс байт в файл
    _, err = jsonFile.Write(jsonData)
    if err != nil {
        return fmt.Errorf("ошибка при записи данных в файл: %s", err.Error())
    }

	return nil
}
