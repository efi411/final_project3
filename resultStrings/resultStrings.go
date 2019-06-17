package resultStrings

import (
	"sync"
)

type ResultStrings struct {
	result string
	leader string
	sync.Mutex
}

func NewResultString() ResultStrings {
	return ResultStrings{
		result: "",
		leader: "",
	}
}

func (r *ResultStrings) GetResultString() (string, string) {
	if len(r.result) > 0 && len(r.leader) > 0 {
		return r.result[:len(r.result)-1], r.leader[:len(r.leader)-1]
	}
	return "error", "error"
}

func (r *ResultStrings) AddMessage(str string) {
	r.Lock()
	r.result = r.result + str + "-"
	r.Unlock()
}

func (r *ResultStrings) AddLeader(str string) {
	r.leader = str + ":" + r.leader
}
