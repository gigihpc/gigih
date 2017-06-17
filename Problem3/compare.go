package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    searchDirsrc := "source"
    searchDirtgt := "target"

    fileListsrc := []string{}
    fileListtgt := []string{}
    err := filepath.Walk(searchDirsrc, func(path string, f os.FileInfo, err error) error {
        fileListsrc = append(fileListsrc, path)
        return nil
    })
    _=err

    err1 := filepath.Walk(searchDirtgt, func(path string, f os.FileInfo, err error) error {
        fileListtgt = append(fileListtgt, path)
        return nil
    })
    _=err1


     for _, filetgt := range fileListtgt {
       //fmt.Println(filetgt)
       CompareWithTarget(filetgt,fileListsrc)
     }
    for _, filesrc := range fileListsrc {
      CompareWithSrc(filesrc,fileListtgt)
    }

}

func CompareWithSrc(filesrc string, fileListtgt []string)  {
  var src []string
  src = strings.SplitAfter(filesrc,"/")
  for _, filetgt := range fileListtgt {
    var tgt_src []string
    tgt_src = strings.SplitAfter(filetgt,"/")
    if src[len(src)-1] != tgt_src[len(tgt_src)-1] {
      fmt.Print(filesrc)
      fmt.Print(" NEW\n")
    }
  }
}

func CompareWithTarget(filetgt string, fileListsrc []string)  {
  var tgtsrc []string
  tgtsrc = strings.SplitAfter(filetgt,"/")
    for _, filesrc := range fileListsrc {
      var src []string
      src = strings.SplitAfter(filesrc,"/")
      if tgtsrc[len(tgtsrc)-1] != src[len(src)-1] {
        fmt.Print(filesrc)
        fmt.Print(" DELETE\n")
      }
    }
}
