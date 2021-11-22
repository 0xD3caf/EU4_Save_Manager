package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*

*************************************
*                                   *
*        EU4 SAVE MANAGER           *
*                                   *
*************************************

Collects eu4 Gamesaves with single button click
Can drop old version back in with 1 click as well
Saves all saves as a tree ideally with time and other info + notes

2 main goals
    1. Assist with backups for long games, maintaining save continuity, eas of export
    2. Assist with save scumming for run starts, correct rivals, general spam, etc
*/

//GLOBAL VARIABLES
var SaveName string
var SaveLocation string
var ExportedSavesLocation string

func init() {
	//set globals from configs file
	currDir, err := os.Getwd()
	CheckErr(err)
	LineList := LineReader(currDir + "\\configs.txt")
	SaveName = strings.Split(LineList[0], "=")[1]
	SaveLocation = strings.Split(LineList[1], "=")[1]
	ExportedSavesLocation = strings.Split(LineList[2], "=")[1]
	/* //VAR CHECK
	fmt.Println("Save Name:", SaveName)
	fmt.Println("Save Location:", SaveLocation)
	fmt.Println("Internal Save Location:", ExportedSavesLocation)
	*/
}

func main() {
	fmt.Println("EU4 SAVE SCUMMER")
	ExportSaves()
}

func ExportSaves() bool {
	//Grabs saves for current game and copys to external folder
	err := os.Mkdir(ExportedSavesLocation+"\\"+SaveName, 0755)
	CheckErr(err)
	DirList, err := os.ReadDir(SaveLocation)
	//!REFACTOR THIS SHIT
	if !CheckErr(err) {
		return false
	}
	for _, line := range DirList {
		if strings.Contains(line.Name(), SaveName) {
			fmt.Println(line.Name())
			oldSaveLoc := SaveLocation + "\\" + SaveName + "\\" + line.Name()
			newSaveLoc := ExportedSavesLocation + "\\" + SaveName + "\\" + line.Name()
			fmt.Println(oldSaveLoc)
			fmt.Println(newSaveLoc)
			copy(oldSaveLoc, newSaveLoc)
		}
	}
	return true
}

func ImportSaves() {
	//Takes saves selected by user and places them back in eu4 gamesave folder + check for saves that are newer, i.e they take form SAVENAME_backup(1)...backup(N).* where N is the number of backups in import sets latest save
	return
}

func ModelSaves() {
	//takes all current saves within export tree and draws a diagram/tree of saves
}

func CheckErr(Error error) bool {
	if Error != nil {
		errString := Error.Error()
		if strings.Contains(errString, "Cannot create") {
			fmt.Println(Error)
		} else {
			fmt.Println(Error)
			return false
		}
	}
	return true
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func LineReader(filename string) []string {
	var LineList []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		LineList = append(LineList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return LineList
}
