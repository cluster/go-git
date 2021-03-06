package object

import (
	"fmt"

	"srcd.works/go-git.v4/utils/merkletrie"
	"srcd.works/go-git.v4/utils/merkletrie/noder"
)

// The folowing functions transform changes types form the merkletrie
// package to changes types from this package.

func newChange(c merkletrie.Change) (*Change, error) {
	ret := &Change{}

	var err error
	if ret.From, err = newChangeEntry(c.From); err != nil {
		return nil, fmt.Errorf("From field: ", err)
	}

	if ret.To, err = newChangeEntry(c.To); err != nil {
		return nil, fmt.Errorf("To field: ", err)
	}

	return ret, nil
}

func newChangeEntry(p noder.Path) (ChangeEntry, error) {
	if p == nil {
		return empty, nil
	}

	asTreeNoder, ok := p.Last().(*treeNoder)
	if !ok {
		return ChangeEntry{}, fmt.Errorf("cannot transform non-TreeNoders")
	}

	return ChangeEntry{
		Name: p.String(),
		Tree: asTreeNoder.parent,
		TreeEntry: TreeEntry{
			Name: asTreeNoder.name,
			Mode: asTreeNoder.mode,
			Hash: asTreeNoder.hash,
		},
	}, nil
}

func newChanges(src merkletrie.Changes) (Changes, error) {
	ret := make(Changes, len(src))
	var err error
	for i, e := range src {
		ret[i], err = newChange(e)
		if err != nil {
			return nil, fmt.Errorf("change #%d: %s", err)
		}
	}

	return ret, nil
}
