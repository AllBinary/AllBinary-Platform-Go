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

type PCGameInputMapping struct {
    PersistentInputMapping
}

// The Java class PCGameInputMapping has subclasses so its variables are this interface
type IPCGameInputMapping interface {
    IPersistentInputMapping
    PCGameInputMapping_() *PCGameInputMapping
    Init(abeClientInformation AbeClientInformationInterface)
    IsDelete(input IInput) bool
    IsSystemInput(input IInput) bool
}

func (this *PCGameInputMapping) PCGameInputMapping_() *PCGameInputMapping {
    return this
}



func NewPCGameInputMapping() *PCGameInputMapping {
    this := &PCGameInputMapping{}
        this.PersistentInputMapping = *NewPersistentInputMapping(GamePersistanceStrings_GetInstance().SAVED_INPUT_CONFIGURATION_RECORD_ID)
    return this
}


// throws Exception
func (this *PCGameInputMapping) Init(abeClientInformation AbeClientInformationInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.INIT)
        PCKeyFactory_GetInstance()
        this.PersistentInputMapping.Init(abeClientInformation)
        if this.GetTotalMapped() == 0 || this.IsDefaultNew() {
            this.GetInputMapping().AddMapping(this.GetDefault())
            this.Save(abeClientInformation)
        }
}


func (this *PCGameInputMapping) IsDelete(input IInput) bool {
        if input == PCKeyFactory_GetInstance().DEL {
            return true
        } else {
            return false
        }
}


func (this *PCGameInputMapping) IsSystemInput(input IInput) bool {
        return false
}

