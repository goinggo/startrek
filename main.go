// http://www.climagic.org/mirrors/VT100_Escape_Codes.html

package main

/*
#cgo CFLAGS: -I../Keyboard/DyLib
#cgo LDFLAGS: -L. -lkeyboard
#include <keyboard.h>
*/
import "C"
import (
	"github.com/goinggo/StarTrek/board"
	"github.com/goinggo/StarTrek/klingonbattle"
)

func main() {
	C.InitKeyboard()

	board.DrawGameBoard()

	klingonbattle.Startup()

	AcceptInput()

	klingonbattle.Shutdown()

	C.CloseKeyboard()
}

func AcceptInput() {
	for {
		board.SetCommandPrompt()

		//left := fmt.Sprintf("%c[D", ASCII_ESC)
		//right := fmt.Sprintf("%c[C", ASCII_ESC)

		command := C.GetCharacter()

		switch command {
		case 'q':
			return

		case 'p':
			board.PhotonShot()
			break

			//case left[0]:
			//	//MoveGunnerLeft()
			//	break

			//case right[0]:
			//	//MoveGunnerRight()
			//	break
		}
	}
}
