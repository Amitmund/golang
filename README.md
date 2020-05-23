# golang
golang codes

```
For Mac
`wget -O AmitMundWebServer https://github.com/Amitmund/golang/blob/master/amundSimpleHTTPServer/AmitMundWebServer_Mac?raw=true`

For Linux
`wget -O AmitMundWebServer https://github.com/Amitmund/golang/blob/master/amundSimpleHTTPServer/AmitMundWebServer_Linux?raw=true`

For Windows
`wget -O AmitMundWebServer https://github.com/Amitmund/golang/blob/master/amundSimpleHTTPServer/AmitMundWebServer_Windows?raw=true`
```


### workspace
```
amund@blr-mptlw:[master][~/github/golang/goworkspace]:$ tree
.
├── bin
├── pkg
└── src
    └── github.com
        └── amund

```

### Other example:

https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos
```
.
├── bin
│   ├── buffalo                                      # command executable
│   ├── dlv                                          # command executable
│   └── packr                                        # command executable
└── src
    └── github.com
        └── digitalocean
            └── godo
                ├── .git                            # Git reposistory metadata
                ├── account.go                      # package source
                ├── account_test.go                 # test source
                ├── ...
                ├── timestamp.go
                ├── timestamp_test.go
                └── util
                    ├── droplet.go
                    └── droplet_test.go


```


### .gitignore - Example

```
# Ignore all
*

# Unignore all with extensions
!*.*

# Unignore all dirs
!*/

### Above combination will ignore all files without extension ###

# Ignore files with extension `.class` & `.sm`
*.class
*.sm

# Ignore `bin` dir
bin/
# or
*/bin/*

# Unignore all `.jar` in `bin` dir
!*/bin/*.jar

# Ignore all `library.jar` in `bin` dir
*/bin/library.jar

# Ignore a file with extension
relative/path/to/dir/filename.extension

# Ignore a file without extension
relative/path/to/dir/anotherfile
```
