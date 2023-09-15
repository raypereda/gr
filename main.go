// This command is a rudimentary grep-like command that works on paragraphs.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// Define the command-line flag
	// blank_line_regex := "^[[:blank:]]*$" // matches UNIX or Windows blank lines
	blank_line_regex := `(\n\n)|(\r\n\r\n)` // matches UNIX or Windows blank lines
	separator_str := flag.String("r", blank_line_regex, "Record separator")

	usage := func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "    gr [-r <string>] <pattern> [<file>]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    gr [-h]\n")

		fmt.Fprintf(flag.CommandLine.Output(), "If no files names, reads from standard input.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Print records that match golang regex pattern.\n")
		flag.PrintDefaults()
	}
	flag.Usage = usage
	flag.Parse()
	var separator *regexp.Regexp
	var err error
	separator, err = regexp.Compile(*separator_str)
	if err != nil {
		log.Fatalln("Error compiling separator regex:", err)
	}

	var content string
	var pattern *regexp.Regexp

	if flag.NArg() == 1 || flag.NArg() == 2 {
		pattern_arg := flag.Arg(0)
		pattern, err = regexp.Compile(pattern_arg)
		if err != nil {
			log.Fatal("Error compiling pattern:", err)
		}
	}

	if flag.NArg() == 1 {
		content = stdin_text()
	} else if flag.NArg() == 2 {
		fmt.Println(flag.Arg((1)))
		data, err := os.ReadFile(flag.Arg(1))
		if err != nil {
			log.Fatal("Error reading file:", err)
		}
		content = string(data)
	} else {
		if flag.NArg() > 2 {
			fmt.Fprintf(flag.CommandLine.Output(), "Too many arguments.\n")
		}
		usage()
		os.Exit(1)
	}

	fmt.Println("record separator: ", *separator_str)

	records := separator.Split(content, -1)

	fmt.Println(len(records), " paragraphs")
	var match_count int
	for _, record := range records {
		if pattern.MatchString(record) {
			match_count++
			fmt.Println("Match # ", match_count)
			fmt.Println(record)
		}
	}
	fmt.Println(match_count, " matches")
}

// read all of standard input into a string
func stdin_text() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
