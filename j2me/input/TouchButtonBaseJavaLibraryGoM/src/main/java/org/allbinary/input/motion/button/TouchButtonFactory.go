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

type TouchButtonFactory struct {
    LIST IBasicArrayList
    List IBasicArrayList
    SavedList IBasicArrayList
    Saved bool
}

var TouchButtonFactory_SINGLETON *TouchButtonFactory = NewTouchButtonFactory()

func (this *TouchButtonFactory) initFields() {
    this.LIST = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    this.List = this.LIST
    this.SavedList = this.LIST
    this.Saved = false
}


func TouchButtonFactory_GetInstance() *TouchButtonFactory {
        return TouchButtonFactory_SINGLETON
}


func NewTouchButtonFactory() *TouchButtonFactory {
    this := &TouchButtonFactory{}
    this.initFields()
    return this
}


func (this *TouchButtonFactory) DefaultList() {
        this.List = this.LIST
        this.Toggle(false, this.LIST)
}


// synchronized
func (this *TouchButtonFactory) Toggle(save bool, list IBasicArrayList) {
        if list != BasicArrayListUtil_GetInstance().GetImmutableInstance() {
            this.SavedList = list
            if !this.Saved {
                this.List = this.SavedList
            }
        } else {
            if save {
                if !this.Saved {
                    this.Saved = true
                    this.SavedList = this.List
                    this.List = this.LIST
                } else {
                }
            } else {
                this.List = this.SavedList
                this.Saved = false
            }
        }
}


// throws Exception
func (this *TouchButtonFactory) SetList(list IBasicArrayList) {
        if list == nil {
            panic(errors.New(fmt.Sprint("Null List")))
        }
        this.List = list
}


func (this *TouchButtonFactory) GetList() IBasicArrayList {
        return this.List
}

