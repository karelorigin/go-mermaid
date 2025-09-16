package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/flowchart"
)

func main() {
	diagram := flowchart.NewFlowchart()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Software Development Process")

	// Define nodes with various shapes
	start := diagram.NewNode("Start Project")
	start.SetShape(flowchart.NodeShapeTerminal)

	requirements := diagram.NewNode("Gather Requirements")
	requirements.SetShape(flowchart.NodeShapeDocument)

	design := diagram.NewNode("System Design")
	design.SetShape(flowchart.NodeShapeProcess)

	database := diagram.NewNode("Database Design")
	database.SetShape(flowchart.NodeShapeDatabase)

	coding := diagram.NewNode("Implementation")
	coding.SetShape(flowchart.NodeShapeProcess)

	testing := diagram.NewNode("Testing")
	testing.SetShape(flowchart.NodeShapePrepare)

	bugs := diagram.NewNode("Bugs Found?")
	bugs.SetShape(flowchart.NodeShapeDecision)

	deploy := diagram.NewNode("Deployment")
	deploy.SetShape(flowchart.NodeShapeInternalStorage)

	monitor := diagram.NewNode("Monitoring")
	monitor.SetShape(flowchart.NodeShapeDisplay)

	end := diagram.NewNode("End")
	end.SetShape(flowchart.NodeShapeTerminal)

	// Add links
	diagram.NewLink(start, requirements)
	diagram.NewLink(requirements, design)
	diagram.NewLink(design, database)
	diagram.NewLink(database, coding)
	diagram.NewLink(coding, testing)
	diagram.NewLink(testing, bugs)
	diagram.NewLink(bugs, coding)
	diagram.NewLink(bugs, deploy)
	diagram.NewLink(deploy, monitor)
	diagram.NewLink(monitor, end)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
