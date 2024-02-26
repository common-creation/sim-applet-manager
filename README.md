# SIM Applet Manager

[![build](https://github.com/common-creation/sim-applet-manager/actions/workflows/build.yml/badge.svg)](https://github.com/common-creation/sim-applet-manager/actions/workflows/build.yml)

An [English version](./README_EN.md) is available.

![](https://i.imgur.com/N5RbVPH.png)

## 概要

[NTTコミュニケーションズ株式会社](https://www.ntt.com/)より提供されている、[IoT Connect Mobile Type S](https://sdpf.ntt.com/services/icms/)のeSIM内のSIMアプレットを管理するためのGUIアプリケーションです。

[martinpaljak/GlobalPlatformPro](https://github.com/martinpaljak/GlobalPlatformPro)のフロントエンドなので、その他キャリアから提供されているeSIMでも使用できる場合があります。

## 動作環境

### macOS

- Mac OS X 10.13 またはそれ以降
- homebrewがインストールされていること
- martinpaljak/GlobalPlatformProがインストールされていること
  - `brew install martinpaljak/brew/gppro --HEAD`

### Windows

- Windows 10 またはそれ以降
- Windows 10 の場合は [WebView2](https://developer.microsoft.com/ja-jp/microsoft-edge/webview2) がインストールされていること
- martinpaljak/GlobalPlatformPro v20.08.12 以降がインストールされていること
  - https://github.com/martinpaljak/GlobalPlatformPro/releases/tag/v20.08.12
  - `PATHが通っている場所` または `%USERPROFILE%\.simappletmanager\gp.exe` に置くこと

## LICENSE

MIT
