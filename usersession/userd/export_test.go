// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package userd

import (
	"github.com/godbus/dbus"

	"github.com/snapcore/snapd/sandbox/cgroup"
)

var (
	SnapFromPid = snapFromPid
)

func MockSnapFromSender(f func(*dbus.Conn, dbus.Sender) (string, error)) func() {
	origSnapFromSender := snapFromSender
	snapFromSender = f
	return func() {
		snapFromSender = origSnapFromSender
	}
}

func MockProcGroup(f func(pid int, match cgroup.GroupMatcher) (string, error)) (restore func()) {
	old := cgroupProcGroup
	cgroupProcGroup = f
	return func() {
		cgroupProcGroup = old
	}
}
