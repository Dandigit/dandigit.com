@echo off
echo Task Manager Manager - Dandigit.com 2018
echo Enabling...
REG add HKCU\Software\Microsoft\Windows\CurrentVersion\Policies\System  /v  DisableTaskMgr  /t REG_DWORD  /d /0 /f
echo Done!
pause