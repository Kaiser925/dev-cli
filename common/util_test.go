package common

import (
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TemplateStruct struct {
	A string
	B string
}

func TestRenderTemplateToExistFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := &TemplateStruct{
		A: "A",
		B: "B",
	}

	err = RenderTemplateFile(`{{.A}} is {{.B}}`, content, tmpFile.Name())
	assert.Equal(t, err, nil)

	fileContent, err := ioutil.ReadAll(tmpFile)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(fileContent), "A is B")
}

func TestRenderTemplateToNewFile(t *testing.T) {
	fileName := "new-file"
	defer os.Remove(fileName)

	content := &TemplateStruct{
		A: "A",
		B: "B",
	}
	_, err := os.Stat(fileName)
	assert.Equal(t, os.IsNotExist(err), true)

	err = RenderTemplateFile(`{{.A}} is {{.B}}`, content, fileName)
	assert.Equal(t, err, nil)

	file, err := os.Open(fileName)
	assert.Equal(t, err, nil)

	fileContent, err := ioutil.ReadAll(file)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(fileContent), "A is B")
}
