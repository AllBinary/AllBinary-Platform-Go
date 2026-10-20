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

type AuthorizationTypeFactory struct {
    NONE *AuthorizationType
    VOICE *AuthorizationType
    RETINAL *AuthorizationType
    IMAGE *AuthorizationType
    DNA *AuthorizationType
    PRINT *AuthorizationType
    CHIP *AuthorizationType
}

var AuthorizationTypeFactory_Instance *AuthorizationTypeFactory = NewAuthorizationTypeFactory()

func (this *AuthorizationTypeFactory) initFields() {
    this.NONE = NewAuthorizationType("None")
    this.VOICE = NewAuthorizationType("Voice")
    this.RETINAL = NewAuthorizationType("Retinal")
    this.IMAGE = NewAuthorizationType("Image")
    this.DNA = NewAuthorizationType("DNA")
    this.PRINT = NewAuthorizationType("Print")
    this.CHIP = NewAuthorizationType("Chip")
}

func NewAuthorizationTypeFactory() *AuthorizationTypeFactory {
    this := &AuthorizationTypeFactory{}
    this.initFields()
    return this
}


func AuthorizationTypeFactory_GetInstance() *AuthorizationTypeFactory {
        return AuthorizationTypeFactory_Instance
}

