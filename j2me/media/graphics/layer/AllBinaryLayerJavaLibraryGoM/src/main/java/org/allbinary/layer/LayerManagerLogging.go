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

type LayerManagerLogging struct {
    LayerManagerLoggingBase
    LogUtil *LogUtil
    StringBuilder *StringMaker
    APPEND string
    APPEND_ string
    _AT_ string
    ATTEMPT_REMOVE_ string
    REMOVE_ string
    REMOVE string
    DID_NOT_REMOVE string
    CLEAR string
}

var LayerManagerLogging_Instance *LayerManagerLogging = NewLayerManagerLogging()
var LayerManagerLogging_RemoveFailed bool = false

func (this *LayerManagerLogging) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StringBuilder = NewStringMaker()
    this.APPEND = "append"
    this.APPEND_ = " append: "
    this._AT_ = " at: "
    this.ATTEMPT_REMOVE_ = " Remove Attempt: "
    this.REMOVE_ = " Remove: "
    this.REMOVE = "remove"
    this.DID_NOT_REMOVE = " Did not remove: "
    this.CLEAR = " Clear List"
}

func NewLayerManagerLogging() *LayerManagerLogging {
    this := &LayerManagerLogging{}
    this.LayerManagerLoggingBase = *NewLayerManagerLoggingBase()
    this.initFields()
    return this
}


func LayerManagerLogging_GetInstance() *LayerManagerLogging {
        return LayerManagerLogging_Instance
}


// throws Exception
func (this *LayerManagerLogging) Append(layerInterface IAllBinaryLayer) {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.APPEND_).Append(layerInterface.GetName())), this, this.APPEND)
}


func (this *LayerManagerLogging) AppendAt(layerInterface IAllBinaryLayer, index int) {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.APPEND_).Append(layerInterface.GetName()).Append(this._AT_).Appendint(index)), this, this.APPEND)
}


func (this *LayerManagerLogging) Remove(layerInterface IAllBinaryLayer) {
        if layerInterface == nil {
            this.StringBuilder.Delete(0, this.StringBuilder.Length())
            this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.ATTEMPT_REMOVE_).Append(StringUtil_GetInstance().ToString(layerInterface))), this, this.REMOVE)
        } else {
            this.StringBuilder.Delete(0, this.StringBuilder.Length())
            this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.ATTEMPT_REMOVE_).Append(layerInterface.GetName())), this, this.REMOVE)
        }
}


func (this *LayerManagerLogging) RemoveResult(layerManager ILayerManager, layerInterface IAllBinaryLayer, result bool) {
        if layerInterface == nil {
            this.StringBuilder.Delete(0, this.StringBuilder.Length())
            this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.REMOVE_).Append(StringUtil_GetInstance().ToString(layerInterface))), this, this.REMOVE)
        } else if result {
            if LayerManagerLogging_RemoveFailed {
                this.StringBuilder.Delete(0, this.StringBuilder.Length())
                this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.REMOVE_).Append(layerInterface.GetName())), this, this.REMOVE)
            }
        } else {
            this.StringBuilder.Delete(0, this.StringBuilder.Length())
            this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.DID_NOT_REMOVE).Append(layerInterface.GetName())), this, this.REMOVE)
            LayerManagerLogging_RemoveFailed = true
        }
}


func (this *LayerManagerLogging) Log(layerManager ILayerManager) {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        var size int = layerManager.GetSize(); _ = size
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        this.StringBuilder.Append(strconv.FormatInt(int64(TsUtil_GetInstance().HashCode(layerManager)), 16)).Append(commonSeps.COLON_SEP)
        this.StringBuilder.Appendint(size).Append(commonSeps.COLON_SEP)
        var allBinaryLayer IAllBinaryLayer; _ = allBinaryLayer
        for index := int(0); index < size; index++ {
            allBinaryLayer = (any(layerManager.GetLayerAt(index)).(IAllBinaryLayer))
            this.StringBuilder.Append(allBinaryLayer.GetName()).Append(commonSeps.COMMA)
        }
        this.LogUtil.PutF(fmt.Sprint(this.StringBuilder), this, this.REMOVE)
}


func (this *LayerManagerLogging) Clear() {
        this.StringBuilder.Delete(0, this.StringBuilder.Length())
        this.LogUtil.PutF(fmt.Sprint(this.StringBuilder.Appendint(TsUtil_GetInstance().HashCode(this)).Append(this.CLEAR)), this, this.CLEAR)
}

