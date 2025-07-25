package dto

import (
	Interface "main/web/app/types"
)

type WebConfig struct {
	Config Interface.WebConfig `inject:"webconfig"`
}
