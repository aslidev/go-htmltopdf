package main

import (
	"flag"
	"fmt"

	"github.com/jackdon/go-htmltopdf/image"
)

func main() {
	var in = flag.String("in", "", "input html file path")
	var output = flag.String("output", "sample.png", "output html file path")
	var quality = flag.String("quality", "94", "set the output image quality: 0 - 100, default 94")
	var width = flag.String("width", "200", "set the output image width")
	var height = flag.String("height", "100", "set the output image height")
	flag.Parse()
	imgGen := image.New()
	imgGen.OnProgressChanged(func(a int) { fmt.Println("Progress::", a, "%") })

	imgGen.OnPhaseChanged(func(msg string) { fmt.Println("Phase::", msg) })

	imgGen.OnError(func(msg string) { fmt.Println("Error::", msg) })

	imgGen.OnWarning(func(msg string) { fmt.Println("Warning::", msg) })
	imgGen.SetGlobalSettings([][2]string{
		{"screenWidth", *width},
		{"screenHeight", *height},
		{"fmt", "png"},
		{"out", *output},
		{"in", *in},
		{"quality", *quality},
	})
	err, _ := imgGen.CreateImage("")
	if err != nil {
		fmt.Println(err)
	}
}
