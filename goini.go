package goini

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	sectionNameRE = regexp.MustCompile("^\\[([^\\]]+)\\]$")
	keyValueRE    = regexp.MustCompile("^([^=]+)\\s*=\\s*(.+)$")
)

// Parse reads text from the specified `io.Reader` and parses it into a two-level map.  The top
// level of the map is the section name, and the second layer map is the key-value map for that
// section.
//
//     config, err := goini.Parse(someReader)
//     if err != nil {
//         panic(err)
//     }
//
//     for sectionName, sectionMap := range config {
//         fmt.Printf("[%s]\n", sectionName)
//         for key, value := range sectionMap {
//             fmt.Printf("%s = %s\n", key, value)
//         }
//     }
func Parse(ior io.Reader) (map[string]map[string]string, error) {
	conf := make(map[string]map[string]string)
	sectionName := "General"

	scanner := bufio.NewScanner(ior)
	for scanner.Scan() {
		line := scanner.Text()

		// strip trailing comments and surrounding whitespace
		if comment := strings.IndexByte(line, ';'); comment >= 0 {
			line = line[:comment]
		}
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			// skip lines which are effectively blank
			continue
		}

		if md := sectionNameRE.FindStringSubmatch(line); md != nil {
			sectionName = md[1]
		} else if md := keyValueRE.FindStringSubmatch(line); md != nil {
			section, ok := conf[sectionName]
			if !ok {
				section = make(map[string]string)
				conf[sectionName] = section
			}
			section[md[1]] = md[2]
		} else {
			return nil, fmt.Errorf("cannot parse line: %q", line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return conf, nil
}

// ParseFile reads text from the file specified by the `pathname` and parses it into a two-level
// map.  The top level of the map is the section name, and the second layer map is the key-value map
// for that section.
//
//     config, err := goini.ParseFile(configPathname)
//     if err != nil {
//         panic(err)
//     }
//
//     for sectionName, sectionMap := range config {
//         fmt.Printf("[%s]\n", sectionName)
//         for key, value := range sectionMap {
//             fmt.Printf("%s = %s\n", key, value)
//         }
//     }
func ParseFile(pathname string) (map[string]map[string]string, error) {
	fh, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	return Parse(fh)
}
