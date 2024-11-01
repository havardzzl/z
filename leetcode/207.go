package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	preReqNum := map[int]int{}
	preWorks := map[int][]int{}
	for _, pre := range prerequisites {
		preReqNum[pre[0]] += 1
		preWorks[pre[1]] = append(preWorks[pre[1]], pre[0])
	}
	candi := []int{}
	for i := 0; i < numCourses; i++ {
		if preReqNum[i] == 0 {
			candi = append(candi, i)
		}
	}
	order := []int{}
	var getOrder func()
	getOrder = func() {
		if len(candi) == 0 {
			return
		}
		work := candi[len(candi)-1]
		candi = candi[:len(candi)-1]
		order = append(order, work)
		for _, w := range preWorks[work] {
			preReqNum[w]--
			if preReqNum[w] == 0 {
				candi = append(candi, w)
			}
		}
		getOrder()
	}
	getOrder()
	return len(order) == numCourses
}
