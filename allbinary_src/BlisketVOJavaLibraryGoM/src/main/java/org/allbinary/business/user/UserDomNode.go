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

type UserDomNode struct {
    DataMappingInterface TableMappingInterface
}

// Implements DomNodeInterface


func (this *UserDomNode) initFields() {
    
}


func NewUserDomNode(dataMappingInterface TableMappingInterface) *UserDomNode {
    this := &UserDomNode{}
    this.initFields()
        this.DataMappingInterface = dataMappingInterface
    return this
}


// throws Exception
func (this *UserDomNode) ToXmlNode(document *Document) *Node {
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        var hashMap *HashMap = this.DataMappingInterface.ToHashMap(); _ = hashMap
        hashMap.Put(EntryData_GetInstance().LASTMODIFIED, EMPTY_STRING)
        hashMap.Put(EntryData_GetInstance().ENCRYPTION, EMPTY_STRING)
        hashMap.Put(UserData_PASSWORD, EMPTY_STRING)
        return ModDomHelper_CreateNameValueNodes(document, UserData_NAME, hashMap)
}


// throws Exception
func (this *UserDomNode) ToXmlDoc() *Document {
        return nil
}


// throws Exception
func (this *UserDomNode) View() string {
        return nil
}

