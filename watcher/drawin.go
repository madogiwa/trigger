// +build darwin

package watcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/fsnotify/fsevents"
)

var (
	es *fsevents.EventStream
	cmd string
	args []string
)

func Watch(path string, command string, arg []string) error {
	dev, err := fsevents.DeviceForPath(path)
	if err != nil {
		return fmt.Errorf("failed to retrieve device for %s: %w", path, err)
	}

	es = &fsevents.EventStream{
		Paths:   []string{path},
		Latency: 1000 * time.Millisecond,
		Device:  dev,
		Flags:   fsevents.FileEvents | fsevents.WatchRoot}

	cmd = command
	args = arg
	return nil
}

func Start() {
	es.Start()
	ec := es.Events
	go func() {
		for msg := range ec {
			for _, event := range msg {
				handleEvent(event)
			}
		}
	}()
}

func Stop() {
	es.Stop()
}

func handleEvent(event fsevents.Event) {
	if event.Flags&fsevents.ItemCreated != 0 ||
		event.Flags&fsevents.ItemModified != 0 ||
		event.Flags&fsevents.ItemInodeMetaMod != 0 ||
		event.Flags&fsevents.ItemChangeOwner != 0 {

		arg := append(args, "/" + event.Path)
		log.Printf("exec: %s %s\n", cmd, strings.Join(arg, " "))

		command := exec.Command(cmd, arg...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Run()
	}
}
