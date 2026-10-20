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

type TestModuleFactory struct {
    AbstractInputAutomationFactory
    BlankModuleConfigurationJPanel *BlankModuleConfigurationJPanel
}

var TestModuleFactory_NAME string = "Test New Game" + fmt.Sprint(InputAutomationModuleData_MODULE_NAME_END)

func (this *TestModuleFactory) initFields() {
    
}


func NewTestModuleFactory() *TestModuleFactory {
    this := &TestModuleFactory{}
        this.AbstractInputAutomationFactory = *NewAbstractInputAutomationFactory(TestModuleFactory_NAME, NewBlankModuleConfigurationJPanel())
        this.initFields()
    return this
}


func (this *TestModuleFactory) GetName() string {
        return TestModuleFactory_NAME
}


func (this *TestModuleFactory) GetConfigurationJPanel() *JPanel {
        return this.BlankModuleConfigurationJPanel
}


// throws Exception
// synchronized
func (this *TestModuleFactory) GetInstance() RunnableInterface {
        return nil
}

