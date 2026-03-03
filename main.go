package main

import (

	"github.com/mmaxemm/internship_tasks/concurrency"
)

/*
Для данной задачи можно использовать различные наборы параметров, однако рекоммендуется попробовать
следующие комбинации:
	- K = 10_000, N = 1, M = 1
	- K = 10_000, N = 8, M = 5
	- K = 10_000, N = 100, M = 50
*/

func main() {
	K := 10000
	M1, N1 := 1, 1
	M2, N2 := 5, 8
	M3, N3 := 50, 100
	concurrency.StartBakery(N1, K, M1)
	concurrency.StartBakery(N2, K, M2)
	concurrency.StartBakery(N3, K, M3)
}
