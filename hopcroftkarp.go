package HopcroftKarp

type Vertex struct {
	vertex string
}
type HopcroftKarp struct {
	graph     map[*Vertex][]*Vertex // 原始的图
	matching  map[*Vertex]*Vertex   // 匹配的图
	dfsPaths  [][]*Vertex           // 走过的路径
	dfsParent map[*Vertex]*Vertex   // 上级节点
	left      map[*Vertex]bool      // 顶点
	right     map[*Vertex]bool      // 顶点的邻居
}

func NewHK(graph map[*Vertex][]*Vertex) *HopcroftKarp {
	hk := &HopcroftKarp{
		graph:    graph,
		matching: make(map[*Vertex]*Vertex, 2*len(graph)),
		left:     make(map[*Vertex]bool, len(graph)),
		right:    make(map[*Vertex]bool, len(graph)),
	}
	hk.init()
	return hk
}

// init 初始化左、右节点
func (self *HopcroftKarp) init() {
	for k, v := range self.graph {
		self.left[k] = true
		for _, v1 := range v {
			self.right[v1] = true
		}
	}
	for v := range self.left {
		for _, neighbour := range self.graph[v] {
			if _, ok := self.graph[neighbour]; !ok {
				self.graph[neighbour] = []*Vertex{v}
			} else {
				self.graph[neighbour] = append(self.graph[neighbour], v)
			}
		}
	}
}

// bfs 广度优先搜索
func (self *HopcroftKarp) bfs() []map[*Vertex]bool {
	var layers = make([]map[*Vertex]bool, 0, len(self.graph))
	var layer = make(map[*Vertex]bool, len(self.graph))

	for v := range self.left {
		if _, ok := self.matching[v]; !ok {
			layer[v] = true
		}
	}
	layers = append(layers, layer)
	var visited = make(map[*Vertex]bool, len(self.graph))
	for {
		layer = layers[len(layers)-1]
		var newLayer = make(map[*Vertex]bool, len(self.graph))
		for v := range layer {
			flag := false
			for v2 := range self.left {
				if v == v2 {
					flag = true
					break
				}
			}
			n, ok := self.matching[v]
			visited[v] = true
			if flag {
				//visited.Add(v)
				for _, neighbour := range self.graph[v] {
					if !visited[neighbour] && (!ok || neighbour != n) {
						newLayer[neighbour] = true
					}
				}
			} else {
				//visited.Add(v)
				for _, neighbour := range self.graph[v] {
					if !visited[neighbour] && (ok && neighbour == n) {
						newLayer[neighbour] = true
					}
				}
			}
		}
		layers = append(layers, newLayer)
		if len(newLayer) == 0 {
			return layers
		}
		for v := range newLayer {
			_, ok := self.matching[v]
			if self.right[v] && !ok {
				return layers
			}
		}
	}
}

// dfs 深度优先搜索
func (self *HopcroftKarp) dfs(v *Vertex, index int, layers []map[*Vertex]bool) bool {
	if index == 0 {
		var path = make([]*Vertex, 1, len(layers))
		path[0] = v
		for {
			if self.dfsParent[v] != v {
				path = append(path, self.dfsParent[v])
				v = self.dfsParent[v]
			} else {
				break
			}
		}
		self.dfsPaths = append(self.dfsPaths, path)
		return true
	}

	for _, neighbour := range self.graph[v] {
		if layers[index-1][neighbour] {
			if _, ok := self.dfsParent[neighbour]; ok {
				continue
			}
			n, ok := self.matching[v]
			if (self.left[neighbour] && (!ok || neighbour != n)) ||
				(self.right[neighbour] && (ok && neighbour == n)) {
				self.dfsParent[neighbour] = v
				if self.dfs(neighbour, index-1, layers) {
					return true
				}
			}
		}
	}
	return false
}

// MaximumMatching 最大匹配
func (self *HopcroftKarp) MaximumMatching() map[*Vertex]*Vertex {
	for {
		layers := self.bfs()
		if len(layers[len(layers)-1]) == 0 {
			break
		}
		var freeVertex = make(map[*Vertex]bool, len(self.graph))
		for v := range layers[len(layers)-1] {
			if _, ok := self.matching[v]; !ok {
				freeVertex[v] = true
			}
		}
		self.dfsPaths = make([][]*Vertex, 0, len(layers))
		self.dfsParent = make(map[*Vertex]*Vertex, len(layers))

		for v := range freeVertex {
			self.dfsParent[v] = v
			self.dfs(v, len(layers)-1, layers)
		}

		if len(self.dfsPaths) == 0 {
			break
		}

		for _, path := range self.dfsPaths {
			for i := 0; i < len(path); i++ {
				if i%2 == 0 {
					self.matching[path[i]] = path[i+1]
					self.matching[path[i+1]] = path[i]
				}
			}
		}
	}
	return self.matching
}
