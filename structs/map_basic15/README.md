## **Explanation of the Code**
This Go program defines an `Engineer` struct, which contains information about a **software engineer** and their assigned **project**. It also includes methods to print the engineer's details and retrieve the project’s priority.

---

## **1. Defining the `Engineer` Struct**
```go
type Engineer struct {
	Name    string
	Age     int
	Project Project
}
```
- The struct contains:
  - `Name`: **string** → Stores the name of the engineer.
  - `Age`: **int** → Stores the age of the engineer.
  - `Project`: **Project struct** → Stores the project the engineer is working on.

---

## **2. Defining the `Project` Struct**
```go
type Project struct {
	Name         string
	Priority     string
	Technologies []string
}
```
- The struct contains:
  - `Name`: **string** → Stores the project name.
  - `Priority`: **string** → Stores the priority of the project.
  - `Technologies`: **slice of strings** → Stores the technologies used in the project.

---

## **3. Method: `Print()`**
```go
func (this Engineer) Print() {
	fmt.Printf("%+v\n", this)
}
```
- This method prints all details of an `Engineer` object.
- `%+v` **formats the struct with field names**.

---

## **4. Method: `GetProjectPriority()`**
```go
func (this Engineer) GetProjectPriority() {
	fmt.Printf("%+v\n", this.Project.Priority)
}
```
- Prints the **priority** of the engineer's project.

---

## **5. Main Function**
```go
func main() {
	engineer := Engineer{
		Name: "Mirzohid",
		Age:  23,
		Project: Project{
			Name:         "Onlayin taksi app yaratish",
			Priority:     "Vaqtni tejash",
			Technologies: []string{
				"Golang",
				"Flutter",
				"Gin",
				"Doker",
				"Microserves",
			},
		},
	}

	engineer.Print()
	engineer.GetProjectPriority()
}
```
### **Step-by-Step Execution**
1. Creates an `Engineer` object named **Mirzohid**.
2. Assigns a project: **"Onlayin taksi app yaratish"** with priority **"Vaqtni tejash"**.
3. Uses `.Print()` to **display all details** of the engineer.
4. Uses `.GetProjectPriority()` to **print the project's priority**.

---

## **Example Output**
```
{Name:Mirzohid Age:23 Project:{Name:Onlayin taksi app yaratish Priority:Vaqtni tejash Technologies:[Golang Flutter Gin Doker Microserves]}}
Vaqtni tejash
```
- The **first line** shows the engineer's **full details**.
- The **second line** shows the **priority** of the project.

---

## **How to Improve the Code?**
✅ **Better formatting in `Print()`**
```go
func (this Engineer) Print() {
	fmt.Println("Engineer Info:")
	fmt.Println("Name: ", this.Name)
	fmt.Println("Age: ", this.Age)
	fmt.Println("Project Name: ", this.Project.Name)
	fmt.Println("Priority: ", this.Project.Priority)
	fmt.Println("Technologies: ", this.Project.Technologies)
}
```
✅ **Return priority instead of printing**
```go
func (this Engineer) GetProjectPriority() string {
	return this.Project.Priority
}
```
✅ **Using it in `main()`**
```go
fmt.Println("Project Priority: ", engineer.GetProjectPriority())
```

---

## **Final Optimized Code**
```go
package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

// Print engineer details
func (this Engineer) Print() {
	fmt.Println("Engineer Info:")
	fmt.Println("Name: ", this.Name)
	fmt.Println("Age: ", this.Age)
	fmt.Println("Project Name: ", this.Project.Name)
	fmt.Println("Priority: ", this.Project.Priority)
	fmt.Println("Technologies: ", this.Project.Technologies)
}

// Get project priority
func (this Engineer) GetProjectPriority() string {
	return this.Project.Priority
}

func main() {
	engineer := Engineer{
		Name: "Mirzohid",
		Age:  23,
		Project: Project{
			Name:         "Onlayin taksi app yaratish",
			Priority:     "Vaqtni tejash",
			Technologies: []string{"Golang", "Flutter", "Gin", "Doker", "Microservices"},
		},
	}

	engineer.Print()
	fmt.Println("Project Priority: ", engineer.GetProjectPriority())
}
```