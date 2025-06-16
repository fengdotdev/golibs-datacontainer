package gomap

import (
	"context"
	"fmt"
)

func (gm *GoMap[K, V]) Populate(ctx context.Context, data map[K]V) chan error {
	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done

		gm.jobs <- func() {
			select {

			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				for key, value := range data {
					gm.mapdata[key] = value
				}
				errChan <- nil // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Overwrite(ctx context.Context, key K, value V) chan error {
	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				gm.mapdata[key] = value
				errChan <- nil // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Set(ctx context.Context, key K, value V) chan error {

	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				gm.mapdata[key] = value
				errChan <- nil // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Get(ctx context.Context, key K) (chan V, chan error) {
	errChan := make(chan error, 1)
	valueChan := make(chan V, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				value, exists := gm.mapdata[key]
				if !exists {

					errChan <- fmt.Errorf("key %v not found", key)
					var zero V

					valueChan <- zero // Return zero value
					return
				}
				errChan <- nil     // No error
				valueChan <- value // Send the value
			}
		}

	}()

	return valueChan, errChan
}

func (gm *GoMap[K, V]) Delete(ctx context.Context, key K) chan error {
	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				if _, exists := gm.mapdata[key]; !exists {
					errChan <- fmt.Errorf("key %v not found", key)
					return
				}
				delete(gm.mapdata, key)
				errChan <- nil // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Clear(ctx context.Context) chan error {
	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				gm.mapdata = make(map[K]V) // Clear the map
				errChan <- nil             // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Contains(ctx context.Context, key K) (chan bool, chan error) {
	containsChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				_, exists := gm.mapdata[key]
				errChan <- nil         // No error
				containsChan <- exists // Key exists
			}
		}

	}()
	return containsChan, errChan
}

func (gm *GoMap[K, V]) Size(ctx context.Context) (chan int, chan error) {
	sizeChan := make(chan int, 1)
	errChan := make(chan error, 1)
gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				size := len(gm.mapdata)
				errChan <- nil   // No error
				sizeChan <- size // Send the size
			}
		}

	}()

	return sizeChan, errChan
}

func (gm *GoMap[K, V]) Keys(ctx context.Context) (chan []K, chan error) {
	keysChan := make(chan []K, 1)
	errChan := make(chan error, 1)
gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				keys := make([]K, 0, len(gm.mapdata))
				for key := range gm.mapdata {
					keys = append(keys, key)
				}
				errChan <- nil   // No error
				keysChan <- keys // Send the keys
			}
		}

	}()

	return keysChan, errChan
}

func (gm *GoMap[K, V]) Values(ctx context.Context) (chan []V, chan error) {
	valuesChan := make(chan []V, 1)
	errChan := make(chan error, 1)
gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				values := make([]V, 0, len(gm.mapdata))
				for _, value := range gm.mapdata {
					values = append(values, value)
				}
				errChan <- nil       // No error
				valuesChan <- values // Send the values
			}
		}

	}()

	return valuesChan, errChan
}

func (gm *GoMap[K, V]) ForEach(ctx context.Context, fn func(K, V) error) chan error {
	errChan := make(chan error, 1)
gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				for key, value := range gm.mapdata {
					if err := fn(key, value); err != nil {
						errChan <- err
						return
					}
				}
				errChan <- nil // No error
			}
		}

	}()

	return errChan
}

func (gm *GoMap[K, V]) Clone(ctx context.Context) (chan *GoMap[K, V], chan error) {
	errChan := make(chan error, 1)
	cloneChan := make(chan *GoMap[K, V], 1)
gm.adquireLimiter() // Acquire a slot in the limiter
	go func() {
		defer gm.releaseLimiter() // Ensure the slot is released when done
		gm.jobs <- func() {

			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:

				clone := NewGoMapFromMap(gm.mapdata) // Create a new GoMap with the same data

				cloneChan <- clone // Send the cloned map
				errChan <- nil     // No error
			}
		}

	}()

	return cloneChan, errChan
}

func (gm *GoMap[K, V]) adquireLimiter() {
	select {
	case gm.limiter <- struct{}{}: // Acquire a slot in the limiter
	default:
		// If the channel is full, block until a slot is available
		gm.limiter <- struct{}{}
	}
}
func (gm *GoMap[K, V]) releaseLimiter() {
	<-gm.limiter // Release the slot in the limiter
}
