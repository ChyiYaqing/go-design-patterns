## Design Patterns in Go

* Design patterns identify solutions to common programming problems, but they are not actual code or algorithm themseleves

### Go language Features That Affect Design Patterns

* No support for traditional OOP classes like in C# or Java
  - No static members or constructors -- affect patterns like Singleton
* No support for class-based inheritance or method overloading
  - Affects patterns like Visitor
* No support for exceptions - error handling is via return values
  - Affects patterns like Builder
* No support for abstract classes
  - Affects patterns like Abstract Factory and Bridge.

### Design Pattern Categories

* Creational
  - Abstract Factory
  - Builder
  - Factory
  - Lazy Initialization
  - Multition
  - Object Pool
  - Prototype
  - Singleton

* Structural
  - Adapter
  - Bridge
  - Composite
  - Decorator
  - Facade
  - Flyweight
  - Proxy

* Behavioral
  - Chain of Responsibility
  - Command
  - Interpreter
  - Iterator
  - Mediator
  - Memento
  - Observer
  - Strategy
  - Visitor
