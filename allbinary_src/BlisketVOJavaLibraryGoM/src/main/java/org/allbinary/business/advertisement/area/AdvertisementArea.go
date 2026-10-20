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

type AdvertisementArea struct {
    Name string
    StoreName string
    Description string
    AdvertisementAreaConstraintsInterface AdvertisementAreaConstraintsInterface
    Campaign string
}

// Implements AdvertisementAreaInterface


func (this *AdvertisementArea) initFields() {
    
    
    
    
    
}


func NewAdvertisementArea(hashMap *HashMap) *AdvertisementArea {
    this := &AdvertisementArea{}
    this.initFields()
        this.Name = any(hashMap.Get(AdvertisementAreaData_GetInstance().NAME)).(string)
        this.StoreName = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string)
        this.Description = any(hashMap.Get(AdvertisementAreaData_GetInstance().DESCRIPTION)).(string)
        var constraintsDomString string = any(hashMap.Get(AdvertisementAreaData_GetInstance().CONSTRAINTS)).(string); _ = constraintsDomString
        var document *Document = DomDocumentHelper_Create(constraintsDomString); _ = document
        this.AdvertisementAreaConstraintsInterface = NewAdvertisementAreaConstraints(document)
        this.Campaign = any(hashMap.Get(AdvertisementCampaignData_GetInstance().NAME)).(string)
    return this
}


func (this *AdvertisementArea) GetName() string {
        return this.Name
}


func (this *AdvertisementArea) GetStoreName() string {
        return this.StoreName
}


func (this *AdvertisementArea) GetDescription() string {
        return this.Description
}


func (this *AdvertisementArea) GetConstraints() AdvertisementAreaConstraintsInterface {
        return this.AdvertisementAreaConstraintsInterface
}


func (this *AdvertisementArea) GetCampaign() string {
        return this.Campaign
}


func (this *AdvertisementArea) SetName(name string) {
        this.Name = name
}


func (this *AdvertisementArea) SetStoreName(storeName string) {
        this.StoreName = storeName
}


func (this *AdvertisementArea) SetDescription(description string) {
        this.Description = description
}


func (this *AdvertisementArea) SetConstraints(advertisementConstraintsInterface AdvertisementAreaConstraintsInterface) {
        this.AdvertisementAreaConstraintsInterface = this.AdvertisementAreaConstraintsInterface
}


func (this *AdvertisementArea) SetCampaign(campaign string) {
        this.Campaign = campaign
}

