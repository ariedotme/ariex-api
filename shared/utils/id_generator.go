package utils

import "github.com/aidarkhanov/nanoid/v2"

func GenerateID() string {
	id, _ := nanoid.New()
	return id
}
