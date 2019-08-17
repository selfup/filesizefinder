package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
)

func main() {
	var size string
	flag.StringVar(&size, "s", "", "size: 1MB,10MB,100MB,1GB,10GB,100GB,1TB")

	var directory string
	flag.StringVar(&directory, "d", "", "directory to start scanning recursively from")

	flag.Parse()

	p := NewFileSizeFinder(size)

	p.Scan(directory)

	for _, file := range p.Files {
		fmt.Println(file)
	}
}

// FileSizeFinder struct contains needed data to perform concurrent operations
type FileSizeFinder struct {
	mutex     sync.Mutex
	Files     []string
	Direction string
	Size      int64
}

// NewFileSizeFinder creates a pointer to FileSizeFinder with default values
func NewFileSizeFinder(size string) *FileSizeFinder {
	lff := new(FileSizeFinder)

	if runtime.GOOS == "windows" {
		lff.Direction = "\\"
	} else {
		lff.Direction = "/"
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

	_, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		panic("cannot read entry point - invalid directory!")
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

	for _, file := range files {
		if file.Size() >= lff.Size {
			lff.mutex.Lock()
			lff.Files = append(lff.Files, directory+lff.Direction+file.Name())
			lff.mutex.Unlock()
		}
	}

	dirLen := len(dirs)
	if dirLen > 0 {
		var dirGroup sync.WaitGroup
		dirGroup.Add(dirLen)

		for _, dir := range dirs {
			go func(diR os.FileInfo, direcTory string, direcTion string) {
				lff.findFiles(direcTory+direcTion+diR.Name(), direcTory)
				dirGroup.Done()
			}(dir, directory, lff.Direction)
		}

		dirGroup.Wait()
	}
}
