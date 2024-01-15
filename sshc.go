package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func removeHost(data string) string {
	pattern := `\bHost\b\s*`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(data, "")
	return result
}

func getConfig() string {
	// Open the ~/.ssh/config file
	file, err := os.Open(fmt.Sprintf("%s/.ssh/config", os.Getenv("HOME")))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var fileContent string

	// Iterate through each line of the file and concatenate to fileContent
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}

	return fileContent
}

type OutputRow struct {
	Index string
	HostName string
}

func getHosts(data string) {

	// Compile the regular expression with the case-insensitive flag
	re := regexp.MustCompile(`\bHost\s+(\w+)`)

	// Find the first match in the text
	linesArray := strings.Split(data, "\n")
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "SSH Host Name"})
	var hosts []string
	for _, line := range(linesArray) {
		if re.MatchString(line) {
			result := removeHost(line)
			hosts = append(hosts, result)
			t.AppendRows([]table.Row{{fmt.Sprintf("%v",len(hosts)), result}})
		}
	}
	fmt.Println(t.Render())
}

func main() {
	data := getConfig()
	getHosts(data)
}