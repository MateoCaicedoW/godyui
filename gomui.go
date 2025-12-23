// Package gomui provides a comprehensive Go UI components library built on top of gomponents.
// It offers 40+ pre-built Basecoat UI components for creating beautiful, type-safe
// user interfaces in pure Go without any templating languages.
//
// Key features:
//   - Pure Go components with no runtime overhead
//   - Type-safe component API with compile-time checking
//   - Built-in dark mode support with localStorage persistence
//   - Complete component library covering common UI patterns
//   - Framework agnostic - works with any Go web framework
//
// Usage:
//
//	import g "maragu.dev/gomponents"
//	h "maragu.dev/gomponents/html"
//	gm "github.com/wawandco/gomui"
//
//	page := gm.Card(
//		gm.CardHeader(gm.CardTitle(g.Text("Hello World"))),
//		gm.CardContent(g.Text("Welcome to GomUI!")),
//	)
//
// Component Organization:
//   - actions/: Button, Badge and other action components
//   - forms/: Form, Input, Textarea and other form components
//   - data/: Table, Disclosure and other data display components
//   - overlays/: Dialog, Sheet, Overlay and other overlay components
//   - feedback/: Alert, Feedback and other feedback components
//   - theme/: Theme-related utilities and components
package gomui
