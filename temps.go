package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	hwmon := "/sys/class/hwmon"

	entries, err := os.ReadDir(hwmon)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		subHwmon := hwmon + "/" + entry.Name()

		name, err := os.ReadFile(subHwmon + "/name")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(strings.TrimSpace(string(name)), ": ")

		temp, err := os.ReadFile(subHwmon + "/temp1_input")
		if err == nil {
			fmt.Println(strings.TrimSpace(string(temp)))
		} else {
			fmt.Println("X")
		}
	}
}
