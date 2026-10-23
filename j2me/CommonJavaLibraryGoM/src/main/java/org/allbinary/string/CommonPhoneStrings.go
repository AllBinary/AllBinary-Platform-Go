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

type CommonPhoneStrings struct {
    UP string
    DOWN string
    LEFT string
    RIGHT string
    ZERO string
    ONE string
    TWO string
    THREE string
    FOUR string
    FIVE string
    SIX string
    SEVEN string
    EIGHT string
    NINE string
    POUND string
    STAR string
    FIRE string
}

var CommonPhoneStrings_Instance *CommonPhoneStrings = NewCommonPhoneStrings()

func (this *CommonPhoneStrings) initFields() {
    this.UP = "UP"
    this.DOWN = "DOWN"
    this.LEFT = "LEFT"
    this.RIGHT = "RIGHT"
    this.ZERO = "0"
    this.ONE = "1"
    this.TWO = "2"
    this.THREE = "3"
    this.FOUR = "4"
    this.FIVE = "5"
    this.SIX = "6"
    this.SEVEN = "7"
    this.EIGHT = "8"
    this.NINE = "9"
    this.POUND = "#"
    this.STAR = "*"
    this.FIRE = "FIRE"
}

func NewCommonPhoneStrings() *CommonPhoneStrings {
    this := &CommonPhoneStrings{}
    this.initFields()
    return this
}


func CommonPhoneStrings_GetInstance() *CommonPhoneStrings {
        return CommonPhoneStrings_Instance
}

