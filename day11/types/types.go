package types

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"sync"
)

type MonkeyOrchestrator struct {
	Items      [][]int
	Inspected  []int
	Operations []func(v int) int
	Tests      []func(v int) int
}

func (m *MonkeyOrchestrator) Begin(rounds int) {
	// for each round
	for r := 0; r < rounds; r++ {
		// for each monkey
		for i, items := range m.Items {
			// go through their items and toss them as needed
			for _, v := range items {
				m.Inspected[i]++
				worry := int(math.Floor(float64(m.Operations[i](v)) / 3))
				fmt.Println(worry)
				to := m.Tests[i](worry)
				fmt.Println(to)
				m.Items[to] = append(m.Items[to], worry)
			}
			m.Items[i] = []int{} // empty the list when finished
		}
	}
}

func (m *MonkeyOrchestrator) Score() int {

	sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
	fmt.Println(m.Inspected)
	return m.Inspected[0] * m.Inspected[1]
}

type MonkeyOrchestrator2 struct {
	Items      [][]*big.Int
	Inspected  []int
	Operations []func(v *big.Int) *big.Int
	Tests      []func(v *big.Int) int
}

func (m *MonkeyOrchestrator2) Begin(rounds int) {
	// for each round
	for r := 0; r < rounds; r++ {
		// for each monkey
		for i, items := range m.Items {

			var wg sync.WaitGroup
			wg.Add(len(items))
			// go through their items and toss them as needed
			for _, value := range items {
				m.Inspected[i]++
				go func(v *big.Int) {
					defer wg.Done()
					worry := m.Operations[i](v)
					// fmt.Println(worry)
					to := m.Tests[i](worry)
					// fmt.Println(to)
					m.Items[to] = append(m.Items[to], worry)
				}(value)
			}
			wg.Wait()
			m.Items[i] = []*big.Int{} // empty the list when finished
		}
	}
}

func (m *MonkeyOrchestrator2) Score() int {

	sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
	fmt.Println(m.Inspected)
	return m.Inspected[0] * m.Inspected[1]
}
