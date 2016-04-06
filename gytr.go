package main

import (
	"flag"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func parseErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// command line args
	templateFile := flag.String("t", "template", "Template file")
	dataFile := flag.String("d", "data.yml", "Data YAML file")
	flag.Parse()

	// read template
	file, err := ioutil.ReadFile(*templateFile)
	parseErr(err)
	text := string(file)

	// read yaml data
	file, err = ioutil.ReadFile(*dataFile)
	parseErr(err)

	// parse yaml data
	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &data)
	parseErr(err)

	// init the template engine
	tmpl, err := template.New("template").Parse(text)
	parseErr(err)

	// render and write output
	err = tmpl.Execute(os.Stdout, data)
	parseErr(err)
}
