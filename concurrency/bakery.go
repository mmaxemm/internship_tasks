package concurrency

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
С помощью горутин и каналов смоделировать работу пекарни:

Пекарня состоит из 2+1 этапов: торт выпекается, торт запаковывается, все торты смотрятся. Связь между
этапами происходит через каналы, этапы выполняются с помощью горутин. Торт представляет из себя объект
с полями типа int: BakedBy, BakeTime, PackedBy, PackTime Программа должна работать следующим образом:
    Запускается N горутин, каждая из которых за время T1 = i +-t1 (i - номер рутины, t1 - выбранный
	вами параметр) создаёт объект тортика, заполняет поля BackedBy=i, BakeTime=T1 и отправляет его в канал

    Существует пул из M горутин. когда в канал приходят торитики с предыдущего этапа, свободная горутина
	из пула начинает его упаковывать, то есть за время T2 = j+-t2 (j - номер рутины из пула,
	t2 - выбранный вами параметр, причём t2>=t1) выставляет тортику параметры PackedBy=j,
	PackTime = T2 и отправляет в канал

    Мы ждём пока придут все тортики, или мы получим сигнал о завершении работы(тогда дожидаемся завершения
	работы текущих рутин), после чего просто выводим тортики в порядке, в котором они пришли к нам. При
	полном отрабатывании программы должно быть K тортиков.
*/

type Cake struct {
	BakedBy int
	BakedTime int
	PackedBy int
	PackedTime int
}

// messages about the Cakes Packed are stored in tmpConcurrency.log, while benchmarks are standard output
func StartBakery(N, M, K int) {
	start := time.Now()
	const t1 = 10 // ms
	const t2 = 15
	
	bakedCakes := make(chan Cake, N)
	packedCakes := make(chan Cake, M)
	
	var bakersWG sync.WaitGroup
	
	bakersWG.Add(N)
	for i := range N {
		go func(bakerID int) {
			defer bakersWG.Done()
			var distributedCakes int // amount of cakes which should be done by this goroutine
			if bakerID < K%N {
				distributedCakes = K/N + 1
			} else {
				distributedCakes = K/N
			}
			for range distributedCakes {
				variance := rand.Intn(2*t1+1) - t1
				bakeTime := max(bakerID + variance, 5)

				time.Sleep(time.Duration(bakeTime) * time.Millisecond)

				cake := Cake{
					BakedBy:   bakerID,
					BakedTime: bakeTime,
				}
				bakedCakes <- cake
			}
		}(i)
	}
	
	go func() {
		bakersWG.Wait()
		close(bakedCakes)
	}()	

	var packersWG sync.WaitGroup
	packersWG.Add(M)
	for i := range M {
		go func(packerID int) {
			defer packersWG.Done()
			for cake := range bakedCakes {
				variance := rand.Intn(2*t2+1) - t2
				packTime := max(packerID + variance, 5)
				
				time.Sleep(time.Duration(packTime) * time.Millisecond)
				
				cake.PackedBy = packerID
				cake.PackedTime = packTime
				packedCakes <- cake
			}
		}(i)
	}

	go func() {
		packersWG.Wait()
		close(packedCakes)
	}()

	f, err := os.OpenFile("tmp_concurrency.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("NEW LOG ENTRY on %v\n", time.Now())
	
	cakeNum := 1
	for cake := range packedCakes {
		log.Printf("Cake N%d: BakedBy: %d, BakedTime: %d, PackedBy: %d, PackedTime: %d\n", 
			cakeNum, cake.BakedBy, cake.BakedTime, cake.PackedBy, cake.PackedTime)
		cakeNum++
	}
	
	f.Sync()
	t := time.Now()
	fmt.Printf("Bakery with params N = %v, M = %v, K = %v took %v\n", N, M, K, t.Sub(start))
}
