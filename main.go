package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
)

const VERSION = "1.0.0"
const usage =
` Usage with hashcat.pot: checkpot -u user_hashes.txt -p hashcat.pot -o user_passes.txt

  Options:
    -h, --help		Show usage and exit.
    -v			Show version and exit.
    -u			The line separated list of user:hash entries.
    -p			The line separated list of hash:pass entries.
    -o			Path to save file containing user:pass line entries.
`

func main() {
	var (
		flVersion	= flag.Bool("v", false, "")
		flHashfile	= flag.String("u", "", "")
		flPotfile	= flag.String("p", "", "")
		flOutfile	= flag.String("o", "", "")
	)
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	if *flVersion {
		fmt.Printf("checkpot version: %s\n", VERSION)
		os.Exit(0)
	}

	if *flHashfile == "" || *flPotfile == "" {
		fmt.Println("Error: Both user hash file and pot file are required...\n")
		fmt.Print(usage)
		os.Exit(1)
	}

	fmt.Println("Loading hashes and passwords from:", *flPotfile)
	potMap, err := readSplitMap(*flPotfile)
	if err != nil {
		fmt.Println("Error occured while reading pot file:", err)
	}

	fmt.Printf("Loading users and hashes from: %s\n", *flHashfile)
	hashMap, err := readSplitMap(*flHashfile)
	if err != nil {
		fmt.Println("Error occured while reading hash file:", err)
	}
	fmt.Println("Number of Passwords loaded:", len(potMap))
	fmt.Println("Number of Users loaded:", len(hashMap))

	fmt.Println("Searching password pot for user hashes...")
	found := make(map[string]string)
	for user, hash := range hashMap {
		if _, exists := potMap[hash]; exists {
			found[user] = potMap[hash]
		}
	}
	fmt.Printf("Number of User credentials found: %d\n", len(found))
	fmt.Printf("Percentage of user credentials found: %.1f%%\n", (float32(len(found)) / float32(len(hashMap))) * 100)

	if *flOutfile != "" {
		fmt.Println("Saving user:pass list to:", *flOutfile)
		saveUserPass(*flOutfile, found)
	}
	fmt.Println("All done! Exiting...")
}

func readSplitMap(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hashes := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ":")
		if len(fields) == 2 {
			hashes[fields[0]] = fields[1]
		}
	}

	return hashes, scanner.Err()
}

func saveUserPass(path string, found map[string]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for user, pass := range found {
		file.WriteString(fmt.Sprintf("%s:%s\n", user, pass))
	}
	file.Sync()

	return nil
}