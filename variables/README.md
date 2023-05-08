`var name type = expression`

Either the typ e or the = expression part may be omitt ed, but not bot h. If the typ e is omitt ed,
it is deter mined by the initializer expression. If the expression is omitt ed, the initial value is
the zero value for the typ e, which is 0 for numbers, false for boole ans, "" for str ings, and nil
for int erfaces and reference typ es (slice, point er, map, channel, function). The zero value of an
ag gregate typ e li ke an array orastr uct has the zero value of all of its elements or fields.


## Simple Data Types: [They contain a sinlgle value]
1. Strings: One or more UTF-8 code. 
a. Interpreted string: “this is a string”
b. Raw string: `this is also a string`. In this formath it will ignore \n character. 
2. Numbers: 
a. Integers: int, int8, int16
b. Unsigned integers: uint , minimum value 0
c. Floating point numbers: float32, float64 [32 bit and 64 bit floating point number
d. Complex numbers: complex64, complex128. Use for math and statistics.  
3. Booleans: true/false. 
4. Errors: The error built-in interface type is the conventional interface for resenting an error condition, with the nil value representing no error. 

type error interface {
   Error() string
}