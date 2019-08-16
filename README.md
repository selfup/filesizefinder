# File Size Finder

Find files, big or small. Much faster than File Explorer in Windows.

Was built to clean up SSDs.

### Install

```
go get github.com/selfup/filesizefinder
go install github.com/selfup/filesizefinder
```

### Use

**Windows**

filesizefinder -d=backslash -s=1GB -f="C:\\Users"

**Unix/Linux**

filesizefinder -d=forwardslash -s=1GB -f="\$HOME"

### Output

New line delimited file paths

Example on Windows:

```
$ filesizefinder -d=backslash -s=1GB -f="C:\\Users"
C:\Users\selfup\Videos\2019-08-15 18-26-12.flv
C:\Users\selfup\Videos\OBS\2019-08-12_12-10-46.flv
```

### Help

```
$ filesizefinder -h
Usage of filesizefinder:
  -d string
        direction of paths backslash (windows) or forwardslash (unix/linux)
  -f string
        folder/directory to start scanning recursively from
  -s string
        size: 1MB,10MB,100MB,1GB,10GB,100GB,1TB
```
