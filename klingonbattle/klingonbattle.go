package klingonbattle

import (
	"github.com/goinggo/StarTrek/board"
	"math/rand"
	"time"
)

type klingonBattle struct {
	Shutdown        bool        // Set when the singleton is shutting down
	ShutdownChannel chan string // Channel used to shutdown the work routine
}

type MoveKlingonShip struct {
	KV, KH                     int
	TopLeftV, TopLeftH         int
	BottomRightV, BottomRightH int
}

var _This *klingonBattle // Reference to the singleton

func Startup() (err error) {
	// Create the singleton
	_This = &klingonBattle{
		Shutdown:        false,
		ShutdownChannel: make(chan string),
	}

	rand.Seed(time.Now().UnixNano())

	// Start the go routine
	go _This.GoRoutine_Battle()

	return err
}

func Shutdown() (err error) {
	_This.Shutdown = true

	_This.ShutdownChannel <- "Down"
	<-_This.ShutdownChannel

	close(_This.ShutdownChannel)

	return err
}

func StartBattle() {
}

func EndBattle() {
}

func (this *klingonBattle) GoRoutine_Battle() {
	for {

		select {
		case <-this.ShutdownChannel:
			this.ShutdownChannel <- "Down"
			return

		case <-time.After(250 * time.Millisecond):
			this.MoveKlingonShip()
			break
		}
	}
}

func (this *klingonBattle) MoveKlingonShip() {
	var kV, kH int

	topLeftV := 14
	topLeftH := 3
	bottomRightV := board.BOARD_HEIGHT - 3
	bottomRightH := board.BOARD_WIDTH - 2

	rand.Seed(time.Now().UnixNano())

	for {
		kH = rand.Intn(bottomRightH + 1)

		if kH < topLeftH {
			continue
		}
		break
	}

	for {
		kV = rand.Intn(bottomRightV + 1)

		if kV < topLeftV {
			continue
		}
		break
	}

	direction := rand.Intn(9)
	distance := rand.Intn(10)

	for move := 0; move < distance; move++ {

		switch direction {
		case 0:
			kV = kV - 1
			kH = kH - 1
			break

		case 1:
			kV = kV - 1
			break

		case 2:
			kV = kV - 1
			kH = kH + 1
			break

		case 3:
			kH = kH - 1
			break

		case 5:
			kH = kH + 1
			break

		case 7:
			kV = kV + 1
			kH = kH - 1
			break

		case 8:
			kV = kV + 1
			break

		case 9:
			kV = kV + 1
			kH = kH + 1
			break
		}

		if kV < topLeftV {
			kV = topLeftV
			direction = 8
		} else if kV > bottomRightV {
			kV = bottomRightV
			direction = 2
		}

		if kH < topLeftH {
			kH = topLeftH
			direction = 6
		} else if kH > bottomRightH {
			kH = bottomRightH
			direction = 3
		}

		board.PrintChar(kV, kH, "K")
		time.Sleep(300 * time.Millisecond)
		board.PrintChar(kV, kH, " ")
	}
}
