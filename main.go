package main

import (
	"encoding/json"
	"fmt"
	"io"
	task "mymodule/tasks"
	"os"
)


func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	fmt.Println(tasks)


}