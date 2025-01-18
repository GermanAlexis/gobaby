package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Clock struct {
	Stop      chan struct{}
	CountDown string
	mu        sync.Mutex
	runnig    bool
}

var currentDuration = 0

func SetDuration(duration int) {
	currentDuration = duration
}

func FormatCountDownToTimestamp(countDown string) int {

	parts := strings.Split(countDown, ":")

	if len(parts) <= 3 {
		return 0
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0
	}
	totalSeconds := (hours * 3600) + (minutes * 6000) + seconds

	return totalSeconds
}

func FormatCountDownToInt(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func NewClock() *Clock {
	return &Clock{Stop: make(chan struct{}, 1), CountDown: "04:00:00", mu: sync.Mutex{}}
}

func StartClock(clock *Clock, duration int) {
	clock.mu.Lock()
	defer clock.mu.Unlock()

	if clock.runnig {
		fmt.Println("Clock Alreay Running")
		return
	}

	clock.runnig = false

	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		clock.CountDown = FormatCountDownToInt(currentDuration)
		currentDuration--

		if currentDuration < 0 {
			currentDuration = duration
		}
		select {
		case <-clock.Stop:
			clock.runnig = false
			return
		default:
		}
	}
}

func StopCountDown(clock *Clock) {
	clock.Stop <- struct{}{}
}
