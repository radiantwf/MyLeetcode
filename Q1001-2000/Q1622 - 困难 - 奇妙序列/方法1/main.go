package main

import "log"

func main() {
	methods := []string{"Fancy", "append", "getIndex", "multAll", "multAll", "getIndex", "addAll", "append", "append", "getIndex", "append", "append", "addAll", "getIndex", "multAll", "addAll", "append", "addAll", "getIndex", "getIndex", "multAll", "multAll", "multAll", "append", "addAll", "getIndex", "getIndex", "getIndex", "append", "getIndex", "addAll", "multAll", "append", "multAll", "addAll", "getIndex", "append", "append", "addAll", "getIndex", "multAll", "getIndex", "addAll", "getIndex", "multAll", "addAll", "getIndex", "addAll", "append", "append", "append", "multAll", "multAll", "append", "multAll", "addAll", "getIndex", "addAll", "multAll", "multAll", "multAll", "append", "multAll", "append", "multAll", "addAll", "append", "append", "getIndex", "getIndex", "getIndex", "addAll", "multAll", "multAll", "append", "append", "getIndex", "append", "append", "append", "getIndex", "getIndex"}
	params := [][]int{{}, {5}, {0}, {14}, {10}, {0}, {12}, {10}, {4}, {2}, {4}, {2}, {1}, {1}, {8}, {11}, {15}, {12}, {0}, {3}, {4}, {11}, {11}, {10}, {8}, {2}, {3}, {0}, {7}, {3}, {2}, {6}, {10}, {6}, {8}, {7}, {9}, {9}, {12}, {0}, {13}, {7}, {3}, {4}, {8}, {14}, {2}, {9}, {9}, {9}, {7}, {5}, {12}, {9}, {3}, {8}, {10}, {14}, {14}, {14}, {6}, {1}, {3}, {11}, {12}, {6}, {7}, {13}, {12}, {5}, {6}, {1}, {11}, {11}, {4}, {9}, {7}, {11}, {1}, {3}, {1}, {0}}

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

const serial_opts_unit = 100

type operation struct {
	method string
	param  int
}

type Fancy struct {
	len                    int
	vals                   []int
	serials                []int
	current_opt_serial_num int
	serial_opts            []*operation
}

func Constructor() Fancy {
	return Fancy{vals: make([]int, 100), serials: make([]int, 100), serial_opts: make([]*operation, serial_opts_unit)}
}

func (this *Fancy) Append(val int) {
	this.len++
	if this.len%100 == 0 {
		this.vals = append(this.vals, make([]int, 100)...)
		this.serials = append(this.serials, make([]int, 100)...)
	}
	if val < 0 {
		val += 1000000007
	}
	this.vals[this.len-1] = val
	this.serials[this.len-1] = this.current_opt_serial_num

	this.current_opt_serial_num++
	if this.current_opt_serial_num%serial_opts_unit == 0 {
		this.serial_opts = append(this.serial_opts, make([]*operation, serial_opts_unit)...)
	}
	this.serial_opts[this.current_opt_serial_num-1] = nil
}

func (this *Fancy) AddAll(inc int) {
	if this.current_opt_serial_num > 0 && this.serial_opts[this.current_opt_serial_num-1] != nil {
		if this.serial_opts[this.current_opt_serial_num-1].method == "AddAll" {
			this.serial_opts[this.current_opt_serial_num-1].param = this.getMod(this.serial_opts[this.current_opt_serial_num-1].param + inc)
			return
		}
	}
	this.current_opt_serial_num++
	if this.current_opt_serial_num%serial_opts_unit == 0 {
		this.serial_opts = append(this.serial_opts, make([]*operation, serial_opts_unit)...)
	}
	this.serial_opts[this.current_opt_serial_num-1] = &operation{method: "AddAll", param: inc}
}

func (this *Fancy) MultAll(m int) {
	if this.current_opt_serial_num > 0 && this.serial_opts[this.current_opt_serial_num-1] != nil {
		if this.serial_opts[this.current_opt_serial_num-1].method == "MultAll" {
			this.serial_opts[this.current_opt_serial_num-1].param = this.getMod(this.serial_opts[this.current_opt_serial_num-1].param * m)
			return
		} else if this.serial_opts[this.current_opt_serial_num-1].method == "AddAll" {
			if this.current_opt_serial_num > 1 && this.serial_opts[this.current_opt_serial_num-2] != nil && this.serial_opts[this.current_opt_serial_num-2].method == "MultAll" {
				this.serial_opts[this.current_opt_serial_num-2].param = this.getMod(this.serial_opts[this.current_opt_serial_num-2].param * m)
				this.serial_opts[this.current_opt_serial_num-1].method = "AddAll"
				this.serial_opts[this.current_opt_serial_num-1].param = this.getMod(this.serial_opts[this.current_opt_serial_num-1].param * m)
			} else {
				this.current_opt_serial_num++
				if this.current_opt_serial_num%serial_opts_unit == 0 {
					this.serial_opts = append(this.serial_opts, make([]*operation, serial_opts_unit)...)
				}
				this.serial_opts[this.current_opt_serial_num-1] = &operation{method: "AddAll", param: this.getMod(this.serial_opts[this.current_opt_serial_num-2].param * m)}
				this.serial_opts[this.current_opt_serial_num-2].method = "MultAll"
				this.serial_opts[this.current_opt_serial_num-2].param = this.getMod(m)
			}
			return
		}
	}
	this.current_opt_serial_num++
	if this.current_opt_serial_num%serial_opts_unit == 0 {
		this.serial_opts = append(this.serial_opts, make([]*operation, serial_opts_unit)...)
	}
	this.serial_opts[this.current_opt_serial_num-1] = &operation{method: "MultAll", param: m}
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= this.len {
		return -1
	}
	val := this.vals[idx]
	serial := this.serials[idx]
	for _, opt := range this.serial_opts[serial:] {
		if opt == nil {
			continue
		}
		switch opt.method {
		case "AddAll":
			val += opt.param
		case "MultAll":
			val *= opt.param
		}
		val = this.getMod(val)
	}
	return val
}

func (this *Fancy) getMod(value int) int {
	if value >= 0 && value < 1000000007 {
		return value
	} else if value < 0 && value > -1000000007 {
		return value + 1000000007
	} else if value >= 1000000007 && value < 1000000007*2 {
		return value - 1000000007
	}
	val := value % 1000000007
	if val < 0 {
		val += 1000000007
	}
	return val
}
