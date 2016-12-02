//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// )
//
// // Rankings struct
// type Rankings struct {
// 	Keyword  string
// 	GetCount uint32
// 	Engine   string
// 	Locale   string
// 	Mobile   bool
// }
//
// func main() {
// 	var jsonBlob = []byte(`
//     {"keyword":"hipaa compliance form", "get_count":157, "engine":"google", "locale":"en-us", "mobile":false}
//   `)
// 	rankings := Rankings{}
// 	err := json.Unmarshal(jsonBlob, &rankings)
// 	if err != nil {
// 		// nozzle.printError("opening config file", err.Error())
// 	}
//
// 	rankingsJSON, _ := json.Marshal(rankings)
// 	err = ioutil.WriteFile("output.json", rankingsJSON, 0644)
// 	fmt.Printf("%+v", rankings)
// }

package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
