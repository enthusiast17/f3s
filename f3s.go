package f3s

// Flag a struct with tag, content, des params.
type Flag struct {
	tag     string // Flag's tag
	content string // Flag's content
	des     string // Flag's description
}

// Flags a struct with Flag's array param.
type Flags struct {
	flags map[string]*Flag
}

// SetNewFlags returns a new Flags's object with parsed flags.
func SetNewFlags(args []string) *Flags {
	flags := make(map[string]*Flag)
	for _, arg := range args {
		hyphen := -1
	r:
		for j, r := range []rune(arg) {
			switch r {
			case '-':
				if j <= 1 {
					hyphen = j
				} else {
					break r
				}
			case '=':
				if j+1 < len(arg) && hyphen > -1 && hyphen+1 < len(arg) {
					flag := &Flag{tag: arg[hyphen+1 : j], content: arg[j+1:]}
					flags[flag.tag] = flag
				}
				break r
			}
		}
	}
	return &Flags{flags: flags}
}

// SetFlag returns a Flag's content value.
func (flags *Flags) SetFlag(tag string, des string) string {
	flag := flags.flags[tag]
	var content string
	if flag != nil {
		flag.des = des
		content = flag.content
	} else {
		flags.flags[tag] = &Flag{tag: tag, des: des}
	}
	return content
}

// GetContents returns an array with contents.
func GetContents(flags *Flags) []string {
	var contents []string
	if flags != nil {
		for _, flag := range flags.flags {
			if flag != nil && flag.des != "" {
				contents = append(contents, flag.content)
			}
		}
	}
	return contents
}

// GetTags returns an array with tags.
func GetTags(flags *Flags) []string {
	var tags []string
	if flags != nil {
		for _, flag := range flags.flags {
			if flag != nil && flag.des != "" {
				tags = append(tags, flag.tag)
			}
		}
	}
	return tags
}

// GetDescriptions returns an array with descriptions.
func GetDescriptions(flags *Flags) []string {
	var dess []string
	if flags != nil {
		for _, flag := range flags.flags {
			if flag != nil && flag.des != "" {
				dess = append(dess, flag.des)
			}
		}
	}
	return dess
}

// GetHelper returns a string with tags and descriptions.
func GetHelper(flags *Flags) string {
	var helper string
	if flags != nil {
		var iter int
		lenFlags := len(flags.flags)
		for _, flag := range flags.flags {
			if flag != nil && flag.des != "" {
				helper += ("--" + flag.tag + "\n\t" + flag.des)
				if iter != lenFlags-1 {
					helper += "\n"
				}
				iter++
			} else {
				lenFlags--
			}
		}
	}
	return helper
}
