// 2. Write a C/C++/Go program to iterate a linux directory tree and count the number
// of files and number of subdirectories in each directory. Also Count the total number
// of files and directories in the directory tree. Implement the logic using
// multi-threading to improve the scanning of the directory.

package linuxtree

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type CountResult struct {
	Files        int
	SubDirs      int
	TotalFiles   int
	TotalSubDirs int
}

func countFilesAndDirs(root string, wg *sync.WaitGroup, results chan<- CountResult) {
	defer wg.Done()

	files, subDirs := 0, 0

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if info.IsDir() {
			subDirs++
		} else {
			files++
		}

		return nil
	})

	results <- CountResult{Files: files, SubDirs: subDirs}
}

func Ques2() {
	_ = godotenv.Load() // Load environment variables from .env file

	rootDir := os.Getenv("DIRECTORY")
	if rootDir == "" {
		fmt.Println("Env variable is not set, assuming the root directory as the present working directory")
		rootDir, _ = os.Getwd() // Use current working directory if DIRECTORY is not set
	}

	var wg sync.WaitGroup
	results := make(chan CountResult)

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		fmt.Println(path, info.Size())

		if info.IsDir() {
			wg.Add(1)
			go countFilesAndDirs(path, &wg, results)
		}

		return nil
	})

	go func() {
		wg.Wait()
		close(results)
	}()

	totalFiles := 0
	totalSubDirs := 0

	for result := range results {
		totalFiles += result.Files
		totalSubDirs += result.SubDirs
	}

	fmt.Printf("Total Files: %d\n", totalFiles)
	fmt.Printf("Total Subdirectories: %d\n", totalSubDirs)
}
