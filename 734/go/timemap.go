package test

import (
	"math"
)

type timeBasedMap map[string]timeBasedVal
type timeBasedVal map[int]string

func newTimeBasedMap(s int) timeBasedMap {
	return make(map[string]timeBasedVal, s)
}

func (m timeBasedMap) set(key, val string, time int) {
	if _, ok := m[key]; !ok {
		m[key] = make(timeBasedVal, 1)
	}

	m[key][time] = val
}

func (m timeBasedMap) get(key string, time int) string {
	v, ok := m[key]
	if !ok {
		return ""
	}

	return v.get(time)
}

func (v timeBasedVal) get(time int) string {
	closest := math.MinInt32
	for tm, val := range v {
		if tm == time {
			return val
		}

		if tm < time {
			if tm > closest {
				closest = tm
			}
		}
	}

	return v[closest]
}
