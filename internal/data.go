package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlRepo struct {
	db *gorm.DB
}

var _ Repo = (*MysqlRepo)(nil)

func NewMysqlRepo(i *do.Injector) (*MysqlRepo, error) {
	config := do.MustInvoke[*Config](i)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard,

		DisableNestedTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&mysqlTimeZoneInfo{}); err != nil {
		return nil, err
	}

	return &MysqlRepo{db: db}, nil
}

type mysqlTimeZoneInfo struct {
	ID string `gorm:"column:ID;type:varchar(26);primarykey"`

	TimeZone     string    `gorm:"column:TIME_ZONE;type:varchar(50);not null;index:idx_time_zone_time_zone"`
	DisplayName  string    `gorm:"column:DISPLAY_NAME;type:varchar(100);not null"`
	Abbreviation string    `gorm:"column:ABBREVIATION;type:varchar(6);not null"`
	StartTime    time.Time `gorm:"column:START_TIME;type:datetime;not null;index:idx_time_zone_start_time"`
	EndTime      time.Time `gorm:"column:END_TIME;type:datetime;not null"`
	Offset       int32     `gorm:"column:OFFSET;type:int;not null"`
	IsDST        int8      `gorm:"column:IS_DST;type:tinyint;not null"`
	Source       string    `gorm:"column:SOURCE;type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t mysqlTimeZoneInfo) TableName() string {
	return "time_zone"
}

func (t mysqlTimeZoneInfo) toTimeZone() *TimeZone {
	return &TimeZone{
		ID:           t.ID,
		Name:         t.TimeZone,
		DisplayName:  t.DisplayName,
		Abbreviation: t.Abbreviation,
		StartTime:    TimestampToTimeString(t.StartTime.Unix(), int(t.Offset)),
		EndTime:      TimestampToTimeString(t.EndTime.Unix(), int(t.Offset)),
		IsDST:        t.IsDST == 1,
		Offset:       int(t.Offset),
		Source:       t.Source,
	}
}

func (r *MysqlRepo) CreateTimeZone(ctx context.Context, tz *TimeZone) (*TimeZone, error) {
	tx := r.db

	model, err := r.timeZone2MysqlTimeZoneInfo(tz)
	if err != nil {
		return nil, err
	}
	if err := tx.Create(model).Error; err != nil {
		return nil, err
	}
	return model.toTimeZone(), nil
}

func (r *MysqlRepo) UpdateTimeZone(ctx context.Context, tz *TimeZone, fields []string) (*TimeZone, error) {
	tx := r.db

	model := new(mysqlTimeZoneInfo)
	if err := tx.Where("ID = ?", tz.ID).First(model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("timezone %s not found: %w", tz.ID, err)
		}
		return nil, err
	}

	if lo.Contains(fields, "name") {
		model.TimeZone = tz.Name
	}
	if lo.Contains(fields, "displayName") {
		model.DisplayName = tz.DisplayName
	}
	if lo.Contains(fields, "abbreviation") {
		model.Abbreviation = tz.Abbreviation
	}
	if lo.Contains(fields, "offset") {
		model.Offset = int32(tz.Offset) // 优先更新 offset
	}
	if lo.Contains(fields, "startTime") {
		sec, err := TimeString2Timestamp(tz.StartTime, int(model.Offset))
		if err != nil {
			return nil, err
		}
		model.StartTime = time.Unix(sec, 0)
	}
	if lo.Contains(fields, "endTime") {
		sec, err := TimeString2Timestamp(tz.EndTime, int(model.Offset))
		if err != nil {
			return nil, err
		}
		model.EndTime = time.Unix(sec, 0)
	}
	if lo.Contains(fields, "isDST") {
		model.IsDST = cast.ToInt8(tz.IsDST)
	}
	if err := tx.Updates(model).Error; err != nil {
		return nil, err
	}
	return model.toTimeZone(), nil
}

func (r *MysqlRepo) DeleteTimeZone(ctx context.Context, id string) error {
	tx := r.db

	return tx.Where("ID = ?", id).Delete(&mysqlTimeZoneInfo{}).Error
}

func (r *MysqlRepo) ListTimeZones(ctx context.Context) ([]*TimeZone, error) {
	tx := r.db

	var res []*mysqlTimeZoneInfo
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return lo.Map(res, func(item *mysqlTimeZoneInfo, _ int) *TimeZone {
		return item.toTimeZone()
	}), nil
}

func (r *MysqlRepo) ListTimeZonesByTimestamp(ctx context.Context, sec int64) ([]*TimeZone, error) {
	tx := r.db

	t := time.Unix(sec, 0)
	var res []*mysqlTimeZoneInfo
	if err := tx.Where("START_TIME <= ?", t).Where("END_TIME > ?", t).Find(&res).Error; err != nil {
		return nil, err
	}

	return lo.Map(res, func(item *mysqlTimeZoneInfo, _ int) *TimeZone {
		return item.toTimeZone()
	}), nil
}

func (r *MysqlRepo) Count(ctx context.Context) (int64, error) {
	tx := r.db

	var count int64
	if err := tx.Model(&mysqlTimeZoneInfo{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *MysqlRepo) ReplaceAllBySource(ctx context.Context, tzs []*TimeZone, source string) error {
	tx := r.db

	return tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("SOURCE = ?", source).Delete(&mysqlTimeZoneInfo{}).Error; err != nil {
			return err
		}
		models, err := r.timeZones2MysqlTimeZoneInfos(tzs)
		if err != nil {
			return err
		}
		return tx.CreateInBatches(models, 100).Error
	})
}

func (r *MysqlRepo) timeZone2MysqlTimeZoneInfo(tz *TimeZone) (*mysqlTimeZoneInfo, error) {
	startTime, err := TimeString2Timestamp(tz.StartTime, tz.Offset)
	if err != nil {
		return nil, err
	}
	endTime, err := TimeString2Timestamp(tz.EndTime, tz.Offset)
	if err != nil {
		return nil, err
	}
	return &mysqlTimeZoneInfo{
		ID:           tz.ID,
		TimeZone:     tz.Name,
		DisplayName:  tz.DisplayName,
		Abbreviation: tz.Abbreviation,
		StartTime:    time.Unix(startTime, 0),
		EndTime:      time.Unix(endTime, 0),
		Offset:       int32(tz.Offset),
		IsDST:        cast.ToInt8(tz.IsDST),
		Source:       tz.Source,
	}, nil
}

func (r *MysqlRepo) timeZones2MysqlTimeZoneInfos(tzs []*TimeZone) ([]*mysqlTimeZoneInfo, error) {
	var res []*mysqlTimeZoneInfo
	for _, tz := range tzs {
		model, err := r.timeZone2MysqlTimeZoneInfo(tz)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}
	return res, nil
}
