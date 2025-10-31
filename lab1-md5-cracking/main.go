package main

import (
	"bufio"
	"flag"
	"fmt"
	"lab1-md5-cracking/utils/crack"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	wordlistPath := flag.String("w", "", "Path to the wordlist file")
	hashValue := flag.String("h", "6a85dfd77d9cb35770c9dc6728d73d3f", "Target MD5 hash to crack")

	flag.Usage = func() {
		fmt.Printf("Usage:\n  %s -w <wordlist> -h <hash>\n", filepath.Base(os.Args[0]))
		fmt.Println("\nExample:")
		fmt.Println("  go run main.go -w nord_vpn.txt -h 6a85dfd77d9cb35770c9dc6728d73d3f")
	}

	flag.Parse()

	if *wordlistPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*wordlistPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening wordlist '%s': %v\n", *wordlistPath, err)
		os.Exit(1)
	}
	defer f.Close()

	vf, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating verbose file: %v\n", err)
		os.Exit(1)
	}
	defer vf.Close()

	vw := bufio.NewWriter(vf)
	defer vw.Flush()

	start := time.Now()
	scanner := bufio.NewScanner(f)

	var tried int64 = 0
	progressInterval := int64(10000)

	fmt.Printf("Starting MD5 crack.\nWordlist: %s\nTarget: %s\nVerbose: verbose.txt\n\n",
		*wordlistPath, *hashValue)

	found := false
	var foundPassword string
	var foundLine int64 = 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tried++
		if line == "" {
			continue
		}

		h := crack.MD5String(line)
		fmt.Fprintf(vw, "[%s] #%d\t%s\t%s\n", time.Now().Format(time.RFC3339), tried, line, h)

		if tried%1000 == 0 {
			_ = vw.Flush()
		}

		if h == *hashValue {
			found = true
			foundPassword = line
			foundLine = tried
			break
		}

		if tried%progressInterval == 0 {
			fmt.Printf("Tried: %d passwords — elapsed: %v\n", tried, time.Since(start).Truncate(time.Second))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading wordlist: %v\n", err)
	}

	_ = vw.Flush()

	if found {
		elapsed := time.Since(start)
		fmt.Printf("\nFOUND!\nPassword: %q\nTried: %d\nLine: %d\nElapsed: %v\n", foundPassword, tried, foundLine, elapsed)
		fmt.Fprintf(vw, "\nFOUND: password=%s tried=%d line=%d elapsed=%s\n", foundPassword, tried, foundLine, elapsed)
	} else {
		fmt.Printf("\nNot found. Tried %d candidates.\n", tried)
		fmt.Fprintf(vw, "\nNOT FOUND: tried=%d\n", tried)
	}
}
