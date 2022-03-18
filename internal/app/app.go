package app

import (
	"fmt"
	"log"
	"time"

	"github.com/bartmika/cameraticker/internal/camera"
	"github.com/bartmika/timekit"
)

type App struct {
	cam    camera.CameraStillReader
	timer  *time.Timer
	ticker *time.Ticker
	done   chan bool // The channel which terminaton boolean value is delivered.
}

func New(cam camera.CameraStillReader) (*App, error) {
	s := &App{
		cam:    cam,
		timer:  nil,
		ticker: nil,
		done:   make(chan bool, 1), // Create a execution blocking channel.
	}
	return s, nil
}

// RunMainRuntimeLoop will consume the main runtime loop and run the business logic of the application.
func (s *App) RunMainRuntimeLoop() {
	defer s.shutdown()

	// DEVELOPERS NOTE:
	// The purpose of this block of code is to find the future date where
	// the minute just started, ex: 5:00 AM, 5:05, 5:10, etc, and then start
	// our main runtime loop to run along for every minute afterwords.

	now := time.Now()
	future := timekit.GetFutureDateByFiveMinuteIntervalPattern(now)
	diff := future.Sub(now)
	t := time.NewTimer(diff)
	s.timer = t

	// DEVELOPERS NOTE:
	// The purpose of this switch is to wait for either a timer tick or an
	// termination code from the system. If a timer tick happens then we will
	// create our 5 minute interval ticker. If our application gets terminated
	// by the user or system then we terminate our timer and exit our main
	// runtime loop.

	log.Printf("Synching with local time...")
	select {
	case <-s.timer.C:
		log.Printf("Synchronized with local time.")
		s.ticker = time.NewTicker(5 * time.Minute)
	case <-s.done:
		s.timer.Stop()
		log.Printf("Interrupted timer.")
		return
	}

	// THE CODE BELOW IS FOR TESTING PURPOSES AND SHOULD ONLY BE USED IF YOU DO
	// NOT WANT TO WAIT FOR THE TIME SYNCH CODE ABOVE. DO NOT USE THE CODE BELOW
	// FOR PRODUCTION SYSTEM!
	// s.ticker = time.NewTicker(1 * time.Minute) // Tick every minute

	// DEVELOPERS NOTE:
	// (1) The purpose of this block of code is to run as a goroutine in the
	//     background as an anonymous function waiting to get either the
	//     ticker chan or app termination chan response.
	// (2) Main runtime loop's execution is blocked by the `done` chan which
	//     can only be triggered when this application gets a termination signal
	//     from the operating system.
	log.Printf("App is now running.")
	go func() {
		for {
			select {
			case <-s.ticker.C:
				// DEVELOPERS NOTE:
				// Golang's tickers are somewhat not accurate according to
				// Denis Bernard's Stackoverlows answer via
				// https://stackoverflow.com/a/51424566. As a result we will
				// have to perform a very slight rounding to get nice numbers.

				// Round out the seconds as we do not want any seconds. So for
				// example the following rounding will happen.
				// 2022-03-16 23:04:59.99754 rounds to 2022-03-16 23:05:00
				d := (60 * time.Second)
				tickDT := time.Now().Round(d)

				s.executeAtTick(tickDT)
			case <-s.done:
				log.Printf("Interrupted ticker.")
				s.ticker.Stop()
				return
			}
		}
	}()
	<-s.done
}

func (s *App) executeAtTick(dt time.Time) {
	fmt.Print(dt, "Tick")
	_, err := s.cam.Snapshot()
	if err != nil {
		log.Fatal(err)
	}
}

// StopMainRuntimeLoop will tell the application to stop the main runtime loop when the process has been finished.
func (s *App) StopMainRuntimeLoop() {
	s.done <- true
}

func (s *App) shutdown() {

}
