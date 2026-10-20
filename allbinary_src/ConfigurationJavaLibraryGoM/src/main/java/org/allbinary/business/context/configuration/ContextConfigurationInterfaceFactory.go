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

type ContextConfigurationInterfaceFactory struct {
}

var ContextConfigurationInterfaceFactory_Instance *ContextConfigurationInterfaceFactory = NewContextConfigurationInterfaceFactory()


func ContextConfigurationInterfaceFactory_GetInstance() *ContextConfigurationInterfaceFactory {
        return ContextConfigurationInterfaceFactory_Instance
}


func NewContextConfigurationInterfaceFactory() *ContextConfigurationInterfaceFactory {
    this := &ContextConfigurationInterfaceFactory{}
    return this
}


// throws Exception
func (this *ContextConfigurationInterfaceFactory) GetInstance_1(document *Document) ContextConfigurationInterface {
        var contextConfigurationData *ContextConfigurationData = ContextConfigurationData_GetInstance(); _ = contextConfigurationData
        var contextConfigurationNode *Node = DomSearchHelper_GetNode(contextConfigurationData.NAME, document.GetChildNodes()); _ = contextConfigurationNode
        return any(NewContextConfigurationView(contextConfigurationNode).GetContextConfigurationInterface()).(ContextConfigurationInterface)
}

// Java overloads of ContextConfigurationInterfaceFactory.getInstance
func (this *ContextConfigurationInterfaceFactory) GetInstance(args ...any) ContextConfigurationInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return this.GetInstance_1(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetInstance_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of ContextConfigurationInterfaceFactory.getInstance for the arguments")
}


// throws Exception
func (this *ContextConfigurationInterfaceFactory) GetInstance_2(contextName string) ContextConfigurationInterface {
        var abPath IAbPath = ContextConfigurationPathUtil_GetAbPath(contextName); _ = abPath
        var contextConfigurationData *ContextConfigurationData = ContextConfigurationData_GetInstance(); _ = contextConfigurationData
        var documentString string = NewCryptFileReader(contextConfigurationData.UNCRYPTED_EXTENSION, contextConfigurationData.ENCRYPTED_EXTENSION).Get(abPath); _ = documentString
        var document *Document = DomDocumentHelper_Create(documentString); _ = document
        return this.GetInstance(document)
}

