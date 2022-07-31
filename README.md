# Kodoku

## What is kodoku?
Kodoku is a derivative term of sudoku, which in Japanese stands for "suuji wa dokushin ni kagiru". This phrase means “the numbers (or digits) must remain single.” (Dakowski, 2022) To achieve a sense of uniqueness and make our project stand out, we chose to exchange the word “suuji” (numbers) for “kotoba” (words), which gave birth to the name: "Kodoku", by using the same derivation method that was used to create the name Sudoku.

The goal of any round of Sudoku is to find a solution to a given 4x4 Sudoku puzzle that has 0 duplicate numbers (or in our case letters) in its columns, rows, or 2x2 sub-grids.

## What are genetic software systems?
Genetic software systems use principles commonly associated with the biology field of genetics. This practice allows for coding information into strings of bits, which can then be combined, recombined, and decoded in ways that mimic rapid evolution in nature. Some operators used in genetic algorithms include crossover, which cuts the two parents at a certain point in the string, and then exchanges the 'tails' of remaining bits after that point to create two children.), and mutation, which is an operator that has a small chance to activate after the crossover operator has finished, to encourage biodiversity in subsequent generations/populations. (Gupta, 2022)

## What is a genetic algorithm?
Genetic algorithms have been around for a long time; however, a definition of a genetic algorithm may be necessary to fully comprehend and explore their uses and the application of this approach in our project. Genetic algorithms are a type of (meta) heuristic that is modelled after natural evolution in nature. Scientifically, they are employed to generate high-quality solutions, based on iterative generational populations of candidate solutions that are assigned fitness values, and selected for procreation accordingly. Unique operators including crossover and mutation are used to formulate novel solutions/genes over a pre-determined number of generations, or the process could be terminated if the pre-set number of solutions have been found. 

## User Input
There are two different input mechanisms integrated into the software, one is through a REPL by the terminal, using a CSV standard, and the second method is through a CSV file that you can define upon initialization of the application through a Command Line Flag (-file [filename]).

## Initial Configuration
The software package uses flags to assign internally the right arguments for the Genetic Algorithm. As we know, beforehand we must define Generation Size, Population Size, and for some selection algorithms, iteration size, we also can define probability variables for the mutation of the genes, all these configuration options can be accessed by using the flag -help from the binary file or if you have Go installed, you can use ‘go run ./main.go -help’.

## Genetic Algorithm
A grid of rows and tiles is generated, and the user is prompted to input the initial population of numbers and "lock them in" to their respective slots on the grid. The program will check each row, each column, and each sub-grid of 2x2, to match the remaining fields with suitable letters.
The program will then examine the resulting grid for duplicates, which will then be considered when calculating the fitness score of each gene/solution and putting it into the population.

Suitable parents are selected from the (current) population, and crossover/mutation operators are applied to generate two new children who can be assessed and put into the new population.
This process is repeated until a suitable solution is found to the puzzle, or the maximum number of generations has been met.

## Fitness Function
A fitness value will be assigned based on the number of duplicates in each row, column, and sub-grid. That value is then divided by 12 to produce a decimal value to be taken and used as the fitness value. The closer this value gets to 1, the higher the fitness value is, and the closer it is to produce a suitable solution.
