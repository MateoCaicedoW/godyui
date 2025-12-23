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
go get github.com/MateoCaicedoW/godyUI
```

Then include Basecoat UI CSS in your HTML:

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/basecoat.cdn.min.css">
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.9/dist/js/all.min.js" defer></script>
```

## Quick Start

```go
package main

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
    bc "github.com/MateoCaicedoW/godyUI"
)

func MyPage() g.Node {
    return Html(
        Head(
            TitleEl(g.Text("My App")),
            Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/basecoat-ui@latest/dist/basecoat.css")),
            bc.DarkModeScript(), // Enable dark mode
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                bc.ThemeToggle("theme-btn"),
                bc.Card(
                    bc.CardHeader(
                        bc.CardTitle(g.Text("Welcome!")),
                        bc.CardDescription(g.Text("Get started with Basecoat UI")),
                    ),
                    bc.CardContent(
                        bc.Button(bc.ButtonPrimary, bc.ButtonDefault,
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
    bc.DarkModeScript(),
    // ... other head elements
)
```

### Theme Toggle Button

Add a theme toggle button anywhere in your UI:

```go
bc.ThemeToggle("theme-toggle")
```

The theme preference is automatically saved to `localStorage` and persists across sessions.

## Component Reference

### Buttons

```go
// Primary button
bc.Button(bc.ButtonPrimary, bc.ButtonDefault, g.Text("Primary"))

// Secondary button
bc.Button(bc.ButtonSecondary, bc.ButtonDefault, g.Text("Secondary"))

// Outline button
bc.Button(bc.ButtonOutline, bc.ButtonDefault, g.Text("Outline"))

// Ghost button
bc.Button(bc.ButtonGhost, bc.ButtonDefault, g.Text("Ghost"))

// Destructive button
bc.Button(bc.ButtonDestructive, bc.ButtonDefault, g.Text("Delete"))

// Different sizes
bc.Button(bc.ButtonPrimary, bc.ButtonSm, g.Text("Small"))
bc.Button(bc.ButtonPrimary, bc.ButtonLg, g.Text("Large"))
bc.Button(bc.ButtonPrimary, bc.ButtonIcon, g.Text("🔍"))
```

### Forms

```go
bc.Form(
    bc.FormField(
        bc.FormLabel("email", g.Text("Email")),
        bc.Input(Type("email"), ID("email"), Placeholder("you@example.com")),
        bc.FormDescription(g.Text("We'll never share your email.")),
    ),
    bc.FormField(
        bc.FormLabel("message", g.Text("Message")),
        bc.Textarea(ID("message"), Placeholder("Your message...")),
    ),
    bc.Button(bc.ButtonPrimary, bc.ButtonDefault, g.Text("Submit")),
)
```

### Cards

```go
bc.Card(
    bc.CardHeader(
        bc.CardTitle(g.Text("Card Title")),
        bc.CardDescription(g.Text("Card description goes here")),
    ),
    bc.CardContent(
        g.Text("Your content here"),
    ),
    bc.CardFooter(
        bc.Button(bc.ButtonPrimary, bc.ButtonDefault, g.Text("Action")),
    ),
)
```

### Alerts

```go
// Default alert
bc.Alert(bc.AlertDefault,
    bc.AlertTitle(g.Text("Heads up!")),
    bc.AlertDescription(g.Text("This is an alert message.")),
)

// Success alert
bc.Alert(bc.AlertSuccess,
    bc.AlertTitle(g.Text("Success!")),
    bc.AlertDescription(g.Text("Operation completed.")),
)

// Destructive alert
bc.Alert(bc.AlertDestructive,
    bc.AlertTitle(g.Text("Error")),
    bc.AlertDescription(g.Text("Something went wrong.")),
)
```

### Badges

```go
bc.Badge(bc.BadgeDefault, g.Text("Default"))
bc.Badge(bc.BadgeSecondary, g.Text("Secondary"))
bc.Badge(bc.BadgeDestructive, g.Text("Destructive"))
bc.Badge(bc.BadgeOutline, g.Text("Outline"))
bc.Badge(bc.BadgeSuccess, g.Text("Success"))
bc.Badge(bc.BadgeWarning, g.Text("Warning"))
```

### Tables

```go
bc.Table(
    bc.TableHeader(
        bc.TableRow(
            bc.TableHead(g.Text("Name")),
            bc.TableHead(g.Text("Email")),
            bc.TableHead(g.Text("Role")),
        ),
    ),
    bc.TableBody(
        bc.TableRow(
            bc.TableCell(g.Text("John Doe")),
            bc.TableCell(g.Text("john@example.com")),
            bc.TableCell(g.Text("Admin")),
        ),
    ),
)
```

### Dialogs/Modals

```go
bc.Dialog(true, // open state
    bc.DialogContent(
        bc.DialogHeader(
            bc.DialogTitle(g.Text("Confirm Action")),
            bc.DialogDescription(g.Text("Are you sure?")),
        ),
        bc.DialogFooter(
            bc.Button(bc.ButtonOutline, bc.ButtonDefault, g.Text("Cancel")),
            bc.Button(bc.ButtonPrimary, bc.ButtonDefault, g.Text("Confirm")),
        ),
    ),
)
```

### Tabs

```go
bc.Tabs(
    bc.TabsList(
        bc.TabsTrigger("account", g.Text("Account")),
        bc.TabsTrigger("password", g.Text("Password")),
    ),
    bc.TabsContent("account",
        bc.Card(
            bc.CardHeader(
                bc.CardTitle(g.Text("Account Settings")),
            ),
            bc.CardContent(
                g.Text("Account content here"),
            ),
        ),
    ),
    bc.TabsContent("password",
        bc.Card(
            bc.CardHeader(
                bc.CardTitle(g.Text("Password Settings")),
            ),
            bc.CardContent(
                g.Text("Password content here"),
            ),
        ),
    ),
)
```

### Navigation Menu

```go
bc.NavigationMenu(
    bc.NavigationMenuList(
        bc.NavigationMenuItem(
            bc.NavigationMenuLink("/", g.Text("Home")),
        ),
        bc.NavigationMenuItem(
            bc.NavigationMenuLink("/about", g.Text("About")),
        ),
        bc.NavigationMenuItem(
            bc.NavigationMenuTrigger(g.Text("Products")),
            bc.NavigationMenuContent(
                bc.NavigationMenuLink("/products/web", g.Text("Web")),
                bc.NavigationMenuLink("/products/mobile", g.Text("Mobile")),
            ),
        ),
    ),
)
```

### Accordion

```go
bc.Accordion(
    bc.AccordionItem(
        bc.AccordionTrigger(g.Text("What is Basecoat UI?")),
        bc.AccordionContent(
            g.Text("Basecoat UI is a component library..."),
        ),
    ),
    bc.AccordionItem(
        bc.AccordionTrigger(g.Text("How do I use it?")),
        bc.AccordionContent(
            g.Text("Simply import the components..."),
        ),
    ),
)
```

### Dropdown Menu

```go
bc.DropdownMenu(
    bc.DropdownMenuTrigger(g.Text("Options")),
    bc.DropdownMenuContent(
        bc.DropdownMenuLabel(g.Text("My Account")),
        bc.DropdownMenuSeparator(),
        bc.DropdownMenuItem(g.Text("Profile")),
        bc.DropdownMenuItem(g.Text("Settings")),
        bc.DropdownMenuSeparator(),
        bc.DropdownMenuItem(g.Text("Logout")),
    ),
)
```

### Breadcrumbs

```go
bc.Breadcrumb(
    bc.BreadcrumbList(
        bc.BreadcrumbItem(
            bc.BreadcrumbLink("/", g.Text("Home")),
        ),
        bc.BreadcrumbSeparator(),
        bc.BreadcrumbItem(
            bc.BreadcrumbLink("/products", g.Text("Products")),
        ),
        bc.BreadcrumbSeparator(),
        bc.BreadcrumbItem(g.Text("Current Page")),
    ),
)
```

### Pagination

```go
bc.Pagination(
    bc.PaginationContent(
        bc.PaginationItem(bc.PaginationPrevious("/page/1")),
        bc.PaginationItem(bc.PaginationLink("/page/1", false, g.Text("1"))),
        bc.PaginationItem(bc.PaginationLink("/page/2", true, g.Text("2"))),
        bc.PaginationItem(bc.PaginationLink("/page/3", false, g.Text("3"))),
        bc.PaginationItem(bc.PaginationNext("/page/3")),
    ),
)
```

### Avatar

```go
bc.Avatar(
    bc.AvatarImage("/avatar.jpg", "User avatar"),
    bc.AvatarFallback(g.Text("JD")),
)
```

### Progress Bar

```go
bc.Progress(65) // 65% complete
```

### Skeleton Loaders

```go
bc.Card(
    bc.CardContent(
        bc.Skeleton(StyleEl(g.Attr("style", "height: 20px; margin-bottom: 8px"))),
        bc.Skeleton(StyleEl(g.Attr("style", "height: 20px; margin-bottom: 8px"))),
        bc.Skeleton(StyleEl(g.Attr("style", "height: 20px; width: 60%"))),
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

## Framework Integration

### Standard Library (net/http)

```go
package main

import (
    "net/http"
    g "maragu.dev/gomponents"
    bc "github.com/MateoCaicedoW/godyUI"
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
    bc "github.com/MateoCaicedoW/godyUI"
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
    bc "github.com/MateoCaicedoW/godyUI"
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
    bc "github.com/MateoCaicedoW/godyUI"
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
    return bc.Card(
        bc.CardHeader(
            bc.CardTitle(g.Text(name)),
            bc.CardDescription(g.Text(email)),
        ),
        bc.CardFooter(
            bc.Button(bc.ButtonPrimary, bc.ButtonDefault, g.Text("View Profile")),
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
            bc.DarkModeScript(),
        ),
        Body(
            Div(Class("container mx-auto p-4"),
                bc.ThemeToggle("theme"),
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
        bc.Alert(bc.AlertDestructive,
            bc.AlertTitle(g.Text("Error")),
            bc.AlertDescription(g.Text(message)),
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