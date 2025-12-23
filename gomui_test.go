package gomui_test

import (
	"testing"

	gm "github.com/wawandco/gomui"
	g "maragu.dev/gomponents"
)

func Test(t *testing.T) {
	// Test that components are accessible from main package
	page := gm.Card(
		gm.CardHeader(
			gm.CardTitle(g.Text("Test")),
			gm.CardDescription(g.Text("Testing components")),
		),

		gm.CardContent(
			gm.ButtonEl(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Click me")),
			gm.Badge(gm.BadgeSuccess, g.Text("Success")),
		),
	)

	_ = page // Just verify compilation
}
