package main

import "fmt"

func main() {
	countryCapitalMap := make(map[string]string) // new map
	countryCapitalMap["China"] = "BeiJing"       // map赋值
	countryCapitalMap["France"] = "Paris"

	for country := range countryCapitalMap { //遍历
		fmt.Println("Capital of", country, "is ", countryCapitalMap[country])
	}
	printKV(countryCapitalMap, "France")

	// delete 元素
	delete(countryCapitalMap, "France")

	printKV(countryCapitalMap, "France")
}

func printKV(m map[string]string, key string) {
	value, ok := m[key] // 如果存在ok=true,反之则ok=false
	if ok {
		fmt.Printf("%s -> %s\n", key, value)
	} else {
		fmt.Printf("key: %s is not present\n", key)
	}
}
