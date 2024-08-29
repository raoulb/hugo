// Copyright 2024 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package page

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

var pagePartitionTestSources = []string{
	"/section1/testpage1.md",
	"/section1/testpage2.md",
	"/section2/testpage3.md",
	"/section2/testpage4.md",
	"/section3/testpage5.md",
}

func preparePagePartitionTestPages(t *testing.T) Pages {
	var pages Pages
	for _, src := range pagePartitionTestSources {
		p := newTestPage()
		p.path = src
		p.section = strings.Split(strings.TrimPrefix(p.path, "/"), "/")[0]
		pages = append(pages, p)
	}
	return pages
}

var comparePagePartition = qt.CmpEquals(cmp.Comparer(func(a, b Page) bool {
	return a == b
}))

func TestPartitionWithCase1(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)
	expect := PagesPartition{
		Pages{pages[0], pages[1], pages[2], pages[3], pages[4]},
	}

	part, err := pages.PartitionWith(5)
	c.Assert(err, qt.IsNil)
	c.Assert(part, comparePagePartition, expect)
}

func TestPartitionWithCase2(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)
	expect := PagesPartition{
		Pages{pages[0], pages[1], pages[2]},
		Pages{pages[3], pages[4]},
	}

	part, err := pages.PartitionWith(3)
	c.Assert(err, qt.IsNil)
	c.Assert(part, comparePagePartition, expect)
}

func TestPartitionWithCase3(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)
	expect := PagesPartition{
		Pages{pages[0], pages[1]},
		Pages{pages[2], pages[3]},
		Pages{pages[4]},
	}

	part, err := pages.PartitionWith(2)
	c.Assert(err, qt.IsNil)
	c.Assert(part, comparePagePartition, expect)
}

func TestPartitionWithCase4(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)
	expect := PagesPartition{
		Pages{pages[0]},
		Pages{pages[1]},
		Pages{pages[2]},
		Pages{pages[3]},
		Pages{pages[4]},
	}

	part, err := pages.PartitionWith(1)
	c.Assert(err, qt.IsNil)
	c.Assert(part, comparePagePartition, expect)
}

func TestPartitionWithCase5(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)
	expect := PagesPartition{
		Pages{pages[0], pages[1], pages[2], pages[3], pages[4]},
	}

	// Partition size is large: all items fall within a single partition
	part, err := pages.PartitionWith(20)
	c.Assert(err, qt.IsNil)
	c.Assert(part, comparePagePartition, expect)
}

func TestPartitionWithCase6(t *testing.T) {
	c := qt.New(t)

	t.Parallel()
	pages := preparePagePartitionTestPages(t)

	// Partition size is < 1: there is no valid partition
	part, err := pages.PartitionWith(0)
	c.Assert(err, qt.IsNotNil)
	c.Assert(part, qt.IsNil)
}
