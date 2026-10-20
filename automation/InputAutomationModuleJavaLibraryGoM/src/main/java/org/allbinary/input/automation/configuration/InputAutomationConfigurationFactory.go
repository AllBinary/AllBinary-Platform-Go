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

type InputAutomationConfigurationFactory struct {
    LogUtil *LogUtil
    InputAutomationConfiguration *InputAutomationConfiguration
}

var InputAutomationConfigurationFactory_Instance *InputAutomationConfigurationFactory = NewInputAutomationConfigurationFactory()

func (this *InputAutomationConfigurationFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.InputAutomationConfiguration = nil
}


func InputAutomationConfigurationFactory_GetInstance() *InputAutomationConfigurationFactory {
        return InputAutomationConfigurationFactory_Instance
}


func NewInputAutomationConfigurationFactory() *InputAutomationConfigurationFactory {
    this := &InputAutomationConfigurationFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *InputAutomationConfigurationFactory) Init(abeClientInformation AbeClientInformationInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var INPUT_AUTOMATION_CONFIGURATION string = "InputAutomationConfiguration"; _ = INPUT_AUTOMATION_CONFIGURATION
        var file *File = InputAutomationConfiguration_GetFile(); _ = file
        if file.IsFile() {
            logUtil.PutF("LoadingConfiguration", INPUT_AUTOMATION_CONFIGURATION, commonStrings.INIT)
            var jaxbContext *JAXBContext = JAXBContext_NewInstance(reflect.TypeOf((*InputAutomationConfiguration)(nil)).Elem()); _ = jaxbContext
            var unmarshaller *Unmarshaller = jaxbContext.CreateUnmarshaller(); _ = unmarshaller
            var root *JAXBElement[*InputAutomationConfiguration] = unmarshaller.Unmarshal(NewStreamSource(NewFileInputStream(file)), reflect.TypeOf((*InputAutomationConfiguration)(nil)).Elem()); _ = root
            this.InputAutomationConfiguration = any(root.GetValue()).(*InputAutomationConfiguration)
            var inputAutomationModuleConfigurationList IList[*InputAutomationModuleConfiguration] = this.InputAutomationConfiguration.GetInputAutomationModuleConfigurationList(); _ = inputAutomationModuleConfigurationList
            logUtil.PutF("isInstalled: " + fmt.Sprint(this.InputAutomationConfiguration.IsInstalled()) + " inputAutomationModuleConfigurationList: " + fmt.Sprint(inputAutomationModuleConfigurationList), INPUT_AUTOMATION_CONFIGURATION, commonStrings.INIT)
            if inputAutomationModuleConfigurationList != nil {
                var size int = inputAutomationModuleConfigurationList.Size(); _ = size
                var inputAutomationModuleConfiguration *InputAutomationModuleConfiguration; _ = inputAutomationModuleConfiguration
                for index := int(0); index < size; index++ {
                    inputAutomationModuleConfiguration = inputAutomationModuleConfigurationList.Get(index)
                    inputAutomationModuleConfiguration.Init(abeClientInformation)
                }
            }
            logUtil.PutF("LoadedConfiguration", INPUT_AUTOMATION_CONFIGURATION, commonStrings.INIT)
        } else {
            logUtil.PutF("New Configuration", INPUT_AUTOMATION_CONFIGURATION, commonStrings.INIT)
            this.InputAutomationConfiguration = NewInputAutomationConfiguration()
        }
}

