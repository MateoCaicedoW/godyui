# Basecoat UI for Gomponents

A comprehensive Go wrapper library for [Basecoat UI](https://www.basecoat.dev/) components, built on top of [gomponents](https://maragu.dev/gomponents). Write beautiful, type-safe UI components in pure Go with no templating languages required.

## Features

✨ **Pure Go** - No templating languages, just Go functions  
🎨 **Complete Component Library** - 40+ pre-built Basecoat UI components  
🌓 **Dark Mode Support** - Built-in dark mode with localStorage persistence  
🔒 **Type-Safe** - Compile-time checking for your UI components  
⚡ **Zero Runtime** - Renders to pure HTML with minimal JavaScript  
🎯 **Framework Agnostic** - Works with any Go web framework  

## Installation

```bash
go get maragu.dev/gomponents
go get github.com/wawandco/gomui
```

Then include Basecoat UI CSS and JavaScript using our helper components in your HTML:

```html
<!-- Option 1: Include both CSS and JS (recommended) -->
<!-- In your Go code: gm.BasecoatAssets() -->

<!-- Option 2: Include CSS and JS separately -->
<!-- In your Go code: gm.BasecoatCSS() and gm.BasecoatJS() -->

<!-- Or manually include Basecoat assets -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/basecoat.cdn.min.css">
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/js/all.min.js" defer></script>
```

## Quick Start

```go
package main

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
    gm "github.com/wawandco/gomui"
)

func MyPage() g.Node {
    return Html(
        Head(
            TitleEl(g.Text("My App")),
            gm.BasecoatAssets(), // Include Basecoat CSS and JS
            gm.DarkModeScript(), // Enable dark mode
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                gm.ThemeToggle("theme-btn"),
                gm.Card(
                    gm.CardHeader(
                        gm.CardTitle(g.Text("Welcome!")),
                        gm.CardDescription(g.Text("Get started with Basecoat UI")),
                    ),
                    gm.CardContent(
                        gm.Button(gm.ButtonPrimary, gm.ButtonDefault,
                            g.Text("Click me"),
                        ),
                    ),
                ),
            ),
        ),
    )
}
```

## Dark Mode

### Setup

Add the dark mode script in your `<head>` to initialize the theme:

```go
Head(
    gm.DarkModeScript(),
    // ... other head elements
)
```

### Theme Toggle Button

Add a theme toggle button anywhere in your UI:

```go
gm.ThemeToggle("theme-toggle")
```

The theme preference is automatically saved to `localStorage` and persists across sessions.

## Component Reference

### Buttons

```go
// Primary button
gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Primary"))

// Secondary button
gm.Button(gm.ButtonSecondary, gm.ButtonDefault, g.Text("Secondary"))

// Outline button
gm.Button(gm.ButtonOutline, gm.ButtonDefault, g.Text("Outline"))

// Ghost button
gm.Button(gm.ButtonGhost, gm.ButtonDefault, g.Text("Ghost"))

// Destructive button
gm.Button(gm.ButtonDestructive, gm.ButtonDefault, g.Text("Delete"))

// Different sizes
gm.Button(gm.ButtonPrimary, gm.ButtonSm, g.Text("Small"))
gm.Button(gm.ButtonPrimary, gm.ButtonLg, g.Text("Large"))
gm.Button(gm.ButtonPrimary, gm.ButtonIcon, g.Text("🔍"))
```

### Forms

```go
gm.Form(
    gm.FormField(
        gm.FormLabel("email", g.Text("Email")),
        gm.Input(Type("email"), ID("email"), Placeholder("you@example.com")),
        gm.FormDescription(g.Text("We'll never share your email.")),
    ),
    gm.FormField(
        gm.FormLabel("message", g.Text("Message")),
        gm.Textarea(ID("message"), Placeholder("Your message...")),
    ),
    gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Submit")),
)
```

### Cards

```go
gm.Card(
    gm.CardHeader(
        gm.CardTitle(g.Text("Card Title")),
        gm.CardDescription(g.Text("Card description goes here")),
    ),
    gm.CardContent(
        g.Text("Your content here"),
    ),
    gm.CardFooter(
        gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Action")),
    ),
)
```

### Alerts

```go
// Default alert
gm.Alert(gm.AlertDefault,
    gm.AlertTitle(g.Text("Heads up!")),
    gm.AlertDescription(g.Text("This is an alert message.")),
)

// Success alert
gm.Alert(gm.AlertSuccess,
    gm.AlertTitle(g.Text("Success!")),
    gm.AlertDescription(g.Text("Operation completed.")),
)

// Destructive alert
gm.Alert(gm.AlertDestructive,
    gm.AlertTitle(g.Text("Error")),
    gm.AlertDescription(g.Text("Something went wrong.")),
)
```

### Badges

```go
gm.Badge(gm.BadgeDefault, g.Text("Default"))
gm.Badge(gm.BadgeSecondary, g.Text("Secondary"))
gm.Badge(gm.BadgeDestructive, g.Text("Destructive"))
gm.Badge(gm.BadgeOutline, g.Text("Outline"))
gm.Badge(gm.BadgeSuccess, g.Text("Success"))
gm.Badge(gm.BadgeWarning, g.Text("Warning"))
```

### Tables

```go
gm.Table(
    gm.TableHeader(
        gm.TableRow(
            gm.TableHead(g.Text("Name")),
            gm.TableHead(g.Text("Email")),
            gm.TableHead(g.Text("Role")),
        ),
    ),
    gm.TableBody(
        gm.TableRow(
            gm.TableCell(g.Text("John Doe")),
            gm.TableCell(g.Text("john@example.com")),
            gm.TableCell(g.Text("Admin")),
        ),
    ),
)
```

### Dialogs/Modals

```go
gm.Dialog(true, // open state
    gm.DialogContent(
        gm.DialogHeader(
            gm.DialogTitle(g.Text("Confirm Action")),
            gm.DialogDescription(g.Text("Are you sure?")),
        ),
        gm.DialogFooter(
            gm.Button(gm.ButtonOutline, gm.ButtonDefault, g.Text("Cancel")),
            gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("Confirm")),
        ),
    ),
)
```

### Tabs

```go
gm.Tabs(
    gm.TabsList(
        gm.TabsTrigger("account", g.Text("Account")),
        gm.TabsTrigger("password", g.Text("Password")),
    ),
    gm.TabsContent("account",
        gm.Card(
            gm.CardHeader(
                gm.CardTitle(g.Text("Account Settings")),
            ),
            gm.CardContent(
                g.Text("Account content here"),
            ),
        ),
    ),
    gm.TabsContent("password",
        gm.Card(
            gm.CardHeader(
                gm.CardTitle(g.Text("Password Settings")),
            ),
            gm.CardContent(
                g.Text("Password content here"),
            ),
        ),
    ),
)
```

### Navigation Menu

```go
gm.NavigationMenu(
    gm.NavigationMenuList(
        gm.NavigationMenuItem(
            gm.NavigationMenuLink("/", g.Text("Home")),
        ),
        gm.NavigationMenuItem(
            gm.NavigationMenuLink("/about", g.Text("About")),
        ),
        gm.NavigationMenuItem(
            gm.NavigationMenuTrigger(g.Text("Products")),
            gm.NavigationMenuContent(
                gm.NavigationMenuLink("/products/web", g.Text("Web")),
                gm.NavigationMenuLink("/products/mobile", g.Text("Mobile")),
            ),
        ),
    ),
)
```

### Accordion

```go
gm.Accordion(
    gm.AccordionItem(
        gm.AccordionTrigger(g.Text("What is Basecoat UI?")),
        gm.AccordionContent(
            g.Text("Basecoat UI is a component library..."),
        ),
    ),
    gm.AccordionItem(
        gm.AccordionTrigger(g.Text("How do I use it?")),
        gm.AccordionContent(
            g.Text("Simply import the components..."),
        ),
    ),
)
```

### Dropdown Menu

```go
gm.DropdownMenu(
    gm.DropdownMenuTrigger(g.Text("Options")),
    gm.DropdownMenuContent(
        gm.DropdownMenuLabel(g.Text("My Account")),
        gm.DropdownMenuSeparator(),
        gm.DropdownMenuItem(g.Text("Profile")),
        gm.DropdownMenuItem(g.Text("Settings")),
        gm.DropdownMenuSeparator(),
        gm.DropdownMenuItem(g.Text("Logout")),
    ),
)
```

### Breadcrumbs

```go
gm.Breadcrumb(
    gm.BreadcrumbList(
        gm.BreadcrumbItem(
            gm.BreadcrumbLink("/", g.Text("Home")),
        ),
        gm.BreadcrumbSeparator(),
        gm.BreadcrumbItem(
            gm.BreadcrumbLink("/products", g.Text("Products")),
        ),
        gm.BreadcrumbSeparator(),
        gm.BreadcrumbItem(g.Text("Current Page")),
    ),
)
```

### Pagination

```go
gm.Pagination(
    gm.PaginationContent(
        gm.PaginationItem(gm.PaginationPrevious("/page/1")),
        gm.PaginationItem(gm.PaginationLink("/page/1", false, g.Text("1"))),
        gm.PaginationItem(gm.PaginationLink("/page/2", true, g.Text("2"))),
        gm.PaginationItem(gm.PaginationLink("/page/3", false, g.Text("3"))),
        gm.PaginationItem(gm.PaginationNext("/page/3")),
    ),
)
```

### Avatar

```go
gm.Avatar(
    gm.AvatarImage("/avatar.jpg", "User avatar"),
    gm.AvatarFallback(g.Text("JD")),
)
```

### Progress Bar

```go
gm.Progress(65) // 65% complete
```

### Skeleton Loaders

```go
gm.Card(
    gm.CardContent(
        gm.Skeleton(StyleEl(g.Attr("style", "height: 20px; margin-bottom: 8px"))),
        gm.Skeleton(StyleEl(g.Attr("style", "height: 20px; margin-bottom: 8px"))),
        gm.Skeleton(StyleEl(g.Attr("style", "height: 20px; width: 60%"))),
    ),
)
```

## Complete Component List

### Layout & Structure
- `Card`, `CardHeader`, `CardTitle`, `CardDescription`, `CardContent`, `CardFooter`
- `Separator`
- `ScrollArea`
- `AspectRatio`

### Navigation
- `NavigationMenu`, `NavigationMenuList`, `NavigationMenuItem`, `NavigationMenuTrigger`, `NavigationMenuContent`, `NavigationMenuLink`
- `Breadcrumb`, `BreadcrumbList`, `BreadcrumbItem`, `BreadcrumbLink`, `BreadcrumbSeparator`
- `Pagination`, `PaginationContent`, `PaginationItem`, `PaginationLink`, `PaginationPrevious`, `PaginationNext`
- `Tabs`, `TabsList`, `TabsTrigger`, `TabsContent`
- `Menubar`, `MenubarMenu`, `MenubarTrigger`, `MenubarContent`, `MenubarItem`

### Forms & Inputs
- `Form`, `FormField`, `FormLabel`, `FormDescription`, `FormMessage`
- `Input`
- `Textarea`
- `Select`
- `Checkbox`
- `Radio`, `RadioGroup`
- `Switch`
- `Slider`
- `Label`

### Buttons & Actions
- `Button` (with variants: Primary, Secondary, Outline, Ghost, Link, Destructive)

### Data Display
- `Table`, `TableHeader`, `TableBody`, `TableFooter`, `TableRow`, `TableHead`, `TableCell`, `TableCaption`
- `Badge` (with variants: Default, Secondary, Destructive, Outline, Success, Warning)
- `Avatar`, `AvatarImage`, `AvatarFallback`
- `Progress`
- `Skeleton`

### Feedback
- `Alert`, `AlertTitle`, `AlertDescription` (with variants: Default, Destructive, Success, Warning, Info)
- `Toast`, `ToastTitle`, `ToastDescription`, `ToastAction`

### Overlays
- `Dialog`, `DialogContent`, `DialogHeader`, `DialogTitle`, `DialogDescription`, `DialogFooter`
- `Sheet`, `SheetContent`, `SheetHeader`, `SheetTitle`, `SheetDescription`, `SheetFooter`
- `Popover`, `PopoverTrigger`, `PopoverContent`
- `Tooltip`, `TooltipTrigger`, `TooltipContent`
- `HoverCard`, `HoverCardTrigger`, `HoverCardContent`

### Menus
- `DropdownMenu`, `DropdownMenuTrigger`, `DropdownMenuContent`, `DropdownMenuItem`, `DropdownMenuSeparator`, `DropdownMenuLabel`
- `ContextMenu`, `ContextMenuTrigger`, `ContextMenuContent`, `ContextMenuItem`

### Disclosure
- `Accordion`, `AccordionItem`, `AccordionTrigger`, `AccordionContent`
- `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent`

### Utility
- `Command`, `CommandInput`, `CommandList`, `CommandEmpty`, `CommandGroup`, `CommandItem`
- `Calendar`
- `Carousel`, `CarouselContent`, `CarouselItem`, `CarouselPrevious`, `CarouselNext`

### Theme
- `DarkModeScript()` - Initialize dark mode
- `ThemeToggle(id)` - Theme toggle button
- `BasecoatCSS()` - Include Basecoat CSS from CDN
- `BasecoatJS()` - Include Basecoat JavaScript from CDN  
- `BasecoatAssets()` - Include both Basecoat CSS and JavaScript for convenience

## Framework Integration

### Standard Library (net/http)

```go
package main

import (
    "net/http"
    g "maragu.dev/gomponents"
    gm "github.com/wawandco/gomui"
)

func handler(w http.ResponseWriter, r *http.Request) {
    page := MyPage()
    page.Render(w)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Chi Router

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    g "maragu.dev/gomponents"
    gm "github.com/wawandco/gomui"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        MyPage().Render(w)
    })
    http.ListenAndServe(":8080", r)
}
```

### Fiber

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    gm "github.com/wawandco/gomui"
)

func main() {
    app := fiber.New()
    
    app.Get("/", func(c *fiber.Ctx) error {
        c.Set("Content-Type", "text/html")
        return MyPage().Render(c)
    })
    
    app.Listen(":8080")
}
```

### Echo

```go
package main

import (
    "github.com/labstack/echo/v4"
    gm "github.com/wawandco/gomui"
)

func main() {
    e := echo.New()
    
    e.GET("/", func(c echo.Context) error {
        c.Response().Header().Set("Content-Type", "text/html")
        return MyPage().Render(c.Response())
    })
    
    e.Start(":8080")
}
```

## Best Practices

### 1. Organize Components

Create reusable component functions:

```go
func UserCard(name, email string) g.Node {
    return gm.Card(
        gm.CardHeader(
            gm.CardTitle(g.Text(name)),
            gm.CardDescription(g.Text(email)),
        ),
        gm.CardFooter(
            gm.Button(gm.ButtonPrimary, gm.ButtonDefault, g.Text("View Profile")),
        ),
    )
}
```

### 2. Create Layout Components

```go
func PageLayout(title string, content g.Node) g.Node {
    return Html(
        Head(
            TitleEl(g.Text(title)),
            Link(Rel("stylesheet"), Href("/static/basecoat.css")),
            gm.DarkModeScript(),
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                gm.ThemeToggle("theme"),
                content,
            ),
        ),
    )
}
```

### 3. Use Conditional Rendering

```go
import "maragu.dev/gomponents"

func ConditionalAlert(hasError bool, message string) g.Node {
    return gomponents.If(hasError,
        gm.Alert(gm.AlertDestructive,
            gm.AlertTitle(g.Text("Error")),
            gm.AlertDescription(g.Text(message)),
        ),
    )
}
```

### 4. Map Over Data

```go
func UserList(users []User) g.Node {
    return Div(
        gomponents.Map(users, func(u User) g.Node {
            return UserCard(u.Name, u.Email)
        }),
    )
}
```

## Examples

Check out the `/examples` directory for complete working examples:

- **Basic Page** - Simple page with cards and buttons
- **Form Example** - Complete form with validation
- **Dashboard** - Full dashboard with navigation and charts
- **E-commerce** - Product listing with filters

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Credits

- Built on top of [gomponents](https://maragu.dev/gomponents) by [@maragudk](https://github.com/maragudk)
- Uses [Basecoat UI](https://www.basecoat.dev/) components and styling
- Inspired by [shadcn/ui](https://ui.shadcn.com/)

## Resources

- [Gomponents Documentation](https://www.gomponents.com/)
- [Basecoat UI Documentation](https://www.basecoat.dev/)
- [Go Documentation](https://golang.org/doc/)

---

Made with ❤️ for the Go community