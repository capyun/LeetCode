package main

/**
@author Shitanford
@create 2021-05-29 18:16
*/

func main() {

}

func minDistance(word1 string, word2 string) int {
	distance := make([][]int, len(word1)+1)
	for i:=0; i<len(distance); i++ {
		distance[i] = make([]int, len(word2)+1)
		distance[i][0] = i
		distance[0][i] = i
	}
	for i:=1; i<len(word1)+1; i++ {
		for j:=1; j<len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				distance[i][j] = distance[i-1][j-1]
			} else {
				distance[i][j] = distance[i-1][j-1]
				if distance[i][j] < distance[i-1][j] {
					distance[i][j] = distance[i-1][j]
				}
				if distance[i][j] < distance[i][j-1] {
					distance[i][j] = distance[i][j-1]
				}
				distance[i][j]++
			}
		}
	}
	return distance[len(word1)][len(word2)]
}