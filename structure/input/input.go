package input

import rl "github.com/gen2brain/raylib-go/raylib"

type KeyInterest struct {
	Pressed  bool // The key has been pressed once
	Released bool // The key has been released once
	Down     bool // The key is being pressed
	Up       bool // The key is *NOT* being pressed
}

type InputContext struct {
	List map[int32]KeyInterest
}

// Set a new KeyInterest entry for input context
func (k *InputContext) Set(rlKey int32) {
	k.List[rlKey] = KeyInterest{
		Pressed:  rl.IsKeyPressed(rlKey),
		Released: rl.IsKeyReleased(rlKey),
		Down:     rl.IsKeyDown(rlKey),
		Up:       rl.IsKeyUp(rlKey),
	}
}

// Set a list of rlKeys as KeyInterest into InputContext
func (k *InputContext) BulkSet(rlKeys []int32) {
	for _, rlKey := range rlKeys {
		k.List[rlKey] = KeyInterest{
			Pressed:  rl.IsKeyPressed(rlKey),
			Released: rl.IsKeyReleased(rlKey),
			Down:     rl.IsKeyDown(rlKey),
			Up:       rl.IsKeyUp(rlKey),
		}
	}
}

// Remove a KeyInterest from InputContext by rlKey
func (k *InputContext) Delete(rlKey int32) {
	delete(k.List, rlKey)
}

// Update InputContext keys mapping, should be called on Run()
func (k *InputContext) Update() {
	for key := range k.List {
		k.List[key] = KeyInterest{
			Pressed:  rl.IsKeyPressed(key),
			Released: rl.IsKeyReleased(key),
			Down:     rl.IsKeyDown(key),
			Up:       rl.IsKeyUp(key),
		}
	}
}

// Return the KeyInterest by its rlKey if is present in InputContext,
// if not, return false
func (k *InputContext) GetKey(rlKey int32) KeyInterest {
	return k.List[rlKey]
}
