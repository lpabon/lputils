// Copyright (c) 2019 Luis Pabón <lpabon@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package workers

import (
	"sync"
)

// Workers provides a method for `numWorkers` number of goroutines to act as
// consumers for the content created by the producer closure. What is placed
// in the channel `ch` by the producer will be passed to the consumer function
// as the argument.
func Workers(numWorkers int, consumerf func(any), producerf func(chan<- any)) {
	ch := make(chan any)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for data := range ch {
				consumerf(data)
			}
			wg.Done()
		}()
	}

	// Start producer
	producerf(ch)

	// Done producing
	close(ch)
	wg.Wait()
}
