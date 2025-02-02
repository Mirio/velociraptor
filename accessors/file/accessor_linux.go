// +build linux

/*
   Velociraptor - Dig Deeper
   Copyright (C) 2019-2022 Rapid7 Inc.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package file

import (
	"time"

	"www.velocidex.com/golang/velociraptor/accessors"
)

// On Linux we need xstat() support to get birth time.
func (self *OSFileInfo) Btime() time.Time {
	return time.Time{}
}

func (self *OSFileInfo) Mtime() time.Time {
	ts := int64(self._Sys().Mtim.Sec)
	return time.Unix(ts, 0)
}

func (self *OSFileInfo) Ctime() time.Time {
	ts := int64(self._Sys().Ctim.Sec)
	return time.Unix(ts, 0)
}

func (self *OSFileInfo) Atime() time.Time {
	ts := int64(self._Sys().Atim.Sec)
	return time.Unix(ts, 0)
}

func NewOSFileSystemAccessor() *OSFileSystemAccessor {
	root_path, _ := accessors.NewLinuxOSPath("")
	return &OSFileSystemAccessor{
		root: root_path,
	}
}

func splitDevNumber(dev uint64) (major, minor uint64) {
	// See bits/sysmacros.h (glibc) or sys/sysmacros.h (musl-libc)
	major = ((dev >> 32) & 0xfffff000) | ((dev >> 8) & 0xfff)
	minor = ((dev >> 12) & 0xffffff00) | (dev & 0xff)
	return
}
