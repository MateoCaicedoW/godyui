package gomui_test

import (
	"testing"

	gm "github.com/wawandco/gomui"
	g "maragu.dev/gomponents"
	ghtml "maragu.dev/gomponents/html"
)

func Test(t *testing.T) {
	// Test that components are accessible from main package
	page := gm.Card(
		gm.CardHeader(
			ghtml.H1(g.Text("Card Title")),
		),

		gm.CardContent(
			gm.ButtonEl(gm.ButtonPrimary, gm.ButtonDefault, false, g.Text("Click me")),
			gm.Badge(gm.BadgeSuccess, g.Text("Success")),
		),
	)

	_ = page // Just verify compilation
}
