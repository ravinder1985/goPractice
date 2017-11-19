package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/goPractice/system"
)

var data system.Data
var jsonConfig system.JSON

func metrics(w http.ResponseWriter, r *http.Request) {
	var d, s string
	for _, commands := range jsonConfig.Commands {
		name := commands.Name + "_" + commands.Type
		s = "# HELP " + name + " The total number of HTTP requests.\n"
		s = s + "# TYPE " + name + " counter"
		d += s + "\n"
		//name := commands.Name
		d += name + "{type=" + commands.Lables.Type + "} " + string(data.Result[commands.Name]) + "\n\n"
	}
	fmt.Fprintf(w, d)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!!")
}

func parseConfig() {
	pwd, _ := os.Getwd()
	var filename = "/config.json"
	source, err := ioutil.ReadFile(pwd + filename)
	//fmt.Printf("Value: %s", source)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(source, &jsonConfig)
	//
	// fmt.Println(jsonConfig.Duration)
	// for _, commands := range jsonConfig.Commands {
	// 	fmt.Println(commands.Name)
	// }
}

func main() {
	parseConfig()
	go system.Poll(&data, jsonConfig)
	http.HandleFunc("/", home)
	http.HandleFunc("/metrics", metrics)
	http.ListenAndServe(":"+jsonConfig.Port, nil)
}
