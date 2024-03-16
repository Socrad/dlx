Dancing Links & Alorythm X 
==================================================================

Introduce
=========
This package provides solutions to the exact cover problem using Dancing Links and Algorythm X by Donald E. Knuth, Stanford University.

Usage
=====
dlx.Initialize(matrix [][]bool, columnNames []string) - Prepares nodes and provides header.
dlx.SearchFunction(getAll bool) [][]dlxNode - Provides search function. If getAll is true, function will search all solutions, or function will get a solution. 
dlx.ResolveSolutions(solutions [][]dlxNode) [][][]string - Converts solutions to a string slices, each string are one of columnNames.

Example
=======
package main

import "github.com/Socrad/dlx"

func main() {
  matrix := [][]bool{
		{false, false, true},
		{true, false, true},
		{false, true, false},
		{true, false, false},
	}
	columnNames := []string{"hat", "cap", "vinnie"}

  header, err := dlx.Initialize(matrix, columnNames)
  
  getAll := true
  search := dlx.SearchFunction(getAll)
  solutions := search(header)
  result := dlx.ResolveSolutions(&solutions)

  // ...
}



