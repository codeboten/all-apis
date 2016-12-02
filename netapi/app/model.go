package main

// Network model representation
type Network struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}

// Networks is a slice of Network objects
type Networks []Network

// Command model representation
type Command struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
}

// Commands is a slice of Command objects
type Commands []Command
