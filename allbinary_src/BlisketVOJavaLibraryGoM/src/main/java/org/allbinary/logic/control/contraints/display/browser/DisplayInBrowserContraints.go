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

type DisplayInBrowserContraints struct {
    Frameable bool
    Mountable bool
    QuickFrameable bool
    MountedLink string
}

// Implements DisplayInBrowserContraintsInterface


func (this *DisplayInBrowserContraints) initFields() {
    
    
    
    
}


func NewDisplayInBrowserContraints(node *Node) *DisplayInBrowserContraints {
    this := &DisplayInBrowserContraints{}
    this.initFields()
    return this
}


func (this *DisplayInBrowserContraints) IsMountable() bool {
        return this.Mountable
}


func (this *DisplayInBrowserContraints) SetMountable(mountable bool) {
        this.Mountable = mountable
}


func (this *DisplayInBrowserContraints) IsFrameable() bool {
        return this.Frameable
}


func (this *DisplayInBrowserContraints) SetFrameable(frameable bool) {
        this.Frameable = frameable
}


func (this *DisplayInBrowserContraints) IsQuickFrameable() bool {
        return this.QuickFrameable
}


func (this *DisplayInBrowserContraints) SetQuickFrameable(quickFrameable bool) {
        this.QuickFrameable = quickFrameable
}


func (this *DisplayInBrowserContraints) GetMountedLink() string {
        return this.MountedLink
}


func (this *DisplayInBrowserContraints) SetMountedLink(mountedLink string) {
        this.MountedLink = mountedLink
}

