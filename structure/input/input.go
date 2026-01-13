package input

import rl "github.com/gen2brain/raylib-go/raylib"

type KeyInterest struct {
	Pressed  bool
	Released bool
	Down     bool
	Up       bool
}

type Keys struct {
	List map[int32]KeyInterest
}

func (k *Keys) Set(rlKey int32) {
	k.List[rlKey] = KeyInterest{
		Pressed:  rl.IsKeyPressed(rlKey),
		Released: rl.IsKeyReleased(rlKey),
		Down:     rl.IsKeyDown(rlKey),
		Up:       rl.IsKeyUp(rlKey),
	}
}

func (k *Keys) BulkSet(rlKeys []int32) {
	for _, rlKey := range rlKeys {
		k.List[rlKey] = KeyInterest{
			Pressed:  rl.IsKeyPressed(rlKey),
			Released: rl.IsKeyReleased(rlKey),
			Down:     rl.IsKeyDown(rlKey),
			Up:       rl.IsKeyUp(rlKey),
		}
	}
}

func (k *Keys) Delete(rlKey int32) {
	delete(k.List, rlKey)
}

func (k *Keys) Update() {
	for key := range k.List {
		k.List[key] = KeyInterest{
			Pressed:  rl.IsKeyPressed(key),
			Released: rl.IsKeyReleased(key),
			Down:     rl.IsKeyDown(key),
			Up:       rl.IsKeyUp(key),
		}
	}
}

func (k *Keys) GetKey(rlKey int32) KeyInterest {
	return k.List[rlKey]
}
