package main

import "log"

func main() {
	methods := []string{"Fancy", "append", "addAll", "append", "multAll", "getIndex", "addAll", "append", "multAll", "getIndex", "getIndex", "getIndex"}
	// [[],[2],[3],[7],[2],[0],[3],[10],[2],[0],[1],[2]]
	params := [][]int{{}, {2}, {3}, {7}, {2}, {0}, {3}, {10}, {2}, {0}, {1}, {2}}

	fancy := Constructor()
	for i, method := range methods {
		if i == 0 {
			continue
		}
		switch method {
		case "append":
			fancy.Append(params[i][0])
		case "addAll":
			fancy.AddAll(params[i][0])
		case "multAll":
			fancy.MultAll(params[i][0])
		case "getIndex":
			log.Println(fancy.GetIndex(params[i][0]))
		}
	}
}

// 解题思路
type offset struct {
	mult_value int64
	add_value  int64
}

type Fancy struct {
	len            int
	vals           []int
	offsets        []offset
	current_offset offset
	last_value     int
	last_index     int
}

func Constructor() Fancy {
	return Fancy{vals: make([]int, 100), offsets: make([]offset, 100), current_offset: offset{mult_value: 1, add_value: 0}, last_value: -1, last_index: -1}
}

func (this *Fancy) Append(val int) {
	this.len++
	if this.len%100 == 0 {
		this.vals = append(this.vals, make([]int, 100)...)
		this.offsets = append(this.offsets, make([]offset, 100)...)
	}
	this.vals[this.len-1] = val
	this.offsets[this.len-1] = this.current_offset
	this.current_offset = offset{mult_value: 1, add_value: 0}
}

func (this *Fancy) AddAll(inc int) {
	this.last_value = -1
	this.last_index = -1
	if this.len == 0 {
		return
	}
	this.current_offset.add_value = this.current_offset.add_value + int64(inc)
	if this.current_offset.add_value >= 1000000007 {
		this.current_offset.add_value -= 1000000007
	}
}

func (this *Fancy) MultAll(m int) {
	this.last_value = -1
	this.last_index = -1
	if this.len == 0 {
		return
	}
	this.current_offset.mult_value = this.current_offset.mult_value * int64(m)
	if this.current_offset.mult_value >= 1000000007 {
		this.current_offset.mult_value = this.getMod(this.current_offset.mult_value)
	}
	this.current_offset.add_value = this.current_offset.add_value * int64(m)
	if this.current_offset.add_value >= 1000000007 {
		this.current_offset.add_value = this.getMod(this.current_offset.add_value)
	}
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= this.len {
		return -1
	}
	if idx == this.last_index {
		return this.last_value
	}
	val := int64(this.vals[idx])
	for i, opt := range this.offsets[idx:this.len] {
		if i == 0 {
			continue
		}
		val = val*opt.mult_value + opt.add_value
		if val >= 1000000007 {
			val = this.getMod(val)
		}
	}
	val = val*this.current_offset.mult_value + this.current_offset.add_value
	val2 := int(this.getMod(val))
	this.last_index = idx
	this.last_value = val2
	return val2
}

func (this *Fancy) getMod(value int64) int64 {
	val := value % 1000000007
	if val < 0 {
		val += 1000000007
	}
	return int64(val)
}
