package main

// dfs
// time complexity: O(e + q*e) where e is length of equations, q is length of queries
// space complexity: O(e)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adList := make(map[string]map[string]float64)
	for i, equation := range equations {
		if _, ok := adList[equation[0]]; !ok {
			adList[equation[0]] = make(map[string]float64)
		}
		if _, ok := adList[equation[1]]; !ok {
			adList[equation[1]] = make(map[string]float64)
		}

		adList[equation[0]][equation[1]] = values[i]
		adList[equation[1]][equation[0]] = 1 / values[i]
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := adList[query[0]]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := adList[query[1]]; !ok {
			res[i] = -1
			continue
		}

		res[i] = divide(query[0], query[1], adList, map[string]bool{})
	}

	return res
}

func divide(dividend, divisor string, adList map[string]map[string]float64, visited map[string]bool) float64 {
	if dividend == divisor {
		return 1
	}

	visited[dividend] = true
	for nb := range adList[dividend] {
		if visited[nb] {
			continue
		}

		res := divide(nb, divisor, adList, visited)
		if res > 0 {
			return adList[dividend][nb] * res
		}

	}

	return -1
}

type Group struct {
	group  string
	weight float64
}

// find will return the group id and weight that the variable belongs to.
// moreover, will update the states of variables along the chain, if there is any discrepancy.
func find(groupID string, groups map[string]Group) (string, float64) {
	if _, ok := groups[groupID]; !ok {
		groups[groupID] = Group{groupID, 1}
	}

	grp := groups[groupID]
	if grp.group != groupID {
		parentGroup, parentWeight := find(grp.group, groups)
		groups[groupID] = Group{parentGroup, parentWeight * grp.weight}
	}

	return groups[groupID].group, groups[groupID].weight
}

// union will attach the group of dividend to that of the divisor, if they are not already the same group.
// in addition, it needs to update the weight of the dividend variable accordingly, so that the ratio between
// the dividend and divisor is respected.
func union(dividend, divisor string, quotient float64, groups map[string]Group) {
	dividendGroup, dividendWeight := find(dividend, groups)
	divisorGroup, divisorWeight := find(divisor, groups)
	if dividendGroup != divisorGroup {
		groups[dividendGroup] = Group{divisorGroup, divisorWeight * quotient / dividendWeight}
	}
}

// union find
// time complexity: O(e + q)
// space complexity: O(e)
func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	groups := make(map[string]Group)
	for i, equation := range equations {
		union(equation[0], equation[1], values[i], groups)
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := groups[query[0]]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := groups[query[1]]; !ok {
			res[i] = -1
			continue
		}

		dividendGroup, dividendWeight := find(query[0], groups)
		divisorGroup, divisorWeight := find(query[1], groups)
		if dividendGroup != divisorGroup {
			res[i] = -1
		} else {
			res[i] = dividendWeight / divisorWeight
		}

	}

	return res
}
