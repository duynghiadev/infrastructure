# Website

1. [https://refactoring.guru/design-patterns](https://refactoring.guru/design-patterns)
2. [https://blog.ntechdevelopers.com/gof-design-patterns-no-la-gi-ma-tai-sao-cac-senior-bat-buoc-phai-biet/](https://blog.ntechdevelopers.com/gof-design-patterns-no-la-gi-ma-tai-sao-cac-senior-bat-buoc-phai-biet/)
3. [https://gpcoder.com/4164-gioi-thieu-design-patterns/](https://gpcoder.com/4164-gioi-thieu-design-patterns/)
4. [https://en.wikipedia.org/wiki/Software_design_pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
5. [https://viblo.asia/s/tong-hop-23-mau-design-patterns-tro-thu-dac-luc-cua-developers-Q75wqJ67ZWb](https://viblo.asia/s/tong-hop-23-mau-design-patterns-tro-thu-dac-luc-cua-developers-Q75wqJ67ZWb)

# Introduction to Design Patterns in Golang (and Other Languages)

## 1. What are Design Patterns?

Design patterns are reusable solutions to common software design problems. They provide best practices to structure code, improve maintainability, and enhance scalability. Patterns are not specific to any language but can be implemented in multiple programming languages, including Golang, Java, Python, and JavaScript.

Design patterns are categorized into three main types:

- **Creational Patterns** : Deal with object creation mechanisms.
- **Structural Patterns** : Focus on class and object composition.
- **Behavioral Patterns** : Define communication between objects.

---

## 2. Creational Patterns

### a) Singleton Pattern

Ensures a class has only one instance and provides a global access point to it.

#### **Golang Implementation**

```go
package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	obj1 := GetInstance()
	obj2 := GetInstance()
	fmt.Println(obj1 == obj2) // Output: true
}
```

#### **Java Implementation**

```java
public class Singleton {
    private static Singleton instance;
    private Singleton() {}
    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
```

---

## 3. Structural Patterns

### b) Adapter Pattern

Allows objects with incompatible interfaces to work together.

#### **Golang Implementation**

```go
package main
import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific request"
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}
	fmt.Println(adapter.Request())
}
```

---

## 4. Behavioral Patterns

### c) Observer Pattern

Defines a dependency between objects so that when one object changes state, all dependents are notified.

#### **Golang Implementation**

```go
package main

import "fmt"

type Observer interface {
	Update(string)
}

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify(string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Deregister(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

type ConcreteObserver struct {
	id string
}

func (o *ConcreteObserver) Update(msg string) {
	fmt.Printf("Observer %s received: %s\n", o.id, msg)
}

func main() {
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{id: "1"}
	observer2 := &ConcreteObserver{id: "2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.Notify("Hello Observers!")
}
```

---

## 5. Recommended Books

1. **Design Patterns: Elements of Reusable Object-Oriented Software** – Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides
2. **Head First Design Patterns** – Eric Freeman, Bert Bates, Kathy Sierra
3. **Dive Into Design Patterns** – Alexander Shvets
4. **Go Design Patterns** – Mario Castro Contreras

---

## 6. Conclusion

Understanding design patterns helps you write more efficient and scalable code. While patterns are not a silver bullet, they provide a structured way to solve recurring problems. By applying them in Golang (or any other language), you can build maintainable and extensible applications.

### **Next Steps:**

- Start implementing patterns in your projects.
- Explore advanced patterns like Dependency Injection and CQRS.
- Read more about software architecture to understand when to use patterns effectively.
