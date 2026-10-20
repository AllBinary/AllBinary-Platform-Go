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

type SpecialDemoGameMidlet struct {
    DemoGameMidlet
}

// The Java class SpecialDemoGameMidlet has subclasses so its variables are this interface
type ISpecialDemoGameMidlet interface {
    IDemoGameMidlet
    SpecialDemoGameMidlet_() *SpecialDemoGameMidlet
    InitView()
    PreInit()
}

func (this *SpecialDemoGameMidlet) SpecialDemoGameMidlet_() *SpecialDemoGameMidlet {
    return this
}



func NewSpecialDemoGameMidlet(clientInformationFactory *ClientInformationFactory, licenseLoadingType *LicenseLoadingType, demoSetupFactory *LicensedDemoSetupFactory, licenseCheckRunnableFactory *LicenseCheckRunnableFactory) *SpecialDemoGameMidlet {
    this := &SpecialDemoGameMidlet{}
        this.DemoGameMidlet = *NewDemoGameMidlet(clientInformationFactory)
    return this
}


func (this *SpecialDemoGameMidlet) InitView() {
}


func (this *SpecialDemoGameMidlet) PreInit() {
        NewLicenseServerInitFileUtil().Init()
}

