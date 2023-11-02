package main

func main() {
	println(countPoints("G4"))
}
func countPoints(rings string) int {
	mapping := make(map[uint8]int)
	mapping['R'] = 1
	mapping['G'] = 2
	mapping['B'] = 4
	ring := make([]int, 10)
	for i := 0; i < len(rings); i = i + 2 {
		color := rings[i]
		ringNo := (int)(rings[i+1] - '0')
		ring[ringNo] = ring[ringNo] | mapping[color]
	}
	res := 0
	for _, i2 := range ring {
		if i2 == 1+2+4 {
			res++
		}
	}
	return res
}
