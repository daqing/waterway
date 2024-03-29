package action_api

import (
	"fmt"

	"github.com/daqing/waterway/lib/repo"
)

// relation 被关注/收藏/点赞的对象
// userId 谁发起了这个动作
func ToggleAction(userId repo.IdType, action string, relation repo.PolyModel) (int64, error) {
	var attrs = []repo.KVPair{
		repo.KV("user_id", userId),
		repo.KV("action", action),
		repo.KV("target_id", relation.PolyId()),
		repo.KV("target_type", relation.PolyType()),
	}

	row, err := repo.FindRow[Action]([]string{"id"}, attrs)

	if err != nil {
		return repo.InvalidCount, err
	}

	if row == nil {
		// current  action not found, create a new one
		row, err = repo.Insert[Action](attrs)
		if err != nil {
			return repo.InvalidCount, err
		}

		if row == nil {
			// record not created
			return repo.InvalidCount, fmt.Errorf("action was not created")
		}
	} else {
		// delete current like action
		err = repo.Delete[Action](attrs)
		if err != nil {
			return repo.InvalidCount, err
		}
	}

	count, err := repo.Count[Action]([]repo.KVPair{
		repo.KV("action", action),
		repo.KV("target_id", relation.PolyId()),
		repo.KV("target_type", relation.PolyType()),
	})

	if err != nil {
		return repo.InvalidCount, err
	}

	return count, nil

}
