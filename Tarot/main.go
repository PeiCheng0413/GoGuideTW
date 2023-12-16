package main

func main() {
	ta := newTarot()
	ta.print()
	ta.shuffle()
	tadraw, _ := draw(ta, 4)
	tadraw.print()
}
