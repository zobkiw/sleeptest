# sleeptest
golang sleep test

### Quick Start

Use `go run main.go` to run this code.

Make sure `main.sh` is executable by using `chmod +x main.sh`.

Please tell me if there are any bugs in this code that could be causing this issue.

### Introduction

This test is to help determine how `time.Sleep()` functions in Go in various scenarios.

I wrote a Go program that called a shell script that executed multiple command line tools - each of which accessed the network. The shell script was executed serially using `exec.Command()`, in a loop 3x, being passed a different parameter each time. Once all of the calls completed, the Go program would execute one final shell script and then `time.Sleep(1 * time.Hour)` before continuing the endless for loop. 

The original shell script would gather data from AWS, format it, and then post it to a chat room, although the exact task didn't seem to influence the sleep timing issue since it was outside the scope of the sleep itself. I noticed that in the chat room the first few times this ran, the timestamps would be one hour later each time. If the script finished at 1:03 the next would would begin at 2:03, etc. However, after a few runs the times would increase to 3 minutes or more. One time I even saw a 9 minute discrepancy...lost time...mind you, the timestamps were very course, so I wrote the program here for more granularity and to distill the problem down.

The original script was running in a `screen` session on an Intel-based mac mini (2018 / 3.6 GHz Quad-Core Intel Core i3 / 8GB RAM) running go1.16.2 darwin/amd64. Interestingly it also seemed to go back to exact 1 hour sleeps once the machine was being touched again - it was only overnight - when the machine was being unused - that the delays seemed to occur.

I wrote the code you see here in order to test more carefully and watch more closely, while mimicing the original as close as possible (ie: random sleep times in the shell script to mimic the average network access times of the original.) As I'm writing this, I have it running on the aforementioned Intel machine in a `screen` session and on an M1-based MBP (2020 / 16 GB) running go1.16.2 darwin/arm64 in both a `screen` session and another terminal tab per usual.

It's only been running a few iterations at the moment but some interesting numbers are already beginning to show. What follows is the output that occurs as the loop restarts after a sleep. I store all the previous sleep times in a slice and display them as comparisons.

- M1 terminal tab:      zzz GO ALL SLEEP TIMES: [1h4m29.850342s 1h4m54.004981s 59m59.994061s]
- M1 screen session:    zzz GO ALL SLEEP TIMES: [1h4m29.852317s 1h4m54.001291s 59m59.997482s]
- Intel screen session: zzz GO ALL SLEEP TIMES: [59m59.940643s 1h0m0.110959s 1h0m0.103718s]

Note that the Intel machine is currently spot on, which it usually is for awhile. I currently only have an ssh connection open to it to check once in awhile but I'm about to disconnect and let it run for the remainder of the day to give it time to exhibit the problem. Also note that the M1 is having problems right out of the gate but then recovers once, in both `screen` and non. It is being used actively but just for typing this text and some light web surfing and terminal work in another window. Nothing that should cause the 4.5+ minutes of lost time. Note that this is the same code running on both machines as you see in the initial commit of this repo.

Updates will be posted here.
