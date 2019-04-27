package pack08

import (
	"fmt"
	"time"
)

func task(task_id int)  {
	for {
		fmt.Printf("%d\n",task_id)
		time.Sleep(time.Second)
	}
}

func StartTask() {
	fmt.Println("--- start task ---")
	go task(1)
	go task(2)
	go task(3)
	go task(4)
	go task(5)

	time.Sleep(time.Second * 30)
}
