package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ButtonVariant defines different button style variants
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "btn-primary"
	ButtonSecondary   ButtonVariant = "btn-secondary"
	ButtonOutline     ButtonVariant = "btn-outline"
	ButtonGhost       ButtonVariant = "btn-ghost"
	ButtonLink        ButtonVariant = "btn-link"
	ButtonDestructive ButtonVariant = "btn-destructive"
)

// ButtonSize defines different button sizes
type ButtonSize string

const (
	ButtonDefault ButtonSize = ""
	ButtonSm      ButtonSize = "btn-sm"
	ButtonLg      ButtonSize = "btn-lg"
	ButtonIcon    ButtonSize = "btn-icon"
)

// ButtonEl creates a button element with specified variant and size
func ButtonEl(variant ButtonVariant, size ButtonSize, children ...g.Node) g.Node {
	classes := "btn"
	if variant != "" {
		classes += " " + string(variant)
	}
	if size != "" {
		classes += " " + string(size)
	}
	return h.Button(h.Class(classes), g.Group(children))
}

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
