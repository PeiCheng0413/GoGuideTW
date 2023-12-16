package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func draw(t tarot, num int) (tarot, tarot) {
	return t[:num], t[num:]
}

func (t tarot) toString() string {
	return strings.Join([]string(t), ",")
}

func (t tarot) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(t.toString()), 0666)
}

func newDeckFromFile(filename string) tarot {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return tarot(s)
}

func (t tarot) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range t {
		newPosition := r.Intn(len(t) - 1)

		t[i], t[newPosition] = t[newPosition], t[i]
	}
}
