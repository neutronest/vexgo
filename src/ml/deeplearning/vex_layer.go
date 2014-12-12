/* package vex_layer
2014 12 11
Author: neutronest toosaka

basic implement for layer
*/
package vex_layer

type vex_layer interface {
	forward()

	backward()

	getParamsAndGrads()
}
