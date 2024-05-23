package tools

import (
	"log"
	"os"
	"strings"
)

func Read_Input(infile_name string) string {
	buffer, err := os.ReadFile(infile_name)
	CheckError(err, "Error: failed to read infile: \""+infile_name+"\"!!")
	if len(buffer) == 0 {
		log.Fatalln("Error: infile is empty: \"" + infile_name + "\"!!")
	}
	return string(buffer)
}

func RemoveEmptyString(slice []string) []string {
	var result []string
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			result = append(result, slice[i])
		}
	}
	return result
}

func CHeckTemplate() string {
	ReplaceArgs()
	var data string
	if strings.HasSuffix(os.Args[2], "standard") || strings.HasSuffix(os.Args[2], "standard.txt") {
		data = Read_Input("Templates/standard.txt")
	} else if strings.HasSuffix(os.Args[2], "shadow") || strings.HasSuffix(os.Args[2], "shadow.txt") {
		data = Read_Input("Templates/shadow.txt")
	} else if strings.HasSuffix(os.Args[2], "thinkertoy") || strings.HasSuffix(os.Args[2], "thinkertoy.txt") {
		data = Read_Input("Templates/thinkertoy.txt")
	} else {
		log.Fatalln("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	return data
}

func ReplaceArgs() (string,string) {
	if strings.HasSuffix(os.Args[1], "standard") || strings.HasSuffix(os.Args[1], "standard.txt") || strings.HasSuffix(os.Args[1], "shadow") || strings.HasSuffix(os.Args[1], "shadow.txt") || strings.HasSuffix(os.Args[1], "thinkertoy") || strings.HasSuffix(os.Args[1], "thinkertoy.txt") {
		os.Args[1],os.Args[2] = os.Args[2],os.Args[1]
	} else {
		return os.Args[1],os.Args[2]
	}
	return os.Args[1],os.Args[2]
}

func IsAllNl(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg+"\n")
	}
}
