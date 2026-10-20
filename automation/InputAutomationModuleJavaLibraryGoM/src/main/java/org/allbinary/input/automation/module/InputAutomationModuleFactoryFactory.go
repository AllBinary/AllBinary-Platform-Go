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

type InputAutomationModuleFactoryFactory struct {
    LogUtil *LogUtil
    HashMap *HashMap
    DefaultListModelHelper *DefaultListModelHelper
    HelpSetListenerInterface *HelpSetListener
}


func (this *InputAutomationModuleFactoryFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewInputAutomationModuleFactoryFactory(inputAutomationModuleConfigurations *InputAutomationModuleConfigurations, helpSetListenerInterface *HelpSetListener) *InputAutomationModuleFactoryFactory {
    this := &InputAutomationModuleFactoryFactory{}
    this.initFields()
        this.HelpSetListenerInterface = helpSetListenerInterface
        this.DefaultListModelHelper = NewDefaultListModelHelper()
        this.HashMap = StdUtil_GetInstance().CreateHashMap()
        var collection Collection = inputAutomationModuleConfigurations.GetHashMap().Values(); _ = collection
        var iterator Iterator = collection.Iterator(); _ = iterator
        var inputAutomationModuleConfiguration *InputAutomationModuleConfiguration; _ = inputAutomationModuleConfiguration
        for iterator.HasNext() {
            inputAutomationModuleConfiguration = any(iterator.Next()).(*InputAutomationModuleConfiguration)
            this.Add(inputAutomationModuleConfiguration)
        }
        this.DefaultListModelHelper.InitDefaultModelList()
        this.LogUtil.PutF("Loaded " + fmt.Sprint(this.HashMap.Size()) + "/" + fmt.Sprint(this.DefaultListModelHelper.GetListModel().GetSize()) + " Input Automation Modules", this, "Contructor")
    return this
}


func (this *InputAutomationModuleFactoryFactory) Add(inputAutomationModuleConfiguration *InputAutomationModuleConfiguration) {
        var inputAutomationModuleInterface IInputAutomationModuleFactoryInterface = inputAutomationModuleConfiguration.GetInputAutomationModuleInterface(); _ = inputAutomationModuleInterface
        this.HashMap.Put(inputAutomationModuleConfiguration.GetName(), inputAutomationModuleInterface)
        this.DefaultListModelHelper.Add(inputAutomationModuleConfiguration.GetName())
        var helpSet *HelpSet = inputAutomationModuleInterface.GetHelpSet(); _ = helpSet
        if !JavaHelpSetNotifier_IsNotified(helpSet) {
            var helpSetEvent *HelpSetEvent = NewHelpSetEvent(this, helpSet, HelpSetEvent_HELPSET_ADDED); _ = helpSetEvent
            this.HelpSetListenerInterface.HelpSetAdded(helpSetEvent)
        }
}


func (this *InputAutomationModuleFactoryFactory) GetListModel() *ListModel {
        return any(this.DefaultListModelHelper.GetListModel()).(*ListModel)
}


func (this *InputAutomationModuleFactoryFactory) GetInstance(moduleName string) IInputAutomationModuleFactoryInterface {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF("Getting Module: " + moduleName, this, commonStrings.GET_INSTANCE)
        var inputAutomationModuleFactoryInterface IInputAutomationModuleFactoryInterface = any(this.HashMap.Get(moduleName)).(IInputAutomationModuleFactoryInterface); _ = inputAutomationModuleFactoryInterface
        if inputAutomationModuleFactoryInterface == nil {
            this.LogUtil.PutF("Module: " + moduleName + " was null", this, commonStrings.GET_INSTANCE)
        }
        return inputAutomationModuleFactoryInterface
}

