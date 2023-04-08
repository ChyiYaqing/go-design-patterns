# 工厂方法 Factory Method Pattern

Factory method creational design pattern allows creating objects without having
to specify the exact type of the object that will be created.

## Implementation

The example to demonstrate the benefits of using Factory Methods, as there is not need to use factories at all if you're dealing with one struct (an object in OO world).

### Source Code

```go
package main

import "fmt"

type dataField interface {
	Print()
}

type dataField1 struct {
	name  string
	value string
}

func (df *dataField1) Print() {
	fmt.Println("dataField1 ->", df.name, ":", df.value)
}

type dataField2 struct {
	name  string
	value string
}

func (df *dataField2) Print() {
	fmt.Println("dataField2 ->", df.name, ":", df.value)
}

type dataFieldFactory interface {
	Create(name, value string) dataField
}

type dataField1Factory struct{}

func (factory *dataField1Factory) Create(name, value string) dataField {
	return &dataField1{
		name:  name,
		value: value,
	}
}

type dataField2Factory struct{}

func (factory *dataField2Factory) Create(name, value string) dataField {
	return &dataField2{
		name:  name,
		value: value,
	}
}

type Document struct {
	dataFieldFactories []dataFieldFactory
	allValues          [][]string
}

func (doc *Document) Print() {
	for i, factory := range doc.dataFieldFactories {
		field := factory.Create(doc.allValues[i][0], doc.allValues[i][1])
		field.Print()
	}
}

func main() {
	doc := &Document{
		dataFieldFactories: []dataFieldFactory{
			&dataField1Factory{},
			&dataField2Factory{},
		},
		allValues: [][]string{{"name1", "value1"}, {"name2", "value2"}},
	}
	doc.Print()
}
```

## Usage

With the factory method, the user can specify the type of storage they want.

```shell
➜ go run factory.go
dataField1 -> name1 : value1
dataField2 -> name2 : value2
```
look into main func you don't find any mentions or initializations of concrete types `dataField1` and `dataField2`. All the complexity is hidden begind `dataFieldFactories`. Both `dataField1Factory` and `dataField2Factory` implement `Create` interface and return dataField interface, that both the concrete types also implement, So, you can call `Print()` for each of the concrete types.

`Document.Print()` uses both the Create and Print interfaces for printing out all the fields without any knowledge about how the fields are actually created and printed. We're achieving this by providing a list of Factory Methods (`dataFied1Factory{}` and `dataField2Factory{}`) and corresponding string values to the Document struct (object).