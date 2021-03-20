package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	commands := []string{"uptime", "uptime", "uptime"} // commands to execute via main.sh

	var sleepTimes []time.Duration // slice of sleep times to see how they progress

	var endTime time.Time       // time it took to process all commands
	var sleepTime time.Duration // time it took to sleep (1 hour is the test time)

	for {

		startTime := time.Now().UTC() // loop start time
		if !endTime.IsZero() {        // first time through?
			sleepTime = startTime.Sub(endTime)         // how long did we sleep?
			sleepTimes = append(sleepTimes, sleepTime) // track all sleep times
			fmt.Println("=======================")
			fmt.Println("zzz GO PREV SLEEP TIME:", sleepTime.String())
			fmt.Println("zzz GO ALL SLEEP TIMES:", sleepTimes)
		}

		fmt.Println(">>> GO STARTING...", startTime.String())

		// loop through commands passed to shell script to simulate work
		for i, command := range commands {
			fmt.Println("GO PROCESSING COMMAND", i+1, command)

			cmd := exec.Command("./main.sh", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println("GO ERROR PROCESSING COMMAND: ", command, err)
			}
		}

		// one additional command since the original script had this
		// may not be needed to reproduce the problem but here for similarity
		additionalCommand := "uptime"
		fmt.Println("GO PROCESSING ADDITIONAL COMMAND: ", additionalCommand)
		cmd := exec.Command("./main.sh", additionalCommand)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("GO ERROR PROCESSING ADDITIONAL COMMAND: ", additionalCommand, err)
		}

		endTime = time.Now().UTC()               // loop end time, begin of sleep time (close enough)
		processingTime := endTime.Sub(startTime) // overall command processing time
		fmt.Println("GO PROCESSING TIME:", processingTime.String())

		fmt.Println("<<< GO SLEEPING...", endTime.String())
		time.Sleep(1 * time.Hour) // zzz
		//time.Sleep(10 * time.Second)
	}

}
