package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Creating file
	file, createError := os.Create("file.txt")

	if createError != nil {
		panic(createError)
	}

	// fileSize, writeError := file.WriteString("Hello file") // Escrevendo texto
	fileSize, writeError := file.Write([]byte("File bytes"))

	if writeError != nil {
		panic(writeError)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho de %d bytes\n", fileSize)

	file.Close()

	// Reading file

	fileToRead, readError := os.ReadFile("file.txt")

	if readError != nil {
		panic(readError)
	}

	fmt.Println(string(fileToRead))

	// Reading file parts

	filePart, partError := os.Open("file.txt")

	if partError != nil {
		panic(partError)
	}

	reader := bufio.NewReader(filePart)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	removeError := os.Remove("file.txt")

	if removeError != nil {
		panic(removeError)
	}
}
