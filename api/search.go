package api

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	query       string
	filepath    string
	ignore_case string
}

func Build(args []string) (*Config, error) {
	var query string
	var filepath string
	length := len(args)

	if length < 3 {

		return nil, errors.New("Both a query and a filepath are required")
	}

	query = args[1]
	filepath = args[2]

	ignore_case := os.Getenv("IGNORE_CASE")

	return &Config{
		query:       query,
		filepath:    filepath,
		ignore_case: ignore_case,
	}, nil
}

func (Config) Run(config Config) {
	contents, err := os.Open(config.filepath)
	if err != nil {
		panic(errors.New("Failed to open the file"))
	}
	defer contents.Close()

	content_scanner := bufio.NewScanner(contents)
	var results []string

	if config.ignore_case == "" {
		results = config.search(config.query, content_scanner)
	} else {
		results = config.search_case_insensitive(config.query, content_scanner)
	}

	for _, line := range results {
		fmt.Println(line)
	}
}

func (Config) search(query string, contents *bufio.Scanner) []string {
	var collection []string

	for contents.Scan() {
		line := contents.Text()
		if strings.Contains(line, query) {
			collection = append(collection, line)
		}
	}

	return collection
}

func (Config) search_case_insensitive(query string, contents *bufio.Scanner) []string {
	var collection []string

	for contents.Scan() {
		line := contents.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(query)) {
			collection = append(collection, line)
		}
	}

	return collection
}
