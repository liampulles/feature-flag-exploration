package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("starting")
	err := http.ListenAndServe(":9080", http.HandlerFunc(handler))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println("closing")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Evaluate feature flag
	msg := "Hello!\n"
	if readFlags().UseAfrikaans {
		msg = "Hallo!\n"
	}

	// Respond
	w.WriteHeader(200)
	w.Write([]byte(msg))
}

type flags struct {
	UseAfrikaans bool `yaml:"UseAfrikaans"`
}

var defaultFlags = flags{
	UseAfrikaans: false,
}

func readFlags() flags {
	// Read flags file
	b, err := os.ReadFile("flags/flags.yaml")
	if err != nil {
		fmt.Printf("failed to read flags file, using default: %s\n", err.Error())
		return defaultFlags
	}

	// Parse as yaml
	var f flags
	err = yaml.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("failed to unmarshal flags file, using default: %s\n", err.Error())
		return defaultFlags
	}

	// Done
	return f
}
