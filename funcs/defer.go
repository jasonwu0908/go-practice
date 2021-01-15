package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	text := []string{"Json", "Joe", "Ian", "Teng"}
	logPath := "./loglog"
	writeLog(logPath, text...)
	deleteLog(logPath)
}

func writeLog(logPath string, x ...string) {
	f, err := os.Create(logPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range x {
		f.WriteString(v + "\n")
		fmt.Printf(`Writing "%v" into -> "%v"`+"\n", v, logPath)
	}
	defer f.Close()
	fmt.Println("Close ", logPath)
}

func deleteLog(logPath string) {
	err := os.Remove(logPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete ", logPath)
}
