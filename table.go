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

// TableData creates a table data cell in a table
func TableData(children ...g.Node) g.Node {
	return h.Td(h.Class("table-cell"), g.Group(children))
}
