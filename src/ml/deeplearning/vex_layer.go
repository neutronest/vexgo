/* package vex_layer
2014 12 11
Author: neutronest toosaka

basic implement for layer
*/
package vexlayer

type vexlayer interface {
	forward()

	backward()

	getParamsAndGrads()
}
