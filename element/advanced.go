package element

import "fmt"

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	if e.IsEmpty() == true {
		e.SetValue(value)
	} else {
		e.next.Append(value)
	}

}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (e *Element) Length() int {
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.next.Length()
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (e *Element) Contains(value int) bool {
	if e.IsEmpty() {
		return false
	}
	if e.value == value {
		return true
	}
	return e.next.Contains(value)
}

// Count gibt die Anzahl der Elemente in der Liste zurück, die den gegebenen Wert enthalten.
func (e *Element) Count(value int) int {
	if e.IsEmpty() {
		return 0
	}

	if e.value == value {
		return e.next.Count(value) + 1
	} else {
		return e.next.Count(value)
	}

}

// Sum berechnet die Summe der Werte aller Elemente in der Liste.
func (e *Element) Sum() int {
	if e.IsEmpty() {
		return 0
	}
	return e.next.Sum() + e.value
}

// Min gibt den kleinsten Wert aller Elemente in der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Min() int {
	if e.IsEmpty() {
		panic("List is empty")

	}
	return e.Mini()
}

func (e *Element) Mini() int {
	if e.IsEmpty() {
		return 10000000000
	}
	if e.value < e.next.Mini() {
		return e.value
	} else {
		return e.next.Mini()
	}

}

// Last gibt das letzte Element der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Last() *Element {
	if e.IsEmpty() {
		panic("List is empty")
	}
	if e.next.IsEmpty() {
		return e
	}
	return e.next.Last()
}

// At gibt das Element an der gegebenen Position zurück.
// Falls die Position nicht existiert, wird eine panic ausgelöst.
func (e *Element) At(position int) *Element {

	if e.IsEmpty() {
		panic("index out of bounds")
	}
	if position == 0 {
		return e
	}

	return e.next.At(position - 1)
}

// String gibt eine textuelle Repräsentation der Liste zurück.
// Die Elemente werden durch " -> " getrennt.
// Falls die Liste leer ist, wird ein leerer String zurückgegeben.
func (e *Element) String() string {
	if e.IsEmpty() {
		return ""
	}
	if e.next.IsEmpty() {
		return fmt.Sprintf("%d", e.value)
	}
	return fmt.Sprintf("%d", e.value) + " -> " + e.next.String()
}

// Swap vertauscht die beiden Elemente an den Stellen i und j.
// Falls eine der Positionen nicht existiert, wird eine panic ausgelöst.
// Die Funktion liefert den neuen Kopf der Liste zurück.

// Es sollen nicht die Werte der Elemente vertauscht werden, sondern die Elemente selbst.
func (e *Element) Swap(i, j int) *Element {

	if i >= j {
		if i == 0 {
			return 
		}
		if j == 1 {
			a := e.next
			b := e.next.next
			e.next = b
			e.next.next = a
			return e.next.Swap(i-1, j)

		}

	}

	if j >= i {
		for a := j; a >= 0; a -- {

		

		
		}

	}

}
