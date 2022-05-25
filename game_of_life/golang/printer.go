package main

func printArray(a [][]bool) string {
	s := ""
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == true {
				s = s + "* "
			} else {
				s = s + "  "
			}
		}

		if i < len(a)-1 {
			s = s + "\n"
		}
	}

	return s
}
