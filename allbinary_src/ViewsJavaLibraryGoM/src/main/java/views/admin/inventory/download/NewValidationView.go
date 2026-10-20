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

type NewValidationView struct {
    DownloadableInventoryItemView
}

// Implements ValidationComponentInterface



func NewNewValidationView(transformInfoInterface TransformInfoInterface) *NewValidationView {
    this := &NewValidationView{}
        this.DownloadableInventoryItemView = *NewDownloadableInventoryItemView(transformInfoInterface, StringUtil_GetInstance().EMPTY_STRING)
        this.ItemInterface = any(NewBasicItem()).(ItemInterface)
        this.DownloadableItem = NewDownloadableItem(this.ItemInterface.GetId())
    return this
}


func (this *NewValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *NewValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *NewValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *NewValidationView) ValidationInfo() string {
        return "Error Creating New Downloadable Item"
}

