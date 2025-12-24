package gomui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// DarkModeScript initializes theme detection and applies saved theme preference
func DarkModeScript() g.Node {
	return h.Script(g.Raw(`
		(function() {
			const theme = localStorage.getItem('theme') || 'light';
			document.documentElement.classList.toggle('dark', theme === 'dark');
		})();
	`))
}

// ThemeToggle creates a theme switcher button for toggling between light/dark modes
func ThemeToggle(id string, children ...g.Node) g.Node {
	return h.Button(
		h.ID(id),
		h.Type("button"),
		g.Attr("onclick", `
			const html = document.documentElement;
			const isDark = html.classList.toggle('dark');
			localStorage.setItem('theme', isDark ? 'dark' : 'light');
		`),

		g.Group(children),
		h.Class("btn-icon-outline size-8"),
	)
}

// BasecoatCSS includes Basecoat CSS from CDN
func BasecoatCSS() g.Node {
	return h.Link(
		h.Rel("stylesheet"),
		h.Href("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/basecoat.cdn.min.css"),
	)
}

// BasecoatJS includes Basecoat JavaScript from CDN
func BasecoatJS() g.Node {
	return h.Script(
		h.Src("https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/js/all.min.js"),
		g.Attr("defer"),
	)
}

// BasecoatAssets includes both Basecoat CSS and JavaScript for convenience
func BasecoatAssets() g.Node {
	return g.Group([]g.Node{
		BasecoatCSS(),
		BasecoatJS(),
	})
}
