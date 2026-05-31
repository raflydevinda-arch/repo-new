package main
import "fmt"

func main () {
	var n,int
	var a []int
	var target int
	
}

BINARY-SEARCH-ITERATIVE(A, n, target) {
  lo := 0
  hi := n - 1

  while lo <= hi
    mid = (lo + hi) / 2   // integer division

    if A[mid] == target
      return mid           // found! return index

    else if A[mid] < target
      lo = mid + 1         // search right half

    else
      hi = mid - 1         // search left half

  return -1              // not found
}