package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //匿名成员
	Radius int
}

type Wheel struct {
	Circle //匿名成员
	Spokes int
}

func Struct_main() {

	var w [3]Wheel
	for i := 0; i < 3; i++ {
		w[i].X = 8
		w[i].Y = 9
		w[i].Radius = 10
		w[i].Spokes = 11
	}

	fmt.Printf("%#v\n", w)

	//Translate struct-array to json ,called marshaling
	data1, err := json.Marshal(w)
	if err != nil {
		log.Fatalf("JSON matshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data1)

	//Translate struct-array to json typically
	data2, err := json.MarshalIndent(w, "", "\t")
	if err != nil {
		log.Fatalf("JSON matshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	//Translate json to struct-slice ,called unmarshaling
	var wheel []struct{ X, Y int }
	if err := json.Unmarshal(data1, &wheel); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(wheel)
}
