package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var outDir = flag.String("o", "", "--o ./output-directory")
	flag.Parse()

	var inputDirectory = os.Args[1]
	var skipped int

	if len(*outDir) == 0 {
		update := promptYN_DefaultNo("No output directory specified.  Overwrite?")
		if update {
			*outDir = inputDirectory
		}
	}

	err := filepath.Walk(inputDirectory, convertFile(*outDir, &skipped))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d skipped\n", skipped)
}

func convertFile(outputDir string, skipped *int) filepath.WalkFunc {
	if len(outputDir) > 0  {
		info, err := os.Stat(outputDir)
		if err != nil && !info.IsDir() {
			create := promptYN_DefaultNo(fmt.Sprintf("path %s doesn't exist.  Create?", outputDir))
			if !create {
				panic(fmt.Errorf("cannot write to %s if it does not exist", outputDir))
			}
			err = os.Mkdir(outputDir, 0755)
			if err != nil {
				panic(err)
			}
		}
	}
	var maxSize int64

  return func (p string, info os.FileInfo, err error) error {
		if err != nil {
				return err
		}
		if ! info.IsDir() {
			outfile := ""
			if len(outputDir) > 0 {
				outfile = fmt.Sprintf("%s%c%s", outputDir, os.PathSeparator, info.Name())
				fmt.Println(p, "->", outfile)
			} else { *skipped += 1}

			if info.Size() > maxSize {
				maxSize = info.Size()
				fmt.Println("max size update", maxSize)
				fmt.Println(info.Name())
			}
		}
		return nil
	}
}

func promptYN_DefaultNo(prompt string) bool {
	fmt.Printf("%s %s", prompt, "[y/N]: ")
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}
