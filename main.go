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

const LAYOUT = "2006-01-02T15:04:05"

func parseLines(from time.Time, sortFlag bool) ([]string, error) {
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
		next := sched.Next(from).In(time.Local)
		result = append(result, fmt.Sprintf("Next: %s # %s", next, line))
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
	now := time.Now().UTC()
	sortFlag := flag.Bool("sort", false, "Sort by next execution time.")
	fromStr := flag.String("from", now.Format(LAYOUT), "Show TOC from its time.")
	flag.Parse()

	from, err := time.Parse(LAYOUT, *fromStr)
	if err != nil {
		log.Fatal(err)
	}

	result, err := parseLines(from, *sortFlag)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range result {
		fmt.Println(l)
	}
}
