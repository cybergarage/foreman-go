// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

// Storing represents an abstract interface of register.
type Storing interface {
	Open() error
	Close() error
	Clear() error

	SetObject(obj *Object) error
	GetObject(key string) (*Object, bool)

	Size() int
}

// Store represents an register store.
type Store struct {
	Storing
	registers map[string]Object
}

// NewStore returns a new store instance.
func NewStore() *Store {
	store := &Store{}
	store.Clear()
	return store
}

// Open opens the store.
func (store *Store) Open() error {
	return nil
}

// Close closes the store.
func (store *Store) Close() error {
	return nil
}

// Clear clears all registers.
func (store *Store) Clear() error {
	store.registers = make(map[string]Object)
	return nil
}

// SetObject sets a object into the store.
func (store *Store) SetObject(obj Object) error {
	key, err := obj.GetName()
	if err != nil {
		return err
	}
	store.registers[key] = obj

	return nil
}

// GetObject gets the specified object.
func (store *Store) GetObject(key string) (Object, bool) {
	obj, ok := store.registers[key]
	return obj, ok
}

// Size returns the object count.
func (store *Store) Size() int {
	return len(store.registers)
}
