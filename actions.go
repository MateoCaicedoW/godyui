package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// BadgeVariant defines different badge style variants
type BadgeVariant string

const (
	BadgeDefault     BadgeVariant = ""
	BadgeSecondary   BadgeVariant = "badge-secondary"
	BadgeDestructive BadgeVariant = "badge-destructive"
	BadgeOutline     BadgeVariant = "badge-outline"
	BadgeSuccess     BadgeVariant = "badge-success"
	BadgeWarning     BadgeVariant = "badge-warning"
)

// Badge creates a colored status label
func Badge(variant BadgeVariant, children ...g.Node) g.Node {
	classes := "badge"
	if variant != "" {
		classes += " " + string(variant)
	}
	return h.Span(h.Class(classes), g.Group(children))
}
