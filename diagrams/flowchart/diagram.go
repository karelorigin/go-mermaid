// Package flowchart provides functionality for creating Mermaid flowcharts
package flowchart

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// FlowChartDirection represents the direction of a flowchart.
type FlowchartDirection string

// CurveStyle represents a curve style.
type CurveStyle string

// List of possible Flowchart directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#direction
const (
	FlowchartDirectionTopToBottom FlowchartDirection = "TB"
	FlowchartDirectionTopDown     FlowchartDirection = "TD"
	FlowchartDirectionBottomUp    FlowchartDirection = "BT"
	FlowchartDirectionRightLeft   FlowchartDirection = "RL"
	FlowchartDirectionLeftRight   FlowchartDirection = "LR"
)

// List of possible Flowchart directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#styling-line-curves
const (
	CurveStyleNone       CurveStyle = ""
	CurveStyleBasis      CurveStyle = "basis"
	CurveStyleBumpX      CurveStyle = "bumpX"
	CurveStyleBumpY      CurveStyle = "bumpY"
	CurveStyleCardinal   CurveStyle = "cardinal"
	CurveStyleCatmullRom CurveStyle = "catmullRom"
	CurveStyleLinear     CurveStyle = "linear"
	CurveStyleMonotoneX  CurveStyle = "monotoneX"
	CurveStyleMonotoneY  CurveStyle = "monotoneY"
	CurveStyleNatural    CurveStyle = "natural"
	CurveStyleStep       CurveStyle = "step"
	CurveStyleStepAfter  CurveStyle = "stepAfter"
	CurveStyleStepBefore CurveStyle = "stepBefore"
)

const (
	baseFlowchartDirectionString string = "flowchart %s\n"
)

// Flowcharts are composed of nodes (geometric shapes) and links (arrows or lines).
// The Mermaid code defines how nodes and links are made and accommodates different arrow types,
// multi-directional arrows, and any linking to and from subgraphs.
// Reference: https://mermaid.js.org/syntax/flowchart.html
type Flowchart struct {
	basediagram.BaseDiagram[FlowchartConfigurationProperties]
	Direction   FlowchartDirection
	CurveStyle  CurveStyle
	classes     []*Class
	nodes       []*Node
	subgraphs   []*Subgraph
	links       []*Link
	idGenerator utils.IDGenerator
}

// NewFlowchart creates a new flowchart diagram
func NewFlowchart() *Flowchart {
	return &Flowchart{
		BaseDiagram: basediagram.NewBaseDiagram(NewFlowchartConfigurationProperties()),
		Direction:   FlowchartDirectionTopToBottom,
		CurveStyle:  CurveStyleNone,
		classes:     make([]*Class, 0),
		nodes:       make([]*Node, 0),
		subgraphs:   make([]*Subgraph, 0),
		links:       make([]*Link, 0),
		idGenerator: utils.NewIDGenerator(),
	}
}

// SetDirection sets the flowchart direction and returns the flowchart for chaining
func (f *Flowchart) SetDirection(direction FlowchartDirection) *Flowchart {
	f.Direction = direction
	return f
}

// RenderToFile saves the flowchart diagram to a file at the specified path.
func (f *Flowchart) RenderToFile(path string) error {
	return utils.RenderToFile(path, f.String())
}

// AddSubgraph adds a new subgraph to the flowchart and returns the created subgraph.
func (f *Flowchart) AddSubgraph(title string) (newSubgraph *Subgraph) {
	newSubgraph = NewSubgraph(f.idGenerator.NextID(), title)

	f.subgraphs = append(f.subgraphs, newSubgraph)

	return
}

// AddNode adds a given node to the flowchart.
func (f *Flowchart) AddNode(node *Node) {
	f.nodes = append(f.nodes, node)
}

// NewNode adds a new node to the flowchart and returns the created node.
func (f *Flowchart) NewNode(text string) (newNode *Node) {
	newNode = NewNode(f.idGenerator.NextID(), text)

	f.nodes = append(f.nodes, newNode)

	return
}

// AddLink adds a link between two nodes to the flowchart.
func (f *Flowchart) AddLink(link *Link) {
	f.links = append(f.links, link)
}

// NewLink adds a new link between two nodes in the flowchart and returns the created link.
func (f *Flowchart) NewLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	f.links = append(f.links, newLink)

	return
}

// AddClass adds a new class to the flowchart and returns the created class.
func (f *Flowchart) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	f.classes = append(f.classes, newClass)

	return
}

// String generates a Mermaid flowchart string representation
func (f *Flowchart) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseFlowchartDirectionString), string(f.Direction)))

	for _, class := range f.classes {
		sb.WriteString(class.String())
	}

	for _, node := range f.nodes {
		sb.WriteString(node.String())
	}

	for _, subgraph := range f.subgraphs {
		sb.WriteString(subgraph.String("%s"))
	}

	for _, link := range f.links {
		sb.WriteString(link.String())
	}

	return f.BaseDiagram.String(sb.String())
}
