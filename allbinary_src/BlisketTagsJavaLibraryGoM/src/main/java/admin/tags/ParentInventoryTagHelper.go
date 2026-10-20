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

type ParentInventoryTagHelper struct {
}

var ParentInventoryTagHelper_Instance *ParentInventoryTagHelper = NewParentInventoryTagHelper()

func NewParentInventoryTagHelper() *ParentInventoryTagHelper {
    this := &ParentInventoryTagHelper{}
    return this
}


func ParentInventoryTagHelper_GetInstance() *ParentInventoryTagHelper {
        return ParentInventoryTagHelper_Instance
}


// throws Exception
func (this *ParentInventoryTagHelper) IsValid(tagClass any, parentTag *Tag) {
        if parentTag == nil {
            panic(errors.New(fmt.Sprint("Must have parent tag.")))
        } else if !(func() bool { _, ok := any(parentTag).(IHelperTag); return ok }()) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Must have at least a ")
            stringBuffer.Append("tags.HelperTag")
            stringBuffer.Append(" as parent")
            stringBuffer.Append("and not: ")
            stringBuffer.Append(reflect.TypeOf(parentTag).String())
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        }
}

