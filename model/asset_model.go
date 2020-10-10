package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Asset{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Asset struct {
	AssetId  string          `json:"asset_id,omitempty" gorm:"type:varchar(36);primary_key;not null;"`
	Symbol   string          `json:"symbol,omitempty" gorm:"type:varchar(512);not null"`
	Name     string          `json:"name,omitempty" gorm:"type:varchar(512);not null"`
	IconUrl  string          `json:"icon_url,omitempty" gorm:"type:varchar(1024);not null"`
	PriceBtc decimal.Decimal `json:"price_btc,omitempty" gorm:"type:varchar(128);not null"`
	PriceUsd decimal.Decimal `json:"price_usd,omitempty" gorm:"type:varchar(128);not null"`
}

type AssetIdToPriceUsd struct {
	AssetId  string          `json:"asset_id,omitempty" gorm:"type:varchar(36);primary_key;not null;"`
	PriceUsd decimal.Decimal `json:"price_usd,omitempty" gorm:"type:varchar(128);not null"`
}

var ASSET *Asset

func (asset *Asset) UpdateAsset(assetData *Asset) (err error) {
	err = db.Save(assetData).Error
	return
}

func (asset *Asset) GetAssetById(assetId string) (assetData *Asset, err error) {
	assetData = &Asset{}
	err = db.Where("asset_id=?", assetId).Find(assetData).Error
	return
}

func (asset *Asset) ListAssetsAllByDB() (assets *[]Asset, err error) {
	assets = &[]Asset{}
	err = db.Find(assets).Error
	return
}

func (asset *Asset) GetPriceUsdByAssetId(assetId string) (priceUsd *AssetIdToPriceUsd, err error) {
	priceUsd = &AssetIdToPriceUsd{}
	err = db.Debug().Table("asset").Select("asset_id,price_usd").Where("asset_id=?", assetId).Scan(priceUsd).Error
	return
}

//
//var (
//	// Profits represent Profit Store implementation
//	Assets = &assetStore{}
//)
//
//type assetStore struct {
//	tx *gorm.DB
//}
//
//func (s *assetStore) DB() *gorm.DB {
//	if s.tx != nil {
//		return s.tx
//	}
//
//	return db
//}
//
//func (s *assetStore) WithTx(tx *gorm.DB) *assetStore {
//	return &assetStore{tx: tx}
//}
//
//func (s *assetStore) FindByPks(ids []string) ([]*Asset, error) {
//	var assets []*Asset
//	if result := db.Where("asset_id in (?)", ids).Find(&assets); result.Error != nil {
//		return nil, result.Error
//	}
//	return assets, nil
//}
//
//// func (s *assetStore) FindByPk(id string) (*Asset, error) {
//// 	return FindAssetByPk(id)
//// }
//
//func (s *assetStore) FindByPk(id string) (*Asset, error) {
//	var asset Asset
//	if result := db.Where("asset_id = ?", id).First(&asset); result.Error != nil {
//		if gorm.IsRecordNotFoundError(result.Error) {
//			return nil, nil
//		} else {
//			return nil, result.Error
//		}
//	}
//
//	return &asset, nil
//}
//
//func (s *assetStore) MapAll() (map[string]*Asset, error) {
//	var assets []*Asset
//	if err := db.Find(&assets).Error; err != nil {
//		return map[string]*Asset{}, err
//	}
//
//	m := make(map[string]*Asset, len(assets))
//	for _, asset := range assets {
//		m[asset.AssetID] = asset
//	}
//
//	return m, nil
//}
//
//func (s *assetStore) ListAll(priceOnly bool) (assets []*Asset, err error) {
//	if priceOnly {
//		if err = db.Where("price_btc <> '0'").Order("price_btc DESC").Find(&assets).Error; err != nil {
//			return
//		}
//	} else {
//		if err = db.Order("price_btc DESC").Find(&assets).Error; err != nil {
//			return
//		}
//	}
//	return
//}
//
//func (s *assetStore) ReadCnyPrices() (map[string]decimal.Decimal, error) {
//	var (
//		assets []*Asset
//		prices = map[string]decimal.Decimal{}
//	)
//
//	if err := db.Select("asset_id, price_cny").Find(&assets).Error; err != nil {
//		return prices, err
//	}
//
//	for _, asset := range assets {
//		prices[asset.AssetID] = asset.PriceCny
//	}
//
//	return prices, nil
//}
//
//func (s *assetStore) BulkCreateAndUpdates(assets []Asset) error {
//	var (
//		valueStrings []string
//		valueArgs    []interface{}
//	)
//
//	for _, item := range assets {
//		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
//
//		valueArgs = append(valueArgs, item.AssetID)
//		valueArgs = append(valueArgs, item.Symbol)
//		valueArgs = append(valueArgs, item.Name)
//		valueArgs = append(valueArgs, item.IconURL)
//		valueArgs = append(valueArgs, item.PriceBtc)
//		valueArgs = append(valueArgs, item.PriceUsd)
//	}
//
//	smt := `INSERT INTO assets(asset_id, symbol, name, icon_url, price_btc, price_usd) VALUES %s ON CONFLICT (asset_id) DO UPDATE SET (icon_url,price_btc,price_usd)=(EXCLUDED.icon_url,EXCLUDED.price_btc,EXCLUDED.price_usd)`
//
//	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
//
//	tx := db.Begin()
//	if err := tx.Exec(smt, valueArgs...).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
//
//func (asset *Asset) Update() error {
//	updates := map[string]interface{}{
//		"symbol":    asset.Symbol,
//		"name":      asset.Name,
//		"icon_url":  asset.IconURL,
//		"chain_id":  asset.ChainID,
//		"price_btc": asset.PriceBtc,
//		"price_usd": asset.PriceUsd,
//		"price_cny": asset.PriceCny,
//	}
//
//	return db.Model(asset).Updates(updates).Error
//}
//