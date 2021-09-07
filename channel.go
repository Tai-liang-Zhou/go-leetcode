package main

// func main() {
// 	outChan := make(chan string, 10)
// 	finishChan := make(chan struct{})
// 	errChan := make(chan error)
// 	wg := sync.WaitGroup{}
// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {

// 		go func(val int, wg *sync.WaitGroup, out chan string, err chan error) {
// 			// defer wg.Done()

// 			// time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
// 			// fmt.Println("finish jod id  :", val)
// 			out <- fmt.Sprintf("finish job is : %d", val)

// 			if val == 300 {
// 				errChan <- errors.New("error in 15")
// 			}

// 		}(i, &wg, outChan, errChan)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(finishChan)
// 	}()
// Loop:
// 	for {
// 		select {
// 		case out := <-outChan:
// 			log.Println(out)
// 			wg.Done()
// 		case <-finishChan:
// 			fmt.Println("finish out")
// 			break Loop
// 		case err := <-errChan:
// 			fmt.Println(err)
// 			break Loop
// 		case <-time.After(1000 * time.Millisecond):
// 			fmt.Println("time out")
// 			break Loop
// 		}
// 	}

// }
