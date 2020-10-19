package LeetCode731

type MyCalendarTwo struct {
	calendars       []time
	calendarsRepeat []time
}
type time struct {
	start int
	end   int
}

func Constructor() MyCalendarTwo {
	var myCalendarTwo = MyCalendarTwo{calendars: make([]time, 0), calendarsRepeat: make([]time, 0)}
	return myCalendarTwo
}

func getRepeatTime(s0, e0, s1, e1 int) (start, end int) {
	start, end = s0, e0
	if start < s1 {
		start = s1
	}
	if end > e1 {
		end = e1
	}
	return
}
func getSplitTime(s0, e0, s1, e1 int) (res []time) {
	if s0 < s1 {
		res = append(res, time{s0, s1})
	}
	if e0 > e1 {
		res = append(res, time{e1, e0})
	}
	return
}
func (this *MyCalendarTwo) Book(start int, end int) bool {
	oldCalendars, oldCalendarsRepeat := this.calendars, this.calendarsRepeat
	if this.BookFirst(start, end) {
		return true
	} else {
		this.calendars, this.calendarsRepeat = oldCalendars, oldCalendarsRepeat
		return false
	}
}
func (this *MyCalendarTwo) BookFirst(start int, end int) bool {
	for _, t := range this.calendars {
		if end <= t.start || start >= t.end {
			continue
		}
		s, e := getRepeatTime(start, end, t.start, t.end)
		if !this.BookRepeat(s, e) {
			return false
		}
		this.calendarsRepeat = append(this.calendarsRepeat, time{s, e})
		tmp := getSplitTime(start, end, t.start, t.end)
		if nil == tmp {
			return true
		}
		for _, tp := range tmp {
			if !this.BookFirst(tp.start, tp.end) {
				return false
			}
		}
		return true
	}
	this.calendars = append(this.calendars, time{start: start, end: end})
	return true
}

func (this *MyCalendarTwo) BookRepeat(start int, end int) bool {
	for _, v := range this.calendarsRepeat {
		if !(start >= v.end || end <= v.start) {
			return false
		}
	}
	return true
}
