package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CommandList(w http.ResponseWriter, r *http.Request) {
	commands := Commands{
		Command{Name: "list-network", Parameters: []string{}},
		Command{Name: "show-network"},
		Command{Name: "create-network"},
		Command{Name: "delete-network"},
	}

	if err := json.NewEncoder(w).Encode(commands); err != nil {
		panic(err)
	}
}

func NetworkCreate(w http.ResponseWriter, r *http.Request) {
	networks := Networks{
		Network{Name: "Subnet 1", CIDR: "10.10.10.0/24"},
		Network{Name: "Subnet 2", CIDR: "192.168.10.0/24"},
	}

	networksJSON, _ := json.Marshal(networks)
	ioutil.WriteFile("/app/networks.json", networksJSON, 0644)
	fmt.Fprintln(w, "Network added:")
}

func NetworkList(w http.ResponseWriter, r *http.Request) {

	raw, err := ioutil.ReadFile("/app/networks.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var networks Networks
	json.Unmarshal(raw, &networks)

	if err := json.NewEncoder(w).Encode(networks); err != nil {
		panic(err)
	}
}

func NetworkShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	networkName := vars["name"]
	fmt.Fprintln(w, "Show:", networkName)
}
