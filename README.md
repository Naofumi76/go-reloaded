
# Go Reloaded Project

  

## Overview

  

This project provides a Go application that processes a text file, applies various transformations, and writes the result to a new file. The application consists of two main Go files: `main.go` and `go-reloaded.go`.

  

Only the basic go packages were allowed for this project.

  

## Project Structure

  

-  `main.go`: This is the entry point of the application. It handles command-line arguments, reads input from a file, processes the file content using the functions defined in `go-reloaded.go`, and writes the output to another file.

-  `go-reloaded.go`: This file contains the core logic for processing the file content. It includes functions for reading the file, applying transformations, and handling various text manipulations.

  

## Features

  

-  **Text Transformation Commands**:

-  `(hex)` - Converts a preceding hexadecimal number to decimal.

-  `(bin)` - Converts a preceding binary number to decimal.

-  `(up)` - Converts the preceding word to uppercase.

-  `(up, n)` - Converts the preceding `n` words to uppercase.

-  `(cap)` - Capitalizes the preceding word.

-  `(cap, n)` - Capitalizes the preceding `n` words.

-  `(low)` - Converts the preceding word to lowercase.

-  `(low, n)` - Converts the preceding `n` words to lowercase.

-  **Punctuation Handling**: Proper handling of punctuation and contractions to ensure grammatical correctness.

  

## Usage

  

### Prerequisites

  

- Go 1.15+ installed on your machine (if using directly the program with the 'main.go' file).

  

### Building the Project

/i\ The project is already built if you want to skip this step, just use the 'go-reloaded' executable.

  
  

To build the project, navigate to the project directory and run:

  

```sh

go  build  -o  go-reloaded  main.go  go-reloaded.go

```

  

### Running the Application

  

The application takes two command-line arguments:

  

1. Path to the input file.

2. Path to the output file.

  

Example usage:

  

```sh

./go-reloaded  input.txt  output.txt

```

  

### Input File Format

  

The input file should contain text with embedded commands for transformations as described in the features section.

  
  

## Functions

  

main.go

  

- main: The entry point function that reads input file path and output file path from command-line arguments, processes the file, and writes the result to the output file.

  

go-reloaded.go

  

- fileSelect: Reads the content of the specified file and passes it for transformation.

- goReloaded: Applies all transformations to the content and prepares the final output.

- applyChanges: Applies text transformation commands like (up), (cap), (low), (hex), and (bin).

- applyPonct: Adjusts punctuation and handles contractions.

- removeCommand: Removes transformation commands from the processed text.

- SplitWhiteSpaces: Splits the input string by white spaces while respecting parentheses.

- ToLower, ToUpper, Capitalize: Helper functions to change the case of text.

- IsAlpha: Checks if a character is alphanumeric.

- AtoiBase: Converts a string from a given base to an integer.

- iterativePower: Computes the power of a number iteratively.

- Atoi: Converts a string to an integer.

  

### Error Handling

  

The application includes basic error handling for file operations. If an error occurs while reading or writing files, it will print the error message to the console.

  

This project is licensed under the MIT License.
