package main

import "fmt"

type tarot []string

func newTarot() tarot {
	tarotSuits := []string{"金幣", "權杖", "聖杯", "寶劍"}
	tarotValues := []string{"王牌", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	var t tarot
	for _, suit := range tarotSuits {
		for _, value := range tarotValues {
			t = append(t, suit+value)
		}
	}
	return t
}

func (ta tarot) print() {
	for i, t := range ta {
		fmt.Println(i, t)
	}
}
