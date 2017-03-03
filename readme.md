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

Example
-------

Build with the solution of Visual Studio.

Batchfile:

```
for /F "tokens=*" %%D in ('apppath devenv.exe') do "%%D" example.sln /build Release
```

is substituted to :

```
"c:\Program Files\Microsoft Visual Studio 10.0\Common7\IDE\devenv.exe" example.sln /build Debug
```

We need not to run the heavy batchfile vcvarsall.bat so-called "Visual Studio Command Prompt."

Misc.
----

`apppath.exe` is same as:

```
for /F "skip=2 tokens=2*" %%I in ('reg query "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\%1" /ve') do echo %%J
```
