package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func convertToColumnName(n string) string {
	sb := strings.Builder{}

	first := true
	count := 0
	for _, c := range n {
		upper := unicode.IsUpper(c)

		if upper && (count == 1 || count == 2 || count == 3) {
			count = count + 1
		}

		if upper && (count == 0 || count == 4) {
			if !first {
				sb.WriteRune('_')
			}
			count = 1
		}

		if !upper {
			count = 0
		}
		sb.WriteRune(unicode.ToUpper(c))
		first = false
	}
	return sb.String()
}

func main() {
	file, err := os.Open("./infoschema.go.copy")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outfile, err := os.Create("./infoschema.go.new")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "\t") {
			fmt.Fprintf(outfile, "%s\n", line)
			continue
		}

		mainLine := strings.TrimSpace(line)
		if mainLine == "" {
			fmt.Fprintf(outfile, "\n")
			continue
		}

		parts := strings.Split(mainLine, " ")
		fmt.Fprintf(outfile, "\t%s `db: \"%s\"`\n", mainLine, convertToColumnName(parts[0]))
	}
}
