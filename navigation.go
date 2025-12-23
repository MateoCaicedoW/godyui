package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Breadcrumb components for navigation trails
// Breadcrumb creates a navigation trail container
func Breadcrumb(children ...g.Node) g.Node {
	return h.Nav(h.Class("breadcrumb"), g.Group(children))
}

// BreadcrumbList creates a list of breadcrumb items
func BreadcrumbList(children ...g.Node) g.Node {
	return h.Ol(h.Class("breadcrumb-list"), g.Group(children))
}

// BreadcrumbItem creates a single breadcrumb element
func BreadcrumbItem(children ...g.Node) g.Node {
	return h.Li(h.Class("breadcrumb-item"), g.Group(children))
}

// BreadcrumbLink creates a clickable breadcrumb item
func BreadcrumbLink(href string, children ...g.Node) g.Node {
	return h.A(h.Class("breadcrumb-link"), h.Href(href), g.Group(children))
}

// BreadcrumbSeparator creates a visual separator between items
func BreadcrumbSeparator() g.Node {
	return h.Span(h.Class("breadcrumb-separator"), g.Text("/"))
}

// Pagination components for page navigation
// Pagination creates a page navigation container
func Pagination(children ...g.Node) g.Node {
	return h.Nav(h.Class("pagination"), g.Group(children))
}

// PaginationContent creates a list of page items
func PaginationContent(children ...g.Node) g.Node {
	return h.Ul(h.Class("pagination-content"), g.Group(children))
}

// PaginationItem creates a single pagination element
func PaginationItem(children ...g.Node) g.Node {
	return h.Li(h.Class("pagination-item"), g.Group(children))
}

// PaginationLink creates a clickable page link
func PaginationLink(href string, active bool, children ...g.Node) g.Node {
	classes := "pagination-link"
	if active {
		classes += " pagination-link-active"
	}
	return h.A(h.Class(classes), h.Href(href), g.Group(children))
}

// PaginationPrevious creates a previous page link
func PaginationPrevious(href string) g.Node {
	return h.A(h.Class("pagination-previous"), h.Href(href), g.Text("Previous"))
}

// PaginationNext creates a next page link
func PaginationNext(href string) g.Node {
	return h.A(h.Class("pagination-next"), h.Href(href), g.Text("Next"))
}

// Tabs components for content switching
// Tabs creates a tab container
func Tabs(children ...g.Node) g.Node {
	return h.Div(h.Class("tabs"), g.Group(children))
}

// TabsList creates a list of tab triggers
func TabsList(children ...g.Node) g.Node {
	return h.Div(h.Class("tabs-list"), g.Group(children))
}

// TabsTrigger creates a clickable tab button
func TabsTrigger(value string, children ...g.Node) g.Node {
	return h.Button(
		h.Class("tabs-trigger"),
		g.Attr("data-value", value),
		g.Group(children),
	)
}

// TabsContent creates the content area for a tab
func TabsContent(value string, children ...g.Node) g.Node {
	return h.Div(
		h.Class("tabs-content"),
		g.Attr("data-value", value),
		g.Group(children),
	)
}

// DropdownMenu creates a dropdown menu container
func DropdownMenu(children ...g.Node) g.Node {
	return h.Div(h.Class("dropdown"), g.Group(children))
}

// DropdownMenuTrigger creates a button to open dropdown
func DropdownMenuTrigger(children ...g.Node) g.Node {
	return h.Button(h.Class("dropdown-trigger"), g.Group(children))
}

// DropdownMenuContent creates the content area of dropdown
func DropdownMenuContent(children ...g.Node) g.Node {
	return h.Div(h.Class("dropdown-content"), g.Group(children))
}

// DropdownMenuItem creates a clickable menu item
func DropdownMenuItem(children ...g.Node) g.Node {
	return h.Div(h.Class("dropdown-item"), g.Group(children))
}

// DropdownMenuSeparator creates a visual divider between items
func DropdownMenuSeparator() g.Node {
	return h.Div(h.Class("dropdown-separator"))
}

// DropdownMenuLabel creates non-interactive label text
func DropdownMenuLabel(children ...g.Node) g.Node {
	return h.Div(h.Class("dropdown-label"), g.Group(children))
}

// ContextMenu creates a right-click menu container
func ContextMenu(children ...g.Node) g.Node {
	return h.Div(h.Class("context-menu"), g.Group(children))
}

// ContextMenuTrigger creates an element to trigger context menu
func ContextMenuTrigger(children ...g.Node) g.Node {
	return h.Div(h.Class("context-menu-trigger"), g.Group(children))
}

// ContextMenuContent creates the content area of context menu
func ContextMenuContent(children ...g.Node) g.Node {
	return h.Div(h.Class("context-menu-content"), g.Group(children))
}

// ContextMenuItem creates a clickable context menu item
func ContextMenuItem(children ...g.Node) g.Node {
	return h.Div(h.Class("context-menu-item"), g.Group(children))
}

// Menubar components for horizontal navigation bars
// Menubar creates a horizontal menu bar
func Menubar(children ...g.Node) g.Node {
	return h.Div(h.Class("menubar"), g.Group(children))
}

// MenubarMenu creates a menu within the menubar
func MenubarMenu(children ...g.Node) g.Node {
	return h.Div(h.Class("menubar-menu"), g.Group(children))
}

// MenubarTrigger creates a button to trigger menubar menu
func MenubarTrigger(children ...g.Node) g.Node {
	return h.Button(h.Class("menubar-trigger"), g.Group(children))
}

// MenubarContent creates the content area of menubar menu
func MenubarContent(children ...g.Node) g.Node {
	return h.Div(h.Class("menubar-content"), g.Group(children))
}

// MenubarItem creates a clickable menubar menu item
func MenubarItem(children ...g.Node) g.Node {
	return h.Div(h.Class("menubar-item"), g.Group(children))
}

// NavigationMenu creates a site navigation menu
func NavigationMenu(children ...g.Node) g.Node {
	return h.Nav(h.Class("navigation-menu"), g.Group(children))
}

// NavigationMenuList creates a list of navigation items
func NavigationMenuList(children ...g.Node) g.Node {
	return h.Ul(h.Class("navigation-menu-list"), g.Group(children))
}

// NavigationMenuItem creates a single navigation menu item
func NavigationMenuItem(children ...g.Node) g.Node {
	return h.Li(h.Class("navigation-menu-item"), g.Group(children))
}

// NavigationMenuTrigger creates a button for navigation submenu
func NavigationMenuTrigger(children ...g.Node) g.Node {
	return h.Button(h.Class("navigation-menu-trigger"), g.Group(children))
}

// NavigationMenuContent creates the content area of navigation menu
func NavigationMenuContent(children ...g.Node) g.Node {
	return h.Div(h.Class("navigation-menu-content"), g.Group(children))
}

// NavigationMenuLink creates a navigation link item
func NavigationMenuLink(href string, children ...g.Node) g.Node {
	return h.A(h.Class("navigation-menu-link"), h.Href(href), g.Group(children))
}
