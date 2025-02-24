// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd March 2019
 * Updated: 24th February 2025
 */

package libclimate

const (
	VersionMajor int16 = 0
	VersionMinor int16 = 6
	VersionPatch int16 = 0
	Version      int64 = (int64(VersionMajor) << 48) + (int64(VersionMinor) << 32) + (int64(VersionPatch) << 16)

	VersionRevision int16 = VersionPatch
)

/* ///////////////////////////// end of file //////////////////////////// */
