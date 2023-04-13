package philosophers_dining

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	id int
	sync.Mutex
}

type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eat(wg *sync.WaitGroup, channel chan bool, delay time.Duration) {
	defer wg.Done()

	// pick up right chopstick
	p.rightChopstick.Lock()

	time.Sleep(time.Millisecond * delay)
	fmt.Println(fmt.Sprintf("Philosopher %d is picked up 'right' chopstick %d", p.id, p.rightChopstick.id))

	// pick up left chopstick
	p.leftChopstick.Lock()
	fmt.Println(fmt.Sprintf("Philosopher %d is picked up 'left' chopstick %d", p.id, p.leftChopstick.id))
	fmt.Println(fmt.Sprintf("Philosopher %d is eating", p.id))

	// put down left chop stick
	p.leftChopstick.Unlock()
	fmt.Println(fmt.Sprintf("Philosopher %d is put down 'left' chopstick %d", p.id, p.leftChopstick.id))

	// put down riht chopstick
	p.rightChopstick.Unlock()
	fmt.Println(fmt.Sprintf("Philosopher %d is put down 'right' chopstick %d", p.id, p.rightChopstick.id))

	channel <- true
	fmt.Println(fmt.Sprintf("Philosopher %d is thinking", p.id))

	<-channel
}

func Dine(delay time.Duration) bool {
	numPhilosophers := 5

	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
		chopsticks[i].id = i + 1
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
		}
	}

	var wg sync.WaitGroup
	channel := make(chan bool, numPhilosophers-1)

	for _, p := range philosophers {
		wg.Add(1)
		go p.eat(&wg, channel, delay)
	}

	wg.Wait()
	return true
}
