package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// InputEl creates a styled text input field
func InputEl(attrs ...g.Node) g.Node {
	return h.Input(h.Class("input"), g.Group(attrs))
}

// TextareaEl creates a multi-line text input field
func TextareaEl(attrs ...g.Node) g.Node {
	return h.Textarea(h.Class("textarea"), g.Group(attrs))
}

// FormComponent creates a form wrapper element
func FormComponent(attrs ...g.Node) g.Node {
	return h.Form(h.Class("form"), g.Group(attrs))
}

// FormField creates a container for form inputs and labels
func FormField(children ...g.Node) g.Node {
	return h.Div(h.Class("form-field"), g.Group(children))
}

// FormLabel creates a label element for form inputs
func FormLabel(forID string, children ...g.Node) g.Node {
	return h.Label(
		h.Class("form-label"),
		h.For(forID),
		g.Group(children),
	)
}

// FormDescription creates a descriptive text for form fields
func FormDescription(children ...g.Node) g.Node {
	return h.P(h.Class("form-description"), g.Group(children))
}

// FormMessage creates validation or error message text
func FormMessage(children ...g.Node) g.Node {
	return h.P(h.Class("form-message"), g.Group(children))
}

// Checkbox component for toggle selection
// Checkbox creates a checkbox input element
func Checkbox(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Type("checkbox"),
		h.Class("checkbox"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Input(allAttrs...)
}

// Radio component for single selection
// Radio creates a radio input element
func Radio(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		h.Type("radio"),
		h.Class("radio"),
	}
	allAttrs = append(allAttrs, attrs...)
	return h.Input(allAttrs...)
}

// RadioGroup component for grouping radio buttons
// RadioGroup creates a container for radio inputs
func RadioGroup(children ...g.Node) g.Node {
	return h.Div(h.Class("radio-group"), g.Group(children))
}

// Switch component for binary toggles
// Switch creates a toggle switch input
func Switch(checked bool, attrs ...g.Node) g.Node {
	return h.Input(
		g.If(checked, h.Checked()),
		g.Group(attrs),
		h.Type("checkbox"),
		h.Class("input"),
		h.Role("switch"),
	)
}

// SwitchContainer creates a container for switch and label
func SwitchLabel(forID string, children ...g.Node) g.Node {
	return h.Label(
		h.Class("label"),
		h.For(forID),
		g.Group(children),
	)
}

// InputLabel creates a label for form inputs
func InputLabel(forID string, children ...g.Node) g.Node {
	return h.Label(
		h.Class("label"),
		h.For(forID),
		g.Group(children),
	)
}
