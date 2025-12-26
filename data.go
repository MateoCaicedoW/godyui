package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

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
