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

type AdvertisementCampaign struct {
    HashMap *HashMap
}

// Implements AdvertisementCampaignInterface


func (this *AdvertisementCampaign) initFields() {
    
}


func NewAdvertisementCampaign(hashMap *HashMap) *AdvertisementCampaign {
    this := &AdvertisementCampaign{}
    this.initFields()
        this.HashMap = hashMap
    return this
}


func (this *AdvertisementCampaign) GetComponent() any {
        return nil
}


func (this *AdvertisementCampaign) GetName() string {
        return any(this.HashMap.Get(AdvertisementCampaignData_GetInstance().NAME)).(string)
}


func (this *AdvertisementCampaign) SetComponentName(name string) {
        this.HashMap.Put(DynamicObjectData_NAME, name)
}


func (this *AdvertisementCampaign) SetName(name string) {
        this.HashMap.Put(AdvertisementCampaignData_GetInstance().NAME, name)
}


func (this *AdvertisementCampaign) Search(advertisementSearchInterface AdvertisementSearchInterface) AdvertisementsInterface {
        return nil
}

