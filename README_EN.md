# SIM Applet Manager

[![build](https://github.com/common-creation/sim-applet-manager/actions/workflows/build.yml/badge.svg)](https://github.com/common-creation/sim-applet-manager/actions/workflows/build.yml)

![](https://i.imgur.com/N5RbVPH.png)

## Overview

This is a GUI application for managing SIM applets in the eSIM of [IoT Connect Mobile Type S](https://sdpf.ntt.com/services/icms/) provided by [NTT Communications Corporation](https://www.ntt.com/).

Since it is a frontend for [martinpaljak/GlobalPlatformPro](https://github.com/martinpaljak/GlobalPlatformPro), it may also be used with eSIMs provided by other carriers.

## Operating Environment

### macOS

- Mac OS X 10.13 or later
- homebrew is installed
- martinpaljak/GlobalPlatformPro is installed
  - `brew install martinpaljak/brew/gppro --HEAD`

### Windows

- Windows 10 or later
- For Windows 10, [WebView2](https://developer.microsoft.com/ja-jp/microsoft-edge/webview2) is installed
- martinpaljak/GlobalPlatformPro v20.08.12 or later is installed
  - https://github.com/martinpaljak/GlobalPlatformPro/releases/tag/v20.08.12
  - Place it in `PATH` or `%USERPROFILE%\.simappletmanager\gp.exe`

## LICENSE

MIT