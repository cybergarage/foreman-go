// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "strings"
import "fmt"

// A Quality represents a quality formula in a QoS clause.
type Quality struct {
	Variable  Variable
	Operator  Operator
	Objective Objective
}

// NewQuality returns a new quality.
func NewQuality() *Quality {
	quality := &Quality{
		Variable: nil,
	}
	return quality
}

// ParseString parses a specified quality string.
func (quality *Quality) ParseString(factory Factory, qualityString string) error {
	qualityStrings := strings.Split(qualityString, SpaceSeparator)

	if len(qualityStrings) != 3 {
		return fmt.Errorf(errorInvalidQualityString, qualityString)
	}

	// Variable
	variableString := qualityStrings[0]
	variable, err := factory.CreateVariable(variableString)
	if err != nil {
		return err
	}
	quality.Variable = variable

	// Operator
	operatorString := qualityStrings[1]
	operator, err := factory.CreateOperator(operatorString)
	if err != nil {
		return err
	}
	quality.Operator = operator

	// Objective
	objectiveString := qualityStrings[2]
	objective, err := factory.CreateObjective(objectiveString)
	if err != nil {
		return err
	}
	quality.Objective = objective

	return nil
}
