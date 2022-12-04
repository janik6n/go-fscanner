package fscanner

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func RunCLI() {
	fmt.Println("______ ______")
	fmt.Println("|  ____/ ____|")
	fmt.Println("| |__ | (___   ___ __ _ _ __  _ __   ___ _ __ ")
	fmt.Println("|  __| \\___ \\ / __/ _` | '_ \\| '_ \\ / _ \\ '__|")
	fmt.Println("| |    ____) | (_| (_| | | | | | | |  __/ |   ")
	fmt.Println("|_|   |_____/ \\___\\__,_|_| |_|_| |_|\\___|_| ")
	fmt.Println("")
	fmt.Println("FScanner v0.0.1")
	fmt.Println("")

	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	pathPtr := fset.String("path", "", "Filesystem path to scan")
	dirsOnlyPtr := fset.Bool("dirsonly", false, "Scan only for directories")
	err := fset.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Path to scan:", *pathPtr)
	if *pathPtr == "" {
		fmt.Println("Mandatory --path is missing. Got:", *pathPtr)
		os.Exit(1)
	}

	fmt.Println("dirsonly:", *dirsOnlyPtr)

	// fmt.Println(Scan(os.DirFS(*pathPtr)))

	for _, v := range Scan(os.DirFS(*pathPtr), *dirsOnlyPtr) {
		fmt.Println("Found:", v)
	}
}

func Scan(fsys fs.FS, dirsonly bool) (objectList []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		fmt.Println(p)
		if p != "." {
			if d.IsDir() {
				// fmt.Println("Directory:", p)
				objectList = append(objectList, "dir: "+p)
			} else if !dirsonly {
				objectList = append(objectList, p)
			}

		}

		// if filepath.Ext(p) == ".txt" {
		// 	count++
		// }
		return nil
	})
	return objectList
}
