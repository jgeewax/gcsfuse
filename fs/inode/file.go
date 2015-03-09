// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inode

import (
	"errors"

	"github.com/jacobsa/fuse"
	"github.com/jacobsa/gcloud/syncutil"
	"golang.org/x/net/context"
)

type FileInode struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	id fuse.InodeID

	// The name of the GCS object backing the inode. This may or may not yet
	// exist.
	//
	// INVARIANT: name != ""
	// INVARIANT: name[len(name)-1] != '/'
	name string

	/////////////////////////
	// Mutable state
	/////////////////////////

	// A mutex that must be held when calling certain methods. See documentation
	// for each method.
	Mu syncutil.InvariantMutex

	// The generation number of the GCS object from which this inode was
	// branched, or zero if it is newly created. This is used as a precondition
	// in object write requests.
	//
	// GUARDED_BY(Mu)
	srcGeneration uint64
}

var _ Inode = &FileInode{}

////////////////////////////////////////////////////////////////////////
// Public interface
////////////////////////////////////////////////////////////////////////

func (f *FileInode) ID() fuse.InodeID {
	return f.id
}

func (f *FileInode) Name() string {
	return f.name
}

func (f *FileInode) Attributes(
	ctx context.Context) (attrs fuse.InodeAttributes, err error) {
	err = errors.New("TODO(jacobsa): Implement FileInode.Attributes.")
	return
}