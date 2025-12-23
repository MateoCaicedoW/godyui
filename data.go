package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// TableEl creates a styled table container
func TableEl(children ...g.Node) g.Node {
	return h.Table(h.Class("table"), g.Group(children))
}

// TableHeader creates the header section of a table
func TableHeader(children ...g.Node) g.Node {
	return h.THead(h.Class("table-header"), g.Group(children))
}

// TableBody creates the main content section of a table
func TableBody(children ...g.Node) g.Node {
	return h.TBody(h.Class("table-body"), g.Group(children))
}

// TableFooter creates the footer section of a table
func TableFooter(children ...g.Node) g.Node {
	return h.TFoot(h.Class("table-footer"), g.Group(children))
}

// TableRow creates a single row in a table
func TableRow(children ...g.Node) g.Node {
	return h.Tr(h.Class("table-row"), g.Group(children))
}

// TableHead creates a header cell in a table
func TableHead(children ...g.Node) g.Node {
	return h.Th(h.Class("table-head"), g.Group(children))
}

// TableCell creates a data cell in a table
func TableCell(children ...g.Node) g.Node {
	return h.Td(h.Class("table-cell"), g.Group(children))
}

// TableCaption creates a caption for a table
func TableCaption(children ...g.Node) g.Node {
	return h.Caption(h.Class("table-caption"), g.Group(children))
}

// Accordion components for collapsible sections
// Accordion creates a collapsible section container
func Accordion(children ...g.Node) g.Node {
	return h.Div(h.Class("accordion"), g.Group(children))
}

// AccordionItem creates a single accordion section
func AccordionItem(children ...g.Node) g.Node {
	return h.Div(h.Class("accordion-item"), g.Group(children))
}

// AccordionTrigger creates a button to expand/collapse accordion
func AccordionTrigger(children ...g.Node) g.Node {
	return h.Button(h.Class("accordion-trigger"), g.Group(children))
}

// AccordionContent creates the expandable content area
func AccordionContent(children ...g.Node) g.Node {
	return h.Div(h.Class("accordion-content"), g.Group(children))
}

// Collapsible components for show/hide functionality
// Collapsible creates a generic show/hide container
func Collapsible(children ...g.Node) g.Node {
	return h.Div(h.Class("collapsible"), g.Group(children))
}

// CollapsibleTrigger creates a button to toggle visibility
func CollapsibleTrigger(children ...g.Node) g.Node {
	return h.Button(h.Class("collapsible-trigger"), g.Group(children))
}

// CollapsibleContent creates the toggleable content area
func CollapsibleContent(children ...g.Node) g.Node {
	return h.Div(h.Class("collapsible-content"), g.Group(children))
}
