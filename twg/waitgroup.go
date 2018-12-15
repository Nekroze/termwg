package twg

import (
	"os"
)

type WaitGroup struct {
	Name string
}

func (wg WaitGroup) Wait() {
	dir := topicToDir(wg.Name)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if getCountFromDir(dir) < 1 {
		addCountToDir(dir)
	}

	awaitDirEmpty(dir)
}

func (wg WaitGroup) Done() {
	subCountFromDir(topicToDir(wg.Name))
}

func (wg WaitGroup) Add(delta int) {
	for i := 1; i <= delta; i++ {
		addCountToDir(topicToDir(wg.Name))
	}
}
