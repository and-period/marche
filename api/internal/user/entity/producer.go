package entity

import (
	"time"

	"github.com/and-period/furumaru/api/internal/codes"
	"github.com/and-period/furumaru/api/internal/common"
	"github.com/and-period/furumaru/api/pkg/set"
	"gorm.io/datatypes"
)

// Producer - 生産者情報
type Producer struct {
	Admin             `gorm:"-"`
	AdminID           string         `gorm:"primaryKey;<-:create"`           // 管理者ID
	CoordinatorID     string         `gorm:"default:null"`                   // コーディネータID
	PhoneNumber       string         `gorm:"default:null"`                   // 電話番号
	Username          string         `gorm:""`                               // 表示名
	Profile           string         `gorm:""`                               // 紹介文
	ThumbnailURL      string         `gorm:""`                               // サムネイルURL
	Thumbnails        common.Images  `gorm:"-"`                              // サムネイル一覧(リサイズ済み)
	ThumbnailsJSON    datatypes.JSON `gorm:"default:null;column:thumbnails"` // サムネイル一覧(JSON)
	HeaderURL         string         `gorm:""`                               // ヘッダー画像URL
	Headers           common.Images  `gorm:"-"`                              // ヘッダー画像一覧(リサイズ済み)
	HeadersJSON       datatypes.JSON `gorm:"default:null;column:headers"`    // ヘッダー画像一覧(JSON)
	PromotionVideoURL string         `gorm:""`                               // 紹介動画URL
	BonusVideoURL     string         `gorm:""`                               // 購入特典動画URL
	InstagramID       string         `gorm:""`                               // SNS(Instagram)アカウント名
	FacebookID        string         `gorm:""`                               // SNS(Facebook)アカウント名
	PostalCode        string         `gorm:"default:null"`                   // 郵便番号
	Prefecture        string         `gorm:"-"`                              // 都道府県
	PrefectureCode    int32          `gorm:"default:null;column:prefecture"` // 都道府県コード
	City              string         `gorm:"default:null"`                   // 市区町村
	AddressLine1      string         `gorm:"default:null"`                   // 町名・番地
	AddressLine2      string         `gorm:"default:null"`                   // ビル名・号室など
	CreatedAt         time.Time      `gorm:"<-:create"`                      // 登録日時
	UpdatedAt         time.Time      `gorm:""`                               // 更新日時
}

type Producers []*Producer

type NewProducerParams struct {
	Admin             *Admin
	CoordinatorID     string
	PhoneNumber       string
	Username          string
	Profile           string
	ThumbnailURL      string
	HeaderURL         string
	PromotionVideoURL string
	BonusVideoURL     string
	InstagramID       string
	FacebookID        string
	PostalCode        string
	PrefectureCode    int32
	City              string
	AddressLine1      string
	AddressLine2      string
}

func NewProducer(params *NewProducerParams) (*Producer, error) {
	producer := &Producer{
		AdminID:           params.Admin.ID,
		CoordinatorID:     params.CoordinatorID,
		PhoneNumber:       params.PhoneNumber,
		Username:          params.Username,
		Profile:           params.Profile,
		ThumbnailURL:      params.ThumbnailURL,
		HeaderURL:         params.HeaderURL,
		PromotionVideoURL: params.PromotionVideoURL,
		BonusVideoURL:     params.BonusVideoURL,
		InstagramID:       params.InstagramID,
		FacebookID:        params.FacebookID,
		PostalCode:        params.PostalCode,
		PrefectureCode:    params.PrefectureCode,
		City:              params.City,
		AddressLine1:      params.AddressLine1,
		AddressLine2:      params.AddressLine2,
		Admin:             *params.Admin,
	}
	if params.PrefectureCode == 0 {
		return producer, nil
	}
	var err error
	producer.Prefecture, err = codes.ToPrefectureJapanese(params.PrefectureCode)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (p *Producer) Fill(admin *Admin) (err error) {
	p.Thumbnails, err = common.NewImagesFromBytes(p.ThumbnailsJSON)
	if err != nil {
		return err
	}
	p.Headers, err = common.NewImagesFromBytes(p.HeadersJSON)
	if err != nil {
		return err
	}
	p.Admin = *admin
	p.Prefecture, _ = codes.ToPrefectureJapanese(p.PrefectureCode)
	return nil
}

func (ps Producers) IDs() []string {
	res := make([]string, len(ps))
	for i := range ps {
		res[i] = ps[i].AdminID
	}
	return res
}

func (ps Producers) CoordinatorIDs() []string {
	return set.UniqBy(ps, func(p *Producer) string {
		return p.CoordinatorID
	})
}

func (ps Producers) Unrelated() Producers {
	res := make(Producers, 0, len(ps))
	for _, p := range ps {
		if p.CoordinatorID != "" {
			continue
		}
		res = append(res, p)
	}
	return res
}

func (ps Producers) Fill(admins map[string]*Admin) error {
	for _, p := range ps {
		admin, ok := admins[p.AdminID]
		if !ok {
			admin = &Admin{ID: p.AdminID, Role: AdminRoleProducer}
		}
		if err := p.Fill(admin); err != nil {
			return err
		}
	}
	return nil
}
