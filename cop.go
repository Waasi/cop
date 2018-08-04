package main

import(
  "os"
  "io"
  "fmt"
  "time"
  "strconv"
  "strings"
  "io/ioutil"
  "path/filepath"
)

func main() {
  if len(os.Args) != 7 {
    fmt.Println("missing arguments")
    os.Exit(1)
  } else if os.Args[3] != "--from" || os.Args[5] != "--to" {
    fmt.Println("invalid arguments:", os.Args[3], os.Args[5])
    os.Exit(1)
  } else {
    source := validPath(os.Args[1])
    destination := validPath(os.Args[2])
    from := validTimeStamp(os.Args[4])
    to := validTimeStamp(os.Args[6])
    files, _ := ioutil.ReadDir(source)

    for _, f := range files {
      hour, min, sec := time.Time.Clock(f.ModTime())
      time := [3]int{hour, min, sec}

      if inInterval(time, from, to) {
        fmt.Println("copying", f.Name(), "to", destination)
        copyFile(filepath.Join(source, f.Name()), filepath.Join(destination, f.Name()))
      }
    }
  }
}

func validPath(pathName string) string {
  _, err := os.Stat(pathName)

  if os.IsNotExist(err) {
    fmt.Println(pathName, "directory does not exist")
    os.Exit(1)
  } else if err != nil {
    fmt.Println("unexpected error with path", pathName)
    os.Exit(1)
  }

  return pathName
}

func validTimeStamp(time string) [3]int  {
  units := strings.Split(time, ":")
  hour, herr := strconv.Atoi(units[0])
  min, merr := strconv.Atoi(units[1])
  sec, serr := strconv.Atoi(units[2])

  if herr != nil || merr != nil || serr != nil {
    fmt.Println("unable to use timestamp", time)
    os.Exit(1)
  }

  return [3]int{hour, min, sec}
}

func inInterval(timestamp [3]int, from [3]int, to [3]int) bool {
  if(timestamp[0] == from[0] || timestamp[0] == to[0]) {
    if(timestamp[1] == from[1] || timestamp[1] == to[1]) {
      return timestamp[2] >= from[2] && timestamp[2] <= to[2]
    } else {
      return timestamp[1] >= from[1] && timestamp[1] <= to[1]
    }
  } else {
    return timestamp[0] >= from[0] && timestamp[0] <= to[0]
  }
}

func copyFile(source string, destination string) bool {
  from, err := os.Open(source)
  if err != nil {
    return false
  }
  defer from.Close()

  to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
    return false
  }
  defer to.Close()

  _, err = io.Copy(to, from)
  if err != nil {
    return false
  }

  return true
}
