package flowchart

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// A list of subgraph constants.
const (
	BaseSubgraphString          string = basediagram.Indentation + "subgraph %s [%s]\n"
	BaseSubgraphDirectionString string = basediagram.Indentation + "direction %s\n"
	BaseSubgraphEndString       string = basediagram.Indentation + "end\n"
	BaseSubgraphLinkString      string = basediagram.Indentation + "%s"
	BaseSubgraphSubgraphString  string = basediagram.Indentation + "%s"
)

// List of possible Subgraph directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#direction
const (
	SubgraphDirectionNone        SubgraphDirection = ""
	SubgraphDirectionTopToBottom SubgraphDirection = "TB"
	SubgraphDirectionBottomUp    SubgraphDirection = "BT"
	SubgraphDirectionRightLeft   SubgraphDirection = "RL"
	SubgraphDirectionLeftRight   SubgraphDirection = "LR"
)

// SubgraphDirection represents a subgraph direction.
type SubgraphDirection string

// Subgraph represents a subgraph.
type Subgraph struct {
	ID          string
	Title       string
	Direction   SubgraphDirection
	subgraphs   []*Subgraph
	links       []*Link
	idGenerator utils.IDGenerator
}

// NewSubgraph creates a new Subgraph with the given ID and title,
// setting the default direction to none.
func NewSubgraph(id string, title string) (newSubgraph *Subgraph) {
	newSubgraph = &Subgraph{
		ID:        id,
		Title:     title,
		Direction: SubgraphDirectionNone,
	}

	return
}

// AddSubgraph adds a new Subgraph to the current Subgraph and returns the created subgraph.
func (s *Subgraph) AddSubgraph(title string) (newSubgraph *Subgraph) {
	if s.idGenerator == nil {
		s.idGenerator = utils.NewIDGenerator()
	}

	newSubgraph = NewSubgraph(s.idGenerator.NextID(), title)
	newSubgraph.idGenerator = s.idGenerator

	s.subgraphs = append(s.subgraphs, newSubgraph)

	return
}

// AddLink adds a new Link to the Subgraph and returns the created link.
func (s *Subgraph) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	s.links = append(s.links, newLink)

	return
}

// String generates a Mermaid string representation of the Subgraph,
// including its subgraphs, direction, and links with the specified indentation.
func (s *Subgraph) String(curIndentation string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(BaseSubgraphString), s.ID, s.Title)))

	direction := ""
	if s.Direction != SubgraphDirectionNone {
		direction = fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(BaseSubgraphDirectionString), string(s.Direction)))
	}

	sb.WriteString(direction)

	for _, subgraph := range s.subgraphs {
		nextIndentation := fmt.Sprintf(string(BaseSubgraphSubgraphString), string(curIndentation))
		sb.WriteString(subgraph.String(nextIndentation))
	}

	for _, link := range s.links {
		sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(BaseSubgraphLinkString), link.String())))
	}

	sb.WriteString(fmt.Sprintf(string(curIndentation), BaseSubgraphEndString))

	return sb.String()
}
