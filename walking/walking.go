package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"slices"
)

func main() {
	last, err := walkRestart("/usr", "lib/golang/src/cmd/compile/internal/inline/inlheur", 1000)
	if err != nil {
		log.Fatalf("walkRestart got err: %s", err)
	}
	log.Printf("Last: %s\n", last)
}

func walkRestart(root string, restartDir string, limit int) (string, error) {
	ckpt := NewCheckpoint(restartDir)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if ckpt.Skip(path, d) {
			fmt.Printf("Skip[%d]: %s\n", limit, d.Name())
			return filepath.SkipDir
		}

		limit--
		if limit == 0 {
			return filepath.SkipAll
		}

		relpath, err := filepath.Rel(root, path)
		if err == nil && d.IsDir() {
			fmt.Printf("Path[%d]: %s\n", limit, relpath)
		}
		return err
	})
	if err != nil {
		return "", err
	}
	return ckpt.last, nil
}

type checkpoint struct {
	stack []string
	last  string
}

func (s *checkpoint) Skip(path string, d fs.DirEntry) bool {
	if !d.IsDir() {
		return false
	}
	if len(s.stack) == 0 {
		s.last = path
		return false
	}
	if d.Name() < s.stack[0] {
		return true
	}
	if d.Name() == s.stack[0] {
		fmt.Printf("Start: %s\n", d.Name())
		s.stack = s.stack[1:]
	}
	return false
}

func NewCheckpoint(path string) *checkpoint {
	c := &checkpoint{last: path}
	for path != "." {
		dir, file := filepath.Dir(path), filepath.Base(path)
		path, c.stack = dir, append(c.stack, file)
	}
	slices.Reverse(c.stack)
	return c
}
