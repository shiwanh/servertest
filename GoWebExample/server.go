package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SiteData struct {
	Name     string
	Count    string
	ServerIP string
	MapData  string
}

var count int = 0

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))

	m.Get("/", func(r render.Render, x *http.Request) {
		place := string(x.FormValue("place"))
		place = strings.Replace(place, " ", "+", -1)

		if len(place) <= 0 {
			place = "Universiteter+I+Norge"
		}
		r.HTML(200, "index", SiteData{"Universteter i Norge", strconv.Itoa(count), getServerIP(), place})
	})

	m.Get("/getCount", func() string {
		print(count)
		return strconv.Itoa(count)
	})

	go countNumber()
	m.RunOnAddr(":1122")
	m.Run()

}

func getServerIP() string {
	readApi, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(readApi.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func countNumber() {
	for true {
		count += 1
		time.Sleep(1 * time.Second)
	}

}
