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

type DestroyableBaseBehavior struct {
}

// Implements DestroyableInterface

// The Java class DestroyableBaseBehavior has subclasses so its variables are this interface
type IDestroyableBaseBehavior interface {
    DestroyableBaseBehavior_() *DestroyableBaseBehavior
    GetName() string
    IsDestroyed() bool
    SetDestroyed(destroyed bool)
}

func (this *DestroyableBaseBehavior) DestroyableBaseBehavior_() *DestroyableBaseBehavior {
    return this
}

var DestroyableBaseBehavior_Instance IDestroyableBaseBehavior = NewDestroyableBaseBehavior()

func NewDestroyableBaseBehavior() *DestroyableBaseBehavior {
    this := &DestroyableBaseBehavior{}
    return this
}


func DestroyableBaseBehavior_GetInstance() IDestroyableBaseBehavior {
        return DestroyableBaseBehavior_Instance
}


func (this *DestroyableBaseBehavior) GetName() string {
        return reflect.TypeOf(this).String()
}


// throws Exception
func (this *DestroyableBaseBehavior) IsDestroyed() bool {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}


func (this *DestroyableBaseBehavior) SetDestroyed(destroyed bool) {
}

