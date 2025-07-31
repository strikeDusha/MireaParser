package parser

import "time"

type Page struct {
	Applicants int
	Time       time.Time
	Planned    int
	Title      string
	List       List
}
type Student struct {
	Id         int
	Priority   int
	Acceptance bool
	Sum        int
	IHP        bool // low p
	IHPO       bool // high p
}
type List []Student

// funcs for later usage, now i just wanna work with excel
// maybe later project will have gui in web but now ehh - but in this iteration i will have same functionality in excel

// func (s List) isSorted() bool {

// 	for i := 1; i < len(s); i++ {
// 		if !(s[i-1].Sum < s[i].Sum) {
// 			return false
// 		}
// 	}
// 	return true
// }
// func (s List) sort() {
// 	for !s.isSorted() {
// 		for i := 1; i < len(s); i++ {
// 			if s[i-1].Sum > s[i].Sum {
// 				s[i], s[i-1] = s[i-1], s[i]
// 			}
// 		}
// 	}
// }

// func Acceptionts(s List) int {
// 	i := 0
// 	for _, v := range s {
// 		if v.Acceptance {
// 			i++
// 		}
// 	}
// 	return i
// }

// func FilterByPriority(s List, pr int) List {
// 	j := make(List, 0, len(s))
// 	for _, v := range s {
// 		if v.Priority <= pr {
// 			j = append(j, v)
// 		}
// 	}
// 	return j
// }

// func (s List) StudentPlace(st Student, pr int) (place int) {

// 	if pr == 0 {
// 		for _, v := range s {
// 			if st.Sum < v.Sum {
// 				place++
// 			}
// 		}
// 		return place
// 	} else {
// 		u := FilterByPriority(s, pr)
// 		for _, v := range u {
// 			if st.Sum < v.Sum {
// 				place++
// 			}
// 		}
// 		return place
// 	}
// }
