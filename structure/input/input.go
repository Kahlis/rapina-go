package input

import rl "github.com/gen2brain/raylib-go/raylib"

type Key struct {
	Pressed  bool // Key has been pressed once
	Released bool // Key has been released once
	Down     bool // Key is being pressed
	Up       bool // Key is *NOT* being pressed
}

type InputContext struct {
	List map[int32]Key
}

// Set a new Key entry for input context
func (k *InputContext) Set(rlKey int32) {
	k.List[rlKey] = Key{
		Pressed:  rl.IsKeyPressed(rlKey),
		Released: rl.IsKeyReleased(rlKey),
		Down:     rl.IsKeyDown(rlKey),
		Up:       rl.IsKeyUp(rlKey),
	}
}

// Set a list of rlKeys as Key into InputContext
func (k *InputContext) BulkSet(rlKeys []int32) {
	for _, rlKey := range rlKeys {
		k.List[rlKey] = Key{
			Pressed:  rl.IsKeyPressed(rlKey),
			Released: rl.IsKeyReleased(rlKey),
			Down:     rl.IsKeyDown(rlKey),
			Up:       rl.IsKeyUp(rlKey),
		}
	}
}

// Remove a Key from InputContext by rlKey
func (k *InputContext) Delete(rlKey int32) {
	delete(k.List, rlKey)
}

// Update InputContext keys mapping, should be called on Run()
func (k *InputContext) Update() {
	for key := range k.List {
		k.List[key] = Key{
			Pressed:  rl.IsKeyPressed(key),
			Released: rl.IsKeyReleased(key),
			Down:     rl.IsKeyDown(key),
			Up:       rl.IsKeyUp(key),
		}
	}
}

// Return the Key by its rlKey if is present in InputContext,
// if not, return false
func (k *InputContext) GetKey(rlKey int32) Key {
	return k.List[rlKey]
}
