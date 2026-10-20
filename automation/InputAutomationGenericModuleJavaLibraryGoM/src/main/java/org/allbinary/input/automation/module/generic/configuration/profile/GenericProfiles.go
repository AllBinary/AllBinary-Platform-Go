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

type GenericProfiles struct {
    LogUtil *LogUtil
    HashMap *HashMap
    ProfilesDefaultListModelHelper *DefaultListModelHelper
    FileName string
}

// Implements DomNodeInterface, DomDocumentMappingInterface

var GenericProfiles_DEFAULT_PROFILES_PATH string = "./modules/configs/profiles/"
var GenericProfiles_DEFAULT_PROFILE_ACTIONS_PATH string = fmt.Sprint(GenericProfiles_DEFAULT_PROFILES_PATH) + "actions/"
var GenericProfiles_DEFAULT_FILE string = fmt.Sprint(GenericProfiles_DEFAULT_PROFILES_PATH) + "profiles.xml"

func (this *GenericProfiles) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewGenericProfiles(fileName string) *GenericProfiles {
    this := &GenericProfiles{}
    this.initFields()
        this.FileName = fileName
        this.ProfilesDefaultListModelHelper = NewDefaultListModelHelper()
        this.HashMap = StdUtil_GetInstance().CreateHashMap()
        if NewFile(fileName).IsFile() {
            var bytes []int8 = make([]int8, 100000); _ = bytes
            var idFile *FileInputStream = NewFileInputStream(fileName); _ = idFile
            var length int = idFile.Read(bytes); _ = length
            var data string = string(bytes); _ = data
            var endIndex int = strings.LastIndex(data, string('>')); _ = endIndex
            var document *Document = DomDocumentHelper_Create(data[0:endIndex + 1]); _ = document
            var nodeList *NodeList = document.GetElementsByTagName(GenericProfilesData_NAME); _ = nodeList
            this.LogUtil.PutF("Number Of Profiles Specified: " + fmt.Sprint(nodeList.GetLength()), this, "Contructor")
            if nodeList != nil {
                this.InitProfiles(nodeList.Item(0))
            } else {
                panic(errors.New(fmt.Sprint(fmt.Sprint(GenericProfileActionData_NAME) + " Name Node Node Children")))
            }
            this.LogUtil.PutF("Loaded: " + fmt.Sprint(this.HashMap.Size()) + " Configuration Profile Actions", this, "Contructor")
        } else {
            this.LogUtil.PutF("No Generic Profile: " + fileName, this, "Contructor")
        }
        this.GetDefaultListModelHelper().InitDefaultModelList()
    return this
}


// throws Exception
func (this *GenericProfiles) InitProfiles(node *Node) {
        var nodeList *NodeList = node.GetChildNodes(); _ = nodeList
        for index := int(0); index < nodeList.GetLength(); index++ {
            var profileNameNode *Node = nodeList.Item(index); _ = profileNameNode
            if profileNameNode.GetNodeName().CompareTo(GenericProfileData_NAME) == 0 {
                var genericProfile *GenericProfile = NewGenericProfile(profileNameNode); _ = genericProfile
                this.HashMap.Put(genericProfile.GetName(), genericProfile)
                this.GetDefaultListModelHelper().Add(genericProfile.GetName())
            }
        }
}


func (this *GenericProfiles) GetDefaultListModelHelper() *DefaultListModelHelper {
        return this.ProfilesDefaultListModelHelper
}


// throws Exception
func (this *GenericProfiles) Get(name string) *GenericProfile {
        var genericProfile *GenericProfile = any(this.HashMap.Get(name)).(*GenericProfile); _ = genericProfile
        if genericProfile == nil {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF("No Generic Profile Named: " + name + " availability was: " + fmt.Sprint(this.HashMap), this, commonStrings.GET)
        }
        return genericProfile
}


// throws Exception
func (this *GenericProfiles) Add(name string) {
        this.HashMap.Put(name, NewGenericProfile(name))
        this.GetDefaultListModelHelper().Add(name)
        this.GetDefaultListModelHelper().InitDefaultModelList()
        this.Save()
}


// throws Exception
func (this *GenericProfiles) Remove(name string) {
        this.HashMap.Remove(name)
        this.GetDefaultListModelHelper().Remove(name)
        this.GetDefaultListModelHelper().InitDefaultModelList()
        this.Save()
}


// throws Exception
func (this *GenericProfiles) ToXmlNode(document *Document) *Node {
        var node *Node = document.CreateElement(GenericProfilesData_NAME); _ = node
        var set Set = this.HashMap.KeySet(); _ = set
        var nameArray []any = set.ToArray(); _ = nameArray
        var size int = len(nameArray); _ = size
        for index := int(0); index < size; index++ {
            var genericProfile *GenericProfile = any(this.HashMap.Get(any(nameArray[index]).(string))).(*GenericProfile); _ = genericProfile
            node.AppendChild(genericProfile.ToXmlNode(document))
        }
        return node
}


// throws Exception
func (this *GenericProfiles) ToXmlDoc() *Document {
        var document *Document = DomDocumentHelper_Create(); _ = document
        var node *Node = this.ToXmlNode(document); _ = node
        document.AppendChild(node)
        return document
}


// throws Exception
func (this *GenericProfiles) Save() {
        var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
        var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
        var documentString string = XmlDocumentHelper_ToString(this.ToXmlDoc()); _ = documentString
        idOutData.WriteBytes(documentString)
}

