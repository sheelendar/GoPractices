package example2

type Mackbook struct {
	KeyBoard KeyBoard
	Mouse    Mouse
}

// Keyboards --------------------------------------|

// WiredKeyBoard Setup
type WiredKeyBoard struct {
}

func (wkb WiredKeyBoard) PressButton() {

}

// WireLessKeyBoard Setup
type WireLessKeyBoard struct {
}

func (wlkb WireLessKeyBoard) PressButton() {

}

// Keyboards --------------------------------------|

// Mouse ------------------------------------------|

// WiredMouse setup
type WiredMouse struct {
}

func (wm WiredMouse) ClickButton() {

}

// WireLessMouse setup
type WireLessMouse struct {
}

func (wlm WireLessMouse) ClickButton() {

}

// Mouse ------------------------------------------|
