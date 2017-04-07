package main

import (
	"os"
	"text/template"
	"bufio"
	"fmt"
	"regexp"
	"errors"
)

type Method struct {
	ModuleName   string
	MethodName   string
	MethodParams []string
	Comment      string
	MethodBody   string
}

type TemplateReader struct {
	gotMeta bool
}

func main() {
	m := &Method{"TestModule", "prepareStatement", []string{"arg1"}, "Some Comment", "Body"}
	t, _ := template.ParseFiles("template_examples/method.txt")
	t.Execute(os.Stdout, m)
	fmt.Println("===============")
	reader := &TemplateReader{false}
	reader.readTemplates("template_examples/method.txt")
}


func (t *TemplateReader) readTemplates(path string) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read meta, then set flag
		for !t.gotMeta {
			t.getTemplateMeta(scanner.Text())
		}
	}
}

func (t *TemplateReader) getTemplateMeta(metaLine string) (map[string]string, error) {
	re := regexp.MustCompile("(^.*): (.*)")
	re.SubexpNames()
	match := re.FindStringSubmatch(metaLine)
	if len(match) < 2 {
		emptyMap := map[string]string{}
		return emptyMap, errors.New("match: this is not a meta string")
	}
	key := match[1]
	value := match[2]
	return map[string]string{key: value}, nil
}