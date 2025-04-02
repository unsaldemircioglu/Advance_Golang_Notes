package resful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, todo)
	fmt.Println(todo)

}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş Yapılacak", false}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

}
