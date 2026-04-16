package element

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	// TODO
	return false
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	// TODO
	return 0
}

// Next liefert das Nachfolger-Element.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Next() *Element {
	// TODO
	return nil
}

// SetValue setzt den Wert des Elements.
// Falls das Element bisher leer war, wird es mit dem gegebenen Wert initialisiert.
func (e *Element) SetValue(value int) {
	// TODO
}
