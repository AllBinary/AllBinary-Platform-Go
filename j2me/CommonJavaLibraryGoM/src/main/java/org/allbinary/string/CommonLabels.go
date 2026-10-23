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

type CommonLabels struct {
    WIDTH_LABEL string
    HEIGHT_LABEL string
    END_LABEL string
    TOTAL_LABEL string
    INDEX_LABEL string
    START_LABEL string
    COMMAND_LABEL string
    NAME_LABEL string
    START string
    ELAPSED string
    CURRENT string
    NEXT string
    ITEM_LABEL string
    RESULT_ string
    COLON_SEP string
}

var CommonLabels_Instance *CommonLabels = NewCommonLabels()

func (this *CommonLabels) initFields() {
    this.WIDTH_LABEL = " Width: "
    this.HEIGHT_LABEL = " Height: "
    this.END_LABEL = "End: "
    this.TOTAL_LABEL = "Total: "
    this.INDEX_LABEL = "index: "
    this.START_LABEL = "Start: "
    this.COMMAND_LABEL = "Command: "
    this.NAME_LABEL = "Name: "
    this.START = this.START_LABEL
    this.ELAPSED = " Elapsed: "
    this.CURRENT = " Current: "
    this.NEXT = "Next"
    this.ITEM_LABEL = "Item: "
    this.RESULT_ = "Result: "
    this.COLON_SEP = ": "
}

func NewCommonLabels() *CommonLabels {
    this := &CommonLabels{}
    this.initFields()
    return this
}


func CommonLabels_GetInstance() *CommonLabels {
        return CommonLabels_Instance
}

