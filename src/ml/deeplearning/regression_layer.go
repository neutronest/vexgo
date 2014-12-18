/* package vex_layer
2014 12 11
Author: neutronest toosaka

basic implement for layer
*/
package vex_layer

import (
	"../base"
)

/*
type VexLayer struct {

	// the type of layer: SVM, Sigmoid, Relu ...
	layerType

	//
	numInputs int

	outSx int

	outSy int

	inSx int

	inSy int
}

*/

/*-----------    RegressionLayer        -----------*/

type RegressionLayer struct {
	layerType string

	inAction DataVol

	outAction DataVol

	numInputs int

	outSx int

	outSy int

	outDepth int

	inSx int

	inSy int

	intDepth int
}

func (RL *RegressionLayer) Init(layer) error {

	if opt == nil {
		return nil
	}

	RL.numInputs = layer.inSx * layer.inSy * layer.inDepth
	RL.outDepth = layer.numInputs
	RL.outSx = 1
	RL.outSy = 1
	RL.layerType = "regression"
}

func (RL RegressionLayer) forward(dataVol DataVol, isTraining bool) (dataResult DataVol) {

	RL.inAction = dataVol
	RL.outAction = dataVol
	return dataVol
}

func (RL RegressonLayer) backward(y []float64) (loss float64) {

	var x = RL.inAction
	var dwArray [len(x.w)]float64
	x.dw = dwArray

	var loss = 0.0
	if len(y) != 1 {

		for i := 0; i < RL.outDepth; i++ {

			var dy = x.w[i] - y[i]
			x.dw[i] = dy
			lost += 0.5 * dy * dy
		}
	} else {

		var dy = x.w[0] - y[0]
		x.dw[0] = dy
		lost += 0.5 * dy * dy
	}

	return loss
}

/*---------------    End RegressionLayer   -----------------*/
