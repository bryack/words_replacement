package main

import (
	"io/fs"
	"log"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`(?i)(\P{L}|^)поддел(к[аиуе]|ко(й|ю)|ок|кам|ками|ках)(\P{L}|$)`)

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

func main() {
	fsys := os.DirFS("/home/bryack/Documents/Obsidian/bryack/014_Go/to_do_list/contracts")
	data, err := ReadAndReplace(fsys, "Task.md")
	if err != nil {
		log.Fatal(err)
	}
	filename := "/home/bryack/Documents/Obsidian/bryack/014_Go/to_do_list/contracts/out.md"
	err = WriteFile(filename, data)
}
