package main

import "fmt"

// Box represents a box that can contain items or other boxes.
type Box struct {
	items []interface{} // Can contain nested Boxes or a "key"
}

// isKey checks if an item is a key.
func isKey(item interface{}) bool {
	key, ok := item.(string)
	return ok && key == "key"
}

// isBox checks if an item is a box.
func isBox(item interface{}) bool {
	_, ok := item.(Box)
	return ok
}

// LookForKey searches for a key inside nested boxes using an iterative approach.
func LookForKey(mainBox Box) {
	stack := []Box{mainBox} // Initialize stack with main box

	for len(stack) > 0 {
		// Pop a box from the stack
		box := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Iterate through the box items
		for _, item := range box.items {
			if isKey(item) {
				fmt.Println("Found the key!")
				return
			} else if isBox(item) {
				// Convert item to Box type and push it onto the stack
				stack = append(stack, item.(Box))
			}
		}
	}

	fmt.Println("Key not found!")
}


func main() {
	// Creating nested boxes
	mainBox := Box{
		items: []interface{}{
			Box{
				items: []interface{}{
					"apple",
					Box{
						items: []interface{}{
							"toy",
							Box{items: []interface{}{"key", "book"}},
						},
					},
					"pen",
				},
			},
			Box{items: []interface{}{"notebook"}},
			"scissors",
		},
	}

	// Search for the key
	LookForKey(mainBox)
}
