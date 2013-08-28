// http://www.climagic.org/mirrors/VT100_Escape_Codes.html

package board

import (
	"fmt"
	"time"
)

const (
	ASCII_ESC    = 27
	BOARD_WIDTH  = 120
	BOARD_HEIGHT = 30
)

func ClearScreen() {

	fmt.Printf("%c[2J", ASCII_ESC)
}

func SetCommandPrompt() {

	fmt.Printf("%c[%d;0H>:", ASCII_ESC, (BOARD_HEIGHT + 1))
	fmt.Printf("%c[K", ASCII_ESC)
}

func MoveCursor(v int, h int) {

	fmt.Printf("%c[%d;%dH", ASCII_ESC, v, h)
}

func PrintChar(v int, h int, ch string) {

	fmt.Printf("%c[%d;%dH%s", ASCII_ESC, v, h, ch)
	fmt.Printf("%c[%d;3H", ASCII_ESC, (BOARD_HEIGHT + 1))
}

func PrintHorizontalBorder(v int, h int, total int) {

	fmt.Printf("%c[%d;%dH", ASCII_ESC, v, h)

	for index := 0; index < total; index++ {
		fmt.Print("=")
	}
}

func PrintVeritcalBorder(v int, h int, total int) {

	total = total + v
	for line := v; line < total; line++ {

		fmt.Printf("%c[%d;%dH=", ASCII_ESC, line, h)
	}
}

func PrintDiagForward(v int, h int, total int) {

	v1 := v
	h1 := h

	for line := 0; line < total; line++ {

		fmt.Printf("%c[%d;%dH=", ASCII_ESC, v1, h1)

		v1 = v1 + 1
		h1 = h1 - 2
	}
}

func PrintDiagBackward(v int, h int, total int) {

	v1 := v
	h1 := h

	for line := 0; line < total; line++ {

		fmt.Printf("%c[%d;%dH=", ASCII_ESC, v1, h1)

		v1 = v1 + 1
		h1 = h1 + 2
	}
}

func DrawGameBoard() {

	DrawOuterFrame()
	DrawFrontScreen()
	DrawCompass()
	DrawGunnerBox()
}

func DrawOuterFrame() {

	ClearScreen()

	PrintHorizontalBorder(0, 0, BOARD_WIDTH)

	PrintVeritcalBorder(0, 0, BOARD_HEIGHT)
	PrintVeritcalBorder(0, BOARD_WIDTH, BOARD_HEIGHT)

	PrintHorizontalBorder(BOARD_HEIGHT, 0, BOARD_WIDTH)

	PrintHorizontalBorder(13, 0, BOARD_WIDTH)

	PrintHorizontalBorder(BOARD_HEIGHT-2, 0, BOARD_WIDTH)
}

func DrawFrontScreen() {

	PrintVeritcalBorder(0, 50, 13)
	PrintVeritcalBorder(0, 80, 13)
}

func DrawGunnerBox() {

	PrintHorizontalBorder(18, 60, 11)
	PrintHorizontalBorder(22, 60, 11)
	PrintVeritcalBorder(19, 60, 4)
	PrintVeritcalBorder(19, 70, 4)
}

func DrawCompass() {

	PrintHorizontalBorder(7, 58, 15)
	PrintVeritcalBorder(4, 65, 7)

	PrintChar(3, 65, "0")
	PrintChar(3, 72, "45")
	PrintChar(7, 74, "90")
	PrintChar(11, 72, "135")
	PrintChar(11, 64, "180")
	PrintChar(11, 56, "225")
	PrintChar(7, 54, "270")
	PrintChar(3, 56, "325")

	PrintDiagBackward(4, 59, 3)
	PrintDiagBackward(8, 67, 3)
	PrintDiagForward(4, 71, 3)
	PrintDiagForward(8, 63, 3)
}

func PhotonShot() {

	v1 := 28
	h1 := 80
	v2 := 28
	h2 := 50

	for line := 0; line < 8; line++ {

		fmt.Printf("%c[%d;%dH=", ASCII_ESC, v1, h1)

		v1 = v1 - 1
		h1 = h1 - 2

		fmt.Printf("%c[%d;%dH=", ASCII_ESC, v2, h2)

		v2 = v2 - 1
		h2 = h2 + 2

		fmt.Printf("%c[%d;3H", ASCII_ESC, (BOARD_HEIGHT + 1))

		time.Sleep(50 * time.Millisecond)
	}

	PrintChar(20, 65, "X")

	time.Sleep(50 * time.Millisecond)

	v1 = 28
	h1 = 80
	v2 = 28
	h2 = 50

	for line := 0; line < 8; line++ {

		fmt.Printf("%c[%d;%dH ", ASCII_ESC, v1, h1)

		v1 = v1 - 1
		h1 = h1 - 2

		fmt.Printf("%c[%d;%dH ", ASCII_ESC, v2, h2)

		v2 = v2 - 1
		h2 = h2 + 2

		PrintChar(28, 50, "=")
		PrintChar(28, 80, "=")
		PrintChar(22, 62, "=")
		PrintChar(22, 68, "=")

		fmt.Printf("%c[%d;3H", ASCII_ESC, (BOARD_HEIGHT + 1))

		time.Sleep(50 * time.Millisecond)
	}

	PrintChar(20, 65, " ")
}

func MoveGunnerLeft() {

}
