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

type DesktopBundle struct {
    Bundle Bundle
}

var DesktopBundle_DESKTOP_SYMBOLIC_NAME string = "org.knopflerfish.bundle.desktop"

func (this *DesktopBundle) initFields() {
    
}


func NewDesktopBundle() *DesktopBundle {
    this := &DesktopBundle{}
    this.initFields()
        var bundleContext BundleContext = InputAutomationBundleActivator_GetBundleContext(); _ = bundleContext
        if bundleContext != nil {
            var bundleArray []Bundle = bundleContext.GetBundles(); _ = bundleArray
            for index := int(0); index < len(bundleArray); index++ {
                var bundle Bundle = bundleArray[index]; _ = bundle
                var symbolicName string = bundle.GetSymbolicName(); _ = symbolicName
                if strings.Compare(symbolicName, DesktopBundle_DESKTOP_SYMBOLIC_NAME) == 0 {
                    this.Bundle = bundle
                }
            }
        } else {
            var textJDialog *JDialog = NewBasicTextJDialog("BundlContext is Null"); _ = textJDialog
            textJDialog.SetVisible(true)
        }
    return this
}


// throws Exception
func (this *DesktopBundle) Start() {
        this.Bundle.Start()
}

