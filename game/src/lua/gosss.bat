@Echo Off
SetLocal EnableDelayedExpansion
For %%i In (*.lua) Do (
    For /F "Usebackq Delims=" %%j In ("%%~nxi") Do (
        Set Str=%%j
        Echo !Str:return=%%~ni! >>New_%%~nxi
    )
    Del %%~nxi >nul
    Ren New_%%~nxi %%~nxi
)