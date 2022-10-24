package mysql

import (
	"bullbeckt/models"
	"database/sql"
	"go.uber.org/zap"
)

func CommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return

}

// 根据id查询
func CommunityWtihID(id int64) (community *models.CommmunityDatail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	community = new(models.CommmunityDatail)
	if err = db.Get(community, sqlStr, id); err == sql.ErrNoRows {
		err = ErrorInvalidID
	}
	return community, err
}
