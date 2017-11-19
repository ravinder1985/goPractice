package system

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

//ConfigJSON object
type ConfigJSON struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Options string `json:"options"`
}

// SystemJSON has json object
type SystemJSON struct {
	Duration int
	Commands []SystemCommands
}

// SystemCommands has commands onject
type SystemCommands struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Options string `json:"options"`
}

// Data would hole the data
type Data struct {
	Result map[string][]byte
}

var jsonConfig SystemJSON

//var jsonConfig []ConfigJSON

// Data would hold the data
//var Data map[string]string

//Sum would add as many as number you pass
func Sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// DateCommand Function is to run date command on server
func DateCommand() {
	out, err := exec.Command("date").Output()
	if (err) != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

// Poll would parse configs and run them every interval.
func Poll(data *Data, jsonConfig SystemJSON) {
	// pwd, _ := os.Getwd()
	// var filename = "/config.json"
	// source, err := ioutil.ReadFile(pwd + filename)
	// //fmt.Printf("Value: %s", source)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// json.Unmarshal(source, &jsonConfig)
	duration := jsonConfig.Duration
	for {
		<-time.After(time.Duration(duration) * time.Second)
		ConfigCommand(jsonConfig, data)
	}
}

// ConfigCommand Run command from config file
func ConfigCommand(jsonConfig SystemJSON, data *Data) {
	for _, commands := range jsonConfig.Commands {
		name := commands.Name
		command := commands.Command
		options := commands.Options
		if data.Result == nil {
			fmt.Println("Allocate memory")
			data.Result = make(map[string][]byte, 1)
		}
		if options == "" {
			out, err := exec.Command(command).Output()
			if (err) != nil {
				log.Fatal(err)
			}
			data.Result[name] = out
			fmt.Printf("%s: %s", name, out)
		} else {
			out, err := exec.Command(command, options).Output()
			if (err) != nil {
				log.Fatal(err)
			}
			data.Result[name] = out
			fmt.Printf("%s: %s", name, out)
		}
	}
}

// ScriptCustom is to run custom script
func ScriptCustom() {
	out, err := exec.Command("sh", "/Users/harvindersingh/shell_script/checkSpace.sh").Output()
	if (err) != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
