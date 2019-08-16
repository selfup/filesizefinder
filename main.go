package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	var direction string
	flag.StringVar(&direction, "d", "", "direction of paths back (windows) or forward (unix/linux)")

	var size string
	flag.StringVar(&size, "s", "", "size: 1MB,10MB,100MB,1GB,10GB,100GB,1TB")

	var folder string
	flag.StringVar(&folder, "f", "", "folder/directory to start scanning recursively from")

	flag.Parse()

	p := NewFileSizeFinder(direction, size)

	p.Scan(folder)

	for _, file := range p.Files {
		fmt.Println(file)
	}
}

// FileSizeFinder struct contains needed data to perform concurrent operations
type FileSizeFinder struct {
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
	Files     []string
	Direction string
	Size      int64
}

// NewFileSizeFinder creates a pointer to FileSizeFinder with default values
func NewFileSizeFinder(direction string, size string) *FileSizeFinder {
	lff := new(FileSizeFinder)

	switch direction {
	case "backslash":
		lff.Direction = "\\"
		break
	case "forwardslash":
		lff.Direction = "/"
		break
	default:
		panic("please provide backslash or forwardslash for directions")
	}

	switch size {
	case "1MB":
		lff.Size = 1000
		break
	case "10MB":
		lff.Size = 10000
		break
	case "100MB":
		lff.Size = 1000000
		break
	case "1GB":
		lff.Size = 1000000000
		break
	case "10GB":
		lff.Size = 10000000000
		break
	case "100GB":
		lff.Size = 1000000000000
		break
	case "1TB":
		lff.Size = 1000000000000000
		break
	default:
		panic("please provide a size 1MB 10MB 100MB 1GB 10GB 100GB 1TB")
	}

	return lff
}

// Scan is a concurrent/parallel directory walker
func (lff *FileSizeFinder) Scan(directory string) {
	if directory == "" {
		panic("please provide a directory")
	}

	lff.findFiles(directory, "")
}

func (lff *FileSizeFinder) findFiles(directory string, prefix string) {
	paths, _ := ioutil.ReadDir(directory)

	var dirs []os.FileInfo
	var files []os.FileInfo

	for _, path := range paths {
		if path.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
	}

	lff.waitGroup.Add(len(files))

	for _, file := range files {
		func(f os.FileInfo, d string) {
			if f.Size() > lff.Size {
				lff.mutex.Lock()
				lff.Files = append(lff.Files, d+lff.Direction+f.Name())
				lff.mutex.Unlock()
				lff.waitGroup.Done()
			} else {
				lff.waitGroup.Done()
			}
		}(file, directory)
	}

	lff.waitGroup.Wait()

	for _, dir := range dirs {
		lff.findFiles(directory+lff.Direction+dir.Name(), directory)
	}
}
