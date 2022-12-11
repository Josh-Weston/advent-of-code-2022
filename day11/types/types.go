package types

import (
	"fmt"
	"math"
	"math/big"
	"sort"
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

	fmt.Println(m.Inspected)
	sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
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

			/*
				var wg sync.WaitGroup
				wg.Add(len(items))
				// go through their items and toss them as needed
				for _, value := range items {
					m.Inspected[i]++
					go func(v *big.Int) {
						defer wg.Done()
						worry := m.Operations[i](v)
						to := m.Tests[i](worry)
						// TODO: yes, I need to make this thread safe (which may or may not be the cause of my problem)
						// TODO: this takes a really longtime to run in general (even with threading)
						m.Items[to] = append(m.Items[to], worry) // does this need to be thread safe?
					}(value)
				}
				// reset our worry after every round, for those items?
				wg.Wait()
				m.Items[i] = []*big.Int{} // empty the list when finished
			*/
			for _, v := range items {
				m.Inspected[i]++
				worry := m.Operations[i](v)
				to := m.Tests[i](worry)
				// TODO: yes, I need to make this thread safe (which may or may not be the cause of my problem)
				// TODO: this takes a really longtime to run in general (even with threading)
				m.Items[to] = append(m.Items[to], worry) // does this need to be thread safe?
			}
			m.Items[i] = []*big.Int{} // empty the list when finished

		}
	}
}

// the order of appending should not matter, just as long as they are appended!

func (m *MonkeyOrchestrator2) Score() int {

	fmt.Println(m.Inspected)
	sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
	return m.Inspected[0] * m.Inspected[1]
}

// Note: we were using / 3 to keep the numbers fairly reasonable in the first 20 rounds of testing
// Now, I need a new approach, THAT DOES NOT LET THE NUMBERS CONTINUE TO CLIMB
type Item struct {
	Original int
	Worry    int
}

type MonkeyOrchestrator3 struct {
	Items      [][]*Item
	Inspected  []int
	Operations []func(v *Item) int
	Tests      []func(v int) int
}

func (m *MonkeyOrchestrator3) Begin(rounds int) {
	// for each round
	for r := 0; r < rounds; r++ {
		// for each monkey
		for i, items := range m.Items {
			// go through their items and toss them as needed
			for _, v := range items {
				m.Inspected[i]++
				worry := m.Operations[i](v)
				to := m.Tests[i](worry) // this unchanged worry value causes too many to go to monkey 2
				m.Items[to] = append(m.Items[to], v)
			}
			m.Items[i] = []*Item{} // empty the list when finished
		}
		// reset the worry values to their original at the end of each round
		for _, items := range m.Items {
			for _, v := range items {
				v.Worry = v.Original
			}
		}

		fmt.Println(m.Items[0][0].Original, m.Items[0][0].Worry)
		fmt.Printf("Round: %d\n", r)
		fmt.Printf("%+v\n", m.Items)
	}
}

func (m *MonkeyOrchestrator3) Score() int {
	fmt.Println(m.Inspected)
	// sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
	// fmt.Println(m.Inspected)
	return m.Inspected[0] * m.Inspected[1]
}

type MonkeyOrchestrator2C struct {
	Items      [][]int
	Inspected  []int
	Operations []func(v int) int
	Tests      []func(v int) int
}

func (m *MonkeyOrchestrator2C) Begin(rounds int) {
	// for each round
	for r := 0; r < rounds; r++ {
		// for each monkey
		for i, items := range m.Items {
			// go through their items and toss them as needed
			for _, v := range items {
				m.Inspected[i]++
				// worry := int(math.Floor(math.Sqrt(float64(m.Operations[i](v)))))
				worry := m.Operations[i](v)
				// this was a good idea, but it doesn't work
				// if worry%96577 == 0 {
				// 	fmt.Println("here")
				// worry = worry % 96577 // for tests
				worry = worry % 9_699_690 // the LCM of the prime number tests
				// }
				// if worry%12 == 0 {
				// 	worry = worry / 12
				// }
				// worry := int(math.Floor(float64(m.Operations[i](v)) / 4))
				to := m.Tests[i](worry)
				m.Items[to] = append(m.Items[to], worry)
			}
			m.Items[i] = []int{} // empty the list when finished
		}
	}
}

func (m *MonkeyOrchestrator2C) Score() int {

	fmt.Println(m.Inspected)
	sort.Sort(sort.Reverse(sort.IntSlice(m.Inspected)))
	return m.Inspected[0] * m.Inspected[1]
}
