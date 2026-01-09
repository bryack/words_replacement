package replacer

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`(?i)(\P{L}|^)поддел(к[аиуе]|ко(й|ю)|ок|кам|ками|ках)(\P{L}|$)`)

func Run(fsys fs.FS, args []string, out io.Writer) error {
	if len(args) < 1 {
		return fmt.Errorf("missing filename argument")
	}
	filename := args[0]

	result, err := ReadAndReplace(fsys, filename)
	if err != nil {
		return err
	}

	fmt.Fprint(out, result)
	return nil
}

func Replace(input string) string {
	return re.ReplaceAllStringFunc(input, func(s string) string {
		groups := re.FindStringSubmatch(s)

		if len(groups) < 5 {
			return s
		}
		prefix := groups[1]
		ending := groups[2]
		boundary := groups[4]

		replacement := "fake"
		if ending == "ок" || ending == "ками" || ending == "ках" || ending == "ки" {
			replacement = "fakes"
		}
		return prefix + replacement + boundary
	})
}

func ReadAndReplace(fsys fs.FS, filename string) (string, error) {
	data, err := fs.ReadFile(fsys, filename)
	if err != nil {
		return "", err
	}

	repl := Replace(string(data))

	return repl, nil
}

func WriteFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
