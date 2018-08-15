package tool

import "github.com/sony/sonyflake"

func CreateUUID() (uint64, error) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	uuid, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return uuid, nil
}
