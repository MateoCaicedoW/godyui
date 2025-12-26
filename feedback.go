package gomui

import (
	"fmt"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// AlertVariant defines different alert style variants
type AlertVariant string

const (
	AlertDefault     AlertVariant = ""
	AlertDestructive AlertVariant = "alert-destructive"
	AlertSuccess     AlertVariant = "alert-success"
	AlertWarning     AlertVariant = "alert-warning"
	AlertInfo        AlertVariant = "alert-info"
)

// Alert creates a message container with variant styling
func Alert(variant AlertVariant, children ...g.Node) g.Node {
	classes := "alert"
	if variant != "" {
		classes += " " + string(variant)
	}
	return h.Div(h.Class(classes), g.Group(children))
}

// AlertTitle creates a heading for alert content
func AlertTitle(children ...g.Node) g.Node {
	return h.H5(h.Class("alert-title"), g.Group(children))
}

// AlertDescription creates descriptive text for alerts
func AlertDescription(children ...g.Node) g.Node {
	return h.Div(h.Class("alert-description"), g.Group(children))
}

// ProgressEl creates a progress bar element
func ProgressEl(value int, attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Class("progress"),
		g.Attr("value", fmt.Sprintf("%d", value)),
		h.Max("100"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Progress(allAttrs...)
}

// Tooltip components for contextual help
// Tooltip creates a tooltip container
func Tooltip(children ...g.Node) g.Node {
	return h.Div(h.Class("tooltip"), g.Group(children))
}

// TooltipTrigger creates an element to show tooltip
func TooltipTrigger(children ...g.Node) g.Node {
	return h.Div(h.Class("tooltip-trigger"), g.Group(children))
}

// TooltipContent creates the tooltip text content
func TooltipContent(children ...g.Node) g.Node {
	return h.Div(h.Class("tooltip-content"), g.Group(children))
}

// Skeleton component for loading placeholders
// Skeleton creates a loading placeholder element
func Skeleton(attrs ...g.Node) g.Node {
	return h.Div(h.Class("skeleton"), g.Group(attrs))
}

// Toast creates a notification toast container
func Toast(children ...g.Node) g.Node {
	return h.Div(h.Class("toast"), g.Group(children))
}

// ToastTitle creates a heading for toast content
func ToastTitle(children ...g.Node) g.Node {
	return h.Div(h.Class("toast-title"), g.Group(children))
}

// ToastDescription creates descriptive text for toasts
func ToastDescription(children ...g.Node) g.Node {
	return h.Div(h.Class("toast-description"), g.Group(children))
}

// ToastAction creates a button for toast interaction
func ToastAction(children ...g.Node) g.Node {
	return h.Button(h.Class("toast-action"), g.Group(children))
}
