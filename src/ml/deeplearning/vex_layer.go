/* package vex_layer
2014 12 11
Author: neutronest toosaka

basic implement for layer
*/
package vex_layer

import (
	"../base"
)

type VexLayer struct {

	// the type of layer: SVM, Sigmoid, Relu ...
	layerType

	//
	numInputs int

	outSx int

	outSy int

	intSx int

	intSy int

	LayerPrototype
}

type LayerPrototype interface {
	forward(data DataFrame, isTraining bool) (dataResult DataFrame)

	backward(y DataType)

	GetParamsAndGrad() error
}
