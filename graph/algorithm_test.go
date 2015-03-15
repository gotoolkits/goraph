package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestBfs(t *testing.T) {
	for _, graph := range testdata.GraphSlice {
		file, err := os.Open("../testdata/data.json")
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		defer file.Close()
		data, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := data.Bfs(oneNode)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}

func TestDfsStack(t *testing.T) {
	for _, graph := range testdata.GraphSlice {
		file, err := os.Open("../testdata/data.json")
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		defer file.Close()
		data, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := data.DfsStack(oneNode)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}

func TestDfs(t *testing.T) {
	for _, graph := range testdata.GraphSlice {
		file, err := os.Open("../testdata/data.json")
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		defer file.Close()
		data, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := []*Node{}
		data.Dfs(oneNode, &rs)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}

func TestTopologicalDag01(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_01")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag := data.TopologicalDag()
	if rs != nil || isDag {
		t.Errorf("test_graph_01 has a cycle (not a DAG). Expected nil and false but %+v %v", rs, isDag)
	}
}

func TestTopologicalDag02(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag := data.TopologicalDag()
	if rs != nil || isDag {
		t.Errorf("test_graph_02 has a cycle (not a DAG). Expected nil and false but %+v %v", rs, isDag)
	}
}

func TestTopologicalDag06(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_06")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag6 := data.TopologicalDag()
	if rs == nil || !isDag6 {
		t.Errorf("test_graph_06 has no cycle (DAG). Not expected nil or false but %+v %v", rs, isDag6)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestTopologicalDag07(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_07")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag := data.TopologicalDag()
	if rs == nil || !isDag {
		t.Errorf("test_graph_07 has no cycle (DAG). Not expected nil or false but %+v %v", rs, isDag)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestTopologicalDag08(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_08")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag := data.TopologicalDag()
	if rs != nil || isDag {
		t.Errorf("test_graph_08 has a cycle (not a DAG). Expected nil and false but %+v %v", rs, isDag)
	}
}

func TestTopologicalDag09(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data, err := FromJSON(file, "test_graph_09")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs, isDag := data.TopologicalDag()
	if rs != nil || isDag {
		t.Errorf("test_graph_09 has a cycle (not a DAG). Expected nil and false but %+v %v", rs, isDag)
	}
}
