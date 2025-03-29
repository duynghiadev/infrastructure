# Website Design Pattern

1. [https://refactoring.guru/design-patterns](https://refactoring.guru/design-patterns)
2. [https://blog.ntechdevelopers.com/gof-design-patterns-no-la-gi-ma-tai-sao-cac-senior-bat-buoc-phai-biet/](https://blog.ntechdevelopers.com/gof-design-patterns-no-la-gi-ma-tai-sao-cac-senior-bat-buoc-phai-biet/)
3. [https://gpcoder.com/4164-gioi-thieu-design-patterns/](https://gpcoder.com/4164-gioi-thieu-design-patterns/)
4. [https://en.wikipedia.org/wiki/Software_design_pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
5. [https://viblo.asia/s/tong-hop-23-mau-design-patterns-tro-thu-dac-luc-cua-developers-Q75wqJ67ZWb](https://viblo.asia/s/tong-hop-23-mau-design-patterns-tro-thu-dac-luc-cua-developers-Q75wqJ67ZWb)

# Design Patterns Guide: From Newbie to Advanced

## Introduction to Design Patterns

Design patterns are typical solutions to common problems in software design. They represent best practices evolved over time by experienced software developers. Design patterns were popularized by the "Gang of Four" (GoF) in their book "Design Patterns: Elements of Reusable Object-Oriented Software."

Design patterns are typically categorized into three main types:

1. Creational Patterns
2. Structural Patterns
3. Behavioral Patterns

## Recommended Books

Before diving into the patterns, here are some excellent resources:

1. **"Design Patterns: Elements of Reusable Object-Oriented Software"** by Gamma, Helm, Johnson, and Vlissides (Gang of Four)
2. **"Head First Design Patterns"** by Eric Freeman and Elisabeth Robson
3. **"Refactoring to Patterns"** by Joshua Kerievsky
4. **"Go Design Patterns"** by Mario Castro Contreras
5. **"Clean Architecture"** by Robert C. Martin
6. **"Patterns of Enterprise Application Architecture"** by Martin Fowler

## 1. Creational Patterns (5 patterns)

Creational patterns provide mechanisms for object creation that increase flexibility and reuse of existing code.

### 1.1 Singleton

**Purpose**: Ensures a class has only one instance and provides a global point of access to it.

**Golang Implementation**:

```go
package singleton

import (
    "sync"
)

type Singleton struct {
    // fields
    data string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{data: "Singleton instance"}
    })
    return instance
}

func (s *Singleton) GetData() string {
    return s.data
}
```

**Other Languages**:

- **Java**: Uses private constructor and static methods
- **Python**: Uses module-level variables or metaclasses
- **JavaScript**: Can use module patterns or ES6 modules

**When to Use**: When exactly one instance of a class is needed, such as for database connections, configuration managers, or thread pools.

### 1.2 Factory Method

**Purpose**: Defines an interface for creating an object, but lets subclasses decide which class to instantiate.

**Golang Implementation**:

```go
package factory

// Product interface
type PaymentMethod interface {
    Pay(amount float64) string
}

// Concrete products
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Credit Card", amount)
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// Factory method
func GetPaymentMethod(method string) (PaymentMethod, error) {
    switch method {
    case "creditcard":
        return &CreditCard{}, nil
    case "paypal":
        return &PayPal{}, nil
    default:
        return nil, fmt.Errorf("Payment method %s not supported", method)
    }
}
```

**Other Languages**:

- **Java**: Often implemented with abstract classes
- **C#**: Uses interfaces and abstract factory patterns
- **Python**: Uses functions that return objects

**When to Use**: When a class can't anticipate the type of objects it needs to create, or when a class wants its subclasses to specify the objects it creates.

### 1.3 Abstract Factory

**Purpose**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

**Golang Implementation**:

```go
package abstractfactory

// Abstract Products
type Button interface {
    Render() string
    Click() string
}

type Checkbox interface {
    Render() string
    Toggle() string
}

// Concrete Products
type WindowsButton struct{}

func (b *WindowsButton) Render() string {
    return "Rendering a Windows button"
}

func (b *WindowsButton) Click() string {
    return "Windows button clicked"
}

type MacButton struct{}

func (b *MacButton) Render() string {
    return "Rendering a Mac button"
}

func (b *MacButton) Click() string {
    return "Mac button clicked"
}

type WindowsCheckbox struct{}

func (c *WindowsCheckbox) Render() string {
    return "Rendering a Windows checkbox"
}

func (c *WindowsCheckbox) Toggle() string {
    return "Windows checkbox toggled"
}

type MacCheckbox struct{}

func (c *MacCheckbox) Render() string {
    return "Rendering a Mac checkbox"
}

func (c *MacCheckbox) Toggle() string {
    return "Mac checkbox toggled"
}

// Abstract Factory
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// Concrete Factories
type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

func (f *WindowsFactory) CreateCheckbox() Checkbox {
    return &WindowsCheckbox{}
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
    return &MacCheckbox{}
}

// Factory selector
func GetFactory(os string) GUIFactory {
    switch os {
    case "windows":
        return &WindowsFactory{}
    case "mac":
        return &MacFactory{}
    default:
        return nil
    }
}
```

**Other Languages**:

- **Java**: Implemented with interfaces and concrete classes
- **C#**: Uses interface inheritance
- **C++**: Uses virtual functions and derived classes

**When to Use**: When a system needs to be independent of how its products are created, composed, and represented, or when a system should be configured with one of multiple families of products.

### 1.4 Builder

**Purpose**: Separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

**Golang Implementation**:

```go
package builder

// Product
type House struct {
    Windows int
    Doors   int
    Rooms   int
    HasGarage bool
    HasSwimmingPool bool
}

// Builder interface
type HouseBuilder interface {
    SetWindows(count int) HouseBuilder
    SetDoors(count int) HouseBuilder
    SetRooms(count int) HouseBuilder
    AddGarage() HouseBuilder
    AddSwimmingPool() HouseBuilder
    Build() *House
}

// Concrete builder
type ConcreteHouseBuilder struct {
    windows int
    doors   int
    rooms   int
    hasGarage bool
    hasSwimmingPool bool
}

func NewHouseBuilder() *ConcreteHouseBuilder {
    return &ConcreteHouseBuilder{}
}

func (b *ConcreteHouseBuilder) SetWindows(count int) HouseBuilder {
    b.windows = count
    return b
}

func (b *ConcreteHouseBuilder) SetDoors(count int) HouseBuilder {
    b.doors = count
    return b
}

func (b *ConcreteHouseBuilder) SetRooms(count int) HouseBuilder {
    b.rooms = count
    return b
}

func (b *ConcreteHouseBuilder) AddGarage() HouseBuilder {
    b.hasGarage = true
    return b
}

func (b *ConcreteHouseBuilder) AddSwimmingPool() HouseBuilder {
    b.hasSwimmingPool = true
    return b
}

func (b *ConcreteHouseBuilder) Build() *House {
    return &House{
        Windows: b.windows,
        Doors:   b.doors,
        Rooms:   b.rooms,
        HasGarage: b.hasGarage,
        HasSwimmingPool: b.hasSwimmingPool,
    }
}

// Director
type HouseDirector struct {
    builder HouseBuilder
}

func NewHouseDirector(b HouseBuilder) *HouseDirector {
    return &HouseDirector{builder: b}
}

func (d *HouseDirector) BuildSmallHouse() *House {
    return d.builder.SetWindows(2).SetDoors(1).SetRooms(2).Build()
}

func (d *HouseDirector) BuildLuxuryHouse() *House {
    return d.builder.SetWindows(6).SetDoors(3).SetRooms(5).AddGarage().AddSwimmingPool().Build()
}
```

**Usage Example**:

```go
builder := NewHouseBuilder()
director := NewHouseDirector(builder)

smallHouse := director.BuildSmallHouse()
luxuryHouse := director.BuildLuxuryHouse()

// Or using builder directly for custom configuration
customHouse := builder.SetWindows(4).SetDoors(2).SetRooms(3).AddGarage().Build()
```

**Other Languages**:

- **Java**: Uses interfaces and method chaining
- **C#**: Similar to Java with fluent interfaces
- **JavaScript**: Uses object literals and functions

**When to Use**: When the algorithm for creating a complex object should be independent of the parts that make up the object, or when the construction process must allow different representations for the object constructed.

### 1.5 Prototype

**Purpose**: Specifies the kinds of objects to create using a prototypical instance and creating new objects by copying this prototype.

**Golang Implementation**:

```go
package prototype

import "fmt"

// Prototype interface
type Cloneable interface {
    Clone() Cloneable
}

// Concrete prototype
type Document struct {
    Content string
    Styles  map[string]string
    Images  []string
}

func (d *Document) Clone() Cloneable {
    clonedDoc := &Document{
        Content: d.Content,
        Styles:  make(map[string]string),
        Images:  make([]string, len(d.Images)),
    }

    // Deep copy maps
    for k, v := range d.Styles {
        clonedDoc.Styles[k] = v
    }

    // Deep copy slices
    copy(clonedDoc.Images, d.Images)

    return clonedDoc
}

func (d *Document) SetContent(content string) {
    d.Content = content
}

func (d *Document) AddStyle(name, value string) {
    d.Styles[name] = value
}

func (d *Document) AddImage(image string) {
    d.Images = append(d.Images, image)
}

// Prototype registry
type DocumentRegistry struct {
    documents map[string]Cloneable
}

func NewDocumentRegistry() *DocumentRegistry {
    return &DocumentRegistry{
        documents: make(map[string]Cloneable),
    }
}

func (r *DocumentRegistry) Register(name string, doc Cloneable) {
    r.documents[name] = doc
}

func (r *DocumentRegistry) Get(name string) Cloneable {
    if doc, ok := r.documents[name]; ok {
        return doc.Clone()
    }
    return nil
}
```

**Usage Example**:

```go
// Create original document
original := &Document{
    Content: "Original content",
    Styles:  make(map[string]string),
    Images:  []string{},
}
original.AddStyle("font", "Arial")
original.AddImage("header.png")

// Register in prototype registry
registry := NewDocumentRegistry()
registry.Register("original", original)

// Get a clone from registry
clonedDoc := registry.Get("original").(*Document)
clonedDoc.SetContent("Modified content")
clonedDoc.AddStyle("color", "blue")

// Original remains unchanged
fmt.Println(original.Content)  // Output: Original content
fmt.Println(clonedDoc.Content) // Output: Modified content
```

**Other Languages**:

- **Java**: Uses the `Cloneable` interface
- **JavaScript**: Uses `Object.create()` or constructor functions
- **C#**: Implements the `ICloneable` interface

**When to Use**: When the classes to instantiate are specified at runtime, or to avoid building a class hierarchy of factories that parallels the class hierarchy of products, or when instances of a class can have one of only a few different combinations of state.

## 2. Structural Patterns (7 patterns)

Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.

### 2.1 Adapter

**Purpose**: Allows objects with incompatible interfaces to collaborate.

**Golang Implementation**:

```go
package adapter

// Target interface
type Target interface {
    Request() string
}

// Adaptee
type LegacySystem struct{}

func (l *LegacySystem) SpecificRequest() string {
    return "Legacy system response"
}

// Adapter
type Adapter struct {
    legacySystem *LegacySystem
}

func NewAdapter(legacySystem *LegacySystem) *Adapter {
    return &Adapter{legacySystem: legacySystem}
}

func (a *Adapter) Request() string {
    return "Adapter: " + a.legacySystem.SpecificRequest()
}
```

**Other Languages**:

- **Java**: Uses inheritance or composition
- **C#**: Uses interfaces and composition
- **Python**: Uses duck typing and composition

**When to Use**: When you want to use an existing class, but its interface doesn't match the one you need, or when you want to create a reusable class that cooperates with unrelated classes with incompatible interfaces.

### 2.2 Bridge

**Purpose**: Separates an abstraction from its implementation so that the two can vary independently.

**Golang Implementation**:

```go
package bridge

// Implementor
type Renderer interface {
    RenderCircle(radius float64) string
    RenderSquare(side float64) string
}

// Concrete implementors
type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) string {
    return fmt.Sprintf("Drawing a circle of radius %.2f using vector graphics", radius)
}

func (v *VectorRenderer) RenderSquare(side float64) string {
    return fmt.Sprintf("Drawing a square with side %.2f using vector graphics", side)
}

type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) string {
    return fmt.Sprintf("Drawing a circle of radius %.2f using raster graphics", radius)
}

func (r *RasterRenderer) RenderSquare(side float64) string {
    return fmt.Sprintf("Drawing a square with side %.2f using raster graphics", side)
}

// Abstraction
type Shape interface {
    Draw() string
    Resize(factor float64)
}

// Refined Abstraction
type Circle struct {
    renderer Renderer
    radius   float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
    return &Circle{
        renderer: renderer,
        radius:   radius,
    }
}

func (c *Circle) Draw() string {
    return c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float64) {
    c.radius *= factor
}

type Square struct {
    renderer Renderer
    side     float64
}

func NewSquare(renderer Renderer, side float64) *Square {
    return &Square{
        renderer: renderer,
        side:     side,
    }
}

func (s *Square) Draw() string {
    return s.renderer.RenderSquare(s.side)
}

func (s *Square) Resize(factor float64) {
    s.side *= factor
}
```

**Usage Example**:

```go
vectorRenderer := &VectorRenderer{}
rasterRenderer := &RasterRenderer{}

circle1 := NewCircle(vectorRenderer, 5)
circle2 := NewCircle(rasterRenderer, 5)
square := NewSquare(vectorRenderer, 10)

fmt.Println(circle1.Draw()) // Drawing a circle of radius 5.00 using vector graphics
fmt.Println(circle2.Draw()) // Drawing a circle of radius 5.00 using raster graphics
fmt.Println(square.Draw())  // Drawing a square with side 10.00 using vector graphics

circle1.Resize(2)
fmt.Println(circle1.Draw()) // Drawing a circle of radius 10.00 using vector graphics
```

**Other Languages**:

- **Java**: Uses abstract classes and interfaces
- **C#**: Uses interfaces and inheritance
- **C++**: Uses abstract base classes and inheritance

**When to Use**: When you want to avoid a permanent binding between an abstraction and its implementation, or when both the abstractions and their implementations should be extensible through subclasses.

### 2.3 Composite

**Purpose**: Composes objects into tree structures to represent part-whole hierarchies, letting clients treat individual objects and compositions uniformly.

**Golang Implementation**:

```go
package composite

import "fmt"

// Component interface
type Component interface {
    Execute() string
    Add(Component)
    Remove(Component)
    GetChildren() []Component
}

// Leaf
type File struct {
    name string
}

func NewFile(name string) *File {
    return &File{name: name}
}

func (f *File) Execute() string {
    return fmt.Sprintf("Processing file: %s", f.name)
}

func (f *File) Add(c Component) {
    // Leaf nodes don't have children
}

func (f *File) Remove(c Component) {
    // Leaf nodes don't have children
}

func (f *File) GetChildren() []Component {
    return nil
}

// Composite
type Directory struct {
    name       string
    components []Component
}

func NewDirectory(name string) *Directory {
    return &Directory{
        name:       name,
        components: []Component{},
    }
}

func (d *Directory) Execute() string {
    result := fmt.Sprintf("Processing directory: %s\n", d.name)
    for _, component := range d.components {
        result += component.Execute() + "\n"
    }
    return result
}

func (d *Directory) Add(component Component) {
    d.components = append(d.components, component)
}

func (d *Directory) Remove(component Component) {
    for i, c := range d.components {
        if c == component {
            d.components = append(d.components[:i], d.components[i+1:]...)
            break
        }
    }
}

func (d *Directory) GetChildren() []Component {
    return d.components
}
```

**Usage Example**:

```go
// Create file objects (leaf nodes)
file1 := NewFile("file1.txt")
file2 := NewFile("file2.txt")
file3 := NewFile("file3.txt")

// Create directory objects (composite nodes)
mainDir := NewDirectory("main")
subDir := NewDirectory("sub")

// Build hierarchy
subDir.Add(file1)
subDir.Add(file2)
mainDir.Add(subDir)
mainDir.Add(file3)

// Execute operations on the entire structure
fmt.Println(mainDir.Execute())
// Output:
// Processing directory: main
// Processing directory: sub
// Processing file: file1.txt
// Processing file: file2.txt
// Processing file: file3.txt
```

**Other Languages**:

- **Java**: Uses inheritance and polymorphism
- **C#**: Uses interfaces and recursive composition
- **JavaScript**: Uses prototype-based inheritance and composition

**When to Use**: When you want to represent part-whole hierarchies of objects, or when you want clients to be able to ignore the difference between compositions of objects and individual objects.

### 2.4 Decorator

**Purpose**: Attaches additional responsibilities to an object dynamically, providing a flexible alternative to subclassing for extending functionality.

**Golang Implementation**:

```go
package decorator

// Component interface
type Component interface {
    Operation() string
}

// Concrete component
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
    return "Basic component"
}

// Base decorator
type Decorator struct {
    component Component
}

func NewDecorator(c Component) *Decorator {
    return &Decorator{component: c}
}

func (d *Decorator) Operation() string {
    return d.component.Operation()
}

// Concrete decorators
type ConcreteDecoratorA struct {
    Decorator
}

func NewConcreteDecoratorA(c Component) *ConcreteDecoratorA {
    return &ConcreteDecoratorA{Decorator: *NewDecorator(c)}
}

func (d *ConcreteDecoratorA) Operation() string {
    return "DecoratorA(" + d.Decorator.Operation() + ")"
}

type ConcreteDecoratorB struct {
    Decorator
}

func NewConcreteDecoratorB(c Component) *ConcreteDecoratorB {
    return &ConcreteDecoratorB{Decorator: *NewDecorator(c)}
}

func (d *ConcreteDecoratorB) Operation() string {
    return "DecoratorB(" + d.Decorator.Operation() + ")"
}
```

**Usage Example**:

```go
// Create a simple component
component := &ConcreteComponent{}
fmt.Println(component.Operation()) // Output: Basic component

// Decorate it with A
decoratedA := NewConcreteDecoratorA(component)
fmt.Println(decoratedA.Operation()) // Output: DecoratorA(Basic component)

// Decorate it with B
decoratedB := NewConcreteDecoratorB(decoratedA)
fmt.Println(decoratedB.Operation()) // Output: DecoratorB(DecoratorA(Basic component))
```

**Real-world Example (HTTP Middleware in Go)**:

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

// Define a middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logger middleware
func Logger(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
        next(w, r)
        fmt.Printf("Request completed in %v\n", time.Since(start))
    }
}

// Auth middleware
func Auth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        fmt.Println("User authenticated")
        next(w, r)
    }
}

// Chain applies middlewares to a handler
func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, middleware := range middlewares {
        h = middleware(h)
    }
    return h
}

func main() {
    // Our base handler
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    }

    // Apply middleware decorators
    decoratedHandler := Chain(handler, Logger, Auth)

    http.HandleFunc("/", decoratedHandler)
    http.ListenAndServe(":8080", nil)
}
```

**Other Languages**:

- **Java**: Uses inheritance and composition
- **C#**: Uses interfaces and composition
- **JavaScript**: Uses higher-order functions or object composition

**When to Use**: When you need to add responsibilities to individual objects dynamically and transparently without affecting other objects, or when extension by subclassing is impractical.

### 2.5 Facade

**Purpose**: Provides a simplified interface to a complex subsystem of classes, making the subsystem easier to use.

**Golang Implementation**:

```go
package facade

// Complex subsystem components
type CPU struct{}

func (c *CPU) Freeze() string {
    return "CPU: Freezing processor."
}

func (c *CPU) Jump(position string) string {
    return fmt.Sprintf("CPU: Jumping to position %s.", position)
}

func (c *CPU) Execute() string {
    return "CPU: Executing commands."
}

type Memory struct{}

func (m *Memory) Load(position string, data string) string {
    return fmt.Sprintf("Memory: Loading data %s at position %s.", data, position)
}

type HardDrive struct{}

func (h *HardDrive) Read(lba int, size int) string {
    return fmt.Sprintf("HardDrive: Reading %d bytes from sector %d.", size, lba)
}

// Facade
type ComputerFacade struct {
    cpu       *CPU
    memory    *Memory
    hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:       &CPU{},
        memory:    &Memory{},
        hardDrive: &HardDrive{},
    }
}

func (c *ComputerFacade) Start() []string {
    results := []string{}
    results = append(results, c.cpu.Freeze())
    results = append(results, c.hardDrive.Read(0, 1024))
    results = append(results, c.memory.Load("0x00", "BOOT SEQUENCE"))
    results = append(results, c.cpu.Jump("0x00"))
    results = append(results, c.cpu.Execute())
    return results
}
```

**Usage Example**:

```go
// Create the facade
computer := NewComputerFacade()

// Use the simplified interface
startupProcess := computer.Start()

// Output the results
for _, step := range startupProcess {
    fmt.Println(step)
}
// Output:
// CPU: Freezing processor.
// HardDrive: Reading 1024 bytes from sector 0.
// Memory: Loading data BOOT SEQUENCE at position 0x00.
// CPU: Jumping to position 0x00.
// CPU: Executing commands.
```

**Other Languages**:

- **Java**: Uses object composition and delegation
- **C#**: Uses similar composition and delegation
- **Python**: Uses simple functions that coordinate subsystems

**When to Use**: When you want to provide a simple interface to a complex subsystem, or when there are many dependencies between clients and implementation classes of an abstraction.

### 2.6 Flyweight

**Purpose**: Uses sharing to support large numbers of similar objects efficiently.

**Golang Implementation**:

```go
package flyweight

import (
    "fmt"
    "strings"
)

// Flyweight
type TextFormatting interface {
    Render(text string) string
}

// Concrete flyweight
type CharacterStyle struct {
    font     string
    size     int
    isBold   bool
    isItalic bool
}

func (c *CharacterStyle) Render(text string) string {
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("Text: %s, Font: %s, Size: %d", text, c.font, c.size))

    if c.isBold {
        sb.WriteString(", Bold")
    }

    if c.isItalic {
        sb.WriteString(", Italic")
    }

    return sb.String()
}

// Flyweight factory
type StyleFactory struct {
    styles map[string]*CharacterStyle
}

func NewStyleFactory() *StyleFactory {
    return &StyleFactory{
        styles: make(map[string]*CharacterStyle),
    }
}

func (f *StyleFactory) GetStyle(font string, size int, bold, italic bool) *CharacterStyle {
    key := fmt.Sprintf("%s-%d-%t-%t", font, size, bold, italic)

    if style, ok := f.styles[key]; ok {
        return style
    }

    style := &CharacterStyle{
        font:     font,
        size:     size,
        isBold:   bold,
        isItalic: italic,
    }

    f.styles[key] = style
    return style
}

func (f *StyleFactory) GetStyleCount() int {
    return len(f.styles)
}

// Context that uses flyweights
type TextEditor struct {
    factory *StyleFactory
    text    []rune
    styles  []*CharacterStyle
}

func NewTextEditor(factory *StyleFactory) *TextEditor {
    return &TextEditor{
        factory: factory,
        text:    []rune{},
        styles:  []*CharacterStyle{},
    }
}

func (e *TextEditor) AppendText(text string, font string, size int, bold, italic bool) {
    style := e.factory.GetStyle(font, size, bold, italic)

    for range text {
        e.text = append(e.text, []rune(text)...)
        for i := 0; i < len(text); i++ {
            e.styles = append(e.styles, style)
        }
    }
}

func (e *TextEditor) GetFormattedText() []string {
    var result []string

    for i, char := range e.text {
        result = append(result, e.styles[i].Render(string(char)))
    }

    return result
}
```

**Usage Example**:

```go
// Create the flyweight factory
factory := NewStyleFactory()

// Create text editor that uses flyweights
editor := NewTextEditor(factory)

// Add text with different styles
editor.AppendText("Hello", "Arial", 12, true, false)
editor.AppendText(" World", "Arial", 12, true, false)
editor.AppendText("!", "Arial", 14, true, true)

// Get formatted output
formatted := editor.GetFormattedText()
for _, f := range formatted {
    fmt.Println(f)
}

// Check how many style objects were actually created
fmt.Printf("Number of style objects created: %d\n", factory.GetStyleCount())
// Output: Number of style objects created: 2
// Even though we have many characters, we only created 2 style objects
```

**Other Languages**:

- **Java**: Uses immutable objects and factory methods
- **C#**: Uses similar approach with factory methods
- **C++**: Uses static member functions and references

**When to Use**: When an application uses a large number of objects that have some shared state among them, or when the application needs a large number of objects, and storage costs are high.

### 2.7 Proxy

**Purpose**: Provides a surrogate or placeholder for another object to control access to it.

**Golang Implementation**:

```go
package proxy

import "fmt"

// Subject interface
type Image interface {
    Display() string
}

// Real subject
type RealImage struct {
    filename string
}

func NewRealImage(filename string) *RealImage {
    image := &RealImage{filename: filename}
    image.loadFromDisk()
    return image
}

func (r *RealImage) loadFromDisk() {
    fmt.Printf("Loading %s from disk\n", r.filename)
}

func (r *RealImage) Display() string {
    return fmt.Sprintf("Displaying %s", r.filename)
}

// Proxy
type ImageProxy struct {
    filename string
    realImage *RealImage
}

func NewImageProxy(filename string) *ImageProxy {
    return &ImageProxy{filename: filename}
}

func (p *ImageProxy) Display() string {
    if p.realImage == nil {
        p.realImage = NewRealImage(p.filename)
    }
    return p.realImage.Display()
}

// Other types of proxies
// Virtual proxy (lazy loading) - implemented above
// Protection proxy
type ProtectedImage struct {
    image Image
    isAdmin bool
}

func NewProtectedImage(image Image, isAdmin bool) *ProtectedImage {
    return &ProtectedImage{
        image: image,
        isAdmin: isAdmin,
    }
}

func (p *ProtectedImage) Display() string {
    if !p.isAdmin {
        return "Access denied: Admin rights required"
    }
    return p.image.Display()
}

// Logging proxy
type LoggingImageProxy struct {
    image Image
}

func NewLoggingImageProxy(image Image) *LoggingImageProxy {
    return &LoggingImageProxy{image: image}
}

func (p *LoggingImageProxy) Display() string {
    fmt.Println("Logging: Image display request")
    result := p.image.Display()
    fmt.Println("Logging: Image display completed")
    return result
}
```

**Usage Example**:

```go
// Using the virtual proxy (lazy loading)
imageProxy := NewImageProxy("sample.jpg")
// RealImage is not loaded yet

fmt.Println(imageProxy.Display())
// Output:
// Loading sample.jpg from disk
// Displaying sample.jpg

fmt.Println(imageProxy.Display())
// Output:
// Displaying sample.jpg (no loading happens - already cached)

// Using the protection proxy
realImage := NewRealImage("confidential.jpg")
protectedImage := NewProtectedImage(realImage, false)
fmt.Println(protectedImage.Display())
// Output: Access denied: Admin rights required

// Change permissions
protectedImageAdmin := NewProtectedImage(realImage, true)
fmt.Println(protectedImageAdmin.Display())
// Output: Displaying confidential.jpg

// Using the logging proxy
loggingImage := NewLoggingImageProxy(realImage)
fmt.Println(loggingImage.Display())
// Output:
// Logging: Image display request
// Displaying confidential.jpg
// Logging: Image display completed
```

**Other Languages**:

- **Java**: Uses interfaces and object composition
- **C#**: Uses similar interfaces and composition
- **Python**: Uses duck typing and `__getattr__` for dynamic proxies

**When to Use**: When you need a more versatile or sophisticated reference to an object than a simple pointer, such as for lazy initialization, access control, logging, or reference counting.

## 3. Behavioral Patterns (11 patterns)

Behavioral patterns are concerned with algorithms and the assignment of responsibilities between objects.

### 3.1 Chain of Responsibility

**Purpose**: Avoids coupling the sender of a request to its receiver by giving more than one object a chance to handle the request.

**Golang Implementation**:

```go
package chainofresponsibility

import "fmt"

// Handler interface
type Handler interface {
    SetNext(handler Handler) Handler
    Handle(request string) string
}

// Base handler (optional in Go, useful for common functionality)
type BaseHandler struct {
    nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
    h.nextHandler = handler
    return handler
}

func (h *BaseHandler) Handle(request string) string {
    if h.nextHandler != nil {
        return h.nextHandler.Handle(request)
    }
    return ""
}

// Concrete handlers
type AuthHandler struct {
    BaseHandler
}

func (h *AuthHandler) Handle(request string) string {
    if request == "auth" {
        return "AuthHandler: I'll handle the authentication request"
    }
    fmt.Println("AuthHandler: Passing to next handler")
    return h.BaseHandler.Handle(request)
}

type ValidationHandler struct {
    BaseHandler
}

func (h *ValidationHandler) Handle(request string) string {
    if request == "validate" {
        return "ValidationHandler: I'll handle the validation request"
    }
    fmt.Println("ValidationHandler: Passing to next handler")
    return h.BaseHandler.Handle(request)
}

type ProcessingHandler struct {
    BaseHandler
}

func (h *ProcessingHandler) Handle(request string) string {
    if request == "process" {
        return "ProcessingHandler: I'll handle the processing request"
    }
    fmt.Println("ProcessingHandler: Passing to next handler")
    return h.BaseHandler.Handle(request)
}

type DefaultHandler struct {
    BaseHandler
}

func (h *DefaultHandler) Handle(request string) string {
    return fmt.Sprintf("DefaultHandler: I'll handle the unknown request '%s'", request)
}
```

**Usage Example**:

```go
// Create handler instances
authHandler := &AuthHandler{}
validationHandler := &ValidationHandler{}
processingHandler := &ProcessingHandler{}
defaultHandler := &DefaultHandler{}

// Build the chain
authHandler.SetNext(validationHandler).SetNext(processingHandler).SetNext(defaultHandler)

// Process requests
fmt.Println(authHandler.Handle("auth"))
// Output: AuthHandler: I'll handle the authentication request

fmt.Println(authHandler.Handle("validate"))
// Output:
// AuthHandler: Passing to next handler
// ValidationHandler: I'll handle the validation request

fmt.Println(authHandler.Handle("process"))
// Output:
// AuthHandler: Passing to next handler
// ValidationHandler: Passing to next handler
// ProcessingHandler: I'll handle the processing request

fmt.Println(authHandler.Handle("unknown"))
// Output:
// AuthHandler: Passing to next handler
// ValidationHandler: Passing to next handler
// ProcessingHandler: Passing to next handler
// DefaultHandler: I'll handle the unknown request 'unknown'
```

**Other Languages**:

- **Java**: Uses abstract classes and inheritance
- **C#**: Uses interfaces and inheritance
- **JavaScript**: Uses higher-order functions or prototypal inheritance

**When to Use**: When more than one object may handle a request and the handler isn't known in advance, or when you want to issue a request to one of several objects without specifying the receiver explicitly.

### 3.2 Command

**Purpose**: Encapsulates a request as an object, allowing parameterization of clients with different requests, queuing of requests, and logging of the requests.

**Golang Implementation**:

```go
package command

import "fmt"

// Command interface
type Command interface {
    Execute() string
    Undo() string
}

// Receiver
type Light struct {
    isOn bool
    room string
}

func NewLight(room string) *Light {
    return &Light{room: room, isOn: false}
}

func (l *Light) TurnOn() string {
    l.isOn = true
    return fmt.Sprintf("Light in %s is now ON", l.room)
}

func (l *Light) TurnOff() string {
    l.isOn = false
    return fmt.Sprintf("Light in %s is now OFF", l.room)
}

// Concrete commands
type LightOnCommand struct {
    light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
    return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() string {
    return c.light.TurnOn()
}

func (c *LightOnCommand) Undo() string {
    return c.light.TurnOff()
}

type LightOffCommand struct {
    light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
    return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() string {
    return c.light.TurnOff()
}

func (c *LightOffCommand) Undo() string {
    return c.light.TurnOn()
}

// Composite command
type MacroCommand struct {
    commands []Command
}

func NewMacroCommand(commands []Command) *MacroCommand {
    return &MacroCommand{commands: commands}
}

func (c *MacroCommand) Execute() string {
    var result string
    for _, command := range c.commands {
        result += command.Execute() + "\n"
    }
    return result
}

func (c *MacroCommand) Undo() string {
    var result string
    // Execute undo in reverse order
    for i := len(c.commands) - 1; i >= 0; i-- {
        result += c.commands[i].Undo() + "\n"
    }
    return result
}

// Invoker
type RemoteControl struct {
    commands []Command
    history  []Command
}

func NewRemoteControl() *RemoteControl {
    return &RemoteControl{
        commands: make([]Command, 0),
        history:  make([]Command, 0),
    }
}

func (r *RemoteControl) AddCommand(command Command) {
    r.commands = append(r.commands, command)
}

func (r *RemoteControl) ExecuteCommand(index int) string {
    if index < 0 || index >= len(r.commands) {
        return "Invalid command index"
    }

    r.history = append(r.history, r.commands[index])
    return r.commands[index].Execute()
}

func (r *RemoteControl) UndoLastCommand() string {
    if len(r.history) == 0 {
        return "No commands to undo"
    }

    lastIndex := len(r.history) - 1
    lastCommand := r.history[lastIndex]
    r.history = r.history[:lastIndex]

    return lastCommand.Undo()
}
```

**Usage Example**:

```go
// Create receivers
kitchenLight := NewLight("Kitchen")
bedroomLight := NewLight("Bedroom")
livingRoomLight := NewLight("Living Room")

// Create commands
kitchenLightOn := NewLightOnCommand(kitchenLight)
kitchenLightOff := NewLightOffCommand(kitchenLight)
bedroomLightOn := NewLightOnCommand(bedroomLight)
bedroomLightOff := NewLightOffCommand(bedroomLight)
livingRoomLightOn := NewLightOnCommand(livingRoomLight)
livingRoomLightOff := NewLightOffCommand(livingRoomLight)

// Create a macro command (turn all lights on)
allLightsOn := NewMacroCommand([]Command{
    kitchenLightOn,
    bedroomLightOn,
    livingRoomLightOn,
})

// Create a remote control (invoker)
remote := NewRemoteControl()
remote.AddCommand(kitchenLightOn)   // Index 0
remote.AddCommand(kitchenLightOff)  // Index 1
remote.AddCommand(bedroomLightOn)   // Index 2
remote.AddCommand(bedroomLightOff)  // Index 3
remote.AddCommand(allLightsOn)      // Index 4

// Execute commands
fmt.Println(remote.ExecuteCommand(0))  // Turn kitchen light on
// Output: Light in Kitchen is now ON

fmt.Println(remote.ExecuteCommand(2))  // Turn bedroom light on
// Output: Light in Bedroom is now ON

fmt.Println(remote.UndoLastCommand())  // Undo last command
// Output: Light in Bedroom is now OFF

fmt.Println(remote.ExecuteCommand(4))  // Execute macro command
// Output:
// Light in Kitchen is now ON
// Light in Bedroom is now ON
// Light in Living Room is now ON
```

**Other Languages**:

- **Java**: Uses interfaces, abstract classes, and inheritance
- **C#**: Uses delegate types and method references
- **JavaScript**: Uses functions as first-class objects

**When to Use**: When you want to parameterize objects with operations, specify, queue, and execute requests at different times, or support undo functionality.

### 3.3 Interpreter

**Purpose**: Given a language, defines a representation for its grammar and an interpreter that uses the representation to interpret sentences in the language.

**Golang Implementation**:

```go
package interpreter

import (
    "fmt"
    "strconv"
    "strings"
)

// Abstract Expression
type Expression interface {
    Interpret(variables map[string]int) int
}

// Terminal Expressions
type NumberExpression struct {
    value int
}

func NewNumberExpression(value int) *NumberExpression {
    return &NumberExpression{value: value}
}

func (e *NumberExpression) Interpret(variables map[string]int) int {
    return e.value
}

type VariableExpression struct {
    name string
}

func NewVariableExpression(name string) *VariableExpression {
    return &VariableExpression{name: name}
}

func (e *VariableExpression) Interpret(variables map[string]int) int {
    if value, exists := variables[e.name]; exists {
        return value
    }
    return 0
}

// Non-terminal Expressions
type AddExpression struct {
    left, right Expression
}

func NewAddExpression(left, right Expression) *AddExpression {
    return &AddExpression{left: left, right: right}
}

func (e *AddExpression) Interpret(variables map[string]int) int {
    return e.left.Interpret(variables) + e.right.Interpret(variables)
}

type SubtractExpression struct {
    left, right Expression
}

func NewSubtractExpression(left, right Expression) *SubtractExpression {
    return &SubtractExpression{left: left, right: right}
}

func (e *SubtractExpression) Interpret(variables map[string]int) int {
    return e.left.Interpret(variables) - e.right.Interpret(variables)
}

// Context/Client
type ExpressionParser struct{}

func (p *ExpressionParser) Parse(input string) (Expression, error) {
    tokens := strings.Fields(input)
    return p.parseExpression(tokens)
}

func (p *ExpressionParser) parseExpression(tokens []string) (Expression, error) {
    if len(tokens) == 0 {
        return nil, fmt.Errorf("empty expression")
    }

    // Simple recursive descent parser for expressions like "x + y - 5"
    var expr Expression
    token := tokens[0]

    // Is it a number?
    if num, err := strconv.Atoi(token); err == nil {
        expr = NewNumberExpression(num)
    } else {
        // Assume it's a variable
        expr = NewVariableExpression(token)
    }

    // Process the rest of the tokens
    for i := 1; i < len(tokens); i += 2 {
        if i+1 >= len(tokens) {
            return nil, fmt.Errorf("invalid expression syntax")
        }

        operator := tokens[i]
        token = tokens[i+1]

        var rightExpr Expression
        if num, err := strconv.Atoi(token); err == nil {
            rightExpr = NewNumberExpression(num)
        } else {
            rightExpr = NewVariableExpression(token)
        }

        switch operator {
        case "+":
            expr = NewAddExpression(expr, rightExpr)
        case "-":
            expr = NewSubtractExpression(expr, rightExpr)
        default:
            return nil, fmt.Errorf("unsupported operator: %s", operator)
        }
    }

    return expr, nil
}
```

**Usage Example**:

```go
// Create a parser
parser := &ExpressionParser{}

// Parse expressions
expr1, _ := parser.Parse("x + y")
expr2, _ := parser.Parse("10 + x - y")
expr3, _ := parser.Parse("x + y - 5")

// Create a context with variables
variables := map[string]int{
    "x": 10,
    "y": 5,
}

// Interpret the expressions
fmt.Printf("x + y = %d\n", expr1.Interpret(variables))
// Output: x + y = 15

fmt.Printf("10 + x - y = %d\n", expr2.Interpret(variables))
// Output: 10 + x - y = 15

fmt.Printf("x + y - 5 = %d\n", expr3.Interpret(variables))
// Output: x + y - 5 = 10

// Change variables and reinterpret
variables["x"] = 7
variables["y"] = 3

fmt.Printf("x + y = %d (after variable change)\n", expr1.Interpret(variables))
// Output: x + y = 10 (after variable change)
```

**Other Languages**:

- **Java**: Uses abstract syntax trees and visitor pattern
- **C#**: Uses similar approach with abstract syntax trees
- **JavaScript**: Uses functions for evaluating expressions

**When to Use**: When you need to interpret a language with a simple grammar, or when you want to create a domain-specific language (DSL) for a particular application.

### 3.4 Iterator

**Purpose**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

**Golang Implementation**:

```go
package iterator

// Iterator interface
type Iterator interface {
    HasNext() bool
    Next() interface{}
}

// Aggregate interface
type Collection interface {
    CreateIterator() Iterator
}

// Concrete iterator
type BookIterator struct {
    books []*Book
    index int
}

func (it *BookIterator) HasNext() bool {
    return it.index < len(it.books)
}

func (it *BookIterator) Next() interface{} {
    if !it.HasNext() {
        return nil
    }
    book := it.books[it.index]
    it.index++
    return book
}

// Concrete collection
type Book struct {
    Title  string
    Author string
}

type BookCollection struct {
    books []*Book
}

func NewBookCollection() *BookCollection {
    return &BookCollection{
        books: []*Book{},
    }
}

func (c *BookCollection) AddBook(book *Book) {
    c.books = append(c.books, book)
}

func (c *BookCollection) CreateIterator() Iterator {
    return &BookIterator{
        books: c.books,
        index: 0,
    }
}

// Filtered iterator
type FilteredBookIterator struct {
    books       []*Book
    index       int
    filterFunc  func(*Book) bool
}

func NewFilteredBookIterator(books []*Book, filterFunc func(*Book) bool) *FilteredBookIterator {
    return &FilteredBookIterator{
        books:      books,
        index:      0,
        filterFunc: filterFunc,
    }
}

func (it *FilteredBookIterator) HasNext() bool {
    for i := it.index; i < len(it.books); i++ {
        if it.filterFunc(it.books[i]) {
            it.index = i
            return true
        }
    }
    return false
}

func (it *FilteredBookIterator) Next() interface{} {
    if !it.HasNext() {
        return nil
    }
    book := it.books[it.index]
    it.index++
    return book
}

// Extension: Filtering collection
func (c *BookCollection) CreateFilteredIterator(filterFunc func(*Book) bool) Iterator {
    return NewFilteredBookIterator(c.books, filterFunc)
}
```

**Usage Example**:

```go
// Create a collection
library := NewBookCollection()
library.AddBook(&Book{Title: "Design Patterns", Author: "Gang of Four"})
library.AddBook(&Book{Title: "Clean Code", Author: "Robert C. Martin"})
library.AddBook(&Book{Title: "Refactoring", Author: "Martin Fowler"})
library.AddBook(&Book{Title: "Domain-Driven Design", Author: "Eric Evans"})

// Use the iterator
fmt.Println("All books:")
iterator := library.CreateIterator()
for iterator.HasNext() {
    book := iterator.Next().(*Book)
    fmt.Printf("- %s by %s\n", book.Title, book.Author)
}

// Use a filtered iterator
fmt.Println("\nBooks by Martin:")
martinFilter := func(book *Book) bool {
    return strings.Contains(book.Author, "Martin")
}
filteredIterator := library.CreateFilteredIterator(martinFilter)
for filteredIterator.HasNext() {
    book := filteredIterator.Next().(*Book)
    fmt.Printf("- %s by %s\n", book.Title, book.Author)
}
```

**Iterating with Go's Built-in Range**:

```go
// Go's built-in iteration is simpler for basic use cases
fmt.Println("\nUsing Go's built-in range:")
for _, book := range library.books {
    fmt.Printf("- %s by %s\n", book.Title, book.Author)
}
```

**Other Languages**:

- **Java**: Has the `Iterator` interface in the Java Collections Framework
- **C#**: Uses `IEnumerable` and `IEnumerator` interfaces
- **JavaScript**: Has built-in iterators and the `Symbol.iterator` protocol

**When to Use**: When you want to provide a standard way to iterate over a collection, or when you want to have multiple traversal methods for a collection, or when you want to hide the underlying complexity of a data structure.

### 3.5 Mediator

**Purpose**: Defines an object that encapsulates how a set of objects interact, promoting loose coupling by keeping objects from referring to each other explicitly.

**Golang Implementation**:

```go
package mediator

import "fmt"

// Mediator interface
type ChatMediator interface {
    SendMessage(message string, user *User)
    AddUser(user *User)
}

// Concrete mediator
type ChatRoom struct {
    users []*User
}

func NewChatRoom() *ChatRoom {
    return &ChatRoom{
        users: make([]*User, 0),
    }
}

func (c *ChatRoom) AddUser(user *User) {
    c.users = append(c.users, user)
}

func (c *ChatRoom) SendMessage(message string, sender *User) {
    for _, user := range c.users {
        // Don't send the message back to the sender
        if user != sender {
            user.Receive(message, sender.name)
        }
    }
}

// Colleague
type User struct {
    name     string
    mediator ChatMediator
}

func NewUser(name string, mediator ChatMediator) *User {
    user := &User{
        name:     name,
        mediator: mediator,
    }
    mediator.AddUser(user)
    return user
}

func (u *User) Send(message string) {
    fmt.Printf("%s sends: %s\n", u.name, message)
    u.mediator.SendMessage(message, u)
}

func (u *User) Receive(message string, from string) {
    fmt.Printf("%s receives from %s: %s\n", u.name, from, message)
}
```

**Usage Example**:

```go
// Create the mediator
chatRoom := NewChatRoom()

// Create users
alice := NewUser("Alice", chatRoom)
bob := NewUser("Bob", chatRoom)
charlie := NewUser("Charlie", chatRoom)

// Users interact through the mediator
alice.Send("Hello everyone!")
// Output:
// Alice sends: Hello everyone!
// Bob receives from Alice: Hello everyone!
// Charlie receives from Alice: Hello everyone!

bob.Send("Hi Alice, nice to meet you!")
// Output:
// Bob sends: Hi Alice, nice to meet you!
// Alice receives from Bob: Hi Alice, nice to meet you!
// Charlie receives from Bob: Hi Alice, nice to meet you!

charlie.Send("Hey folks!")
// Output:
// Charlie sends: Hey folks!
// Alice receives from Charlie: Hey folks!
// Bob receives from Charlie: Hey folks!
```

**Other Languages**:

- **Java**: Uses interfaces and concrete mediator classes
- **C#**: Uses similar approach with interfaces and delegate events
- **JavaScript**: Uses objects and functions as mediators

**When to Use**: When communication between objects is complex but well defined, or when reusing an object is difficult because it communicates with many other objects, or when you want to have a centralized point of communication control.

### 3.6 Memento

**Purpose**: Captures and externalizes an object's internal state without violating encapsulation, allowing the object to be restored to this state later.

**Golang Implementation**:

```go
package memento

import (
    "fmt"
    "time"
)

// Memento
type EditorMemento struct {
    content string
    timestamp time.Time
}

func NewEditorMemento(content string) *EditorMemento {
    return &EditorMemento{
        content:   content,
        timestamp: time.Now(),
    }
}

func (m *EditorMemento) GetContent() string {
    return m.content
}

func (m *EditorMemento) GetTimestamp() time.Time {
    return m.timestamp
}

// Originator
type TextEditor struct {
    content string
}

func NewTextEditor() *TextEditor {
    return &TextEditor{
        content: "",
    }
}

func (e *TextEditor) Type(text string) {
    e.content += text
}

func (e *TextEditor) GetContent() string {
    return e.content
}

func (e *TextEditor) Save() *EditorMemento {
    return NewEditorMemento(e.content)
}

func (e *TextEditor) Restore(memento *EditorMemento) {
    e.content = memento.GetContent()
}

// Caretaker
type History struct {
    mementos []*EditorMemento
    current  int
}

func NewHistory() *History {
    return &History{
        mementos: make([]*EditorMemento, 0),
        current:  -1,
    }
}

func (h *History) Push(memento *EditorMemento) {
    // If we've gone back in history and are now adding a new state,
    // we need to remove all future states
    if h.current < len(h.mementos)-1 {
        h.mementos = h.mementos[:h.current+1]
    }

    h.mementos = append(h.mementos, memento)
    h.current = len(h.mementos) - 1
}

func (h *History) Undo() *EditorMemento {
    if h.current <= 0 {
        // No more undos available
        return nil
    }

    h.current--
    return h.mementos[h.current]
}

func (h *History) Redo() *EditorMemento {
    if h.current >= len(h.mementos)-1 {
        // No more redos available
        return nil
    }

    h.current++
    return h.mementos[h.current]
}

func (h *History) ShowHistory() {
    fmt.Println("History:")
    for i, memento := range h.mementos {
        marker := " "
        if i == h.current {
            marker = "*"
        }
        fmt.Printf("%s %d: %s (%s)\n", marker, i, memento.GetContent(), memento.GetTimestamp().Format("15:04:05"))
    }
}
```

**Usage Example**:

```go
// Create the editor and history
editor := NewTextEditor()
history := NewHistory()

// Initial save
history.Push(editor.Save())

// Make some changes and save states
editor.Type("Hello ")
history.Push(editor.Save())

editor.Type("World ")
history.Push(editor.Save())

editor.Type("from Memento pattern!")
history.Push(editor.Save())

// Show current state
fmt.Printf("Current text: %s\n", editor.GetContent())
// Output: Current text: Hello World from Memento pattern!

// Show history
history.ShowHistory()

// Undo twice
if memento := history.Undo(); memento != nil {
    editor.Restore(memento)
}
if memento := history.Undo(); memento != nil {
    editor.Restore(memento)
}

fmt.Printf("After undo: %s\n", editor.GetContent())
// Output: After undo: Hello

// Redo once
if memento := history.Redo(); memento != nil {
    editor.Restore(memento)
}

fmt.Printf("After redo: %s\n", editor.GetContent())
// Output: After redo: Hello World

// Make a new change from this point
editor.Type("from a new timeline!")
history.Push(editor.Save())

fmt.Printf("New text: %s\n", editor.GetContent())
// Output: New text: Hello World from a new timeline!

// Show updated history
history.ShowHistory()
```

**Other Languages**:

- **Java**: Uses serialization and inner classes
- **C#**: Uses similar approach with serialization
- **JavaScript**: Uses object literals and JSON

**When to Use**: When you need to create snapshots of an object's state to be able to restore it later, or when direct access to an object's fields would break encapsulation.

### 3.7 Observer

**Purpose**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

**Golang Implementation**:

```go
package observer

import "fmt"

// Observer interface
type Observer interface {
    Update(data interface{})
}

// Subject interface
type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify(data interface{})
}

// Concrete subject
type NewsPublisher struct {
    observers []Observer
}

func NewNewsPublisher() *NewsPublisher {
    return &NewsPublisher{
        observers: make([]Observer, 0),
    }
}

func (p *NewsPublisher) Attach(observer Observer) {
    p.observers = append(p.observers, observer)
}

func (p *NewsPublisher) Detach(observer Observer) {
    for i, obs := range p.observers {
        if obs == observer {
            p.observers = append(p.observers[:i], p.observers[i+1:]...)
            break
        }
    }
}

func (p *NewsPublisher) Notify(data interface{}) {
    for _, observer := range p.observers {
        observer.Update(data)
    }
}

func (p *NewsPublisher) PublishNews(title string, content string) {
    news := map[string]string{
        "title":   title,
        "content": content,
    }
    fmt.Printf("Publishing news: %s\n", title)
    p.Notify(news)
}

// Concrete observers
type EmailSubscriber struct {
    email string
}

func NewEmailSubscriber(email string) *EmailSubscriber {
    return &EmailSubscriber{
        email: email,
    }
}

func (s *EmailSubscriber) Update(data interface{}) {
    news, ok := data.(map[string]string)
    if !ok {
        return
    }

    fmt.Printf("Sending email to %s - Breaking News: %s\n", s.email, news["title"])
}

type SMSSubscriber struct {
    phoneNumber string
}

func NewSMSSubscriber(phoneNumber string) *SMSSubscriber {
    return &SMSSubscriber{
        phoneNumber: phoneNumber,
    }
}

func (s *SMSSubscriber) Update(data interface{}) {
    news, ok := data.(map[string]string)
    if !ok {
        return
    }

    fmt.Printf("Sending SMS to %s - Alert: %s\n", s.phoneNumber, news["title"])
}

type WebHookSubscriber struct {
    url string
}

func NewWebHookSubscriber(url string) *WebHookSubscriber {
    return &WebHookSubscriber{
        url: url,
    }
}

func (s *WebHookSubscriber) Update(data interface{}) {
    news, ok := data.(map[string]string)
    if !ok {
        return
    }

    fmt.Printf("Triggering webhook at %s with data: %s\n", s.url, news["title"])
}
```

**Usage Example**:

```go
// Create the publisher
publisher := NewNewsPublisher()

// Create subscribers
emailSub1 := NewEmailSubscriber("john@example.com")
emailSub2 := NewEmailSubscriber("jane@example.com")
smsSub := NewSMSSubscriber("+1234567890")
webhookSub := NewWebHookSubscriber("https://example.com/webhook")

// Register subscribers
publisher.Attach(emailSub1)
publisher.Attach(emailSub2)
publisher.Attach(smsSub)
publisher.Attach(webhookSub)

// Publish news - all subscribers will be notified
publisher.PublishNews("Major Weather Alert", "Severe storms expected in the area.")

// Remove a subscriber
publisher.Detach(emailSub2)

// Add a mobile app subscriber
mobileAppSub := NewMobileAppSubscriber("user123")
publisher.Attach(mobileAppSub)

// Publish more news
publisher.PublishNews("Traffic Update", "Highway 101 reopened after accident cleared.")
publisher.PublishNews("Local Event", "Annual festival starts this weekend.")
```

### 3.8 State

**Purpose**: Allows an object to alter its behavior when its internal state changes. The object will appear to change its class.

**Golang Implementation**:

```go
package state

import "fmt"

// State interface
type State interface {
    Handle(context *Context)
    GetName() string
}

// Context that maintains a reference to a state object
type Context struct {
    state State
}

func NewContext() *Context {
    return &Context{
        state: &ConcreteStateA{},
    }
}

func (c *Context) SetState(state State) {
    fmt.Printf("Changing state from %s to %s\n", c.state.GetName(), state.GetName())
    c.state = state
}

func (c *Context) Request() {
    c.state.Handle(c)
}

// Concrete State A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
    fmt.Println("ConcreteStateA handles the request.")
    context.SetState(&ConcreteStateB{})
}

func (s *ConcreteStateA) GetName() string {
    return "StateA"
}

// Concrete State B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
    fmt.Println("ConcreteStateB handles the request.")
    context.SetState(&ConcreteStateA{})
}

func (s *ConcreteStateB) GetName() string {
    return "StateB"
}
```

**Usage Example**:

```go
// Create a new context with default state A
context := NewContext()

// Request will be handled by state A and transition to state B
context.Request()

// Request will be handled by state B and transition back to state A
context.Request()

// Request will be handled by state A and transition to state B again
context.Request()
```

### 3.9 Strategy

**Purpose**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

**Golang Implementation**:

```go
package strategy

import "fmt"

// Strategy interface
type PaymentStrategy interface {
    Pay(amount float64) bool
}

// Concrete strategies
type CreditCardPayment struct {
    cardNumber string
    cvv        string
    name       string
}

func NewCreditCardPayment(cardNumber, cvv, name string) *CreditCardPayment {
    return &CreditCardPayment{
        cardNumber: cardNumber,
        cvv:        cvv,
        name:       name,
    }
}

func (p *CreditCardPayment) Pay(amount float64) bool {
    fmt.Printf("Paying %.2f using Credit Card\n", amount)
    return true
}

type PayPalPayment struct {
    email    string
    password string
}

func NewPayPalPayment(email, password string) *PayPalPayment {
    return &PayPalPayment{
        email:    email,
        password: password,
    }
}

func (p *PayPalPayment) Pay(amount float64) bool {
    fmt.Printf("Paying %.2f using PayPal\n", amount)
    return true
}

type BankTransferPayment struct {
    accountNumber string
    bankCode      string
}

func NewBankTransferPayment(accountNumber, bankCode string) *BankTransferPayment {
    return &BankTransferPayment{
        accountNumber: accountNumber,
        bankCode:      bankCode,
    }
}

func (p *BankTransferPayment) Pay(amount float64) bool {
    fmt.Printf("Paying %.2f using Bank Transfer\n", amount)
    return true
}

// Context
type ShoppingCart struct {
    paymentStrategy PaymentStrategy
    items           []float64
}

func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        items: make([]float64, 0),
    }
}

func (c *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
    c.paymentStrategy = strategy
}

func (c *ShoppingCart) AddItem(price float64) {
    c.items = append(c.items, price)
}

func (c *ShoppingCart) GetTotal() float64 {
    total := 0.0
    for _, price := range c.items {
        total += price
    }
    return total
}

func (c *ShoppingCart) Checkout() bool {
    if c.paymentStrategy == nil {
        fmt.Println("No payment strategy set")
        return false
    }
    return c.paymentStrategy.Pay(c.GetTotal())
}
```

**Usage Example**:

```go
// Create a new shopping cart
cart := NewShoppingCart()

// Add some items
cart.AddItem(79.99)
cart.AddItem(29.50)

// Create payment strategies
creditCard := NewCreditCardPayment("1234-5678-9012-3456", "123", "John Doe")
paypal := NewPayPalPayment("john@example.com", "password123")
bankTransfer := NewBankTransferPayment("987654321", "ABC123")

// Set payment strategy and checkout
cart.SetPaymentStrategy(creditCard)
cart.Checkout()

// Change payment strategy and checkout again
cart.SetPaymentStrategy(paypal)
cart.Checkout()

// Change payment strategy again and checkout
cart.SetPaymentStrategy(bankTransfer)
cart.Checkout()
```

### 3.10 Template Method

**Purpose**: Defines the skeleton of an algorithm in a method, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

**Golang Implementation**:

```go
package template

import "fmt"

// Abstract class (interface with concrete method)
type DataProcessor interface {
    OpenFile()
    ProcessData()
    CloseFile()
    WriteReport()
}

// Template method
func ProcessTemplate(processor DataProcessor) {
    processor.OpenFile()
    processor.ProcessData()
    processor.CloseFile()
    processor.WriteReport()
}

// Abstract class as a struct with embedded interface
type AbstractDataProcessor struct {
    DataProcessor
}

func (p *AbstractDataProcessor) OpenFile() {
    fmt.Println("Opening file...")
}

func (p *AbstractDataProcessor) CloseFile() {
    fmt.Println("Closing file...")
}

func (p *AbstractDataProcessor) WriteReport() {
    fmt.Println("Writing report...")
}

// Concrete implementations
type CSVProcessor struct {
    AbstractDataProcessor
}

func NewCSVProcessor() *CSVProcessor {
    processor := &CSVProcessor{}
    processor.DataProcessor = processor
    return processor
}

func (p *CSVProcessor) ProcessData() {
    fmt.Println("Processing CSV data...")
}

type XMLProcessor struct {
    AbstractDataProcessor
}

func NewXMLProcessor() *XMLProcessor {
    processor := &XMLProcessor{}
    processor.DataProcessor = processor
    return processor
}

func (p *XMLProcessor) ProcessData() {
    fmt.Println("Processing XML data...")
}

type JSONProcessor struct {
    AbstractDataProcessor
}

func NewJSONProcessor() *JSONProcessor {
    processor := &JSONProcessor{}
    processor.DataProcessor = processor
    return processor
}

func (p *JSONProcessor) ProcessData() {
    fmt.Println("Processing JSON data...")
}

func (p *JSONProcessor) WriteReport() {
    fmt.Println("Writing JSON-specific report with additional metrics...")
}
```

**Usage Example**:

```go
// Create processors for different data formats
csvProcessor := NewCSVProcessor()
xmlProcessor := NewXMLProcessor()
jsonProcessor := NewJSONProcessor()

// Process data using the template method
fmt.Println("Processing CSV:")
ProcessTemplate(csvProcessor)

fmt.Println("\nProcessing XML:")
ProcessTemplate(xmlProcessor)

fmt.Println("\nProcessing JSON:")
ProcessTemplate(jsonProcessor)
```

### 3.11 Visitor

**Purpose**: Represents an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.

**Golang Implementation**:

```go
package visitor

import "fmt"

// Element interface
type Element interface {
    Accept(visitor Visitor)
}

// Visitor interface
type Visitor interface {
    VisitCircle(circle *Circle)
    VisitRectangle(rectangle *Rectangle)
    VisitTriangle(triangle *Triangle)
}

// Concrete Elements
type Circle struct {
    Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
    visitor.VisitCircle(c)
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
    visitor.VisitRectangle(r)
}

type Triangle struct {
    Base   float64
    Height float64
}

func (t *Triangle) Accept(visitor Visitor) {
    visitor.VisitTriangle(t)
}

// Concrete Visitors
type AreaVisitor struct {
    TotalArea float64
}

func (v *AreaVisitor) VisitCircle(circle *Circle) {
    area := 3.14 * circle.Radius * circle.Radius
    fmt.Printf("Circle area: %.2f\n", area)
    v.TotalArea += area
}

func (v *AreaVisitor) VisitRectangle(rectangle *Rectangle) {
    area := rectangle.Width * rectangle.Height
    fmt.Printf("Rectangle area: %.2f\n", area)
    v.TotalArea += area
}

func (v *AreaVisitor) VisitTriangle(triangle *Triangle) {
    area := 0.5 * triangle.Base * triangle.Height
    fmt.Printf("Triangle area: %.2f\n", area)
    v.TotalArea += area
}

type PerimeterVisitor struct {
    TotalPerimeter float64
}

func (v *PerimeterVisitor) VisitCircle(circle *Circle) {
    perimeter := 2 * 3.14 * circle.Radius
    fmt.Printf("Circle perimeter: %.2f\n", perimeter)
    v.TotalPerimeter += perimeter
}

func (v *PerimeterVisitor) VisitRectangle(rectangle *Rectangle) {
    perimeter := 2 * (rectangle.Width + rectangle.Height)
    fmt.Printf("Rectangle perimeter: %.2f\n", perimeter)
    v.TotalPerimeter += perimeter
}

func (v *PerimeterVisitor) VisitTriangle(triangle *Triangle) {
    // Simplified calculation (equilateral triangle assumed)
    perimeter := 3 * triangle.Base
    fmt.Printf("Triangle perimeter: %.2f\n", perimeter)
    v.TotalPerimeter += perimeter
}
```

**Usage Example**:

```go
// Create elements
circle := &Circle{Radius: 5}
rectangle := &Rectangle{Width: 4, Height: 6}
triangle := &Triangle{Base: 3, Height: 4}

// Create a shape collection
shapes := []Element{circle, rectangle, triangle}

// Create visitors
areaVisitor := &AreaVisitor{}
perimeterVisitor := &PerimeterVisitor{}

// Calculate areas
fmt.Println("Calculating areas:")
for _, shape := range shapes {
    shape.Accept(areaVisitor)
}
fmt.Printf("Total area: %.2f\n\n", areaVisitor.TotalArea)

// Calculate perimeters
fmt.Println("Calculating perimeters:")
for _, shape := range shapes {
    shape.Accept(perimeterVisitor)
}
fmt.Printf("Total perimeter: %.2f\n", perimeterVisitor.TotalPerimeter)
```

# Conclusion

Design patterns are essential tools in a developer's toolkit. They provide tested solutions to common problems, help create more maintainable code, and establish a common vocabulary for teams. While the examples in this guide are in Go, the concepts apply across programming languages.

Remember that patterns should be used judiciously. Not every problem requires a pattern, and overuse can lead to unnecessarily complex code. Always consider the specific requirements of your application before applying a pattern.

The patterns in this guide are not exhaustive but cover the most widely used ones. As you gain experience, you'll recognize when and how to apply these patterns effectively in your own code.

# Recommended Books

1. **Design Patterns: Elements of Reusable Object-Oriented Software** (Gang of Four book) - The classic reference
2. **Head First Design Patterns** - A more beginner-friendly approach
3. **Learning Go Design Patterns** - Specifically for Go developers
4. **Design Patterns in Go** - Another Go-specific resource
5. **Clean Code** by Robert C. Martin - Includes design pattern concepts

# Learning Path

## Beginner Level

- Start with Singleton, Factory, Builder patterns
- Understand Decorator, Adapter, Facade patterns
- Learn Observer and Strategy patterns

## Intermediate Level

- Master Abstract Factory, Prototype
- Deep dive into Composite, Bridge, Proxy
- Practice Command, Iterator, State patterns

### Advanced Level

- Understand Flyweight, Mediator, Memento
- Master Visitor, Interpreter patterns
- Learn to combine patterns effectively

Remember that design patterns are tools, not rules. Use them when they solve a real problem in your code, not just because they exist. In Go specifically, many patterns can be implemented more simply than in classical OOP languages due to Go's interface system and first-class functions.
