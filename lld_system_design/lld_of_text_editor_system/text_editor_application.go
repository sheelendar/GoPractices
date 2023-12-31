package main

import "fmt"

/*
Application Requirements
The application should be able to support the following
functionalities:
▪ display(): Display the entire content
▪ display(n,m): Display lines from line n to m
▪ insert(n, text): Insert text at line n
▪ delete(n): Delete line n
▪ delete(n,m): Delete all lines from line n to m
▪ copy(n,m): Copy contents from line n to m to clipboard
▪ paste(n): Paste contents from clipboard to line n
▪ undo(): Undo the last command
▪ redo(): Redo the last command
The application is expected to be in memory(not in file/DB)
*/

func main() {
	fmt.Println("sheelendar")
	notePad := GetNotepadConstractor("Hi my name is Sheelendar.I am doing study at mnnit college. My frienfd name is Ram.He is also doing study in mnnit allahabad.I have another friend sunil.He is very smart in his work")

	notePad.display()
	fmt.Println("************************************************************")

	notePad.displayWithLine(1, 2)
	fmt.Println("************************************************************")

	notePad.insert(1, "Ohh!")
	notePad.display()
	fmt.Println("************************************************************")

	notePad.undo()
	notePad.display()
	fmt.Println("************************************************************")

	notePad.redo()
	notePad.display()
	fmt.Println("************************************************************")

	notePad.delete(2)
	notePad.display()
	fmt.Println("************************************************************")

	notePad.deleteWithLine(2, 3)
	notePad.display()
	fmt.Println("************************************************************")

}
