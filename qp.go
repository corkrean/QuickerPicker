package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math/rand"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	//calling functins that request 1) population and 2)number to be picked from population then store them
	totalPool := pool()
	totalPicks := picks()

	//declares string array slice with variable length to store names of people in population
	names := make([]string, totalPool)

	//declaring outside of for header to avoid int64/int type mismatch
	var index int64 = 0
	//loop that populates a name for each person in the pool by calling naming function
	for ; index < totalPool; index++{
		names[index] = naming(index)
	}

	//function that randomly picks names and returns picked names in string array
	fmt.Println(picker(totalPool, totalPicks,names))
}

//prompts user for quantity of selection pool
func pool () int64 {
	fmt.Printf("How many people are in the pool? ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(),10,64) //convert string to base 10, 64 bit integer

	return input
}

//prompts user for number of people to be selected by the totalPool
func picks () int64 {
	fmt.Printf("How many people should be picked by the pool? ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(),10,64) //convert string to base 10, 64 bit integer

	return input
}

//requests user for name, then returns that name
func naming (x int64) string{
	x = x+1 // this eliminates position 0 being requested
	fmt.Printf("Enter the name of person " + strconv.FormatInt(x,10) + ":" )
	scanner.Scan()
	input := scanner.Text()
	return input //returns name x
}

//randomly picks choice
func picker(total int64,numberOfPicks int64,nombres [] string )[]string{
	winners := make([]string, numberOfPicks) //dynamically sized array used to store names of those randomly chosen

		var index int64 = 0
		for ; index < numberOfPicks ; index++{
			winners[index] = nombres [rand.Int63n(total)]

			//this subloop is used to ensure no duplicate picks
			var subindex int64 = 0
			for ; subindex < numberOfPicks ; subindex++{
				if index != subindex{ //avoids checking the same position in the array against one another
					if winners[index] == winners[subindex]{ //this will check every existing pick in the array against the new pick
						index-- //this effectively "redraws" the pick
					}
				}
			}
		}
	return winners
}
