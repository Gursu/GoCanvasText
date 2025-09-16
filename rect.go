package GoCanvasText

import "fmt"

func FillRect(w int, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print("@")
		}
		fmt.Print("\n")
	}
}
