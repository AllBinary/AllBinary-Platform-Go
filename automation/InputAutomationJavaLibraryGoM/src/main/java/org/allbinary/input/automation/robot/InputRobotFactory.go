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

type InputRobotFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Hashtable *ABHashtable[any, any]
    HelpSetListenerInterface *HelpSetListener
}

var InputRobotFactory_InputRobotFactory *InputRobotFactory = NewInputRobotFactory()

func (this *InputRobotFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    
}


// throws Exception
func InputRobotFactory_GetInstance() *InputRobotFactory {
        return InputRobotFactory_InputRobotFactory
}


func NewInputRobotFactory() *InputRobotFactory {
    this := &InputRobotFactory{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR)
                }
            }()
            var graphenv *GraphicsEnvironment = GraphicsEnvironment_GetLocalGraphicsEnvironment(); _ = graphenv
            var screens []*GraphicsDevice = graphenv.GetScreenDevices(); _ = screens
            var inputRobotInterface InputRobotInterface; _ = inputRobotInterface
            for i := int(0); i < len(screens); i++ {
                inputRobotInterface = any(NewInputRobot(screens[i])).(InputRobotInterface)
                this.LogUtil.PutF("Adding Robot: " + inputRobotInterface.GetName(), this, "getRobots")
                this.Get().(*ABHashtable).Put(inputRobotInterface.GetName(), inputRobotInterface)
            }
            this.LogUtil.PutF("Number Of Robots: " + fmt.Sprint(this.Hashtable.Size()), this, "getRobots")
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *InputRobotFactory) AddListener(helpSetListenerInterface *HelpSetListener) {
        this.HelpSetListenerInterface = helpSetListenerInterface
}


// throws Exception
func (this *InputRobotFactory) Add(inputRobotInterface InputRobotInterface) {
        this.LogUtil.PutF("Adding InputRobotInterface: " + inputRobotInterface.GetName(), this, "add")
        this.Get().(*ABHashtable).Put(inputRobotInterface.GetName(), inputRobotInterface)
        var helpSet *HelpSet = inputRobotInterface.GetHelpSet(); _ = helpSet
        if this.HelpSetListenerInterface != nil {
            if helpSet != nil {
                if !JavaHelpSetNotifier_IsNotified(helpSet) {
                    var helpSetEvent *HelpSetEvent = NewHelpSetEvent(this, helpSet, HelpSetEvent_HELPSET_ADDED); _ = helpSetEvent
                    this.HelpSetListenerInterface.HelpSetAdded(helpSetEvent)
                }
            } else {
                this.LogUtil.PutF("Null HelpSet For: " + inputRobotInterface.GetName(), this, "add")
            }
        } else {
            this.LogUtil.PutF("No HelpSet Listener", this, "add")
        }
}


// throws Exception
func (this *InputRobotFactory) LoadLibraries() {
        var set Set = this.Get().(*ABHashtable).KeySet(); _ = set
        this.LogUtil.PutF("Loading Libraries", this, "loadLibraries")
        var nameArray []any = set.ToArray(); _ = nameArray
        var size int = len(nameArray); _ = size
        for index := int(0); index < size; index++ {
            var inputRobotInterface InputRobotInterface = InputRobotFactory_GetInstance().Get(any(nameArray[index]).(string)).(InputRobotInterface); _ = inputRobotInterface
            InputRobotFactory_LoadLibrary(inputRobotInterface)
        }
}


// throws Exception
func InputRobotFactory_LoadLibraries(collection Collection) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        logUtil.PutF("Loading Libraries", "InputRobotFactory", "loadLibraries")
        var iterator Iterator = collection.Iterator(); _ = iterator
        for iterator.HasNext() {
            InputRobotFactory_LoadLibrary(any(iterator.Next()).(InputRobotInterface))
        }
}


// throws Exception
func InputRobotFactory_LoadLibrary(inputRobotInterface InputRobotInterface) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if InterfaceUtil_IsImplemented(reflect.TypeOf((*SecuredNativeLibraryInterface)(nil)).Elem(), inputRobotInterface) {
            logUtil.PutF("Loading Library: " + inputRobotInterface.GetName(), "InputRobotFactory", "loadLibraries")
            var securedNativeLibraryInterface SecuredNativeLibraryInterface = any(inputRobotInterface).(SecuredNativeLibraryInterface); _ = securedNativeLibraryInterface
            securedNativeLibraryInterface.Load()
        }
}


// throws Exception
func (this *InputRobotFactory) UnloadLibraries() {
        this.LogUtil.PutF("Unloading Libraries", this, "unloadLibraries")
        var set Set = this.Get().(*ABHashtable).KeySet(); _ = set
        var inputRobotInterface InputRobotInterface; _ = inputRobotInterface
        var inputRobotArray []any = set.ToArray(); _ = inputRobotArray
        var size int = len(inputRobotArray); _ = size
        for index := int(0); index < size; index++ {
            inputRobotInterface = this.Get(any(inputRobotArray[index]).(string)).(InputRobotInterface)
            if InterfaceUtil_IsImplemented(reflect.TypeOf((*SecuredNativeLibraryInterface)(nil)).Elem(), inputRobotInterface) {
                this.LogUtil.PutF("Unloading Library: " + inputRobotInterface.GetName(), this, "unloadLibraries")
                var securedNativeLibraryInterface SecuredNativeLibraryInterface = any(inputRobotInterface).(SecuredNativeLibraryInterface); _ = securedNativeLibraryInterface
                securedNativeLibraryInterface.Unload()
            }
        }
}


// throws Exception
func (this *InputRobotFactory) Get_1() *ABHashtable[any, any] {
        return this.Hashtable
}

// Java overloads of InputRobotFactory.get
func (this *InputRobotFactory) Get(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.Get_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.Get_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of InputRobotFactory.get for the arguments")
}


func (this *InputRobotFactory) Get_2(name string) InputRobotInterface {
        this.LogUtil.PutF("Getting Robot: " + name, this, "getRobots")
        return any(this.Hashtable.Get(name)).(InputRobotInterface)
}

