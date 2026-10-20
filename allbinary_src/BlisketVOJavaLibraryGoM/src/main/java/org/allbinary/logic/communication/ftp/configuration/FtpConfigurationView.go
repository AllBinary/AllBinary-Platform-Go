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

type FtpConfigurationView struct {
    FtpConfigurationInterface FtpConfigurationInterface
}

// Implements DomNodeInterface


func (this *FtpConfigurationView) initFields() {
    
}




// Java overloads of new FtpConfigurationView
func NewFtpConfigurationView(args ...any) *FtpConfigurationView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewFtpConfigurationView_1(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[FtpConfigurationInterface](args[0], pass) {
            return NewFtpConfigurationView_2(Java_Arg[FtpConfigurationInterface](args[0]))
        }
    }
    panic("No Java overload of new FtpConfigurationView for the arguments")
}
func NewFtpConfigurationView_1(node *Node) *FtpConfigurationView {
    this := &FtpConfigurationView{}
    this.initFields()
        this.SetFtpConfigurationInterface(any(NewFtpConfiguration()).(FtpConfigurationInterface))
        var chileNodeList *NodeList = node.GetChildNodes(); _ = chileNodeList
        var server string = DomNodeHelper_GetTextNodeValue(FtpConfigurationData_SERVER, chileNodeList); _ = server
        this.GetFtpConfigurationInterface().SetServer(server)
        var userName string = DomNodeHelper_GetTextNodeValue(FtpConfigurationData_USERNAME, chileNodeList); _ = userName
        this.GetFtpConfigurationInterface().SetUserName(userName)
        var password string = DomNodeHelper_GetTextNodeValue(FtpConfigurationData_PASSWORD, chileNodeList); _ = password
        this.GetFtpConfigurationInterface().SetPassword(password)
        var path string = DomNodeHelper_GetTextNodeValue(FtpConfigurationData_PATH, chileNodeList); _ = path
        this.GetFtpConfigurationInterface().SetPath(path)
    return this
}


func NewFtpConfigurationView_2(ftpConfigurationInterface FtpConfigurationInterface) *FtpConfigurationView {
    this := &FtpConfigurationView{}
    this.initFields()
        this.SetFtpConfigurationInterface(ftpConfigurationInterface)
    return this
}


// throws Exception
func (this *FtpConfigurationView) ToXmlNode(document *Document) *Node {
        var ftpConfigurationMapping *FtpConfigurationMapping = NewFtpConfigurationMapping(this.GetFtpConfigurationInterface()); _ = ftpConfigurationMapping
        var hashMap *HashMap = ftpConfigurationMapping.ToHashMap(); _ = hashMap
        var node *Node = ModDomHelper_CreateNameValueNodes(document, FtpConfigurationData_NAME, hashMap); _ = node
        return node
}


func (this *FtpConfigurationView) GetFtpConfigurationInterface() FtpConfigurationInterface {
        return this.FtpConfigurationInterface
}


func (this *FtpConfigurationView) SetFtpConfigurationInterface(ftpConfigurationInterface FtpConfigurationInterface) {
        this.FtpConfigurationInterface = ftpConfigurationInterface
}

