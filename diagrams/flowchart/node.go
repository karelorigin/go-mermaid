package flowchart

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// NodeShape represents a node shape.
type NodeShape string

// List of possible Node shapes.
// Reference: https://mermaid.js.org/syntax/flowchart.html#complete-list-of-new-shapes
const (
	// Basic shapes
	NodeShapeProcess          NodeShape = `rect`       // Process (Rectangle)
	NodeShapeEvent            NodeShape = `rounded`    // Event (Rounded rectangle)
	NodeShapeTerminal         NodeShape = `stadium`    // Terminal (Stadium-shaped)
	NodeShapeSubprocess       NodeShape = `fr-rect`    // Subprocess (Framed rectangle)
	NodeShapeDatabase         NodeShape = `cyl`        // Database (Cylinder)
	NodeShapeStart            NodeShape = `circle`     // Start (Circle)
	NodeShapeOdd              NodeShape = `odd`        // Odd shape (Asymmetric)
	NodeShapeDecision         NodeShape = `diam`       // Decision (Diamond)
	NodeShapePrepare          NodeShape = `hex`        // Prepare (Hexagon)
	NodeShapeInputOutput      NodeShape = `lean-r`     // Input/Output (Parallelogram)
	NodeShapeOutputInput      NodeShape = `lean-l`     // Output/Input (Alt Parallelogram)
	NodeShapeManualOperation  NodeShape = `trap-b`     // Manual Operation (Trapezoid)
	NodeShapeManual           NodeShape = `trap-t`     // Manual (Alt Trapezoid)
	NodeShapeStopDouble       NodeShape = `dbl-circ`   // Stop (Double Circle)
	NodeShapeText             NodeShape = `text`       // Text block
	NodeShapeCard             NodeShape = `notch-rect` // Card
	NodeShapeLinedProcess     NodeShape = `lin-rect`   // Lined Process (Rectangle with shadow)
	NodeShapeStartSmall       NodeShape = `sm-circ`    // Start (Small circle)
	NodeShapeStopFramed       NodeShape = `fr-circ`    // Stop (Circle with frame)
	NodeShapeForkJoin         NodeShape = `fork`       // Fork/Join
	NodeShapeCollate          NodeShape = `hourglass`  // Collate (Hourglass)
	NodeShapeComment          NodeShape = `brace`      // Comment (Left brace)
	NodeShapeCommentRight     NodeShape = `brace-r`    // Comment Right (Right brace)
	NodeShapeCommentBothSides NodeShape = `braces`     // Comment (Both braces)
	NodeShapeComLink          NodeShape = `bolt`       // Com Link (Lightning bolt)
	NodeShapeDocument         NodeShape = `doc`        // Document
	NodeShapeDelay            NodeShape = `delay`      // Delay
	NodeShapeStorage          NodeShape = `h-cyl`      // Storage (Horizontal cylinder)
	NodeShapeDiskStorage      NodeShape = `lin-cyl`    // Disk Storage (Lined cylinder)
	NodeShapeDisplay          NodeShape = `curv-trap`  // Display (Curved trapezoid)
	NodeShapeDividedProcess   NodeShape = `div-rect`   // Divided Process
	NodeShapeExtract          NodeShape = `tri`        // Extract (Triangle)
	NodeShapeInternalStorage  NodeShape = `win-pane`   // Internal Storage
	NodeShapeJunction         NodeShape = `f-circ`     // Junction (Filled circle)
	NodeShapeLinedDocument    NodeShape = `lin-doc`    // Lined Document
	NodeShapeLoopLimit        NodeShape = `notch-pent` // Loop Limit
	NodeShapeManualFile       NodeShape = `flip-tri`   // Manual File
	NodeShapeManualInput      NodeShape = `sl-rect`    // Manual Input (Sloped rectangle)
	NodeShapeMultiDocument    NodeShape = `docs`       // Multi-Document
	NodeShapeMultiProcess     NodeShape = `st-rect`    // Multi-Process
	NodeShapePaperTape        NodeShape = `flag`       // Paper Tape
	NodeShapeStoredData       NodeShape = `bow-rect`   // Stored Data
	NodeShapeSummary          NodeShape = `cross-circ` // Summary (Circle with cross)
	NodeShapeTaggedDocument   NodeShape = `tag-doc`    // Tagged Document
	NodeShapeTaggedProcess    NodeShape = `tag-rect`   // Tagged Process
)

const (
	baseNodeShapeString string = basediagram.Indentation + "%s@{ shape: %s, label: \"%s\"}"
	baseNodeClassString string = ":::%s"
	baseNodeStyleString string = basediagram.Indentation + "style %s %s\n"
)

// Node represents a node in a flowchart
type Node struct {
	ID    string
	Shape NodeShape
	Text  string
	Style *NodeStyle
	Class *Class
}

// NewNode creates a new Node with the given ID and text, setting default shape to round edges.
func NewNode(id string, text string) (newNode *Node) {
	newNode = &Node{
		ID:    id,
		Text:  text,
		Shape: NodeShapeProcess,
	}

	return
}

// SetClass sets the node class and returns the node for chaining
func (n *Node) SetClass(class *Class) *Node {
	n.Class = class
	return n
}

// SetText sets the node text and returns the node for chaining
func (n *Node) SetText(text string) *Node {
	n.Text = text
	return n
}

// SetStyle sets the style for the node and returns the node for chaining
func (n *Node) SetStyle(style *NodeStyle) *Node {
	n.Style = style
	return n
}

// SetShape sets the node shape and returns the node for chaining
func (n *Node) SetShape(shape NodeShape) *Node {
	n.Shape = shape
	return n
}

// String generates a Mermaid string representation of the node, including its shape, class, and style.
func (n *Node) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseNodeShapeString), n.ID, string(n.Shape), escape(n.Text)))

	if n.Class != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeClassString), n.Class.Name))
	}

	sb.WriteByte('\n')

	if n.Style != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleString), n.ID, n.Style.String()))
	}

	return sb.String()
}

// escape escapes quotes so that it can be inserted into a mermaid property.
func escape(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
