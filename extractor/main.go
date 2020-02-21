package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	file   = "./poly3.txt"
	output = "./output3.txt"
)

type LatLong struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Coordinates struct {
	Values []LatLong
}

func readData() (data []byte, err error) {
	read, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	data = read
	return
}

func createArgs(d []byte) (args []string, err error) {
	args = strings.Split(string(d), ",")
	if len(args) == 0 {
		err = errors.New("NO DATA WAS FOUND")
	}
	return
}

func createLatLng(args []string) {
	var latLngs []LatLong
	for _, pair := range args {
		ll := strings.Fields(pair)
		latLngs = append(latLngs, LatLong{Lat: ll[1], Lng: ll[0]})
		//fmt.Println(pair)
	}
	//fmt.Println(latLngs)
	coords := &Coordinates{
		Values: latLngs,
	}
	coordsJson, _ := json.Marshal(coords)
	fmt.Println(string(coordsJson))
	writeData(string(coordsJson))
}

func writeData(val string) {
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(val)
}

func main() {
	fmt.Println("extracting text")
	data, err := readData()
	if err != nil {
		panic(err)
	}
	args, err := createArgs(data)
	createLatLng(args)
}
