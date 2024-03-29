package membership_api

import "github.com/daqing/waterway/lib/repo"

type MembershipResp struct {
	Name      string
	ExpiredAt repo.Timestamp
}

func (r MembershipResp) Fields() []string {
	return []string{"name", "expired_at"}
}
