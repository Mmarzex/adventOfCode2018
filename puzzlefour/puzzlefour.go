package puzzlefour

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EventType string

const (
	SLEEP  = "falls asleep"
	WAKEUP = "wakes up"
)

type Event struct {
	T  time.Time
	Ev EventType
}
type EventLog map[int][]Event

var logLineRegexp = regexp.MustCompile(`\[(.+)] (.+)`)
var shiftStartRegexp = regexp.MustCompile(`Guard #(\d+) begins shift`)

func parseLog(lines []string) EventLog {
	sort.Strings(lines)
	var id int
	events := make(EventLog)
	for _, ln := range lines {
		matches := logLineRegexp.FindAllStringSubmatch(ln, -1)
		t, err := time.Parse("2006-01-02 15:04", matches[0][1])
		if err != nil {
			log.Fatal(err)
		}
		logMsg := strings.Trim(matches[0][2], " ")
		if idMatch := shiftStartRegexp.FindAllStringSubmatch(logMsg, -1); len(idMatch) == 1 {
			id, _ = strconv.Atoi(idMatch[0][1])
		} else {
			events[id] = append(events[id], Event{t, EventType(logMsg)})
		}

	}
	return events
}

type GuardMin struct {
	m int
	count int
	guard int
}

func calculateGuardsSleep(events EventLog) {
	minAsleep := make(map[int]int)
	minutes := make(map[int]map[int]int)

	for k, e := range events {
		for i := 0; i < len(e); i += 2 {
			asleep, awake := e[i], e[i+1]
			minAsleep[k] += int(awake.T.Sub(asleep.T).Minutes())
			for j := asleep.T.Minute(); j < awake.T.Minute(); j++ {
				if _, ok := minutes[k]; !ok {
					minutes[k] = make(map[int]int)
				}
				minutes[k][j]++
			}
		}
	}
	fmt.Println(minAsleep)
	fmt.Println(minutes)

	worstGuard := -1
	worstGuardCount := -1
	for k, v := range minAsleep {
		if v > worstGuardCount {
			worstGuard = k
			worstGuardCount = v
		}
	}
	worstMinute := 0
	for m, v := range minutes[worstGuard] {
		if v > minutes[worstGuard][worstMinute] {
			worstMinute = m
		}
	}
	fmt.Println(fmt.Sprintf("Worst Guard %d Worst Minute %d", worstGuard, worstMinute))
	fmt.Println(fmt.Sprintf("Part 1 Answer %d", worstGuard * worstMinute))


	worstGuardTwo := GuardMin{
		count: -1,
	}
	for guardId, mins := range minutes {
		guardMin := GuardMin{
			m: -1,
			count: -1,
		}
		for m, c := range mins {
			if c > guardMin.count {
				guardMin.m = m
				guardMin.count = c
			}
		}
		if guardMin.count > worstGuardTwo.count {
			worstGuardTwo = guardMin
			worstGuardTwo.guard = guardId
		}
	}
	fmt.Println(worstGuardTwo)
	fmt.Println(fmt.Sprintf("Part 2 Answer %d", worstGuardTwo.guard * worstGuardTwo.m))
}

func readInput(path string) EventLog {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(dat), "\n")
	events := parseLog(data)
	return events
}

func RunPuzzle() {
	eventLog := readInput("inputPuzzleFour.txt")
	fmt.Println(eventLog)
	calculateGuardsSleep(eventLog)
}
