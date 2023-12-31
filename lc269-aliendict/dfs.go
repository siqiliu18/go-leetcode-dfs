package lc269aliendict

/* https://www.youtube.com/watch?v=ZU7fiX0WCCY
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"

t -> f

w -> e

r -> t

e -> r

w - e - r - t - f
*/

// This logic uses BFS, by starting with no parent nodes and put them into a queue
// loop through children and if they don't have any more parents (removed), add it to the queue
func AlienOrderBFS(words []string) string {

	// similar to other languages we want to create a map of string to set here
	// but since Go doesn't have set, we create a map of string to map of bool(doesn't matter) instead

	// collect child nodes of an existing node based on giving input
	nodeChildren := make(map[string]map[string]bool)

	// collect parent nodes of an existing node based on giving input
	nodeParents := make(map[string]map[string]bool)

	// init child 'set' and parent 'set' of each node
	for _, word := range words {
		for _, char := range word {
			nodeChildren[string(char)] = make(map[string]bool)
			nodeParents[string(char)] = make(map[string]bool)
		}
	}

	// loop through word list
	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		// ex: word1 = abc, word2 = ab, this claim is not following the sorted lexicographically rule
		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}

		// comparing chars between two words
		for j := 0; j < len(w1); j++ {
			// if not the same, then w2 char is the child or w1 char
			if w1[j] != w2[j] {
				nodeChildren[string(w1[j])][string(w2[j])] = true
				nodeParents[string(w2[j])][string(w1[j])] = true
				break
			}
		}
	}

	// create a array of string for appending only?
	res := ""

	// topological sort is bfs?
	queue := []string{}

	// extract nodes that have no parents
	no_par_node := []string{}
	for node, cs := range nodeParents {
		if len(cs) == 0 {
			no_par_node = append(no_par_node, node)
		}
	}

	// delete them from map and add them to queue to trigger the topological sort
	for _, node := range no_par_node {
		delete(nodeParents, node)
		queue = append(queue, node)
	}

	// if queue isn't empty
	for len(queue) != 0 {
		node := queue[0]  // pop the first node
		queue = queue[1:] // remove the first node from queue
		res += node       // add node to res arr

		// loop through children of the popped node
		for ch := range nodeChildren[node] {
			delete(nodeParents[ch], node)  // delete node from current child's parent list
			if len(nodeParents[ch]) == 0 { // if current child has no parent anymore
				queue = append(queue, ch) // add current child to queue
				delete(nodeParents, ch)   // delete current child from parent map
			}
		}
	}

	// if the map for storing parents of a node isn't empty, meaning there was invalid relationship
	if len(nodeParents) != 0 {
		return ""
	}

	return res
}

func AlienOrderPostOrderDFS(words []string) string {

	// similar to other languages we want to create a map of string to set here
	// but since Go doesn't have set, we create a map of string to map of bool(doesn't matter) instead

	// collect child nodes of an existing node based on giving input
	nodeChildren := make(map[string]map[string]bool)

	// collect parent nodes of an existing node based on giving input
	nodeParents := make(map[string]map[string]bool)

	// init child 'set' and parent 'set' of each node
	for _, word := range words {
		for _, char := range word {
			nodeChildren[string(char)] = make(map[string]bool)
			nodeParents[string(char)] = make(map[string]bool)
		}
	}

	visited := make(map[string]bool)

	// loop through word list
	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		// ex: word1 = abc, word2 = ab, this claim is not following the sorted lexicographically rule
		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}

		// comparing chars between two words
		for j := 0; j < len(w1); j++ {
			// if not the same, then w2 char is the child or w1 char
			if w1[j] != w2[j] {
				visited[string(w1[j])] = false
				visited[string(w2[j])] = false
				nodeChildren[string(w1[j])][string(w2[j])] = true
				nodeParents[string(w2[j])][string(w1[j])] = true
				break
			}
		}
	}

	res := ""

	stack := []string{"0"}
	nodeChildren["0"] = make(map[string]bool)

	// extract nodes that have no parents
	for node, cs := range nodeParents {
		if len(cs) == 0 {
			nodeChildren["0"][node] = true
		}
	}

	if _, ok := nodeChildren["0"]; !ok {
		return ""
	}

	// visited["0"] = true

	// if queue isn't empty
	for len(stack) != 0 {
		node := stack[len(stack)-1] // pop the last node

		if len(nodeChildren[node]) == 0 {
			stack = stack[:len(stack)-1]
		}

		for ch := range nodeChildren[node] {
			if len(nodeChildren[ch]) == 0 {
				if !visited[ch] {
					res += ch // missing the begging node
				}
				delete(nodeParents[ch], node)
				delete(nodeChildren[node], ch) // most important
				delete(nodeChildren, ch)
				delete(nodeParents, ch)
				visited[ch] = true
			} else {
				stack = append(stack, ch)
			}
		}
	}

	return Reverse(res)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
