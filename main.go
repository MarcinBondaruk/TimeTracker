package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Activity struct {
	id         int
	title      string
	startedAt  int64
	finishedAt int64
}

func (a *Activity) start() {
	a.startedAt = time.Now().UnixMilli()
}

func (a *Activity) stop() {
	if a.startedAt == 0 {
		fmt.Println("only ongoing activities can be stopped")
	}

	a.finishedAt = time.Now().UnixMilli()
}

func printHelp() {
	fmt.Printf("available commands:\n\tstart\tstart an activity\n\tstop\tstop an activity\n")
}

func printUnrecognizedCommand(cmd string) {
	fmt.Printf("Unrecognized command %q\n", cmd)
	printHelp()
}

func startActivity(id int, title string) {
	newActivity := Activity{id: id, title: os.Args[2]}
	newActivity.start()
	fmt.Printf("%+v\n", newActivity)

	data := strings.Join(
		[]string{
			strconv.Itoa(newActivity.id),
			newActivity.title,
			strconv.Itoa(int(newActivity.startedAt)),
			strconv.Itoa(int(newActivity.finishedAt)),
		},
		",",
	)
	err := os.WriteFile("activities.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
	}
}

func stopActivity(id int) {
	fmt.Printf("todo: implement me - stopping activity")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "start":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
		}
		startActivity(id, os.Args[3])
	case "stop":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
		}
		stopActivity(id)
	default:
		printUnrecognizedCommand(os.Args[1])
	}
}
