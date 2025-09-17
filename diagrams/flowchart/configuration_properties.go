package flowchart

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// A set of flowchart property constants.
const (
	BaseFlowchartConfigurationProperties string = basediagram.Indentation + "flowchart:\n"
	FlowchartPropertyTitleTopMargin      string = "titleTopMargin"
	FlowchartPropertyDiagramPadding      string = "diagramPadding"
	FlowchartPropertyHtmlLabels          string = "htmlLabels"
	FlowchartPropertyNodeSpacing         string = "nodeSpacing"
	FlowchartPropertyRankSpacing         string = "rankSpacing"
	FlowchartPropertyCurve               string = "curve"
	FlowchartPropertyPadding             string = "padding"
	FlowchartPropertyDefaultRenderer     string = "defaultRenderer"
	FlowchartPropertyWrappingWidth       string = "wrappingWidth"
	FlowchartPropertyArrowMarkerAbsolute string = "arrowMarkerAbsolute"
)

// FlowchartConfigurationProperties holds flowchart-specific configuration
type FlowchartConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewFlowchartConfigurationProperties() FlowchartConfigurationProperties {
	return FlowchartConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *FlowchartConfigurationProperties) SetTitleTopMargin(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyTitleTopMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyTitleTopMargin,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetDiagramPadding(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyDiagramPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyDiagramPadding,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetHtmlLabels(v bool) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyHtmlLabels] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyHtmlLabels,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetNodeSpacing(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyNodeSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyNodeSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetRankSpacing(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyRankSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyRankSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetCurve(v string) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyCurve] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyCurve,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetPadding(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetDefaultRenderer(v string) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyDefaultRenderer] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyDefaultRenderer,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetWrappingWidth(v int) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyWrappingWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyWrappingWidth,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetArrowMarkerAbsolute(v bool) *FlowchartConfigurationProperties {
	c.properties[FlowchartPropertyArrowMarkerAbsolute] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: FlowchartPropertyArrowMarkerAbsolute,
			Val:  v,
		},
	}
	return c
}

func (c FlowchartConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(BaseFlowchartConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
