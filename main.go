package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"kodoku/genetic"
	"kodoku/kodoku"
)

var (
	fileName  = flag.String("file", "", "file name")
	popSize   = flag.Int("population", 100, "population size")
	genSize   = flag.Int("generation", 1000, "generation size")
	iterSize  = flag.Int("iteration", 1000, "selection iteration rate")
	crossRate = flag.Float64("crossover", 0.5, "crossover rate")
	mutRate   = flag.Float64("mutation", 0.1, "mutation rate")
)

func init() {
	flag.Parse()
}

func main() {

	fmt.Print("Welcome to Kodoku, this is a word puzzle solver for a Sudoku-Like puzzle.\n")
	fmt.Print("Please enter the puzzle you want to solve, using CSV standard.\n")
	fmt.Print("You must enter a x by x matrix of characters, where x is the size of the puzzle.\n")
	fmt.Print("You can either leave a value empty or denote it with - for being missing, ex: a, b, -, c \n")
	fmt.Print("You can also enter 'quit' to exit the program.\n")

	var puzzle [][]string
	switch *fileName {
	case "":
		scanner := bufio.NewScanner(os.Stdin)
		var text string
		for {
			fmt.Print("Enter a comma separated list of characters (ex: a, d, c, d): ")

			scanner.Scan()
			scan := scanner.Text()
			if scan == "quit" {
				break
			}

			trim := strings.ReplaceAll(scan, " ", "")
			text += fmt.Sprintf("%s\n", trim)

			puzzle, _ = csv.NewReader(strings.NewReader(text)).ReadAll()
		}
	default:
		f, err := os.Open(*fileName)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		var text string
		for scanner.Scan() {
			trim := strings.ReplaceAll(scanner.Text(), " ", "")
			text += fmt.Sprintf("%s\n", trim)

			puzzle, _ = csv.NewReader(strings.NewReader(text)).ReadAll()
		}
	}

	grid := kodoku.NewGrid(len(puzzle), len(puzzle[0]))
	grid.FillFromCSV(puzzle)
	fmt.Println("\nThe puzzle you entered is:", grid)
	g := genetic.New(1, 4, *popSize, *genSize, *iterSize, *mutRate, *crossRate, grid.Export())
	fmt.Println("Searching for a solution: ")
	solution := g.Solve()
	if solution != nil {
		grid.Import(solution.Export())
		if solution.Fitness() < 1 {
			fmt.Println(fmt.Sprintf("This is an incomplete solution given \n population: %v, selection iteration %v, %v crossrate mutation and %v mutation rate", *genSize, *iterSize, *crossRate, *mutRate))
			fmt.Println("\nThe system returned a solution with a fitness of:", solution.Fitness())
			fmt.Println("\nTo find a better solution try playing with the parameters:")
		}
		fmt.Println("\nThe solution is:", grid)
	} else {
		fmt.Println("No solution found")
	}
}
