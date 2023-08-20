// 3. Write a script on any language to get the list of running PID by parsing “top”
// command
// Explanation:
// top is a Linux command, which shows all the processes in the system. The
// problem is to write a script that parses the output of top and provides the PIDs and
// users in the system . Consider the output of top is given a file input to you.

package topscript

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	PID  int
	User string
}

func Ques3() {
	filePath := "./output.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	processes := make([]Process, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "  PID") {
			continue // Skip header line
		}

		fields := strings.Fields(line)
		if len(fields) < 12 {
			continue
		}

		pid, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		user := fields[1]

		processes = append(processes, Process{
			PID:  pid,
			User: user,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the list of processes
	fmt.Println("Running Processes:")
	for _, process := range processes {
		fmt.Printf("PID: %d, User: %s\n", process.PID, process.User)
	}
}
