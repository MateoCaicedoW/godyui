package gody

import (
	"fmt"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Theme/Dark Mode utilities
func DarkModeScript() g.Node {
	return Script(g.Raw(`
		(function() {
			const theme = localStorage.getItem('theme') || 'light';
			document.documentElement.classList.toggle('dark', theme === 'dark');
		})();
	`))
}

func ThemeToggle(id string) g.Node {
	return Button(
		ID(id),
		Class("btn btn-ghost"),
		g.Attr("onclick", `
			const html = document.documentElement;
			const isDark = html.classList.toggle('dark');
			localStorage.setItem('theme', isDark ? 'dark' : 'light');
		`),
		g.Text("🌓"),
	)
}

// Button creates a Basecoat UI button component
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "btn-primary"
	ButtonSecondary   ButtonVariant = "btn-secondary"
	ButtonOutline     ButtonVariant = "btn-outline"
	ButtonGhost       ButtonVariant = "btn-ghost"
	ButtonLink        ButtonVariant = "btn-link"
	ButtonDestructive ButtonVariant = "btn-destructive"
)

type ButtonSize string

const (
	ButtonDefault ButtonSize = ""
	ButtonSm      ButtonSize = "btn-sm"
	ButtonLg      ButtonSize = "btn-lg"
	ButtonIcon    ButtonSize = "btn-icon"
)

func ButtonEl(variant ButtonVariant, size ButtonSize, children ...g.Node) g.Node {
	classes := "btn"
	if variant != "" {
		classes += " " + string(variant)
	}
	if size != "" {
		classes += " " + string(size)
	}
	return Button(Class(classes), g.Group(children))
}

// Input creates a Basecoat UI input component
func InputEl(attrs ...g.Node) g.Node {
	return Input(Class("input"), g.Group(attrs))
}

// Textarea creates a Basecoat UI textarea component
func TextareaEl(attrs ...g.Node) g.Node {
	return Textarea(Class("textarea"), g.Group(attrs))
}

// Select creates a Basecoat UI select component
func SelectEl(children ...g.Node) g.Node {
	return Select(Class("select"), g.Group(children))
}

// Form components
func FormComponent(attrs ...g.Node) g.Node {
	return Form(Class("form"), g.Group(attrs))
}

func FormField(children ...g.Node) g.Node {
	return Div(Class("form-field"), g.Group(children))
}

func FormLabel(forID string, children ...g.Node) g.Node {
	return LabelEl(
		Class("form-label"),
		For(forID),
		g.Group(children),
	)
}

func FormDescription(children ...g.Node) g.Node {
	return P(Class("form-description"), g.Group(children))
}

func FormMessage(children ...g.Node) g.Node {
	return P(Class("form-message"), g.Group(children))
}

// Card components
func Card(children ...g.Node) g.Node {
	return Div(Class("card"), g.Group(children))
}

func CardHeader(children ...g.Node) g.Node {
	return Div(Class("card-header"), g.Group(children))
}

func CardTitle(children ...g.Node) g.Node {
	return H3(Class("card-title"), g.Group(children))
}

func CardDescription(children ...g.Node) g.Node {
	return P(Class("card-description"), g.Group(children))
}

func CardContent(children ...g.Node) g.Node {
	return Div(Class("card-content"), g.Group(children))
}

func CardFooter(children ...g.Node) g.Node {
	return Div(Class("card-footer"), g.Group(children))
}

// Badge component
type BadgeVariant string

const (
	BadgeDefault     BadgeVariant = ""
	BadgeSecondary   BadgeVariant = "badge-secondary"
	BadgeDestructive BadgeVariant = "badge-destructive"
	BadgeOutline     BadgeVariant = "badge-outline"
	BadgeSuccess     BadgeVariant = "badge-success"
	BadgeWarning     BadgeVariant = "badge-warning"
)

func Badge(variant BadgeVariant, children ...g.Node) g.Node {
	classes := "badge"
	if variant != "" {
		classes += " " + string(variant)
	}
	return Span(Class(classes), g.Group(children))
}

// Alert components
type AlertVariant string

const (
	AlertDefault     AlertVariant = ""
	AlertDestructive AlertVariant = "alert-destructive"
	AlertSuccess     AlertVariant = "alert-success"
	AlertWarning     AlertVariant = "alert-warning"
	AlertInfo        AlertVariant = "alert-info"
)

func Alert(variant AlertVariant, children ...g.Node) g.Node {
	classes := "alert"
	if variant != "" {
		classes += " " + string(variant)
	}
	return Div(Class(classes), g.Group(children))
}

func AlertTitle(children ...g.Node) g.Node {
	return H5(Class("alert-title"), g.Group(children))
}

func AlertDescription(children ...g.Node) g.Node {
	return Div(Class("alert-description"), g.Group(children))
}

// Dialog/Modal components
func DialogEl(open bool, children ...g.Node) g.Node {
	attrs := []g.Node{Class("dialog")}
	if open {
		attrs = append(attrs, g.Attr("open"))
	}
	attrs = append(attrs, g.Group(children))
	return Dialog(attrs...)
}

func DialogContent(children ...g.Node) g.Node {
	return Div(Class("dialog-content"), g.Group(children))
}

func DialogHeader(children ...g.Node) g.Node {
	return Div(Class("dialog-header"), g.Group(children))
}

func DialogTitle(children ...g.Node) g.Node {
	return H2(Class("dialog-title"), g.Group(children))
}

func DialogDescription(children ...g.Node) g.Node {
	return P(Class("dialog-description"), g.Group(children))
}

func DialogFooter(children ...g.Node) g.Node {
	return Div(Class("dialog-footer"), g.Group(children))
}

// Sheet (Drawer/Sidebar) components
func Sheet(open bool, children ...g.Node) g.Node {
	attrs := []g.Node{Class("sheet")}
	if open {
		attrs = append(attrs, g.Attr("open"))
	}
	attrs = append(attrs, g.Group(children))
	return Div(attrs...)
}

func SheetContent(children ...g.Node) g.Node {
	return Div(Class("sheet-content"), g.Group(children))
}

func SheetHeader(children ...g.Node) g.Node {
	return Div(Class("sheet-header"), g.Group(children))
}

func SheetTitle(children ...g.Node) g.Node {
	return H2(Class("sheet-title"), g.Group(children))
}

func SheetDescription(children ...g.Node) g.Node {
	return P(Class("sheet-description"), g.Group(children))
}

func SheetFooter(children ...g.Node) g.Node {
	return Div(Class("sheet-footer"), g.Group(children))
}

// Popover components
func PopoverEl(children ...g.Node) g.Node {
	return Div(Class("popover"), g.Group(children))
}

func PopoverTrigger(children ...g.Node) g.Node {
	return Button(Class("popover-trigger"), g.Group(children))
}

func PopoverContent(children ...g.Node) g.Node {
	return Div(Class("popover-content"), g.Group(children))
}

// Dropdown Menu components
func DropdownMenu(children ...g.Node) g.Node {
	return Div(Class("dropdown"), g.Group(children))
}

func DropdownMenuTrigger(children ...g.Node) g.Node {
	return Button(Class("dropdown-trigger"), g.Group(children))
}

func DropdownMenuContent(children ...g.Node) g.Node {
	return Div(Class("dropdown-content"), g.Group(children))
}

func DropdownMenuItem(children ...g.Node) g.Node {
	return Div(Class("dropdown-item"), g.Group(children))
}

func DropdownMenuSeparator() g.Node {
	return Div(Class("dropdown-separator"))
}

func DropdownMenuLabel(children ...g.Node) g.Node {
	return Div(Class("dropdown-label"), g.Group(children))
}

// Accordion components
func Accordion(children ...g.Node) g.Node {
	return Div(Class("accordion"), g.Group(children))
}

func AccordionItem(children ...g.Node) g.Node {
	return Div(Class("accordion-item"), g.Group(children))
}

func AccordionTrigger(children ...g.Node) g.Node {
	return Button(Class("accordion-trigger"), g.Group(children))
}

func AccordionContent(children ...g.Node) g.Node {
	return Div(Class("accordion-content"), g.Group(children))
}

// Breadcrumb components
func Breadcrumb(children ...g.Node) g.Node {
	return Nav(Class("breadcrumb"), g.Group(children))
}

func BreadcrumbList(children ...g.Node) g.Node {
	return Ol(Class("breadcrumb-list"), g.Group(children))
}

func BreadcrumbItem(children ...g.Node) g.Node {
	return Li(Class("breadcrumb-item"), g.Group(children))
}

func BreadcrumbLink(href string, children ...g.Node) g.Node {
	return A(Class("breadcrumb-link"), Href(href), g.Group(children))
}

func BreadcrumbSeparator() g.Node {
	return Span(Class("breadcrumb-separator"), g.Text("/"))
}

// Pagination components
func Pagination(children ...g.Node) g.Node {
	return Nav(Class("pagination"), g.Group(children))
}

func PaginationContent(children ...g.Node) g.Node {
	return Ul(Class("pagination-content"), g.Group(children))
}

func PaginationItem(children ...g.Node) g.Node {
	return Li(Class("pagination-item"), g.Group(children))
}

func PaginationLink(href string, active bool, children ...g.Node) g.Node {
	classes := "pagination-link"
	if active {
		classes += " pagination-link-active"
	}
	return A(Class(classes), Href(href), g.Group(children))
}

func PaginationPrevious(href string) g.Node {
	return A(Class("pagination-previous"), Href(href), g.Text("Previous"))
}

func PaginationNext(href string) g.Node {
	return A(Class("pagination-next"), Href(href), g.Text("Next"))
}

// Table components
func TableEl(children ...g.Node) g.Node {
	return Table(Class("table"), g.Group(children))
}

func TableHeader(children ...g.Node) g.Node {
	return THead(Class("table-header"), g.Group(children))
}

func TableBody(children ...g.Node) g.Node {
	return TBody(Class("table-body"), g.Group(children))
}

func TableFooter(children ...g.Node) g.Node {
	return TFoot(Class("table-footer"), g.Group(children))
}

func TableRow(children ...g.Node) g.Node {
	return Tr(Class("table-row"), g.Group(children))
}

func TableHead(children ...g.Node) g.Node {
	return Th(Class("table-head"), g.Group(children))
}

func TableCell(children ...g.Node) g.Node {
	return Td(Class("table-cell"), g.Group(children))
}

func TableCaption(children ...g.Node) g.Node {
	return Caption(Class("table-caption"), g.Group(children))
}

// Checkbox component
func Checkbox(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		Type("checkbox"),
		Class("checkbox"),
	}
	allAttrs = append(allAttrs, attrs...)
	return Input(allAttrs...)
}

// Radio component
func Radio(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		Type("radio"),
		Class("radio"),
	}
	allAttrs = append(allAttrs, attrs...)
	return Input(allAttrs...)
}

// RadioGroup component
func RadioGroup(children ...g.Node) g.Node {
	return Div(Class("radio-group"), g.Group(children))
}

// Switch component
func Switch(checked bool, attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		Type("checkbox"),
		Class("switch"),
	}
	if checked {
		allAttrs = append(allAttrs, g.Attr("checked"))
	}
	allAttrs = append(allAttrs, attrs...)
	return Input(allAttrs...)
}

// Label component
func InputLabel(forID string, children ...g.Node) g.Node {
	return Label(
		Class("label"),
		For(forID),
		g.Group(children),
	)
}

// Tabs components
func Tabs(children ...g.Node) g.Node {
	return Div(Class("tabs"), g.Group(children))
}

func TabsList(children ...g.Node) g.Node {
	return Div(Class("tabs-list"), g.Group(children))
}

func TabsTrigger(value string, children ...g.Node) g.Node {
	return Button(
		Class("tabs-trigger"),
		g.Attr("data-value", value),
		g.Group(children),
	)
}

func TabsContent(value string, children ...g.Node) g.Node {
	return Div(
		Class("tabs-content"),
		g.Attr("data-value", value),
		g.Group(children),
	)
}

// Progress component
func ProgressEl(value int, attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		Class("progress"),
		g.Attr("value", fmt.Sprintf("%d", value)),
		Max("100"),
	}
	allAttrs = append(allAttrs, attrs...)
	return Progress(allAttrs...)
}

// Slider component
func Slider(attrs ...g.Node) g.Node {
	allAttrs := []g.Node{
		Type("range"),
		Class("slider"),
	}
	allAttrs = append(allAttrs, attrs...)
	return Input(allAttrs...)
}

// Separator component
func Separator() g.Node {
	return Div(Class("separator"))
}

// Avatar components
func Avatar(children ...g.Node) g.Node {
	return Div(Class("avatar"), g.Group(children))
}

func AvatarImage(src, alt string) g.Node {
	return Img(
		Class("avatar-image"),
		Src(src),
		Alt(alt),
	)
}

func AvatarFallback(children ...g.Node) g.Node {
	return Div(Class("avatar-fallback"), g.Group(children))
}

// Tooltip components
func Tooltip(children ...g.Node) g.Node {
	return Div(Class("tooltip"), g.Group(children))
}

func TooltipTrigger(children ...g.Node) g.Node {
	return Div(Class("tooltip-trigger"), g.Group(children))
}

func TooltipContent(children ...g.Node) g.Node {
	return Div(Class("tooltip-content"), g.Group(children))
}

// Skeleton component for loading states
func Skeleton(attrs ...g.Node) g.Node {
	return Div(Class("skeleton"), g.Group(attrs))
}

// Toast/Notification components
func Toast(children ...g.Node) g.Node {
	return Div(Class("toast"), g.Group(children))
}

func ToastTitle(children ...g.Node) g.Node {
	return Div(Class("toast-title"), g.Group(children))
}

func ToastDescription(children ...g.Node) g.Node {
	return Div(Class("toast-description"), g.Group(children))
}

func ToastAction(children ...g.Node) g.Node {
	return Button(Class("toast-action"), g.Group(children))
}

// Command/Command Palette components
func Command(children ...g.Node) g.Node {
	return Div(Class("command"), g.Group(children))
}

func CommandInput(attrs ...g.Node) g.Node {
	return Input(Class("command-input"), g.Group(attrs))
}

func CommandList(children ...g.Node) g.Node {
	return Div(Class("command-list"), g.Group(children))
}

func CommandEmpty(children ...g.Node) g.Node {
	return Div(Class("command-empty"), g.Group(children))
}

func CommandGroup(children ...g.Node) g.Node {
	return Div(Class("command-group"), g.Group(children))
}

func CommandItem(children ...g.Node) g.Node {
	return Div(Class("command-item"), g.Group(children))
}

// Context Menu components
func ContextMenu(children ...g.Node) g.Node {
	return Div(Class("context-menu"), g.Group(children))
}

func ContextMenuTrigger(children ...g.Node) g.Node {
	return Div(Class("context-menu-trigger"), g.Group(children))
}

func ContextMenuContent(children ...g.Node) g.Node {
	return Div(Class("context-menu-content"), g.Group(children))
}

func ContextMenuItem(children ...g.Node) g.Node {
	return Div(Class("context-menu-item"), g.Group(children))
}

// Hover Card components
func HoverCard(children ...g.Node) g.Node {
	return Div(Class("hover-card"), g.Group(children))
}

func HoverCardTrigger(children ...g.Node) g.Node {
	return Div(Class("hover-card-trigger"), g.Group(children))
}

func HoverCardContent(children ...g.Node) g.Node {
	return Div(Class("hover-card-content"), g.Group(children))
}

// Menubar components
func Menubar(children ...g.Node) g.Node {
	return Div(Class("menubar"), g.Group(children))
}

func MenubarMenu(children ...g.Node) g.Node {
	return Div(Class("menubar-menu"), g.Group(children))
}

func MenubarTrigger(children ...g.Node) g.Node {
	return Button(Class("menubar-trigger"), g.Group(children))
}

func MenubarContent(children ...g.Node) g.Node {
	return Div(Class("menubar-content"), g.Group(children))
}

func MenubarItem(children ...g.Node) g.Node {
	return Div(Class("menubar-item"), g.Group(children))
}

// Navigation Menu components
func NavigationMenu(children ...g.Node) g.Node {
	return Nav(Class("navigation-menu"), g.Group(children))
}

func NavigationMenuList(children ...g.Node) g.Node {
	return Ul(Class("navigation-menu-list"), g.Group(children))
}

func NavigationMenuItem(children ...g.Node) g.Node {
	return Li(Class("navigation-menu-item"), g.Group(children))
}

func NavigationMenuTrigger(children ...g.Node) g.Node {
	return Button(Class("navigation-menu-trigger"), g.Group(children))
}

func NavigationMenuContent(children ...g.Node) g.Node {
	return Div(Class("navigation-menu-content"), g.Group(children))
}

func NavigationMenuLink(href string, children ...g.Node) g.Node {
	return A(Class("navigation-menu-link"), Href(href), g.Group(children))
}

// Scroll Area component
func ScrollArea(children ...g.Node) g.Node {
	return Div(Class("scroll-area"), g.Group(children))
}

// Collapsible components
func Collapsible(children ...g.Node) g.Node {
	return Div(Class("collapsible"), g.Group(children))
}

func CollapsibleTrigger(children ...g.Node) g.Node {
	return Button(Class("collapsible-trigger"), g.Group(children))
}

func CollapsibleContent(children ...g.Node) g.Node {
	return Div(Class("collapsible-content"), g.Group(children))
}

// Aspect Ratio component
func AspectRatio(ratio string, children ...g.Node) g.Node {
	return Div(
		Class("aspect-ratio"),
		StyleEl(g.Attr("style", "aspect-ratio: "+ratio)),
		g.Group(children),
	)
}

// Calendar component (basic structure)
func Calendar(children ...g.Node) g.Node {
	return Div(Class("calendar"), g.Group(children))
}

// Carousel components
func Carousel(children ...g.Node) g.Node {
	return Div(Class("carousel"), g.Group(children))
}

func CarouselContent(children ...g.Node) g.Node {
	return Div(Class("carousel-content"), g.Group(children))
}

func CarouselItem(children ...g.Node) g.Node {
	return Div(Class("carousel-item"), g.Group(children))
}

func CarouselPrevious() g.Node {
	return Button(Class("carousel-previous"), g.Text("‹"))
}

func CarouselNext() g.Node {
	return Button(Class("carousel-next"), g.Text("›"))
}
