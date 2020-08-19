// Package siahe provides a simple API for storing key/value(s) and retrieving the
// stored values by doing efficient prefix search on the keys.
package siahe

import (
	"sort"

	radix "github.com/armon/go-radix"
)

// Siahe provides an interface for implementations of an index which stores related
// terms to certain identifier and allows efficient retrieval of those identifiers
// in the insertion order by searching between that identifier's related terms.
type Siahe interface {
	// Index stores the given term for its identifier. Identifier can be any non-empty
	// string. Please keep in mind that insertion order determines retrieval order.
	Index(identifier string, term string)
	// Find returns all identifiers for stored pairs which their term matches the
	// given prefix. Find guarantees that there would be no duplicates in the returned
	// result.
	Find(term string) []string
}

// New returns the default implementation of the siahe.
func New() Siahe {
	return &siahe{
		tree:        radix.New(),
		idsMaxOrder: 0,
		idsOrder:    make(map[string]int),
	}
}

// siahe is a simple implementation of Siahe interface which uses "Radix Tree" under
// the hood.
type siahe struct {
	tree *radix.Tree

	idsMaxOrder int
	idsOrder    map[string]int
}

func (s *siahe) Index(identifier string, term string) {
	termRelatedIDs := sortableIDs{}

	if previousValues, found := s.tree.Get(term); found {
		termRelatedIDs = previousValues.(sortableIDs)
	}

	var (
		ok    bool
		order int = -1
	)

	if order, ok = s.idsOrder[identifier]; !ok {
		s.idsMaxOrder++
		s.idsOrder[identifier] = s.idsMaxOrder
		order = s.idsOrder[identifier]
	}

	termRelatedIDs = append(termRelatedIDs, sortableID{value: identifier, order: order})

	_, _ = s.tree.Insert(term, termRelatedIDs)
}

func (s siahe) Find(term string) []string {
	idsSet := make(map[string]sortableID)

	s.tree.WalkPrefix(term, func(foundTerm string, relatedIdentifiers interface{}) bool {
		for _, nodeItem := range relatedIdentifiers.(sortableIDs) {
			idsSet[nodeItem.value] = nodeItem
		}

		return false
	})

	strIDs := make([]string, len(idsSet))

	for k, v := range idsSetToSortedSlice(idsSet) {
		strIDs[k] = v.value
	}

	return strIDs
}

func idsSetToSortedSlice(idsSet map[string]sortableID) sortableIDs {
	ids := make(sortableIDs, len(idsSet))
	i := 0

	for _, v := range idsSet {
		ids[i] = v
		i++
	}

	sort.Sort(ids)

	return ids
}

type sortableID struct {
	value string
	order int
}

type sortableIDs []sortableID

func (si sortableIDs) Len() int {
	return len(si)
}

func (si sortableIDs) Less(i, j int) bool {
	return si[i].order < si[j].order
}

func (si sortableIDs) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}
