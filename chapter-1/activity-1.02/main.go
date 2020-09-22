package main
import "fmt"
func main() {
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println( a == 10, b == 5 )
}

func swap(a *int, b *int) {
/*	I spent a lot more time on this activity than I thought I would,
	mostly because pointers confuse me af, and I was doing some investigations ...
	I also found this blog post https://dave.cheney.net/2014/03/17/pointers-in-go
	which has a little more explination on the topic and a link to a book that could be useful.
	I stopped at "Attempt 3" (line 56) before lookin at the course's solution ...
	Uncomment a block, save and run if interested in any of these variations.
	I haven't used pointers since I was learning C 5 years ago,
	so maybe this is not the most practical thing to focus on, as the course says ...
	*/

/*	 Attempt 1 ... 
	At first I thought of using 2 temporary values to do the swap.
	The problem I think is it totally defeats the point of pointers.
	With this implementation, there are two new memory locations needed.
	We might as well just copy the values and forget the pointers. 
	Though the below code does output "true true".

	_a := *a // _a is a new storage for the value that a points to
	_b := *b // _b is a new storage for the value that b points to

	// now that the values are held in _a and _b,
	// the pointers' values can re reassigned
	*a = _b
	*b = _a
	*/

/*	 
	// Attempt 2 ...
	// To make pointers more useful, I thought I could make a new pointer to point to the value of a:
	_a := a	_a := a
	fmt.Printf("_a: %#v\t a: %#v\n", _a, a) 
	// _a and a are identical (they are both pointers, pointing to the same value)

	// and then set a to point to b's value
	a = b
	fmt.Printf("b: %#v\t a: %#v\n", b, a)

	// and set b to point to _a's value
	b = _a
	fmt.Printf("b: %#v\t _a: %#v\n", b, _a)

	// This works in swapping the pointers (as can be seen with the print statements),
	// but it doesnt change the values they point to, which needs to happen outside of the scope of the function.
	// Since this function is not supposed to return anything, we need to edit the values themselves inside it.
	*/
	
/*	// Attempt 3 ...
	// Since the values themselves need to be edited, I went back to using a temporary variable (only one this time):
	aOld := *a 
	// aOld is a variable and has its own memory allocation. It is assigned the value of a before doing swapping.
	*a = *b 
	// the value of a (written *a) is set equal to the value of b (*b)
	*b = aOld 
	// the value of b (written *b) is set equal to the old value of a (variable aOld)
	*/

/*	 
	// Course's solution ...
	// I would not have guessed the course's solution, because apparently Go's ability to do multiple assignments 
	// means that the right-hand side resolves before the left-hand side does.
	// The course's solution:
		*a, *b = *b, *a

	// This seems non-logical to me, but I guess the compiler can manage the needed temporaty memory allocation...
	// ... perhaps this illustrates one of the benefits of GO over C.
	*/
	
}