package main

func findCircleNum(M [][]int) int {
	var visited = make([]int, len(M))
	count := 0
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(M, visited, i)
			count++
		}
	}
	return count
}

func dfs(M [][]int, visited []int, i int) {
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j]
		}
	}
}

// func main() {
// 	list1 := []int{4, 2, 43, 82, 54, 30, 11, 23, 96, 74, 64, 77, 25, 34, 26, 13, 21, 50, 38, 14, 58, 17, 65, 27, 84, 12, 55, 92, 57, 69, 72, 79, 46, 52, 29, 78, 63, 16, 59, 33, 80, 1, 86, 44, 31, 88, 99, 67, 49, 60, 37, 91, 87, 15, 32, 61, 56, 90, 42, 62, 95, 51, 20, 22, 98, 85, 71, 75, 9, 97, 5, 24, 81, 28, 41, 48, 40, 0, 94, 10, 45, 39, 73, 89, 76, 6, 18, 47, 3, 93, 70, 19, 35, 8, 7, 53, 83, 66, 68, 36}
// 	isConneted := make([][]int, len(list1))

// 	for i := range isConneted {
// 		isConneted[i] = make([]int, len(list1))
// 	}

// 	for i, j := range list1 {
// 		isConneted[i][j] = 1
// 		isConneted[j][i] = 1
// 	}

// 	fmt.Println(isConneted)

// }
