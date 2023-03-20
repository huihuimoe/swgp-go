package jsonhelper

import (
	"os"

	"muzzammil.xyz/jsonc"
)

func LoadAndDecode2(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return jsonc.Unmarshal(b, v)
}
