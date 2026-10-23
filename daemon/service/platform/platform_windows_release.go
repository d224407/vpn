//
//  Daemon for IVPN Client Desktop
//  https://github.com/ivpn/desktop-app
//
//  Created by Stelnykovych Alexandr.
//  Copyright (c) 2023 IVPN Limited.
//
//  This file is part of the Daemon for IVPN Client Desktop.
//
//  The Daemon for IVPN Client Desktop is free software: you can redistribute it and/or
//  modify it under the terms of the GNU General Public License as published by the Free
//  Software Foundation, either version 3 of the License, or (at your option) any later version.
//
//  The Daemon for IVPN Client Desktop is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
//  or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
//  details.
//
//  You should have received a copy of the GNU General Public License
//  along with the Daemon for IVPN Client Desktop. If not, see <https://www.gnu.org/licenses/>.
//

//go:build windows && !debug
// +build windows,!debug

package platform

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// initialize all constant values (e.g. servicePortFile) which can be used in external projects (IVPN CLI)
func doInitConstantsForBuild() {
}

func doOsInitForBuild() {
	installDir := getInstallDir()
	wfpDllPath = path.Join(installDir, "IVPN Firewall Native.dll")
	nativeHelpersDllPath = path.Join(installDir, "IVPN Helpers Native.dll")
	splitTunDriverPath = path.Join(installDir, "SplitTunnelDriver", "ivpn-split-tunnel.sys")

	// In release builds, vendor binaries are in a flat structure (no arch subdirs)
	openVpnBinaryPath = path.Join(installDir, "OpenVPN", "openvpn.exe")
	obfsproxyStartScript = path.Join(installDir, "OpenVPN", "obfsproxy", "obfs4proxy.exe")
	v2rayBinaryPath = path.Join(installDir, "v2ray", "v2ray.exe")
	wgBinaryPath = path.Join(installDir, "WireGuard", "wireguard.exe")
	wgToolBinaryPath = path.Join(installDir, "WireGuard", "wg.exe")
	dnscryptproxyBinPath = path.Join(installDir, "dnscrypt-proxy", "dnscrypt-proxy.exe")
	kemHelperBinaryPath = path.Join(installDir, "kem", "kem-helper.exe")
}

// PORTABLE MODE: if a file named "portable" is located next to the daemon executable,
// the daemon uses its own folder as "install dir" (etc, log, mutable ...)
// instead of reading the folder from the registry. No installer required.
func getPortableDir() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	dir := filepath.Dir(exe)
	if _, err := os.Stat(filepath.Join(dir, "portable")); err != nil {
		return ""
	}
	for _, sub := range []string{"etc", "log", "mutable"} {
		_ = os.MkdirAll(filepath.Join(dir, sub), 0o755)
	}
	return filepath.ToSlash(dir)
}

func getInstallDir() string {
	if portableDir := getPortableDir(); portableDir != "" {
		return portableDir
	}

	ret := ""

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\IVPN Client`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer k.Close()

	if err == nil {
		ret, _, err = k.GetStringValue("")
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	}

	if len(ret) == 0 {
		fmt.Println("WARNING: There is no info about IVPN Client install folder in the registry. Is IVPN Client installed?")
		return ""
	}

	return strings.ReplaceAll(ret, `\`, `/`)
}

func getEtcDirCommon() string {
	return getEtcDir()
}
