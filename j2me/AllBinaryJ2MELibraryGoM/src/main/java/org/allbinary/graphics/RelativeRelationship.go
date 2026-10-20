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

type RelativeRelationship struct {
    TypesAllowedList IBasicArrayList
    X int
    Y int
    Z int
}

// The Java class RelativeRelationship has subclasses so its variables are this interface
type IRelativeRelationship interface {
    RelativeRelationship_() *RelativeRelationship
    GetX() int
    GetY() int
    GetZ() int
    GetTypesAllowedList() IBasicArrayList
    SetTypesAllowedList(typesAllowedList IBasicArrayList)
}

func (this *RelativeRelationship) RelativeRelationship_() *RelativeRelationship {
    return this
}

var RelativeRelationship_NULL_RELATIVE_RELATIONSHIP IRelativeRelationship = NewRelativeRelationship(PointFactory_GetInstance().ZERO_ZERO, BasicArrayListUtil_GetInstance().GetImmutableInstance())

func (this *RelativeRelationship) initFields() {
    
    
    
    
}


func NewRelativeRelationship(point IGPoint, typesAllowedList IBasicArrayList) *RelativeRelationship {
    this := &RelativeRelationship{}
    this.initFields()
        this.X = point.GetX()
        this.Y = point.GetY()
        this.Z = point.GetZ()
        this.TypesAllowedList = typesAllowedList
    return this
}


func (this *RelativeRelationship) GetX() int {
        return this.X
}


func (this *RelativeRelationship) GetY() int {
        return this.Y
}


func (this *RelativeRelationship) GetZ() int {
        return this.Z
}


func (this *RelativeRelationship) GetTypesAllowedList() IBasicArrayList {
        return this.TypesAllowedList
}


func (this *RelativeRelationship) SetTypesAllowedList(typesAllowedList IBasicArrayList) {
        this.TypesAllowedList = typesAllowedList
}

