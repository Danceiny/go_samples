package main

import (
	"github.com/buger/jsonparser"
	"fmt"
)

func main()  {

	data := []byte(`{
	  "person": {
		"name": {
		  "first": "Leonid",
		  "last": "Bugaev",
		  "fullName": "Leonid Bugaev"
		},
		"github": {
		  "handle": "buger",
		  "followers": 109
		},
		"avatars": [
		  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
		]
	  },
	  "company": {
		"name": "Acme"
	  }
	}`)

	// You can specify key path by providing arguments to Get function

	fullName := jsonparser.Get(data, "person", "name", "fullName")
	fmt.Print(fullName)

}
