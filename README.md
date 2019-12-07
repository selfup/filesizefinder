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

`filesizefinder -s=1GB -d="C:\\Users"`

**Unix/Linux**

`filesizefinder -s=1GB -d="$HOME"`

### Performance

```
$ time filesizefinder.exe -s=1GB -d="C:\\Users" > .results

real 0m4.456s
user 0m0.000s
sys 0m0.016s
```

### Output

New line delimited file paths

Example on Windows:

```
$ filesizefinder -s=1GB -d="C:\\Users"
C:\Users\selfup\Videos\2019-08-15 18-26-12.flv
C:\Users\selfup\Videos\OBS\2019-08-12_12-10-46.flv
```

### Help

```
$ filesizefinder -h
Usage of filesizefinder:
  -e string
        directory to start scanning recursively from
  -s string
        size: 1GB,10GB,100GB,1TB
```
