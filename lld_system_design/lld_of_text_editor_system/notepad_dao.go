package main

import (
	"fmt"
	"strings"
)

type Notepad struct {
	allContent []string
	undoStack  []string
	redoStack  []string
	buffer     []string
}

func GetNotepadConstractor(text string) *Notepad {
	arr := strings.Split(text, ".")
	return &Notepad{allContent: arr}
}

func (notepad *Notepad) insert(n int, text string) {
	size := len(notepad.allContent)
	if n > size {
		fmt.Println("n line not find in text")
		return
	}
	pushIntoStack(notepad, notepad.allContent)
	line := notepad.allContent[n]
	notepad.allContent[n] = line + " " + text
}

func (notepad *Notepad) display() {
	for _, line := range notepad.allContent {
		fmt.Println(line)
	}
}

func (notepad *Notepad) displayWithLine(n1, n2 int) {
	size := len(notepad.allContent)
	if n1 > size {
		fmt.Println("n1 line not there")
		return
	}
	if n2 > size {
		fmt.Println("n2 line not there")
		return
	}
	if n1 > n2 {
		fmt.Println("n1 can be greater than n1")
		return
	}
	for i := n1 - 1; i < n2; i++ {
		fmt.Println(notepad.allContent[i])
	}
}

func (notepad *Notepad) delete(n int) bool {
	size := len(notepad.allContent)
	if n > size {
		fmt.Println("n1 line not there")
		return false
	}
	pushIntoStack(notepad, notepad.allContent)
	arr1 := notepad.allContent[:n-1]
	arr2 := notepad.allContent[n:size]
	notepad.allContent = nil
	notepad.allContent = append(arr1, arr2...)
	return true
}

func pushIntoStack(notepad *Notepad, s []string) {
	notepad.undoStack = nil
	for _, line := range s {
		notepad.undoStack = append(notepad.undoStack, line)
	}
}
func (notepad *Notepad) deleteWithLine(n1, n2 int) bool {
	size := len(notepad.allContent)
	if n1 > size {
		fmt.Println("n1 line not there")
		return false
	}
	if n2 > size {
		fmt.Println("n2 line not there")
		return false
	}
	if n1 > n2 {
		fmt.Println("n1 can be greater than n1")
		return false
	}
	pushIntoStack(notepad, notepad.allContent)
	arr1 := notepad.allContent[:n1-1]
	arr2 := notepad.allContent[n2:size]
	notepad.allContent = nil
	notepad.allContent = append(arr1, arr2...)
	return true
}

func (notepad *Notepad) copy(n1, n2 int) bool {
	size := len(notepad.allContent)
	if n1 > size {
		fmt.Println("n1 line not there")
		return false
	}
	if n2 > size {
		fmt.Println("n2 line not there")
		return false
	}
	if n1 > n2 {
		fmt.Println("n1 can be greater than n1")
		return false
	}
	for i := n1 - 1; i < n2; i++ {
		notepad.buffer = append(notepad.buffer, notepad.allContent[i])
	}
	return true
}

func (notepad *Notepad) paste(n int) bool {
	size := len(notepad.allContent)
	if n > size {
		fmt.Println("n1 line not there")
		return false
	}
	pushIntoStack(notepad, notepad.allContent)
	arr1 := notepad.allContent[:n]
	arr2 := notepad.allContent[n:size]
	notepad.allContent = nil
	notepad.allContent = append(notepad.allContent, arr1...)
	notepad.allContent = append(notepad.allContent, notepad.buffer...)
	notepad.allContent = append(notepad.allContent, arr2...)
	return true
}
func (notepad *Notepad) undo() bool {
	if notepad.undoStack == nil {
		fmt.Println("Nothing to undo")
		return false
	}
	notepad.redoStack = nil
	notepad.redoStack = append(notepad.redoStack, notepad.allContent...)
	notepad.allContent = notepad.undoStack
	return true
}

func (notepad *Notepad) redo() bool {
	if notepad.redoStack == nil {
		fmt.Println("Nothing to redo")
		return false
	}
	notepad.undoStack = nil
	notepad.undoStack = append(notepad.undoStack, notepad.allContent...)
	notepad.allContent = notepad.redoStack
	return true
}
