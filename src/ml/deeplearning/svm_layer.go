/* package svm_layer
Author: neutronest


*/
package svm_layer

import (
	"../base"
)

type SVMLayer struct {
	outDepth int

	outSx int

	outSy int

	layerType string
}

func (svmLayer *SVMLayer) init() {

	return nil
}

func (svmLayer *SVMLayer) Forward(data DataFrame, isTraining bool) (dataResult DataFrame) {

}

func (svmLayer *SVMLayer) Backward(y dataType) {
	return nil
}

func (svmLayer *SVMLayer) GetParamsAndGrads() {

	return nil
}
