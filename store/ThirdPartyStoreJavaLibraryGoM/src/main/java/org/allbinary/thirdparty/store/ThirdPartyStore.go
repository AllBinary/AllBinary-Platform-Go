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

type ThirdPartyStore struct {
    Authorized bool
    Name string
}

// The Java class ThirdPartyStore has subclasses so its variables are this interface
type IThirdPartyStore interface {
    ThirdPartyStore_() *ThirdPartyStore
    Authorize(clientInformationInterface AbeClientInformationInterface)
    IsAuthorized() bool
    SetAuthorized(authorized bool)
    GetName() string
}

func (this *ThirdPartyStore) ThirdPartyStore_() *ThirdPartyStore {
    return this
}


func (this *ThirdPartyStore) initFields() {
    
    
}


func NewThirdPartyStore(name string) *ThirdPartyStore {
    this := &ThirdPartyStore{}
    this.initFields()
        this.Name = name
    return this
}


func (this *ThirdPartyStore) Authorize(clientInformationInterface AbeClientInformationInterface) {
}


func (this *ThirdPartyStore) IsAuthorized() bool {
        return this.Authorized
}


func (this *ThirdPartyStore) SetAuthorized(authorized bool) {
        this.Authorized = authorized
}


func (this *ThirdPartyStore) GetName() string {
        return this.Name
}

