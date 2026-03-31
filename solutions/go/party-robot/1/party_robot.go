package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tnum string
	if table > 99 {
		tnum = fmt.Sprintf("%v", table)
	} else if table > 9 && table < 100 {
		tnum = fmt.Sprintf("0%v", table)
	} else if table < 10 {
		tnum = fmt.Sprintf("00%v", table)
	}
	return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, tnum, direction, distance, neighbor)
}
