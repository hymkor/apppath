apppath.exe
===========

Print `App Paths`.

```
$ apppath.exe wordpad.exe
"%ProgramFiles%\Windows NT\Accessories\WORDPAD.EXE"
```

Some applications are executable even if it is not on %PATH%.
Because the application's fullpath is written in 
`HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths`.

This program shows it.
