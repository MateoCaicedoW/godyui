package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Card components for creating content containers
// Card creates a main card container
func Card(children ...g.Node) g.Node {
	return h.Div(h.Class("card"), g.Group(children))
}

// CardHeader creates the top section of a card
func CardHeader(children ...g.Node) g.Node {
	return h.Div(h.Class("card-header"), g.Group(children))
}

// CardTitle creates a heading for card content
func CardTitle(children ...g.Node) g.Node {
	return h.H3(h.Class("card-title"), g.Group(children))
}

// CardDescription creates supporting text for card content
func CardDescription(children ...g.Node) g.Node {
	return h.P(h.Class("card-description"), g.Group(children))
}

// CardContent creates the main content area of a card
func CardContent(children ...g.Node) g.Node {
	return h.Div(h.Class("card-content"), g.Group(children))
}

// CardFooter creates the bottom section of a card
func CardFooter(children ...g.Node) g.Node {
	return h.Div(h.Class("card-footer"), g.Group(children))
}

// Separator component for visual dividers
// Separator creates a visual divider line
func Separator() g.Node {
	return h.Div(h.Class("separator"))
}

// Avatar components for user images
// Avatar creates a user avatar container
func Avatar(children ...g.Node) g.Node {
	return h.Div(h.Class("avatar"), g.Group(children))
}

// AvatarImage creates an avatar image element
func AvatarImage(src, alt string) g.Node {
	return h.Img(
		h.Class("avatar-image"),
		h.Src(src),
		h.Alt(alt),
	)
}

// AvatarFallback creates fallback text for avatar
func AvatarFallback(children ...g.Node) g.Node {
	return h.Div(h.Class("avatar-fallback"), g.Group(children))
}

// ScrollArea creates a scrollable content container
func ScrollArea(children ...g.Node) g.Node {
	return h.Div(h.Class("scroll-area"), g.Group(children))
}

// AspectRatio creates a container with fixed aspect ratio
func AspectRatio(ratio string, children ...g.Node) g.Node {
	return h.Div(
		h.Class("aspect-ratio"),
		h.StyleEl(g.Attr("style", "aspect-ratio: "+ratio)),
		g.Group(children),
	)
}

// Calendar component for date selection
// Calendar creates a calendar display container
func Calendar(children ...g.Node) g.Node {
	return h.Div(h.Class("calendar"), g.Group(children))
}

// Carousel components for image/content sliders
// Carousel creates a sliding content container
func Carousel(children ...g.Node) g.Node {
	return h.Div(h.Class("carousel"), g.Group(children))
}

// CarouselContent creates the main carousel content area
func CarouselContent(children ...g.Node) g.Node {
	return h.Div(h.Class("carousel-content"), g.Group(children))
}

// CarouselItem creates a single slide in carousel
func CarouselItem(children ...g.Node) g.Node {
	return h.Div(h.Class("carousel-item"), g.Group(children))
}

// CarouselPrevious creates a button to go to previous slide
func CarouselPrevious() g.Node {
	return h.Button(h.Class("carousel-previous"), g.Text("‹"))
}

// CarouselNext creates a button to go to next slide
func CarouselNext() g.Node {
	return h.Button(h.Class("carousel-next"), g.Text("›"))
}
