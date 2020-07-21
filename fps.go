package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

const (
	KEY  = "Preferences"
	PATH = `SOFTWARE\Microsoft\Windows\CurrentVersion\TaskManager`
	POS  = 68
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Specify fps count")
	}

	fps, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	frameTime := 1000 / fps
	fmt.Printf("Setting %dms frame time\n", frameTime)

	k, err := registry.OpenKey(registry.CURRENT_USER, PATH, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	bytes, _, err := k.GetBinaryValue(KEY)
	if err != nil {
		log.Fatal(err)
	}

	binary.LittleEndian.PutUint32(bytes[POS:], uint32(frameTime))

	err = k.SetBinaryValue(KEY, bytes)
	if err != nil {
		log.Fatal(err)
	}
}
