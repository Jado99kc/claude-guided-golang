package main

import "fmt"

type Animal interface{ Sound() string }
type Dog struct{}

func (d *Dog) Sound() string { return "woof" }

type Logger struct{ prefix string }

func (l Logger) Log(m string) { fmt.Println(l.prefix + m) }

type Service struct {
	Logger
	name string
}

func main() {
	var a Animal = &Dog{} // only *Dog satisfies Animal interface (regular Dog doesn't because the methiod is defined on *Dog)
	fmt.Println(a.Sound())

	var d *Dog
	a = d
	fmt.Println(a == nil, a.Sound())

	svc := Service{Logger: Logger{prefix: "[svc] "}, name: "orders"}
	svc.Log("started " + svc.name)

	var st fmt.Stringer
	fmt.Println(st == nil)
}

//Predictions
// fmt.Println(a.Sound()) ->  will print "woof" -> Dog satisfies the Animal interface which then has a sound method that points to the original dog type for execution.
//correct answer: "*Dog has Sound in its method set, so &Dog{} can be stored in a. The call dispatches to (*Dog).Sound, which returns woof."
// fmt.Println(a == nil, a.Sound()) -> will print "false woof" -> a is not nil since assigning a pointer value to a overwrites it to what d had which is a type of *Dog with a nil value
//correct answer: will print "false woof" -> a is not nil since assigning a pointer value to a overwrites it to what d had which is a type of *Dog with a nil value and a.sound() does not panic eventhough the value of a is nil because the function Sound never acceses the value of a
// svc.Log("started " + svc.name) ->  will print "[svc] started orders" ->  the prefix is printed before the Log() payload since the payload is treated as a single string and also because runs above wahtever logis is inside the parameter passed
/*correct answer:That isn't the mechanism. What makes this line work is embedding and promotion (§3.3), in three steps:
Service has no Log method of its own. It embeds Logger, with no field name, so Logger's methods are promoted. svc.Log(...) is shorthand for svc.Logger.Log(...).
The argument is evaluated before the call, as with any function: "started " + svc.name becomes "started orders".
Inside Log, the receiver l is the embedded Logger, whose prefix is "[svc] ". Log prints l.prefix + m, which gives [svc] started orders. */
// fmt.Println(st == nil) ->  will print "true" -> this is true because decalring interfaces without assigning a value results in (nil, nil) for both its concrete type and value
