package main

type omikuji []string

func newOmikuji() omikuji {
	levels := []string{"大", "中", "小"}
	states := []string{"吉", "平", "凶"}
	omi := omikuji{}

	for _, level := range levels {
		for _, state := range states {
			omi = append(omi, level+state)
		}
	}
	return omi
}
