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

//var jsonConfig []system.ConfigJSON
var jsonConfig system.SystemJSON

func rev(number int) int {
	div := 10
	reminder := 1
	newNumber := 0
	for reminder > 0 {
		reminder = number % div
		number = number - reminder
		number = number / 10
		if reminder > 0 {
			newNumber = newNumber * 10
			newNumber = newNumber + reminder
		}
	}
	return newNumber
}

func metrics(w http.ResponseWriter, r *http.Request) {
	var d string
	for _, commands := range jsonConfig.Commands {
		name := commands.Name
		d += name + " : " + string(data.Result[name]) + "\n"
	}
	fmt.Fprintf(w, d)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!!")
}

func parseConfigBackup() {
	pwd, _ := os.Getwd()
	var filename = "/config.json"
	source, err := ioutil.ReadFile(pwd + filename)
	//fmt.Printf("Value: %s", source)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(source, &jsonConfig)
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

	fmt.Println(jsonConfig.Duration)
	for _, commands := range jsonConfig.Commands {
		fmt.Println(commands.Name)
	}

}
func main() {
	// var dat map[string]interface{}
	// if err := json.Unmarshal(source, &dat); err != nil {
	// 	panic(err)
	// }
	//command := dat["command"].(string)
	//fmt.Println(dat)
	//fmt.Println(command)
	parseConfig()
	// a, _ := strconv.Atoi(os.Args[1])
	// fmt.Println(rev(int(a)))
	// fmt.Println(helper.Add(1, 3))
	// x, _ := helper.Swap(5, 7)
	// fmt.Println(x)
	// compare(1)
	// compare(4)
	// system.Sum(1, 2, 3)
	// system.DateCommand()
	// system.ScriptCustom()
	//system.ConfigCommand()
	go system.Poll(&data, jsonConfig)
	http.HandleFunc("/", home)
	http.HandleFunc("/metrics", metrics)
	http.ListenAndServe(":8080", nil)
}
