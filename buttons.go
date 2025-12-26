package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ButtonVariant defines different button style variants
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "primary"
	ButtonSecondary   ButtonVariant = "secondary"
	ButtonOutline     ButtonVariant = "outline"
	ButtonGhost       ButtonVariant = "ghost"
	ButtonLink        ButtonVariant = "link"
	ButtonDestructive ButtonVariant = "destructive"
)

// ButtonSize defines different button sizes
type ButtonSize string

const (
	ButtonDefault ButtonSize = ""
	ButtonSm      ButtonSize = "sm"
	ButtonLg      ButtonSize = "lg"
)

// ButtonEl creates a button element with specified variant, size, and icon flag
func ButtonEl(variant ButtonVariant, size ButtonSize, isIcon bool, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon)
	return h.Button(h.Class(class), g.Group(children))
}

// LinkButtonEl creates an anchor element with button styling for navigation
// Use this when you need a link that visually appears as a button
func LinkButtonEl(variant ButtonVariant, size ButtonSize, isIcon bool, href string, children ...g.Node) g.Node {
	class := buildButtonClass(variant, size, isIcon)
	return h.A(h.Class(class), h.Href(href), g.Group(children))
}

// buildButtonClass constructs the appropriate button class based on variant, size, and icon flag
// Examples: "btn", "btn-primary", "btn-lg-destructive", "btn-sm-icon-outline"
func buildButtonClass(variant ButtonVariant, size ButtonSize, isIcon bool) string {
	class := "btn"

	// Build combined class: btn-{size}-{icon}-{variant}
	if size != "" {
		class += "-" + string(size)
	}
	if isIcon {
		class += "-icon"
	}
	if variant != "" && variant != ButtonPrimary {
		class += "-" + string(variant)
	}

	return class
}
