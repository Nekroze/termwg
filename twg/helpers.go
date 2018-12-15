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
	return filepath.Join(os.TempDir(), "termwg", topic)
}

func awaitDirEmpty(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	if e := watcher.Add(dir); e != nil {
		panic(e)
	}

	for {
		select {
		case _, ok := <-watcher.Events:
			if !ok || getCountFromDir(dir) < 1 {
				return
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				panic(err)
			}
			fmt.Println("OH NO!", err)
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
	path := filepath.Join(dir, fmt.Sprintf("%s", uuid.NewV4()))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	_, err = os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
}

func subCountFromDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	name := files[rand.Intn(len(files))].Name()
	path := filepath.Join(dir, name)
	if e := os.Remove(path); e != nil {
		panic(e)
	}
}
