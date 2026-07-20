package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var vulnerabilityDB = map[string]string{
	"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855": "Test Vulnerability (Empty File)",
}

func main() {
	dirPtr := flag.String("scan", ".", "The directory path to hunt")
	flag.Parse()

	printBanner()
	fmt.Printf("\n[*] HUNTING IN: %s\n", *dirPtr)
	fmt.Println("[*] Mode: Deep Enumeration & Hash-Based Vulnerability Audit")
	fmt.Println("--------------------------------------------------------------------------------")

	filepath.Walk(*dirPtr, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			processFile(path, info)
		}
		return nil
	})

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("[+] Hunt complete. The Ghost has finished its rounds.")
}

func processFile(path string, info fs.FileInfo) {
	if strings.Contains(info.Name(), ".env") || strings.Contains(info.Name(), "config.") {
		fmt.Printf("[!] ENUMERATED SENSITIVE: %s\n", path)
	}

	if strings.HasSuffix(info.Name(), ".exe") || strings.HasSuffix(info.Name(), ".elf") {
		hash, _ := getFileHash(path)
		if reason, exists := vulnerabilityDB[hash]; exists {
			fmt.Printf("[!!!] VULNERABILITY FOUND: %s | Match: %s\n", path, reason)
		}
	}

	if info.Mode()&0002 != 0 {
		fmt.Printf("[!] INSECURE PERMISSIONS: %s\n", path)
	}
}

func getFileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func printBanner() {
	fmt.Println(`
                             .---.                                    
                            /  _  \                                 
                           |  (o)  |      HUNTDD GHOST           
                            \  _  /                                
                              | | 
							 I don't hunt ghosts, I am the Ghost. I hunt for you.                              
  _    _  _   _  _   _  _____  ____  _____     ____  _   _  ___   ____  _____ 
 | |  | || | | || \ | ||_   _||  _ \|  __ \   / ___|| | | |/ _ \ / ___||_   _|
 | |__| || | | ||  \| |  | |  | | | | |  | | | |  _ | |_| | | | |\___ \  | |  
 |  __  || |_| || |\  |  | |  | |_| | |__| | | |_| ||  _  | |_| | ___) | | |  
 |_|  |_| \___/ |_| \_|  |_|  |____/|_____/   \____||_| |_|\___/|____ /  |_|  
 --------------------------------------------------------------------------------
 [+] CREATOR: Arjun Raj | CONTACT: arjunraj.cyber@gmail.com
 [+] "Even the silent have flaws."`)
}
