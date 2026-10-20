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

type GenericProfileActions struct {
    LogUtil *LogUtil
    CommonFileStrings *CommonDataFileStrings
    Name string
    GenericProfileActionsJPanel *GenericProfileActionsJPanel
    ActionsDefaultListModelHelper *DefaultListModelHelper
    HashMap *HashMap
}

// Implements DomNodeInterface, DomDocumentMappingInterface

var GenericProfileActions_DEFAULT_PROFILE_ACTIONS_PATH string = fmt.Sprint(GenericProfiles_DEFAULT_PROFILES_PATH) + "actions/"

func (this *GenericProfileActions) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonFileStrings = CommonDataFileStrings_GetInstance()
    
    
    
    
}




// Java overloads of new GenericProfileActions
func NewGenericProfileActions(args ...any) *GenericProfileActions {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*GenericProfileActionsJPanel](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewGenericProfileActions_1(Java_Arg[*GenericProfileActionsJPanel](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 3 && Java_IsArg[*GenericProfileActionsJPanel](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewGenericProfileActions_2(Java_Arg[*GenericProfileActionsJPanel](args[0]), Java_Arg[IAbPath](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 3 && Java_IsArg[*GenericProfileActionsJPanel](args[0], pass) && Java_IsArg[*FileInputStream](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewGenericProfileActions_3(Java_Arg[*GenericProfileActionsJPanel](args[0]), Java_Arg[*FileInputStream](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of new GenericProfileActions for the arguments")
}
func NewGenericProfileActions_1(genericProfileActionsJPanel *GenericProfileActionsJPanel, name string) *GenericProfileActions {
    this := &GenericProfileActions{}
    this.initFields()
        this.SetName(name)
        this.Init(genericProfileActionsJPanel)
        this.Load()
    return this
}


func NewGenericProfileActions_2(genericProfileActionsJPanel *GenericProfileActionsJPanel, abPath IAbPath, name string) *GenericProfileActions {
    this := &GenericProfileActions{}
    this.initFields()
        this.SetName(name)
        this.Init(genericProfileActionsJPanel)
        this.Load()
    return this
}


func NewGenericProfileActions_3(genericProfileActionsJPanel *GenericProfileActionsJPanel, fileInputStream *FileInputStream, name string) *GenericProfileActions {
    this := &GenericProfileActions{}
    this.initFields()
        this.SetName(name)
        this.Init(genericProfileActionsJPanel)
        this.FileInit(fileInputStream)
    return this
}


// throws Exception
func (this *GenericProfileActions) Init(genericProfileActionsJPanel *GenericProfileActionsJPanel) {
        this.SetGenericProfileActionsJPanel(genericProfileActionsJPanel)
        this.ActionsDefaultListModelHelper = NewDefaultListModelHelper()
        this.SetHashMap(NewHashMap())
}


// throws Exception
func (this *GenericProfileActions) Save() {
        var idFile *FileOutputStream = NewFileOutputStream(fmt.Sprint(GenericProfileActions_DEFAULT_PROFILE_ACTIONS_PATH) + this.GetName() + fmt.Sprint(this.CommonFileStrings._XML)); _ = idFile
        var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
        idOutData.WriteBytes(XmlDocumentHelper_ToString(this.ToXmlDoc()))
}


func (this *GenericProfileActions) GetFile(name string) *File {
        var fileName string = GenericProfileActions_DEFAULT_PROFILE_ACTIONS_PATH + name + fmt.Sprint(this.CommonFileStrings._XML); _ = fileName
        return NewFile(fileName)
}


// throws Exception
func (this *GenericProfileActions) Load() {
        var file *File = this.GetFile(this.GetName()); _ = file
        if file.IsFile() {
            var idFile *FileInputStream = NewFileInputStream(file); _ = idFile
            this.FileInit(idFile)
        } else {
            this.LogUtil.PutF("No Profile: " + fmt.Sprint(file.GetAbsolutePath()), this, "Contructor")
        }
}


// throws Exception
func (this *GenericProfileActions) FileInit(fileInputStream *FileInputStream) {
        var bytes []int8 = make([]int8, 100000); _ = bytes
        var length int = fileInputStream.Read(bytes); _ = length
        var data string = string(bytes); _ = data
        var endIndex int = strings.LastIndex(data, string('>')); _ = endIndex
        var document *Document = DomDocumentHelper_Create(data[0:endIndex + 1]); _ = document
        var nameNodeList *NodeList = document.GetElementsByTagName(GenericProfileActionsData_NAME); _ = nameNodeList
        this.LogUtil.PutF("Number Of Profiles Specified: " + fmt.Sprint(nameNodeList.GetLength()), this, "Contructor")
        for index := int(0); index < nameNodeList.GetLength(); index++ {
            var node *Node = nameNodeList.Item(index); _ = node
            var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
            if nodeList != nil {
                this.InitActions(nodeList)
            } else {
                panic(errors.New(fmt.Sprint(fmt.Sprint(GenericProfileActionData_NAME) + " Name Node Node Children")))
            }
        }
        this.LogUtil.PutF("Loaded: " + fmt.Sprint(this.GetHashMap().Size()) + " Configuration Profile Action(s)", this, "Contructor")
        this.GetDefaultListModelHelper().InitDefaultModelList()
}


// throws Exception
func (this *GenericProfileActions) InitActions(nodeList *NodeList) {
        for index := int(0); index < nodeList.GetLength(); index++ {
            var actionNode *Node = nodeList.Item(index); _ = actionNode
            if actionNode != nil {
                var genericConfigurationProfileAction *GenericProfileAction = NewGenericProfileAction(this.GetGenericProfileActionsJPanel().GetGenericProfileActionJPanel(), actionNode); _ = genericConfigurationProfileAction
                this.GetHashMap().Put(genericConfigurationProfileAction.GetName(), genericConfigurationProfileAction)
                this.GetDefaultListModelHelper().Add(genericConfigurationProfileAction.GetName())
            } else {
                panic(errors.New(fmt.Sprint(fmt.Sprint(GenericProfileActionData_NAME) + " Node Null")))
            }
        }
}


func (this *GenericProfileActions) GetDefaultListModelHelper() *DefaultListModelHelper {
        return this.ActionsDefaultListModelHelper
}


func (this *GenericProfileActions) GetAction(string_ string) *GenericProfileAction {
        return any(this.GetHashMap().Get(string_)).(*GenericProfileAction)
}


// throws Exception
func (this *GenericProfileActions) Add(name string) {
        var genericProfileAction *GenericProfileAction = NewGenericProfileAction(name); _ = genericProfileAction
        this.GetHashMap().Put(genericProfileAction.GetName(), genericProfileAction)
        this.GetDefaultListModelHelper().Add(genericProfileAction.GetName())
        this.GetDefaultListModelHelper().InitDefaultModelList()
        this.Save()
}


// throws Exception
func (this *GenericProfileActions) Remove(name string) {
        this.GetHashMap().Remove(name)
        this.GetDefaultListModelHelper().Remove(name)
        this.GetDefaultListModelHelper().InitDefaultModelList()
        this.Save()
}


func (this *GenericProfileActions) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, "toHashMap()")
        return hashMap
}


// throws Exception
func (this *GenericProfileActions) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfileActionsData_NAME); _ = node
        var set Set = this.GetHashMap().KeySet(); _ = set
        var actionNameArray []any = set.ToArray(); _ = actionNameArray
        var size int = len(actionNameArray); _ = size
        for index := int(0); index < size; index++ {
            var nextActionName string = any(actionNameArray[index]).(string); _ = nextActionName
            var nextGenericProfileAction *GenericProfileAction = any(this.GetAction(nextActionName)).(*GenericProfileAction); _ = nextGenericProfileAction
            node.AppendChild(nextGenericProfileAction.ToXmlNode(document))
        }
        return node
}


// throws Exception
func (this *GenericProfileActions) ToXmlDoc() *Document {
        var document *Document = DomDocumentHelper_Create(); _ = document
        var node *Node = this.ToXmlNode(document); _ = node
        document.AppendChild(node)
        return document
}


func (this *GenericProfileActions) GetGenericProfileActionsJPanel() *GenericProfileActionsJPanel {
        return this.GenericProfileActionsJPanel
}


func (this *GenericProfileActions) SetGenericProfileActionsJPanel(genericProfileActionsJPanel *GenericProfileActionsJPanel) {
        this.GenericProfileActionsJPanel = genericProfileActionsJPanel
}


func (this *GenericProfileActions) GetHashMap() *HashMap {
        return this.HashMap
}


func (this *GenericProfileActions) SetHashMap(hashMap *HashMap) {
        this.HashMap = hashMap
}


func (this *GenericProfileActions) GetName() string {
        return this.Name
}


func (this *GenericProfileActions) SetName(name string) {
        this.Name = name
}

