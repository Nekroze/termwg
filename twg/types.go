package twg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	uuid "github.com/satori/go.uuid"
)

func topicToDir(topic string) string {
	return filepath.Join("/", "tmp", "termwg", topic)
}

func awaitDirEmpty(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	if e := out.watcher.Add(dir); e != nil {
		panic(e)
	}

	for {
		select {
		case _, ok := <-watcher.Events:
			if !ok || getCountFromDir(dir) < 1 {
				break
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				panic(err)
			}
			fmt.Errorf("OH NO! %s", err)
		}
	}
}

func getCountFromDir(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return len(files)
}

func addCountToDir(dir string) {
	os.OpenFile(
		filepath.Join(dir, fmt.Sprintf("%s", uuid.NewV4())),
		os.O_RDONLY|os.O_CREATE,
		0666,
	)
}

func subCountFromDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	if e := os.Remove(files[rand.Intn(len(files))].Path); e != nil {
		panic(e)
	}
}

type WaitGroup struct {
	name string
}

func (wg WaitGroup) Wait() {
	dir := topicToDir(wg.name)

	os.MkdirAll(dir, os.ModePerm)
	if getCountFromDir(dir) < 1 {
		addCountToDir(dir)
	}

	awaitDirEmpty(dir)
}

func (wg WaitGroup) Done() {
	subCountFromDir(topicToDir(wg.name))
}

func (wg WaitGroup) Add(delta int) {
	addCountToDir(topicToDir(wg.name))
}
