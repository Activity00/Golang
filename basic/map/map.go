package main

import "fmt"

func main() {
	dict := make(map[string]string)
	fmt.Println(dict)

	dict1 := map[string]string{"red": "red", "blue": "blue"}
	fmt.Println(dict1)

	dict2 := map[string][]string{"red": {"1", "2", "3"}}
	fmt.Println(dict2)
	fmt.Println(dict2["red"])

	// nil 映射不能赋值
	colors := map[string]string{"red": "red"}
	colors["yellow"] = "yellow"
	fmt.Println(colors)

	value, exist := colors["blue"]
	if exist {
		fmt.Println(exist, value)
	} else {
		fmt.Println(exist, value)
	}

	dict3 := map[string]string{}
	dict3["Blue"] = "blue"
	fmt.Println(dict3)

	delete(dict3, "blue")
	for key, value := range dict1 {
		fmt.Println(key, value)
	}

}
