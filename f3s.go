package f3s

import (
	"fmt"
	"os"
)

// Flag a struct with tag, content, dec params.
type Flag struct {
	tag     string
	content string
	dec     string
}

// Flags a struct with Flag's array param.
type Flags struct {
	flags []*Flag
}

// SetNewFlags returns a new Flags's object.
func SetNewFlags() *Flags {
	return &Flags{}
}

// Parse returns a Flags object with parsed flags or displays helper if there is no arguments.
func (flags *Flags) Parse(displayHelper bool) *Flags {
	args := os.Args[1:]
	if displayHelper {
		if len(args) == 0 || (len(args) == 1 && (args[0] == "-h" || args[0] == "--help")) {
			for _, flag := range flags.flags {
				fmt.Printf("--%s\n\t%s\n", flag.tag, flag.dec)
			}
			os.Exit(0)
		}
	}
	for _, flag := range flags.flags {
		for _, arg := range args {
			f := ("--" + flag.tag + "=")
			if len(arg) > len(f) && arg[:len(f)] == f {
				flag.content = arg[len(f):]
				break
			}
		}
	}
	return flags
}

/* I will consider about it later ;)
func SetFlag(tag string, dec string) *Flags {
	return parse(tag, dec)
}
*/

// SetFlag returns a Flags object with tag and dec.
func (flags *Flags) SetFlag(tag string, dec string) *Flags {
	var newFlags []*Flag
	if flags != nil {
		newFlags = append(newFlags, flags.flags...)
	}
	newFlag := &Flag{tag: tag, dec: dec}
	newFlags = append(newFlags, newFlag)
	flags.flags = newFlags
	return flags
}

// GetMap returns a map with flags and contents.
func GetMap(flags *Flags) map[string]string {
	var m map[string]string
	for _, flag := range flags.flags {
		if flag != nil {
			m[flag.tag] = flag.tag
		}
	}
	return m
}

// GetContents returns an array with contents.
func GetContents(flags *Flags) []string {
	var contents []string
	for _, flag := range flags.flags {
		if flag != nil {
			contents = append(contents, flag.content)
		}
	}
	return contents
}

// GetTags returns an array with tags.
func GetTags(flags *Flags) []string {
	var tags []string
	for _, flag := range flags.flags {
		if flag != nil {
			tags = append(tags, flag.tag)
		}
	}
	return tags
}

// Content returns a content from given Flags object.
func Content(flags *Flags, tag string) string {
	var content string
	for _, flag := range flags.flags {
		if flag != nil && flag.tag == tag {
			content = flag.content
			break
		}
	}
	return content
}

// HasContent returns a bool, if the given tag has a content returns 'true', else returns 'false'.
func HasContent(flags *Flags, tag string) bool {
	for _, flag := range flags.flags {
		if flag != nil && flag.tag == tag {
			if flag.content == "" || len(flag.content) == 0 {
				return false
			}
		}
	}
	return true
}
