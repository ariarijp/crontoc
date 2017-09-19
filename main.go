package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/robfig/cron"
)

func parseLines(sortFlag bool) ([]string, error) {
	now := time.Now()
	r := regexp.MustCompile(`^[\d/\-,*]`)
	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

	var result []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !r.MatchString(line) {
			continue
		}
		def := strings.Fields(line)
		sched, err := specParser.Parse(strings.Join(def[0:5], " "))
		if err != nil {
			return nil, err
		}
		result = append(result, fmt.Sprintf("Next: %s # %s", sched.Next(now), line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if sortFlag {
		sort.Strings(result)
	}

	return result, nil
}

func main() {
	sortFlag := flag.Bool("sort", false, "Sort by next execution time.")
	flag.Parse()

	result, err := parseLines(*sortFlag)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range result {
		fmt.Println(l)
	}
}
