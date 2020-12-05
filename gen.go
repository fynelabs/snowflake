//+build ignore

package main

import (
	"bufio"
	"bytes"
	"image/png"
	"io/ioutil"
	"os"

	"fyne.io/fyne"
)

func main() {
	pix := captureSnowflake(fyne.NewSize(512, 512), true)
	buf := &bytes.Buffer{}
	write := bufio.NewWriter(buf)
	err := png.Encode(write, pix)
	if err != nil {
		fyne.LogError("Unable to encode icon", err)
	}
	_ = write.Flush()

	fileName := "Icon.png"
	_ = os.Remove(fileName)
	err = ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		fyne.LogError("Unable to create icon file", err)
	}
}
