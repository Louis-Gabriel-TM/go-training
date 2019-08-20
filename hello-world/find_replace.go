package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine searches for 'old' in line to replace it by 'new'.
// Returns 'found=true' if 'old' wad found, 'res' with the resulting string
// and 'occ' with the number of occurrences of 'old'.
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1) // -1 to replace all
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

// FindReplaceFile search for old in src to replace it by new.
// Returns 'occ' with the number of occurrences of 'old'
// and 'lines' with the index of lines containing 'old'.
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Fprintf(writer, res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "///GOLANG///"
	occ, lines, err := FindReplaceFile("to_replace.txt", "to_save.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find/replace: %v\n", err)
	}
	fmt.Println("=== Summary ===")
	defer fmt.Println("=== End of Summary ===")
	fmt.Printf("Nb of occurrences of '%v': %v\n", old, occ)
	fmt.Printf("Nb of lines: %d\n", len(lines))
	fmt.Print("Lines: [")
	len := len(lines)
	for i, line := range lines {
		fmt.Printf("%v", line)
		if i < len-1 {
			fmt.Print(" - ")
		}
	}
	fmt.Println(" ]")
}
