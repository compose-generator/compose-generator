; -------------------------------
; Start
  
  !define PRODUCT_NAME "Compose Generator"
  !define PUBLISHER_NAME "ChilliBits"
  !define VERSION "$%VERSION%"
  !define ARCH "$%ARCH%"
  !define MUI_BRANDINGTEXT "Compose Generator $%VERSION%"
  CRCCheck On

!include "MUI.nsh"

;---------------------------------
;General
 
  OutFile "compose-generator.exe"
  ShowInstDetails "nevershow"
  ShowUninstDetails "nevershow"
 
;--------------------------------
;Folder selection page
 
  InstallDir "$PROGRAMFILES64\${PUBLISHER_NAME}\${PRODUCT_NAME}"

;--------------------------------
;Modern UI Configuration
 
  !define MUI_WELCOMEPAGE
  !define MUI_LICENSEPAGE
  !define MUI_DIRECTORYPAGE
  !define MUI_ABORTWARNING
  !define MUI_UNINSTALLER
  !define MUI_UNCONFIRMPAGE
  !define MUI_FINISHPAGE

;--------------------------------
;Language
 
  !insertmacro MUI_LANGUAGE "English"
 
;-------------------------------- 
;Installer Sections
Section "install" Installation
 
  EnVar::Check "NULL" "NULL"
  Pop $0
  DetailPrint "EnVar::Check write access HKCU returned=|$0|"

  ReadRegStr $0 HKCU "Environment" "Path"
  StrCpy $0 "$PROGRAMFILES64\${PUBLISHER_NAME}\${PRODUCT_NAME};$0"
  WriteRegStr HKCU "Environment" "Path" '$0'

;Add files
  SetOutPath "$INSTDIR"
  File /oname=compose-generator.exe dist/compose-generator_windows_${ARCH}/compose-generator.exe
  File /r predefined-templates

;create start-menu items
  CreateDirectory "$SMPROGRAMS\${PRODUCT_NAME}"
  CreateShortCut "$SMPROGRAMS\${PRODUCT_NAME}\Uninstall.lnk" "$INSTDIR\Uninstall.exe" "" "$INSTDIR\Uninstall.exe" 0
  ;CreateShortCut "$SMPROGRAMS\${PRODUCT_NAME}\${PRODUCT_NAME}.lnk" "$INSTDIR\${MUI_FILE}.exe" "" "$INSTDIR\${MUI_FILE}.exe" 0

;write uninstall information to the registry
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${PRODUCT_NAME}" "DisplayName" "${PRODUCT_NAME}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${PRODUCT_NAME}" "UninstallString" "$INSTDIR\Uninstall.exe"

  WriteUninstaller "$INSTDIR\Uninstall.exe"

SectionEnd

;--------------------------------
;Uninstaller Section
Section "Uninstall"

;Delete Files 
  RMDir /r "$INSTDIR\*.*"

;Remove the installation directory
  RMDir "$INSTDIR"

;Delete Start Menu Shortcuts
  Delete "$DESKTOP\${PRODUCT_NAME}.lnk"
  Delete "$SMPROGRAMS\${PRODUCT_NAME}\*.*"
  RmDir  "$SMPROGRAMS\${PRODUCT_NAME}"
 
;Delete Uninstaller And Unistall Registry Entries
  DeleteRegKey HKEY_LOCAL_MACHINE "SOFTWARE\${PRODUCT_NAME}"
  DeleteRegKey HKEY_LOCAL_MACHINE "SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\${PRODUCT_NAME}"
 
SectionEnd

;eof