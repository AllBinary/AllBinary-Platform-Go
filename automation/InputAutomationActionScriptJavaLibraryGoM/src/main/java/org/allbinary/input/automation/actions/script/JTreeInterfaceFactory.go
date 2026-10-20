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

type JTreeInterfaceFactory struct {
    JTreeInterface JTreeInterface
}

var JTreeInterfaceFactory_Instance *JTreeInterfaceFactory = NewJTreeInterfaceFactory()

func (this *JTreeInterfaceFactory) initFields() {
    
}

func NewJTreeInterfaceFactory() *JTreeInterfaceFactory {
    this := &JTreeInterfaceFactory{}
    this.initFields()
    return this
}


func JTreeInterfaceFactory_GetInstance() *JTreeInterfaceFactory {
        return JTreeInterfaceFactory_Instance
}


func (this *JTreeInterfaceFactory) Set(jTreeInterface JTreeInterface) {
        this.JTreeInterface = jTreeInterface
}


func (this *JTreeInterfaceFactory) GetJTreeInterface() JTreeInterface {
        return this.JTreeInterface
}

