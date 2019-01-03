// Code generated by gotemplate. DO NOT EDIT.

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black Tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package fork_multi_index

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eosspark/container/templates"
	rbt "github.com/eosspark/container/templates/tree"
)

// template type Map(K,V,Compare)

func assertByLibBlockNumIndexImplementation() {
	var _ templates.Map = (*byLibBlockNumIndex)(nil)
}

// Map holds the elements in a red-black Tree
type byLibBlockNumIndex struct {
	*rbt.Tree
}

// NewWith instantiates a Tree map with the custom comparator.
func newByLibBlockNumIndex() *byLibBlockNumIndex {
	return &byLibBlockNumIndex{Tree: rbt.NewWith(byLibBlockNumCompare, false)}
}

func copyFromByLibBlockNumIndex(tm *byLibBlockNumIndex) *byLibBlockNumIndex {
	return &byLibBlockNumIndex{Tree: rbt.CopyFrom(tm.Tree)}
}

type multiByLibBlockNumIndex = byLibBlockNumIndex

func newMultiByLibBlockNumIndex() *multiByLibBlockNumIndex {
	return &byLibBlockNumIndex{Tree: rbt.NewWith(byLibBlockNumCompare, true)}
}

func copyMultiFromByLibBlockNumIndex(tm *byLibBlockNumIndex) *byLibBlockNumIndex {
	return &byLibBlockNumIndex{Tree: rbt.CopyFrom(tm.Tree)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byLibBlockNumIndex) Put(key ByLibBlockNumComposite, value IndexKey) {
	m.Tree.Put(key, value)
}

func (m *byLibBlockNumIndex) Insert(key ByLibBlockNumComposite, value IndexKey) iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.Insert(key, value)}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in Tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byLibBlockNumIndex) Get(key ByLibBlockNumComposite) iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.Get(key)}
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byLibBlockNumIndex) Remove(key ByLibBlockNumComposite) {
	m.Tree.Remove(key)
}

// Keys returns all keys in-order
func (m *byLibBlockNumIndex) Keys() []ByLibBlockNumComposite {
	keys := make([]ByLibBlockNumComposite, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key().(ByLibBlockNumComposite)
	}
	return keys
}

// Values returns all values in-order based on the key.
func (m *byLibBlockNumIndex) Values() []IndexKey {
	values := make([]IndexKey, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value().(IndexKey)
	}
	return values
}

// Each calls the given function once for each element, passing that element's key and value.
func (m *byLibBlockNumIndex) Each(f func(key ByLibBlockNumComposite, value IndexKey)) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		f(Iterator.Key(), Iterator.Value())
	}
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil,nil otherwise if no element
// matches the criteria.
func (m *byLibBlockNumIndex) Find(f func(key ByLibBlockNumComposite, value IndexKey) bool) (k ByLibBlockNumComposite, v IndexKey) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		if f(Iterator.Key(), Iterator.Value()) {
			return Iterator.Key(), Iterator.Value()
		}
	}
	return
}

// String returns a string representation of container
func (m byLibBlockNumIndex) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}

// Iterator holding the Iterator's state
type iteratorByLibBlockNumIndex struct {
	rbt.Iterator
}

// Iterator returns a stateful Iterator whose elements are key/value pairs.
func (m *byLibBlockNumIndex) Iterator() iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{Iterator: m.Tree.Iterator()}
}

// Begin returns First Iterator whose position points to the first element
// Return End Iterator when the map is empty
func (m *byLibBlockNumIndex) Begin() iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.Begin()}
}

// End returns End Iterator
func (m *byLibBlockNumIndex) End() iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.End()}
}

// Value returns the current element's value.
// Does not modify the state of the Iterator.
func (Iterator *iteratorByLibBlockNumIndex) Value() IndexKey {
	return Iterator.Iterator.Value().(IndexKey)
}

// Key returns the current element's key.
// Does not modify the state of the Iterator.
func (Iterator *iteratorByLibBlockNumIndex) Key() ByLibBlockNumComposite {
	return Iterator.Iterator.Key().(ByLibBlockNumComposite)
}

func (m *byLibBlockNumIndex) LowerBound(key ByLibBlockNumComposite) iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.LowerBound(key)}
}

func (m *byLibBlockNumIndex) UpperBound(key ByLibBlockNumComposite) iteratorByLibBlockNumIndex {
	return iteratorByLibBlockNumIndex{m.Tree.UpperBound(key)}

}

// ToJSON outputs the JSON representation of the map.
type pairByLibBlockNumIndex struct {
	Key ByLibBlockNumComposite `json:"key"`
	Val IndexKey               `json:"val"`
}

func (m byLibBlockNumIndex) MarshalJSON() ([]byte, error) {
	elements := make([]pairByLibBlockNumIndex, 0, m.Size())
	it := m.Iterator()
	for it.Next() {
		elements = append(elements, pairByLibBlockNumIndex{it.Key(), it.Value()})
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *byLibBlockNumIndex) UnmarshalJSON(data []byte) error {
	elements := make([]pairByLibBlockNumIndex, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for _, pair := range elements {
			m.Put(pair.Key, pair.Val)
		}
	}
	return err
}