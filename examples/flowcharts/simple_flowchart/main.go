package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/flowchart"
)

func main() {
	// Create a new flowchart
	diagram := flowchart.NewFlowchart()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Simple Process Flow")

	// Add nodes with different shapes
	start := diagram.NewNode("Start")
	start.SetShape(flowchart.NodeShapeTerminal)

	input := diagram.NewNode("Get User Input")
	input.SetShape(flowchart.NodeShapeManualInput)

	process := diagram.NewNode("Process Data")
	process.SetShape(flowchart.NodeShapeProcess)

	decision := diagram.NewNode("Valid?")
	decision.SetShape(flowchart.NodeShapeDecision)

	output := diagram.NewNode("Display Result")
	output.SetShape(flowchart.NodeShapeDisplay)

	end := diagram.NewNode("End")
	end.SetShape(flowchart.NodeShapeTerminal)

	// Add links between nodes
	diagram.NewLink(start, input)
	diagram.NewLink(input, process)
	diagram.NewLink(process, decision)
	diagram.NewLink(decision, output)
	diagram.NewLink(decision, input)
	diagram.NewLink(output, end)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
