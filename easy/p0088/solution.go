package p0088

func merge(lhs *[]int, lhsLen int, rhs []int, rhsLen int) {
	insertIdx := len(*lhs) - 1
	lhsLen--
	rhsLen--

	for lhsLen >= 0 && rhsLen >= 0 {
		if (*lhs)[lhsLen] > rhs[rhsLen] {
			(*lhs)[insertIdx] = (*lhs)[lhsLen]
			lhsLen--
		} else {
			(*lhs)[insertIdx] = rhs[rhsLen]
			rhsLen--
		}
		insertIdx--
	}

	for lhsLen >= 0 {
		(*lhs)[insertIdx] = (*lhs)[lhsLen]
		lhsLen--
		insertIdx--
	}

	for rhsLen >= 0 {
		(*lhs)[insertIdx] = rhs[rhsLen]
		rhsLen--
		insertIdx--
	}
}
