# sleeptest
golang `time.Sleep()` test

IMPORTANT: NOTE: 1.16.3 seems to solve the issue outlined here. Feel free to continue to read if you have nothing better to do.

### TL;DR

`time.Sleep(1 * time.Hour)` can sometimes sleep for longer than the 1 hour that is specified, although I generally see it between 3-5 minutes when it is behaving badly in my experience. Testing against go1.16.2 currently.

macOS has a problem on both M1 and Intel. A Digital Ocean Linux droplet works perfectly fine. See description and data below.

See https://github.com/golang/go/issues/44343 if you've stumbled upon this from some other place.

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

Both the Intel and the M1 Macs are running Big Sur 11.2.3.

It's only been running a few iterations at the moment but some interesting numbers are already beginning to show. What follows is the output that occurs as the loop restarts after a sleep. I store all the previous sleep times in a slice and display them as comparisons.

```
- M1 terminal tab:        [1h4m29.850342s 1h4m54.004981s 59m59.994061s 1h5m25.892846s 2h2m8.529273s 1h5m24.720327s 59m59.980485s]
- M1 screen session:      [1h4m29.852317s 1h4m54.001291s 59m59.997482s 1h5m25.886925s 2h2m8.530685s 1h5m24.720294s 59m59.977141s]
- Intel screen session:   [59m59.940643s 1h0m0.110959s 1h0m0.103718s 1h3m58.080461s 1h6m25.252947s 1h11m23.48131s 1h7m46.400632s]

- DgtlOcn screen session: [1h0m0.081298801s 1h0m0.097626585s 1h0m0.000158801s 1h0m0.000324235s]
  Ubuntu Docker 19.03.12
  go1.16.2 linux/amd64
```

Note that the Intel machine is currently spot on, which it usually is for awhile. I currently only have an ssh connection open to it to check once in awhile but I'm about to disconnect and let it run for the remainder of the day to give it time to exhibit the problem. Also note that the M1 is having problems right out of the gate but then recovers once, in both `screen` and non. It is being used actively but just for typing this text and some light web surfing and terminal work in another window. Nothing that should cause the 4.5+ minutes of lost time. Note that this is the same code running on both machines as you see in the initial commit of this repo.

Updates will be posted here.

### About the 2 hour time

One interesting find, the M1 was left idle for almost two hours. It last posted a "sleeping" message at 17:15:58. Once I returned to the computer (at approximately 19:07:42) the program was still sleeping. This may be how time.Sleep() works? If the machine itself goes to sleep it pauses and doesn't count that time? I would have expected that it might "catch up" as soon as the machine came back to life, however. The `screen` session also exhibited the same issue as one might expect. I'm waiting now to see if it (the Go program) comes back to life on it's own and wakes up...and there it goes, about 10 minutes after I woke up the machine. Note the 5th time in the M1 lists above (2h2m8.529273s and 2h2m8.530685s). Note the next sample "recovered" to a more normal time, albeit, more than an hour.

### Digitial Ocean Droplet

I decided to also do some runs on a Digital Ocean droplet. One of the smallest, configuration shown in the chart above. Note the very strict adherance to the 1 hour time. Impressive, so we know that Go can do it. It would seem that maybe macOS is causing some problems specifically.

### Python

For kicks, I wrote a super simple Python 3.9.1 script to see how sleep behaves. I only ran it on the M1, I think it proves the point. Python can exhibit a similar problem. The results and the script itself follow:

```
ABOUT TO SLEEP:  2021-03-21 14:23:52.171367
AWAKE:           2021-03-21 15:23:52.178701 // one hour = one hour
----------------
ABOUT TO SLEEP:  2021-03-21 15:23:52.181397
AWAKE:           2021-03-21 16:24:43.046964 // one hour = one hour and 51-ish seconds
----------------                            // the machine was on but idle starting sometime in the second iteration
ABOUT TO SLEEP:  2021-03-21 16:24:43.048686 // the machine was being used sometime in the third iteration
AWAKE:           2021-03-21 17:59:33.436385 // one hour = one hour and 35-ish minutes
```

```
import time
from datetime import datetime

while 1:
    print("ABOUT TO SLEEP: ", datetime.utcnow())
    time.sleep(60*60)
    print("AWAKE:          ", datetime.utcnow())
    print("----------------")
```
