package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DialogEl creates a modal dialog overlay
func DialogEl(open bool, children ...g.Node) g.Node {
	attrs := []g.Node{h.Class("dialog")}
	if open {
		attrs = append(attrs, g.Attr("open"))
	}
	attrs = append(attrs, g.Group(children))
	return h.Dialog(attrs...)
}

// DialogContent creates the main content area of a dialog
func DialogContent(children ...g.Node) g.Node {
	return h.Div(h.Class("dialog-content"), g.Group(children))
}

// DialogHeader creates the top section of a dialog
func DialogHeader(children ...g.Node) g.Node {
	return h.Div(h.Class("dialog-header"), g.Group(children))
}

// DialogTitle creates a heading for dialog content
func DialogTitle(children ...g.Node) g.Node {
	return h.H2(h.Class("dialog-title"), g.Group(children))
}

// DialogDescription creates supporting text for dialogs
func DialogDescription(children ...g.Node) g.Node {
	return h.P(h.Class("dialog-description"), g.Group(children))
}

// DialogFooter creates the bottom section of a dialog
func DialogFooter(children ...g.Node) g.Node {
	return h.Div(h.Class("dialog-footer"), g.Group(children))
}

// Sheet (Drawer/Sidebar) components for slide-out panels
// Sheet creates a slide-out drawer component
func Sheet(open bool, children ...g.Node) g.Node {
	attrs := []g.Node{h.Class("sheet")}
	if open {
		attrs = append(attrs, g.Attr("open"))
	}
	attrs = append(attrs, g.Group(children))
	return h.Div(attrs...)
}

// SheetContent creates the main content area of a sheet
func SheetContent(children ...g.Node) g.Node {
	return h.Div(h.Class("sheet-content"), g.Group(children))
}

// SheetHeader creates the top section of a sheet
func SheetHeader(children ...g.Node) g.Node {
	return h.Div(h.Class("sheet-header"), g.Group(children))
}

// SheetTitle creates a heading for sheet content
func SheetTitle(children ...g.Node) g.Node {
	return h.H2(h.Class("sheet-title"), g.Group(children))
}

// SheetDescription creates supporting text for sheets
func SheetDescription(children ...g.Node) g.Node {
	return h.P(h.Class("sheet-description"), g.Group(children))
}

// SheetFooter creates the bottom section of a sheet
func SheetFooter(children ...g.Node) g.Node {
	return h.Div(h.Class("sheet-footer"), g.Group(children))
}
