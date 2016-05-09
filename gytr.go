// Copyright 2016, Michele Sorcinelli
//
// Gytr is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Gytr is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
