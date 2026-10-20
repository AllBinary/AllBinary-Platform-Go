/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type AdvancedRTSLayerFormFactory struct {
    RTSLayerFormFactory
    LogUtil *LogUtil
    ScrollSelectionForm IScrollSelectionForm
}


func (this *AdvancedRTSLayerFormFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewAdvancedRTSLayerFormFactory(label string, itemArrayFactoryInterface ItemArraySingletonFactoryInterface) *AdvancedRTSLayerFormFactory {
    this := &AdvancedRTSLayerFormFactory{}
    this.RTSLayerFormFactory = *NewRTSLayerFormFactory()
    this.initFields()
        var scrollSelectionForm IScrollSelectionForm = ScrollSelectionForm_GetNullScrollSelectionForm(); _ = scrollSelectionForm
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            var items []IABCustomItem = itemArrayFactoryInterface.GetItems(); _ = items
            scrollSelectionForm = RTSScrollSelectionFormFactory_GetInstance(label, items)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.ScrollSelectionForm = scrollSelectionForm
    return this
}


func (this *AdvancedRTSLayerFormFactory) GetInstance(rtsLayer IRTSLayer) IScrollSelectionForm {
        return this.ScrollSelectionForm
}

