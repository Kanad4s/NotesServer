package cli

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	showAll = iota
	showGroups
	addGroup
	removeGroup
	addNote
	removeNote
)

func ShowNavigation() {
	var err error
	desc := ""
	i := 0
	for err != nil {
		desc, err = showOptionDescription(i)
		if err != nil {
			fmt.Printf("%d. %s, %v\n", i, desc, err)
		}
		i++
	}
}

func GetNumber() (int, error) {
	reader := bufio.NewReader(os.Stdin)
readNumber:
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	val, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Print(err.Error())
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "invalid syntax") {
			fmt.Println("Wrong input, enter number:")
			goto readNumber
		}
		panic(err)
	}
	return val, nil
}

func showOptionDescription(option int) (desc string, err error) {
	err = nil
	switch option {
	case showAll:
		desc = "Show all groups and notes"
	case showGroups:
		desc = "Show groups"
	case addGroup:
		desc = "Add group"
	case removeGroup:
		desc = "Remove group"
	case addNote:
		desc = "Add note to group"
	case removeNote:
		desc = "Remove note from group"
	default:
		err = errors.New("unsupported type")
	}
	return
}
