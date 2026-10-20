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

type BasicGameResources struct {
    LogUtil *LogUtil
    RESOURCE string
    RESOURCE_DESTROY string
}

// The Java class BasicGameResources has subclasses so its variables are this interface
type IBasicGameResources interface {
    BasicGameResources_() *BasicGameResources
    Init(ROOT string, SIZE []string)
    Append(ROOT string, sizeString string)
    GetString() string
}

func (this *BasicGameResources) BasicGameResources_() *BasicGameResources {
    return this
}

var BasicGameResources_StringBuffer *StringMaker = NewStringMaker()

func (this *BasicGameResources) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.RESOURCE = StringUtil_GetInstance().EMPTY_STRING
    this.RESOURCE_DESTROY = StringUtil_GetInstance().EMPTY_STRING
}


func NewBasicGameResources() *BasicGameResources {
    this := &BasicGameResources{}
    this.initFields()
    return this
}


func (this *BasicGameResources) Init(ROOT string, SIZE []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
                }
            }()
            var scale int = int(GameConfigurationCentral_GetInstance().SCALE.GetValue()); _ = scale
            if scale > 1 && scale - 1 < len(SIZE) {
                this.Append(ROOT, SIZE[scale - 1])
            } else {
                panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Invalid Value: ").Appendint(scale)))))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *BasicGameResources) Append(ROOT string, sizeString string) {
        var DESTROY string = "_destroy"; _ = DESTROY
        var string_ string = this.GetString(); _ = string_
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(ROOT)
        BasicGameResources_StringBuffer.Append(string_)
        BasicGameResources_StringBuffer.Append(sizeString)
        this.RESOURCE = fmt.Sprint(BasicGameResources_StringBuffer)
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(ROOT)
        BasicGameResources_StringBuffer.Append(DESTROY)
        BasicGameResources_StringBuffer.Append(string_)
        BasicGameResources_StringBuffer.Append(sizeString)
        this.RESOURCE_DESTROY = fmt.Sprint(BasicGameResources_StringBuffer)
}


// throws Exception
func (this *BasicGameResources) GetString() string {
        return GameGraphicsResourceUtil_GetInstance().GetName()
}

