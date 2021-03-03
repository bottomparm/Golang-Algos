package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// fmt.Println(defangIPaddr("1.1.1.1"))
	// fmt.Println(runningSum([]int {1, 2, 3, 4}))
	// fmt.Println(isPalindrome(1212))
	// fmt.Println(isValid("()()()(){}]"))
	// fmt.Println(kidsWithCandies([]int {2,3,5,1,3}, 3))
	// fmt.Println(searchInsert([]int{1, 2, 3, 5}, 3))
	// fmt.Println(fib(100))

	// var root *TreeNode
	// var rootRight *TreeNode
	// var rootLeft *TreeNode
	// rootRight.Val = 10
	// rootLeft.Val = 12
	// root.Val = 8
	// root.Right = rootRight
	// root.Left = rootLeft
	// fmt.Println(levelOrder(root))

	// islandGrid := [][]byte{
	// 	{48, 49, 49, 49, 48},
	// 	{48, 49, 49, 49, 48},
	// 	{48, 48, 48, 48, 48},
	// 	{48, 49, 48, 49, 48},
	// 	{48, 49, 48, 49, 48},
	// }
	// fmt.Println(numIslands(islandGrid))
	
	// fmt.Println(shuffle([]int{1,2,3,4,4,3,2,1}, 4))
	// fmt.Println(diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
}

// Defanging an IP Address
func defangIPaddr(address string) string {
	replacer := strings.NewReplacer(".", "[.]")
	defanged := replacer.Replace(address)
	return defanged
}

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
func runningSum(nums []int) []int {
	currSum := 0
	var runningSums []int
	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		runningSums = append(runningSums, currSum)
	}
	return runningSums
}

// Given an integer x, return true if x is palindrome integer.
func isPalindrome(x int) bool {
	y := strconv.Itoa(x)
	for i, j := 0, len(y)-1; i < j; i, j = i+1, j-1 {
		// how to print a string in ASCII
		fmt.Println(string(y[i]))
		head := y[i]
		tail := y[j]
		if head != tail {
			return false
		}
	}
	return true
}

// check for a valid parens/bracket/curly brace layout.
func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++ {
		if len(stack) < 1 {
			if string(s[i]) == "}" || string(s[i]) == ")" || string(s[i]) == "]" {
				return false
			}
			stack = append(stack, string(s[i]))
		} else {
			if string(s[i]) == "[" || string(s[i]) == "{" || string(s[i]) == "(" {
				stack = append(stack, string(s[i]))
			} else {
				if string(s[i]) == ")" {
					if stack[len(stack)-1] != "(" {
						return false
					}
					stack = stack[:len(stack)-1]
				} else if string(s[i]) == "}" {
					if stack[len(stack)-1] != "{" {
						return false
					}
					stack = stack[:len(stack)-1]
				} else if string(s[i]) == "]" {
					if stack[len(stack)-1] != "[" {
						return false
					}
					stack = stack[:len(stack)-1]
				}
			}

		}
	}
	if len(stack) < 1 {
		return true
	}
	return false
}

// Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
// For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// find largest number
	bigNum := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > bigNum {
			bigNum = candies[i]
		}
	}
	// generate result array with false default value
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= bigNum {
			res[i] = true
		}
	}
	return res
}

// Given a sorted array of distinct integers and a target value, return the index if the target is // found. If not, return the index where it would be if it were inserted in order.
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i], target)
		if target < nums[i] || target == nums[i] {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		}
	}
	return -1
}

// Given n, calculate F(n).
func fib(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i < 2 {
			f = i
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue = append(queue, root)
	res = append(res, []int{queue[0].Val})

	for i := 0; len(queue) > 0; i++ {
		level := []*TreeNode{}
		arrOfInts := []int{}

		for j := 0; j < len(queue); j++ {
			if queue[j].Left != nil {
				left := &queue[j].Left.Val
				arrOfInts = append(arrOfInts, *left)
				level = append(level, queue[j].Left)
			}
			if queue[j].Right != nil {
				right := &queue[j].Right.Val
				arrOfInts = append(arrOfInts, *right)
				level = append(level, queue[j].Right)
			}
		}

		if len(level) > 0 {
			res = append(res, arrOfInts)
		}
		queue = level
	}
	return res
}

// Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
func numIslands(grid [][]byte) int {
	islands := 0
	checked := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		checked[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if checked[i][j] == true {
				continue
			}
			checked[i][j] = true
			if string(grid[i][j]) == "1" {
				islands++
				checked = islandCheck(grid, checked, i, j)
			}
		}
	}
	return islands
}

func islandCheck(grid [][]byte, checked [][]bool, i int, j int) [][]bool {
	idxsToCheck := []int{i, j}
	for keepGoing := true; keepGoing; keepGoing = len(idxsToCheck) > 0 {
		top := ""
		bottom := ""
		left := ""
		right := ""

		i = idxsToCheck[0]
		j = idxsToCheck[1]

		if i != 0 {
			top = string(grid[i-1][j])
			if top == "1" && checked[i-1][j] == false {
				idxsToCheck = append(idxsToCheck, i-1, j)
			}
			checked[i-1][j] = true
		}
		if i != len(grid)-1 {
			bottom = string(grid[i+1][j])
			if bottom == "1" && checked[i+1][j] == false {
				idxsToCheck = append(idxsToCheck, i+1, j)
			}
			checked[i+1][j] = true
		}
		if j != 0 {
			left = string(grid[i][j-1])
			if left == "1" && checked[i][j-1] == false {
				idxsToCheck = append(idxsToCheck, i, j-1)
			}
			checked[i][j-1] = true
		}
		if j != len(grid[0])-1 {
			right = string(grid[i][j+1])
			if right == "1" && checked[i][j+1] == false {
				idxsToCheck = append(idxsToCheck, i, j+1)
			}
			checked[i][j+1] = true
		}

		idxsToCheck = idxsToCheck[2:]
	}

	return checked
}

// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
func shuffle(nums []int, n int) []int {
	res := []int{}
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		res = append(res, nums[i], nums[j])
	}
	return res
}

// A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost row
// or leftmost column and going in the bottom-right direction until reaching the matrix's end.
// For example, the matrix diagonal starting from mat[2][0], where mat is a 6 x 3 matrix,
// includes cells mat[2][0], mat[3][1], and mat[4][2].

// Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return
// the resulting matrix.
func diagonalSort(mat [][]int) [][]int {
	x := 0
	y := 0

	// sort main diagonal and diagonals below main diagonal
	for i := 0; i < len(mat); i++ {
		x = i
		y = 0
		diag := getDiag(mat, x, y)
		sort.Ints(diag)
		mat = updateMatrix(mat, diag, x, y)
	}

	// sort diagonals above main diagonal
	for j := 1; j < len(mat[0]); j++ {
		x = 0
		y = j
		diag := getDiag(mat, x, y)
		sort.Ints(diag)
		mat = updateMatrix(mat, diag, x, y)
	}

	return mat
}

func getDiag(mat [][]int, x int, y int) []int {
	diag := []int{}
	for keepGoing := true; keepGoing; keepGoing = x < len(mat) && y < len(mat[0]) {
		diag = append(diag, mat[x][y])
		x++
		y++
	}
	return diag
}

func updateMatrix(mat [][]int, diag []int, x int, y int) [][]int {
	for i := 0; i < len(diag); i++ {
		mat[x][y] = diag[i]
		x++
		y++
	}
	return mat
}
