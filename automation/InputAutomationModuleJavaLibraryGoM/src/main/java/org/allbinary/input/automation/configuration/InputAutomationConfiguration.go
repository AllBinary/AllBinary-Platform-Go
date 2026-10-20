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

//name="inputAutomationConfiguration", 
type InputAutomationConfiguration struct {
    LogUtil *LogUtil
    Installed bool
    InputAutomationModuleConfigurationList IList[*InputAutomationModuleConfiguration]
}

var InputAutomationConfiguration_FILENAME string = "InputAutomationConfig.xml"

func (this *InputAutomationConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewInputAutomationConfiguration() *InputAutomationConfiguration {
    this := &InputAutomationConfiguration{}
    this.initFields()
    return this
}


func InputAutomationConfiguration_GetFile() *File {
        return NewFile(InputAutomationConfiguration_FILENAME)
}


func (this *InputAutomationConfiguration) IsInstalled() bool {
        return this.Installed
}


func (this *InputAutomationConfiguration) SetInstalled(installed bool) {
        this.Installed = installed
        this.LogUtil.PutF("Installed: " + fmt.Sprint(installed), this, "setIntalled")
}


func (this *InputAutomationConfiguration) GetInputAutomationModuleConfigurationList() IList[*InputAutomationModuleConfiguration] {
        return this.InputAutomationModuleConfigurationList
}


// throws Exception
func (this *InputAutomationConfiguration) Save() {
        var document *Document = DomDocumentHelper_Create(); _ = document
        var file *File = NewFile(InputAutomationConfiguration_FILENAME); _ = file
        FileUtil_GetInstance().Copy(NewAbPath(file.GetAbsolutePath(), StringUtil_GetInstance().EMPTY_STRING), NewAbPath(fmt.Sprint(file.GetAbsolutePath()) + ".bak.xml", StringUtil_GetInstance().EMPTY_STRING))
        var jaxbContext *JAXBContext = JAXBContext_NewInstance(reflect.TypeOf((*InputAutomationConfiguration)(nil)).Elem()); _ = jaxbContext
        var marshaller *Marshaller = jaxbContext.CreateMarshaller(); _ = marshaller
        marshaller.SetProperty(Marshaller_JAXB_FORMATTED_OUTPUT, true)
        marshaller.Marshal(this, document)
        DomDocumentFileHelper_Save(FileWrapperUtil_WrapFile(file), document)
}

