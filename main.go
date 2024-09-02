package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"sdvx-autoscener/pkg/constant"
	"sdvx-autoscener/pkg/utils"

	"github.com/andreykaipov/goobs"
	"github.com/kbinani/screenshot"
)

func main() {
	n := screenshot.NumActiveDisplays()
	fmt.Println("Number of ACTIVE_DISPLAY is", n)

	bounds := screenshot.GetDisplayBounds(constant.ACTIVE_DISPLAY)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	// fileName := fmt.Sprintf("%d_%dx%d.png", constant.ACTIVE_DISPLAY, bounds.Dx(), bounds.Dy())

	file, _ := os.Create(filepath.Join(constant.OUTPUT_PATH, "screen.png"))

	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("#%d : %v \"%s\"\n", constant.ACTIVE_DISPLAY, bounds, "screen.png")

	client, err := goobs.New(constant.OBS_WS_URL, goobs.WithPassword(constant.OBS_WS_PASSWORD))
	if err != nil {
		log.Fatalf("Error connecting to OBS WebSocket: %v", err)
		panic(err)
	}
	defer client.Disconnect()

	version, err := client.General.GetVersion()
	if err != nil {
		panic(err)
	}

	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Server protocol version: %s\n", version.ObsWebSocketVersion)
	fmt.Printf("Client protocol version: %s\n", goobs.ProtocolVersion)
	fmt.Printf("Client library version: %s\n", goobs.LibraryVersion)

	imgDetect := utils.ContainsObject()
	fmt.Println("Result should be true: ", imgDetect)
}
