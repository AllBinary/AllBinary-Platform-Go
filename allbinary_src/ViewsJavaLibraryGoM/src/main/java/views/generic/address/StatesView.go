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

type StatesView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements DomNodeInterface


func (this *StatesView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewStatesView() *StatesView {
    this := &StatesView{}
    this.initFields()
    return this
}


func (this *StatesView) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "toXmlNode", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var node *Node = document.CreateElement(StreetAddressData_STATES); _ = node
            var statesFile string = fmt.Sprint(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().LINEDATAPATH) + "states.txt"; _ = statesFile
            var lineReader *LineReader = NewLineReader(statesFile); _ = lineReader
            for lineReader.HasNext() {
                var option string = lineReader.Next(); _ = option
                node.AppendChild(ModDomHelper_CreateNameValueNodes(document, StreetAddressData_STATE, option))
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

