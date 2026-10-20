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

type AbstractInputAutomationFactory struct {
    InputAutomationModuleFactoryInterface
    Name string
    JPanel *JPanel
    HelpSet *HelpSet
}

// The Java class AbstractInputAutomationFactory has subclasses so its variables are this interface
type IAbstractInputAutomationFactory interface {
    IInputAutomationModuleFactoryInterface
    AbstractInputAutomationFactory_() *AbstractInputAutomationFactory
    GetName() string
    GetConfigurationJPanel() *JPanel
    GetHelpSet() *HelpSet
}

func (this *AbstractInputAutomationFactory) AbstractInputAutomationFactory_() *AbstractInputAutomationFactory {
    return this
}


func (this *AbstractInputAutomationFactory) initFields() {
    
    
    
}




// Java overloads of new AbstractInputAutomationFactory
func NewAbstractInputAutomationFactory(args ...any) *AbstractInputAutomationFactory {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[*JPanel](args[1], pass) && Java_IsArg[*HelpSet](args[2], pass) {
            return NewAbstractInputAutomationFactory_1(Java_Arg[string](args[0]), Java_Arg[*JPanel](args[1]), Java_Arg[*HelpSet](args[2]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*JPanel](args[1], pass) {
            return NewAbstractInputAutomationFactory_2(Java_Arg[string](args[0]), Java_Arg[*JPanel](args[1]))
        }
    }
    panic("No Java overload of new AbstractInputAutomationFactory for the arguments")
}
func NewAbstractInputAutomationFactory_1(name string, jPanel *JPanel, helpSet *HelpSet) *AbstractInputAutomationFactory {
    this := &AbstractInputAutomationFactory{}
    this.InputAutomationModuleFactoryInterface = *NewInputAutomationModuleFactoryInterface()
    this.initFields()
        this.Name = name
        this.JPanel = jPanel
        this.HelpSet = helpSet
    return this
}


func NewAbstractInputAutomationFactory_2(name string, jPanel *JPanel) *AbstractInputAutomationFactory {
    this := &AbstractInputAutomationFactory{}
    this.InputAutomationModuleFactoryInterface = *NewInputAutomationModuleFactoryInterface()
    this.initFields()
        this.Name = name
        this.JPanel = jPanel
        var url *URL = reflect.TypeOf(this).GetResource("/help/Help.hs"); _ = url
        this.HelpSet = JavaHelpUtil_GetInstance().GetHelpSet(url)
    return this
}


func (this *AbstractInputAutomationFactory) GetName() string {
        return this.Name
}


func (this *AbstractInputAutomationFactory) GetConfigurationJPanel() *JPanel {
        return this.JPanel
}


func (this *AbstractInputAutomationFactory) GetHelpSet() *HelpSet {
        return this.HelpSet
}

