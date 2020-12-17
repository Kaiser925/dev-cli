/*
 * Developed by Kaiser925 on 2020/12/17.
 * Lasted modified 2020/11/10.
 * Copyright (c) 2020.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

func TestWriteNewFile(t *testing.T) {
	fileName := "newFile"
	defer os.Remove(fileName)
	content := "test content"
	l, err := WriteFile(fileName, content)

	assert.Equal(t, err, nil)
	assert.Equal(t, l, len(content))
}
