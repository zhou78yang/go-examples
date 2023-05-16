package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	IpListLen          = 3
	maxSamePriorityNum = 2
)

type ipRecord struct {
	Ttl      int
	Priority int
	Address  string
}

func getIpList(txtList []string) []string {
	records := convertToRecord(txtList)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	counter := make(map[int]int, len(records))
	for _, r := range records {
		counter[r.Priority]++
	}
	priorities := make([]int, 0, len(counter))
	for p := range counter {
		priorities = append(priorities, p)
	}
	sort.Ints(priorities)

	need := make(map[int]int, IpListLen)
	rest := IpListLen
	for _, p := range priorities {
		n := maxSamePriorityNum
		if n > counter[p] {
			n = counter[p]
		}

		if n >= rest {
			need[p] = rest
			break
		}
		need[p] = n
		rest -= n
	}

	fmt.Println("info:", counter, priorities, need)

	result := make([]string, 0, IpListLen)
	for _, p := range priorities {
		for _, r := range records {
			if need[p] == 0 {
				break
			}

			if r.Priority == p {
				result = append(result, r.Address)
				need[p]--
			}
		}
	}
	return result
}

func convertToRecord(txtList []string) []*ipRecord {
	var result []*ipRecord
	for _, txt := range txtList {
		if r, err := parseTxt(txt); err != nil {
			fmt.Println("parse error", err)
		} else {
			result = append(result, r)
		}
	}
	return result
}

func parseTxt(txt string) (*ipRecord, error) {
	fields := strings.Fields(txt)
	ttl, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}
	priority, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}

	return &ipRecord{
		Ttl:      ttl,
		Priority: priority,
		Address:  fields[4],
	}, nil
}

func main() {
	txtList := []string{
		"60 101 1 center 192.168.0.11",
		"60 201 1 center 192.168.0.21",
		"60 301 1 center 192.168.0.31",
		"60 301 1 center 192.168.0.32",
		"60 401 1 center 192.168.0.41",
		"60 401 1 center 192.168.0.42",
		"60 301 1 center 192.168.0.33",
		"60 201 1 center 192.168.0.22",
		"60 201 1 center 192.168.0.23",
		"60 101 1 center 192.168.0.12",
		"60 101 1 center 192.168.0.13",
		"60 101 1 center 192.168.0.14",
		"60 101 1 center 192.168.0.15",
	}
	fmt.Println(getIpList(txtList[:5]))
	fmt.Println(getIpList(txtList[:9]))
	fmt.Println(getIpList(txtList[:10])) // 刚好够
	fmt.Println(getIpList(txtList))
	fmt.Println(getIpList(txtList))
}
