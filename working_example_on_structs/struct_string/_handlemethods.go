                          |
// func (p *pp) handleMethods(argument interface{}) (handled bool) {
// 	// ...

// 	// Checks whether a given argument is an error or an fmt.Stringer
// 	switch v := argument.(type) {
// 	// ...
// 	// If the argument is a Stringer, calls its String() method
// 	case Stringer:
// 		// ...
// 		//        pocket.String()
// 		//              ^
// 		//              |
// 		p.fmtString(v.String(), verb)
// 		return
// 	}

// 	// ...
// }
