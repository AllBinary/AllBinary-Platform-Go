/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type AdvertisementAreaConstraints struct {
    DisplayInBrowserContraintsInterface DisplayInBrowserContraintsInterface
    TwoDimensionConstraintsInterface TwoDimensionalConstraintInterface
}

// Implements AdvertisementAreaConstraintsInterface


func (this *AdvertisementAreaConstraints) initFields() {
    
    
}


func NewAdvertisementAreaConstraints(document *Document) *AdvertisementAreaConstraints {
    this := &AdvertisementAreaConstraints{}
    this.initFields()
        var node *Node = DomSearchHelper_GetNode(AdvertisementAreaData_GetInstance().NAME, document.GetChildNodes()); _ = node
        this.DisplayInBrowserContraintsInterface = NewDisplayInBrowserContraints(node)
        this.TwoDimensionConstraintsInterface = NewTwoDimensionalConstraint(node)
    return this
}


func (this *AdvertisementAreaConstraints) SetDisplayInBrowserContraints(displayInBrowserContraintsInterface DisplayInBrowserContraintsInterface) {
        this.DisplayInBrowserContraintsInterface = displayInBrowserContraintsInterface
}


func (this *AdvertisementAreaConstraints) GetTwoDimensionContraints_1(twoDimensionConstraintsInterface TwoDimensionalConstraintInterface) {
        this.TwoDimensionConstraintsInterface = twoDimensionConstraintsInterface
}

// Java overloads of AdvertisementAreaConstraints.getTwoDimensionContraints
func (this *AdvertisementAreaConstraints) GetTwoDimensionContraints(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TwoDimensionalConstraintInterface](args[0], pass) {
            return this.GetTwoDimensionContraints_1(Java_Arg[TwoDimensionalConstraintInterface](args[0]))
        }
        if len(args) == 0 {
            return this.GetTwoDimensionContraints_2()
        }
    }
    panic("No Java overload of AdvertisementAreaConstraints.getTwoDimensionContraints for the arguments")
}


func (this *AdvertisementAreaConstraints) GetDisplayInBrowserContraints() DisplayInBrowserContraintsInterface {
        return this.DisplayInBrowserContraintsInterface
}


func (this *AdvertisementAreaConstraints) GetTwoDimensionContraints_2() TwoDimensionalConstraintInterface {
        return this.TwoDimensionConstraintsInterface
}

