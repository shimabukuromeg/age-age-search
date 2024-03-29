// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"time"

	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MeshiWhereInput represents a where input for filtering Meshi queries.
type MeshiWhereInput struct {
	Predicates []predicate.Meshi  `json:"-"`
	Not        *MeshiWhereInput   `json:"not,omitempty"`
	Or         []*MeshiWhereInput `json:"or,omitempty"`
	And        []*MeshiWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "article_id" field predicates.
	ArticleID             *string  `json:"articleID,omitempty"`
	ArticleIDNEQ          *string  `json:"articleIDNEQ,omitempty"`
	ArticleIDIn           []string `json:"articleIDIn,omitempty"`
	ArticleIDNotIn        []string `json:"articleIDNotIn,omitempty"`
	ArticleIDGT           *string  `json:"articleIDGT,omitempty"`
	ArticleIDGTE          *string  `json:"articleIDGTE,omitempty"`
	ArticleIDLT           *string  `json:"articleIDLT,omitempty"`
	ArticleIDLTE          *string  `json:"articleIDLTE,omitempty"`
	ArticleIDContains     *string  `json:"articleIDContains,omitempty"`
	ArticleIDHasPrefix    *string  `json:"articleIDHasPrefix,omitempty"`
	ArticleIDHasSuffix    *string  `json:"articleIDHasSuffix,omitempty"`
	ArticleIDEqualFold    *string  `json:"articleIDEqualFold,omitempty"`
	ArticleIDContainsFold *string  `json:"articleIDContainsFold,omitempty"`

	// "title" field predicates.
	Title             *string  `json:"title,omitempty"`
	TitleNEQ          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGT           *string  `json:"titleGT,omitempty"`
	TitleGTE          *string  `json:"titleGTE,omitempty"`
	TitleLT           *string  `json:"titleLT,omitempty"`
	TitleLTE          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`

	// "image_url" field predicates.
	ImageURL             *string  `json:"imageURL,omitempty"`
	ImageURLNEQ          *string  `json:"imageURLNEQ,omitempty"`
	ImageURLIn           []string `json:"imageURLIn,omitempty"`
	ImageURLNotIn        []string `json:"imageURLNotIn,omitempty"`
	ImageURLGT           *string  `json:"imageURLGT,omitempty"`
	ImageURLGTE          *string  `json:"imageURLGTE,omitempty"`
	ImageURLLT           *string  `json:"imageURLLT,omitempty"`
	ImageURLLTE          *string  `json:"imageURLLTE,omitempty"`
	ImageURLContains     *string  `json:"imageURLContains,omitempty"`
	ImageURLHasPrefix    *string  `json:"imageURLHasPrefix,omitempty"`
	ImageURLHasSuffix    *string  `json:"imageURLHasSuffix,omitempty"`
	ImageURLEqualFold    *string  `json:"imageURLEqualFold,omitempty"`
	ImageURLContainsFold *string  `json:"imageURLContainsFold,omitempty"`

	// "store_name" field predicates.
	StoreName             *string  `json:"storeName,omitempty"`
	StoreNameNEQ          *string  `json:"storeNameNEQ,omitempty"`
	StoreNameIn           []string `json:"storeNameIn,omitempty"`
	StoreNameNotIn        []string `json:"storeNameNotIn,omitempty"`
	StoreNameGT           *string  `json:"storeNameGT,omitempty"`
	StoreNameGTE          *string  `json:"storeNameGTE,omitempty"`
	StoreNameLT           *string  `json:"storeNameLT,omitempty"`
	StoreNameLTE          *string  `json:"storeNameLTE,omitempty"`
	StoreNameContains     *string  `json:"storeNameContains,omitempty"`
	StoreNameHasPrefix    *string  `json:"storeNameHasPrefix,omitempty"`
	StoreNameHasSuffix    *string  `json:"storeNameHasSuffix,omitempty"`
	StoreNameEqualFold    *string  `json:"storeNameEqualFold,omitempty"`
	StoreNameContainsFold *string  `json:"storeNameContainsFold,omitempty"`

	// "address" field predicates.
	Address             *string  `json:"address,omitempty"`
	AddressNEQ          *string  `json:"addressNEQ,omitempty"`
	AddressIn           []string `json:"addressIn,omitempty"`
	AddressNotIn        []string `json:"addressNotIn,omitempty"`
	AddressGT           *string  `json:"addressGT,omitempty"`
	AddressGTE          *string  `json:"addressGTE,omitempty"`
	AddressLT           *string  `json:"addressLT,omitempty"`
	AddressLTE          *string  `json:"addressLTE,omitempty"`
	AddressContains     *string  `json:"addressContains,omitempty"`
	AddressHasPrefix    *string  `json:"addressHasPrefix,omitempty"`
	AddressHasSuffix    *string  `json:"addressHasSuffix,omitempty"`
	AddressEqualFold    *string  `json:"addressEqualFold,omitempty"`
	AddressContainsFold *string  `json:"addressContainsFold,omitempty"`

	// "site_url" field predicates.
	SiteURL             *string  `json:"siteURL,omitempty"`
	SiteURLNEQ          *string  `json:"siteURLNEQ,omitempty"`
	SiteURLIn           []string `json:"siteURLIn,omitempty"`
	SiteURLNotIn        []string `json:"siteURLNotIn,omitempty"`
	SiteURLGT           *string  `json:"siteURLGT,omitempty"`
	SiteURLGTE          *string  `json:"siteURLGTE,omitempty"`
	SiteURLLT           *string  `json:"siteURLLT,omitempty"`
	SiteURLLTE          *string  `json:"siteURLLTE,omitempty"`
	SiteURLContains     *string  `json:"siteURLContains,omitempty"`
	SiteURLHasPrefix    *string  `json:"siteURLHasPrefix,omitempty"`
	SiteURLHasSuffix    *string  `json:"siteURLHasSuffix,omitempty"`
	SiteURLEqualFold    *string  `json:"siteURLEqualFold,omitempty"`
	SiteURLContainsFold *string  `json:"siteURLContainsFold,omitempty"`

	// "published_date" field predicates.
	PublishedDate      *time.Time  `json:"publishedDate,omitempty"`
	PublishedDateNEQ   *time.Time  `json:"publishedDateNEQ,omitempty"`
	PublishedDateIn    []time.Time `json:"publishedDateIn,omitempty"`
	PublishedDateNotIn []time.Time `json:"publishedDateNotIn,omitempty"`
	PublishedDateGT    *time.Time  `json:"publishedDateGT,omitempty"`
	PublishedDateGTE   *time.Time  `json:"publishedDateGTE,omitempty"`
	PublishedDateLT    *time.Time  `json:"publishedDateLT,omitempty"`
	PublishedDateLTE   *time.Time  `json:"publishedDateLTE,omitempty"`

	// "latitude" field predicates.
	Latitude      *float64  `json:"latitude,omitempty"`
	LatitudeNEQ   *float64  `json:"latitudeNEQ,omitempty"`
	LatitudeIn    []float64 `json:"latitudeIn,omitempty"`
	LatitudeNotIn []float64 `json:"latitudeNotIn,omitempty"`
	LatitudeGT    *float64  `json:"latitudeGT,omitempty"`
	LatitudeGTE   *float64  `json:"latitudeGTE,omitempty"`
	LatitudeLT    *float64  `json:"latitudeLT,omitempty"`
	LatitudeLTE   *float64  `json:"latitudeLTE,omitempty"`

	// "longitude" field predicates.
	Longitude      *float64  `json:"longitude,omitempty"`
	LongitudeNEQ   *float64  `json:"longitudeNEQ,omitempty"`
	LongitudeIn    []float64 `json:"longitudeIn,omitempty"`
	LongitudeNotIn []float64 `json:"longitudeNotIn,omitempty"`
	LongitudeGT    *float64  `json:"longitudeGT,omitempty"`
	LongitudeGTE   *float64  `json:"longitudeGTE,omitempty"`
	LongitudeLT    *float64  `json:"longitudeLT,omitempty"`
	LongitudeLTE   *float64  `json:"longitudeLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "municipality" edge predicates.
	HasMunicipality     *bool                     `json:"hasMunicipality,omitempty"`
	HasMunicipalityWith []*MunicipalityWhereInput `json:"hasMunicipalityWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *MeshiWhereInput) AddPredicates(predicates ...predicate.Meshi) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the MeshiWhereInput filter on the MeshiQuery builder.
func (i *MeshiWhereInput) Filter(q *MeshiQuery) (*MeshiQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyMeshiWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyMeshiWhereInput is returned in case the MeshiWhereInput is empty.
var ErrEmptyMeshiWhereInput = errors.New("ent: empty predicate MeshiWhereInput")

// P returns a predicate for filtering meshis.
// An error is returned if the input is empty or invalid.
func (i *MeshiWhereInput) P() (predicate.Meshi, error) {
	var predicates []predicate.Meshi
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, meshi.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Meshi, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, meshi.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Meshi, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, meshi.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, meshi.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, meshi.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, meshi.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, meshi.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, meshi.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, meshi.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, meshi.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, meshi.IDLTE(*i.IDLTE))
	}
	if i.ArticleID != nil {
		predicates = append(predicates, meshi.ArticleIDEQ(*i.ArticleID))
	}
	if i.ArticleIDNEQ != nil {
		predicates = append(predicates, meshi.ArticleIDNEQ(*i.ArticleIDNEQ))
	}
	if len(i.ArticleIDIn) > 0 {
		predicates = append(predicates, meshi.ArticleIDIn(i.ArticleIDIn...))
	}
	if len(i.ArticleIDNotIn) > 0 {
		predicates = append(predicates, meshi.ArticleIDNotIn(i.ArticleIDNotIn...))
	}
	if i.ArticleIDGT != nil {
		predicates = append(predicates, meshi.ArticleIDGT(*i.ArticleIDGT))
	}
	if i.ArticleIDGTE != nil {
		predicates = append(predicates, meshi.ArticleIDGTE(*i.ArticleIDGTE))
	}
	if i.ArticleIDLT != nil {
		predicates = append(predicates, meshi.ArticleIDLT(*i.ArticleIDLT))
	}
	if i.ArticleIDLTE != nil {
		predicates = append(predicates, meshi.ArticleIDLTE(*i.ArticleIDLTE))
	}
	if i.ArticleIDContains != nil {
		predicates = append(predicates, meshi.ArticleIDContains(*i.ArticleIDContains))
	}
	if i.ArticleIDHasPrefix != nil {
		predicates = append(predicates, meshi.ArticleIDHasPrefix(*i.ArticleIDHasPrefix))
	}
	if i.ArticleIDHasSuffix != nil {
		predicates = append(predicates, meshi.ArticleIDHasSuffix(*i.ArticleIDHasSuffix))
	}
	if i.ArticleIDEqualFold != nil {
		predicates = append(predicates, meshi.ArticleIDEqualFold(*i.ArticleIDEqualFold))
	}
	if i.ArticleIDContainsFold != nil {
		predicates = append(predicates, meshi.ArticleIDContainsFold(*i.ArticleIDContainsFold))
	}
	if i.Title != nil {
		predicates = append(predicates, meshi.TitleEQ(*i.Title))
	}
	if i.TitleNEQ != nil {
		predicates = append(predicates, meshi.TitleNEQ(*i.TitleNEQ))
	}
	if len(i.TitleIn) > 0 {
		predicates = append(predicates, meshi.TitleIn(i.TitleIn...))
	}
	if len(i.TitleNotIn) > 0 {
		predicates = append(predicates, meshi.TitleNotIn(i.TitleNotIn...))
	}
	if i.TitleGT != nil {
		predicates = append(predicates, meshi.TitleGT(*i.TitleGT))
	}
	if i.TitleGTE != nil {
		predicates = append(predicates, meshi.TitleGTE(*i.TitleGTE))
	}
	if i.TitleLT != nil {
		predicates = append(predicates, meshi.TitleLT(*i.TitleLT))
	}
	if i.TitleLTE != nil {
		predicates = append(predicates, meshi.TitleLTE(*i.TitleLTE))
	}
	if i.TitleContains != nil {
		predicates = append(predicates, meshi.TitleContains(*i.TitleContains))
	}
	if i.TitleHasPrefix != nil {
		predicates = append(predicates, meshi.TitleHasPrefix(*i.TitleHasPrefix))
	}
	if i.TitleHasSuffix != nil {
		predicates = append(predicates, meshi.TitleHasSuffix(*i.TitleHasSuffix))
	}
	if i.TitleEqualFold != nil {
		predicates = append(predicates, meshi.TitleEqualFold(*i.TitleEqualFold))
	}
	if i.TitleContainsFold != nil {
		predicates = append(predicates, meshi.TitleContainsFold(*i.TitleContainsFold))
	}
	if i.ImageURL != nil {
		predicates = append(predicates, meshi.ImageURLEQ(*i.ImageURL))
	}
	if i.ImageURLNEQ != nil {
		predicates = append(predicates, meshi.ImageURLNEQ(*i.ImageURLNEQ))
	}
	if len(i.ImageURLIn) > 0 {
		predicates = append(predicates, meshi.ImageURLIn(i.ImageURLIn...))
	}
	if len(i.ImageURLNotIn) > 0 {
		predicates = append(predicates, meshi.ImageURLNotIn(i.ImageURLNotIn...))
	}
	if i.ImageURLGT != nil {
		predicates = append(predicates, meshi.ImageURLGT(*i.ImageURLGT))
	}
	if i.ImageURLGTE != nil {
		predicates = append(predicates, meshi.ImageURLGTE(*i.ImageURLGTE))
	}
	if i.ImageURLLT != nil {
		predicates = append(predicates, meshi.ImageURLLT(*i.ImageURLLT))
	}
	if i.ImageURLLTE != nil {
		predicates = append(predicates, meshi.ImageURLLTE(*i.ImageURLLTE))
	}
	if i.ImageURLContains != nil {
		predicates = append(predicates, meshi.ImageURLContains(*i.ImageURLContains))
	}
	if i.ImageURLHasPrefix != nil {
		predicates = append(predicates, meshi.ImageURLHasPrefix(*i.ImageURLHasPrefix))
	}
	if i.ImageURLHasSuffix != nil {
		predicates = append(predicates, meshi.ImageURLHasSuffix(*i.ImageURLHasSuffix))
	}
	if i.ImageURLEqualFold != nil {
		predicates = append(predicates, meshi.ImageURLEqualFold(*i.ImageURLEqualFold))
	}
	if i.ImageURLContainsFold != nil {
		predicates = append(predicates, meshi.ImageURLContainsFold(*i.ImageURLContainsFold))
	}
	if i.StoreName != nil {
		predicates = append(predicates, meshi.StoreNameEQ(*i.StoreName))
	}
	if i.StoreNameNEQ != nil {
		predicates = append(predicates, meshi.StoreNameNEQ(*i.StoreNameNEQ))
	}
	if len(i.StoreNameIn) > 0 {
		predicates = append(predicates, meshi.StoreNameIn(i.StoreNameIn...))
	}
	if len(i.StoreNameNotIn) > 0 {
		predicates = append(predicates, meshi.StoreNameNotIn(i.StoreNameNotIn...))
	}
	if i.StoreNameGT != nil {
		predicates = append(predicates, meshi.StoreNameGT(*i.StoreNameGT))
	}
	if i.StoreNameGTE != nil {
		predicates = append(predicates, meshi.StoreNameGTE(*i.StoreNameGTE))
	}
	if i.StoreNameLT != nil {
		predicates = append(predicates, meshi.StoreNameLT(*i.StoreNameLT))
	}
	if i.StoreNameLTE != nil {
		predicates = append(predicates, meshi.StoreNameLTE(*i.StoreNameLTE))
	}
	if i.StoreNameContains != nil {
		predicates = append(predicates, meshi.StoreNameContains(*i.StoreNameContains))
	}
	if i.StoreNameHasPrefix != nil {
		predicates = append(predicates, meshi.StoreNameHasPrefix(*i.StoreNameHasPrefix))
	}
	if i.StoreNameHasSuffix != nil {
		predicates = append(predicates, meshi.StoreNameHasSuffix(*i.StoreNameHasSuffix))
	}
	if i.StoreNameEqualFold != nil {
		predicates = append(predicates, meshi.StoreNameEqualFold(*i.StoreNameEqualFold))
	}
	if i.StoreNameContainsFold != nil {
		predicates = append(predicates, meshi.StoreNameContainsFold(*i.StoreNameContainsFold))
	}
	if i.Address != nil {
		predicates = append(predicates, meshi.AddressEQ(*i.Address))
	}
	if i.AddressNEQ != nil {
		predicates = append(predicates, meshi.AddressNEQ(*i.AddressNEQ))
	}
	if len(i.AddressIn) > 0 {
		predicates = append(predicates, meshi.AddressIn(i.AddressIn...))
	}
	if len(i.AddressNotIn) > 0 {
		predicates = append(predicates, meshi.AddressNotIn(i.AddressNotIn...))
	}
	if i.AddressGT != nil {
		predicates = append(predicates, meshi.AddressGT(*i.AddressGT))
	}
	if i.AddressGTE != nil {
		predicates = append(predicates, meshi.AddressGTE(*i.AddressGTE))
	}
	if i.AddressLT != nil {
		predicates = append(predicates, meshi.AddressLT(*i.AddressLT))
	}
	if i.AddressLTE != nil {
		predicates = append(predicates, meshi.AddressLTE(*i.AddressLTE))
	}
	if i.AddressContains != nil {
		predicates = append(predicates, meshi.AddressContains(*i.AddressContains))
	}
	if i.AddressHasPrefix != nil {
		predicates = append(predicates, meshi.AddressHasPrefix(*i.AddressHasPrefix))
	}
	if i.AddressHasSuffix != nil {
		predicates = append(predicates, meshi.AddressHasSuffix(*i.AddressHasSuffix))
	}
	if i.AddressEqualFold != nil {
		predicates = append(predicates, meshi.AddressEqualFold(*i.AddressEqualFold))
	}
	if i.AddressContainsFold != nil {
		predicates = append(predicates, meshi.AddressContainsFold(*i.AddressContainsFold))
	}
	if i.SiteURL != nil {
		predicates = append(predicates, meshi.SiteURLEQ(*i.SiteURL))
	}
	if i.SiteURLNEQ != nil {
		predicates = append(predicates, meshi.SiteURLNEQ(*i.SiteURLNEQ))
	}
	if len(i.SiteURLIn) > 0 {
		predicates = append(predicates, meshi.SiteURLIn(i.SiteURLIn...))
	}
	if len(i.SiteURLNotIn) > 0 {
		predicates = append(predicates, meshi.SiteURLNotIn(i.SiteURLNotIn...))
	}
	if i.SiteURLGT != nil {
		predicates = append(predicates, meshi.SiteURLGT(*i.SiteURLGT))
	}
	if i.SiteURLGTE != nil {
		predicates = append(predicates, meshi.SiteURLGTE(*i.SiteURLGTE))
	}
	if i.SiteURLLT != nil {
		predicates = append(predicates, meshi.SiteURLLT(*i.SiteURLLT))
	}
	if i.SiteURLLTE != nil {
		predicates = append(predicates, meshi.SiteURLLTE(*i.SiteURLLTE))
	}
	if i.SiteURLContains != nil {
		predicates = append(predicates, meshi.SiteURLContains(*i.SiteURLContains))
	}
	if i.SiteURLHasPrefix != nil {
		predicates = append(predicates, meshi.SiteURLHasPrefix(*i.SiteURLHasPrefix))
	}
	if i.SiteURLHasSuffix != nil {
		predicates = append(predicates, meshi.SiteURLHasSuffix(*i.SiteURLHasSuffix))
	}
	if i.SiteURLEqualFold != nil {
		predicates = append(predicates, meshi.SiteURLEqualFold(*i.SiteURLEqualFold))
	}
	if i.SiteURLContainsFold != nil {
		predicates = append(predicates, meshi.SiteURLContainsFold(*i.SiteURLContainsFold))
	}
	if i.PublishedDate != nil {
		predicates = append(predicates, meshi.PublishedDateEQ(*i.PublishedDate))
	}
	if i.PublishedDateNEQ != nil {
		predicates = append(predicates, meshi.PublishedDateNEQ(*i.PublishedDateNEQ))
	}
	if len(i.PublishedDateIn) > 0 {
		predicates = append(predicates, meshi.PublishedDateIn(i.PublishedDateIn...))
	}
	if len(i.PublishedDateNotIn) > 0 {
		predicates = append(predicates, meshi.PublishedDateNotIn(i.PublishedDateNotIn...))
	}
	if i.PublishedDateGT != nil {
		predicates = append(predicates, meshi.PublishedDateGT(*i.PublishedDateGT))
	}
	if i.PublishedDateGTE != nil {
		predicates = append(predicates, meshi.PublishedDateGTE(*i.PublishedDateGTE))
	}
	if i.PublishedDateLT != nil {
		predicates = append(predicates, meshi.PublishedDateLT(*i.PublishedDateLT))
	}
	if i.PublishedDateLTE != nil {
		predicates = append(predicates, meshi.PublishedDateLTE(*i.PublishedDateLTE))
	}
	if i.Latitude != nil {
		predicates = append(predicates, meshi.LatitudeEQ(*i.Latitude))
	}
	if i.LatitudeNEQ != nil {
		predicates = append(predicates, meshi.LatitudeNEQ(*i.LatitudeNEQ))
	}
	if len(i.LatitudeIn) > 0 {
		predicates = append(predicates, meshi.LatitudeIn(i.LatitudeIn...))
	}
	if len(i.LatitudeNotIn) > 0 {
		predicates = append(predicates, meshi.LatitudeNotIn(i.LatitudeNotIn...))
	}
	if i.LatitudeGT != nil {
		predicates = append(predicates, meshi.LatitudeGT(*i.LatitudeGT))
	}
	if i.LatitudeGTE != nil {
		predicates = append(predicates, meshi.LatitudeGTE(*i.LatitudeGTE))
	}
	if i.LatitudeLT != nil {
		predicates = append(predicates, meshi.LatitudeLT(*i.LatitudeLT))
	}
	if i.LatitudeLTE != nil {
		predicates = append(predicates, meshi.LatitudeLTE(*i.LatitudeLTE))
	}
	if i.Longitude != nil {
		predicates = append(predicates, meshi.LongitudeEQ(*i.Longitude))
	}
	if i.LongitudeNEQ != nil {
		predicates = append(predicates, meshi.LongitudeNEQ(*i.LongitudeNEQ))
	}
	if len(i.LongitudeIn) > 0 {
		predicates = append(predicates, meshi.LongitudeIn(i.LongitudeIn...))
	}
	if len(i.LongitudeNotIn) > 0 {
		predicates = append(predicates, meshi.LongitudeNotIn(i.LongitudeNotIn...))
	}
	if i.LongitudeGT != nil {
		predicates = append(predicates, meshi.LongitudeGT(*i.LongitudeGT))
	}
	if i.LongitudeGTE != nil {
		predicates = append(predicates, meshi.LongitudeGTE(*i.LongitudeGTE))
	}
	if i.LongitudeLT != nil {
		predicates = append(predicates, meshi.LongitudeLT(*i.LongitudeLT))
	}
	if i.LongitudeLTE != nil {
		predicates = append(predicates, meshi.LongitudeLTE(*i.LongitudeLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, meshi.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, meshi.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, meshi.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, meshi.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, meshi.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, meshi.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, meshi.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, meshi.CreatedAtLTE(*i.CreatedAtLTE))
	}

	if i.HasMunicipality != nil {
		p := meshi.HasMunicipality()
		if !*i.HasMunicipality {
			p = meshi.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasMunicipalityWith) > 0 {
		with := make([]predicate.Municipality, 0, len(i.HasMunicipalityWith))
		for _, w := range i.HasMunicipalityWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasMunicipalityWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, meshi.HasMunicipalityWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyMeshiWhereInput
	case 1:
		return predicates[0], nil
	default:
		return meshi.And(predicates...), nil
	}
}

// MunicipalityWhereInput represents a where input for filtering Municipality queries.
type MunicipalityWhereInput struct {
	Predicates []predicate.Municipality  `json:"-"`
	Not        *MunicipalityWhereInput   `json:"not,omitempty"`
	Or         []*MunicipalityWhereInput `json:"or,omitempty"`
	And        []*MunicipalityWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "zipcode" field predicates.
	Zipcode             *string  `json:"zipcode,omitempty"`
	ZipcodeNEQ          *string  `json:"zipcodeNEQ,omitempty"`
	ZipcodeIn           []string `json:"zipcodeIn,omitempty"`
	ZipcodeNotIn        []string `json:"zipcodeNotIn,omitempty"`
	ZipcodeGT           *string  `json:"zipcodeGT,omitempty"`
	ZipcodeGTE          *string  `json:"zipcodeGTE,omitempty"`
	ZipcodeLT           *string  `json:"zipcodeLT,omitempty"`
	ZipcodeLTE          *string  `json:"zipcodeLTE,omitempty"`
	ZipcodeContains     *string  `json:"zipcodeContains,omitempty"`
	ZipcodeHasPrefix    *string  `json:"zipcodeHasPrefix,omitempty"`
	ZipcodeHasSuffix    *string  `json:"zipcodeHasSuffix,omitempty"`
	ZipcodeIsNil        bool     `json:"zipcodeIsNil,omitempty"`
	ZipcodeNotNil       bool     `json:"zipcodeNotNil,omitempty"`
	ZipcodeEqualFold    *string  `json:"zipcodeEqualFold,omitempty"`
	ZipcodeContainsFold *string  `json:"zipcodeContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "meshis" edge predicates.
	HasMeshis     *bool              `json:"hasMeshis,omitempty"`
	HasMeshisWith []*MeshiWhereInput `json:"hasMeshisWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *MunicipalityWhereInput) AddPredicates(predicates ...predicate.Municipality) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the MunicipalityWhereInput filter on the MunicipalityQuery builder.
func (i *MunicipalityWhereInput) Filter(q *MunicipalityQuery) (*MunicipalityQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyMunicipalityWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyMunicipalityWhereInput is returned in case the MunicipalityWhereInput is empty.
var ErrEmptyMunicipalityWhereInput = errors.New("ent: empty predicate MunicipalityWhereInput")

// P returns a predicate for filtering municipalities.
// An error is returned if the input is empty or invalid.
func (i *MunicipalityWhereInput) P() (predicate.Municipality, error) {
	var predicates []predicate.Municipality
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, municipality.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Municipality, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, municipality.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Municipality, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, municipality.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, municipality.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, municipality.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, municipality.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, municipality.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, municipality.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, municipality.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, municipality.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, municipality.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, municipality.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, municipality.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, municipality.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, municipality.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, municipality.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, municipality.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, municipality.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, municipality.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, municipality.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, municipality.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, municipality.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, municipality.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, municipality.NameContainsFold(*i.NameContainsFold))
	}
	if i.Zipcode != nil {
		predicates = append(predicates, municipality.ZipcodeEQ(*i.Zipcode))
	}
	if i.ZipcodeNEQ != nil {
		predicates = append(predicates, municipality.ZipcodeNEQ(*i.ZipcodeNEQ))
	}
	if len(i.ZipcodeIn) > 0 {
		predicates = append(predicates, municipality.ZipcodeIn(i.ZipcodeIn...))
	}
	if len(i.ZipcodeNotIn) > 0 {
		predicates = append(predicates, municipality.ZipcodeNotIn(i.ZipcodeNotIn...))
	}
	if i.ZipcodeGT != nil {
		predicates = append(predicates, municipality.ZipcodeGT(*i.ZipcodeGT))
	}
	if i.ZipcodeGTE != nil {
		predicates = append(predicates, municipality.ZipcodeGTE(*i.ZipcodeGTE))
	}
	if i.ZipcodeLT != nil {
		predicates = append(predicates, municipality.ZipcodeLT(*i.ZipcodeLT))
	}
	if i.ZipcodeLTE != nil {
		predicates = append(predicates, municipality.ZipcodeLTE(*i.ZipcodeLTE))
	}
	if i.ZipcodeContains != nil {
		predicates = append(predicates, municipality.ZipcodeContains(*i.ZipcodeContains))
	}
	if i.ZipcodeHasPrefix != nil {
		predicates = append(predicates, municipality.ZipcodeHasPrefix(*i.ZipcodeHasPrefix))
	}
	if i.ZipcodeHasSuffix != nil {
		predicates = append(predicates, municipality.ZipcodeHasSuffix(*i.ZipcodeHasSuffix))
	}
	if i.ZipcodeIsNil {
		predicates = append(predicates, municipality.ZipcodeIsNil())
	}
	if i.ZipcodeNotNil {
		predicates = append(predicates, municipality.ZipcodeNotNil())
	}
	if i.ZipcodeEqualFold != nil {
		predicates = append(predicates, municipality.ZipcodeEqualFold(*i.ZipcodeEqualFold))
	}
	if i.ZipcodeContainsFold != nil {
		predicates = append(predicates, municipality.ZipcodeContainsFold(*i.ZipcodeContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, municipality.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, municipality.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, municipality.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, municipality.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, municipality.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, municipality.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, municipality.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, municipality.CreatedAtLTE(*i.CreatedAtLTE))
	}

	if i.HasMeshis != nil {
		p := municipality.HasMeshis()
		if !*i.HasMeshis {
			p = municipality.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasMeshisWith) > 0 {
		with := make([]predicate.Meshi, 0, len(i.HasMeshisWith))
		for _, w := range i.HasMeshisWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasMeshisWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, municipality.HasMeshisWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyMunicipalityWhereInput
	case 1:
		return predicates[0], nil
	default:
		return municipality.And(predicates...), nil
	}
}
