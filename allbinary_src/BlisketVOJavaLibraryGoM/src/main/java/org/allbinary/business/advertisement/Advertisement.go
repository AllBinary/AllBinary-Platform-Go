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

type Advertisement struct {
    Artist string
    ArtistLink string
    Title string
    DateAdded string
    Directory string
    FramedLink string
    Rank int
    ListPrice *Money
    OurPrice *Money
    Image AdvertisementImageInterface
    Product AdvertisementProductInterface
    Thumbnail AdvertisementThumbnailInterface
    DisplayInBrowserContraintsInterface DisplayInBrowserContraintsInterface
}

// Implements AdvertisementInterface


func (this *Advertisement) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewAdvertisement() *Advertisement {
    this := &Advertisement{}
    this.initFields()
        this.DisplayInBrowserContraintsInterface = nil
    return this
}


func (this *Advertisement) GetArtist() string {
        return this.Artist
}


func (this *Advertisement) SetArtist(artist string) {
        this.Artist = artist
}


func (this *Advertisement) GetArtistLink() string {
        return this.ArtistLink
}


func (this *Advertisement) SetArtistLink(artistLink string) {
        this.ArtistLink = artistLink
}


func (this *Advertisement) GetTitle() string {
        return this.Title
}


func (this *Advertisement) SetTitle(title string) {
        this.Title = title
}


func (this *Advertisement) GetDateAdded() string {
        return this.DateAdded
}


func (this *Advertisement) SetDateAdded(dateAdded string) {
        this.DateAdded = dateAdded
}


func (this *Advertisement) GetDirectory() string {
        return this.Directory
}


func (this *Advertisement) SetDirectory(directory string) {
        this.Directory = directory
}


func (this *Advertisement) GetFramedLink() string {
        return this.FramedLink
}


func (this *Advertisement) SetFramedLink(framedLink string) {
        this.FramedLink = framedLink
}


func (this *Advertisement) GetListPrice() *Money {
        return this.ListPrice
}


func (this *Advertisement) SetListPrice(listPrice *Money) {
        this.ListPrice = listPrice
}


func (this *Advertisement) GetOurPrice() *Money {
        return this.OurPrice
}


func (this *Advertisement) SetOurPrice(ourPrice *Money) {
        this.OurPrice = ourPrice
}


func (this *Advertisement) GetRank() int {
        return this.Rank
}


func (this *Advertisement) SetRank(rank int) {
        this.Rank = rank
}


func (this *Advertisement) GetImage() AdvertisementImageInterface {
        return this.Image
}


func (this *Advertisement) SetImage(image AdvertisementImageInterface) {
        this.Image = image
}


func (this *Advertisement) GetProduct() AdvertisementProductInterface {
        return this.Product
}


func (this *Advertisement) SetProduct(product AdvertisementProductInterface) {
        this.Product = product
}


func (this *Advertisement) GetThumbnail() AdvertisementThumbnailInterface {
        return this.Thumbnail
}


func (this *Advertisement) SetThumbnail(thumbnail AdvertisementThumbnailInterface) {
        this.Thumbnail = thumbnail
}


func (this *Advertisement) GetDisplayInBrowserContraintsInterface() DisplayInBrowserContraintsInterface {
        return this.DisplayInBrowserContraintsInterface
}

