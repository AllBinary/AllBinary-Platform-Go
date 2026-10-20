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

type GroupLayerManagerListener struct {
    LayerManagerEventListener
    LogUtil *LogUtil
    List IBasicArrayList
}

var GroupLayerManagerListener_SINGLETON *GroupLayerManagerListener = NewGroupLayerManagerListener()

func (this *GroupLayerManagerListener) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.List = NewBasicArrayListD()
}


func GroupLayerManagerListener_GetInstance() *GroupLayerManagerListener {
        return GroupLayerManagerListener_SINGLETON
}


func NewGroupLayerManagerListener() *GroupLayerManagerListener {
    this := &GroupLayerManagerListener{}
    this.LayerManagerEventListener = *NewLayerManagerEventListener()
    this.initFields()
        LayerManagerEventHandler_GetInstance().AddListener(this)
    return this
}


func (this *GroupLayerManagerListener) Clear() {
        var groupList IBasicArrayList; _ = groupList
        var size int = this.List.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            groupList = any(this.List.BasicArrayList_().ObjectArray[index]).(IBasicArrayList)
            groupList.Clear()
        }
}


func (this *GroupLayerManagerListener) GetGroupSizeFromInterface(groupInterfaceCompositeInterface GroupInterfaceCompositeInterface) int {
        var groupInterfaceArray []*Group = groupInterfaceCompositeInterface.GetGroupInterface(); _ = groupInterfaceArray
        return this.GetGroupSize(groupInterfaceArray[0])
}


func (this *GroupLayerManagerListener) GetGroupSize(groupInterface *Group) int {
        var id int = int(groupInterface.GetGroupId()); _ = id
        return this.GetGroupSizeById(id)
}


func (this *GroupLayerManagerListener) GetList(groupInterface *Group) IBasicArrayList {
        var id int = int(groupInterface.GetGroupId()); _ = id
        return this.GetListById(id)
}


func (this *GroupLayerManagerListener) GetListById(groupId int) IBasicArrayList {
        var groupList IBasicArrayList = any(this.List.BasicArrayList_().ObjectArray[groupId]).(IBasicArrayList); _ = groupList
        return groupList
}


func (this *GroupLayerManagerListener) GetGroupSizeById(groupId int) int {
        var groupList IBasicArrayList = any(this.List.BasicArrayList_().ObjectArray[groupId]).(IBasicArrayList); _ = groupList
        var size int = groupList.Size(); _ = size
        return size
}


func (this *GroupLayerManagerListener) AreAllOtherGroupsEmpty(groupInterface *Group) bool {
        var id int = int(groupInterface.GetGroupId()); _ = id
        var size int = this.List.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            if id != index {
                var groupSize int = this.GetGroupSizeById(index); _ = groupSize
                if groupSize != 0 {
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Group Size: ").Appendint(groupSize)), this, "areAllOtherGroupsEmpty")
                    return false
                }
            }
        }
        return true
}


func (this *GroupLayerManagerListener) IsIdInList(id int, excludeGroupList IBasicArrayList) bool {
        var size int = excludeGroupList.Size(); _ = size
        var groupInterfaceArray []*Group; _ = groupInterfaceArray
        var groupInterface *Group; _ = groupInterface
        var groupId int; _ = groupId
        for index := int(size - 1); index >= 0; index-- {
            groupInterfaceArray = any(excludeGroupList.BasicArrayList_().ObjectArray[index]).([]*Group)
            var size2 int = len(groupInterfaceArray); _ = size2
            for index2 := int(0); index2 < size2; index2++ {
                groupInterface = groupInterfaceArray[index2]
                groupId = int(groupInterface.GetGroupId())
                if groupId == id {
                    return true
                }
            }
        }
        return false
}


func (this *GroupLayerManagerListener) AreAllOtherGroupsLessThan(excludeGroupList IBasicArrayList, maxSize int) bool {
        var size int = this.List.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            if !this.IsIdInList(index, excludeGroupList) {
                var groupSize int = this.GetGroupSizeById(index); _ = groupSize
                if groupSize >= maxSize {
                    return false
                }
            }
        }
        return true
}


func (this *GroupLayerManagerListener) Init(total int) {
        for this.List.Size() <= total + 1 {
            this.List.Add(NewBasicArrayListD())
        }
}


func (this *GroupLayerManagerListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *GroupLayerManagerListener) OnCreateLayerManagerEvent(layerManagerEvent *LayerManagerEvent) {
        var layerInterface IAllBinaryLayer = layerManagerEvent.GetLayerInterface(); _ = layerInterface
        var groupInterfaceArray []*Group = layerInterface.GetGroupInterface(); _ = groupInterfaceArray
        var size int = len(groupInterfaceArray); _ = size
        var id int; _ = id
        var groupList IBasicArrayList; _ = groupList
        for index := int(0); index < size; index++ {
            id = int(groupInterfaceArray[index].GetGroupId())
            groupList = any(this.List.BasicArrayList_().ObjectArray[id]).(IBasicArrayList)
            if groupList == nil {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("id: ").Appendint(id)), this, "onCreateLayerManagerEvent")
            }
            if !groupList.Contains(layerInterface) {
                groupList.Add(layerInterface)
            } else {
            }
        }
}


// throws Exception
func (this *GroupLayerManagerListener) OnDeleteLayerManagerEvent(layerManagerEvent *LayerManagerEvent) {
        var layerInterface IAllBinaryLayer = layerManagerEvent.GetLayerInterface(); _ = layerInterface
        var groupInterfaceArray []*Group = layerInterface.GetGroupInterface(); _ = groupInterfaceArray
        var size int = len(groupInterfaceArray); _ = size
        var id int; _ = id
        var groupList IBasicArrayList; _ = groupList
        for index := int(0); index < size; index++ {
            id = int(groupInterfaceArray[index].GetGroupId())
            groupList = any(this.List.BasicArrayList_().ObjectArray[id]).(IBasicArrayList)
            groupList.Remove(layerInterface)
        }
}


func (this *GroupLayerManagerListener) Log() {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var GROUP string = "Group: "; _ = GROUP
        var SPACE string = CommonSeps_GetInstance().SPACE; _ = SPACE
        var TOTAL_LABEL string = CommonLabels_GetInstance().TOTAL_LABEL; _ = TOTAL_LABEL
        var size int = this.List.Size(); _ = size
        for index := int(size - 1); index >= 0; index-- {
            var groupList IBasicArrayList = any(this.List.BasicArrayList_().ObjectArray[index]).(IBasicArrayList); _ = groupList
            stringBuffer.Append(GROUP)
            stringBuffer.Appendint(index)
            stringBuffer.Append(SPACE)
            stringBuffer.Append(TOTAL_LABEL)
            stringBuffer.Appendint(groupList.Size())
            stringBuffer.Append(SPACE)
        }
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "log")
}

