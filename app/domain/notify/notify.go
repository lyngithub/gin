package notify

import (
	"sync"
	"xx/forms"
	"xx/global"
)

func SetChan(form *forms.UserStatus) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		global.StatusUpdates <- form
	}()
	wg.Wait()

	return nil
}
