package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"sample/graph"
	"strconv"
)
//リクエストデータの構造
type RequestData struct {
	Num   int
	Paths []Path
	Start int
	End   int
}
//経路の構造
type Path struct {
	From int
	To   int
	Dist int
}
//関数　ホーム
func home(w http.ResponseWriter,
	r *http.Request) {
	t, _ := template.ParseFiles("canvas.html")
	t.Execute(w, nil)
}
//関数　ダイクストラ法
func dijkstra(data RequestData) string {
	n := data.Num
	paths := data.Paths
	start := data.Start
	end := data.End

	visited := make([]bool, n)
	dist := make([]int, n)
	root := make([]int, n)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0
	root[start] = -1

	var pq graph.PQ
	var dummy graph.Node

	dummy.Id = -1
	dummy.Dist = -1

	pq = append(pq, dummy)
	var node graph.Node
	node.Id = start
	node.Dist = dist[start]
	pq = append(pq, node)
	for len(pq) > 1 {
		node = graph.Extract(&pq)
		visited[node.Id] = true
		for i := 0; i < len(paths); i++ {
			if paths[i].From == node.Id {
				to := paths[i].To
				if visited[to] {
					continue
				}
				if dist[to] > dist[node.Id]+paths[i].Dist {
					dist[to] = dist[node.Id] + paths[i].Dist
					var toNode graph.Node
					toNode.Id = to
					toNode.Dist = dist[to]
					graph.Insert(&pq, toNode)
					root[to] = node.Id
				}
			}
		}
	}
	path := root[end]
	rootStr := strconv.Itoa(end)
	for path != -1 {
		rootStr += "," + strconv.Itoa(path)
		path = root[path]
	}
	fmt.Println(rootStr)
	json := "{\"root\":\"" + rootStr + "\",\"dist\":" + strconv.Itoa(dist[end]) + "}"
	return json

}
//関数　計算
func calc(w http.ResponseWriter,
	r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var data RequestData
	json.Unmarshal(body, &data)

	json := dijkstra(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(json))
}
//関数　メイン
func main() {
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
