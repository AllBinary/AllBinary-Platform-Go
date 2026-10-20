/*
 * 
 * To change this license header, choose License Headers in Project Properties. 
 * To change this template file, choose Tools | Templates 
 * and open the template in the editor.  
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

type TouchButtonResource struct {
    RESOURCE string
    HINT string
}

// The Java class TouchButtonResource has subclasses so its variables are this interface
type ITouchButtonResource interface {
    TouchButtonResource_() *TouchButtonResource
}

func (this *TouchButtonResource) TouchButtonResource_() *TouchButtonResource {
    return this
}


func (this *TouchButtonResource) initFields() {
    
    
}


func NewTouchButtonResource(resource string, hint string) *TouchButtonResource {
    this := &TouchButtonResource{}
    this.initFields()
        this.RESOURCE = resource
        this.HINT = hint
    return this
}

