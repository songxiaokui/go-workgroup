// Package quit provides functions for running and cancelling execution using functions.
package quit

// New creates function for running and cancelling execution.
func New(run func() error, quit func()) func(<-chan struct{}) error {
	return func(stop <-chan struct{}) error {
		go func() {
			<-stop
			quit()
		}()
		return run()
	}
}
