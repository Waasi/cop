package main

import(
  "fmt"
  "os"
  "path/filepath"
  "io"
  "io/ioutil"
  "time"
)

func main() {
  //currentUser, _ := user.Current()

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
      if inInterval(f.ModTime(), from, to) {
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

func validTimeStamp(datetime string) time.Time {
  layout := "2006-01-02T15:04 MST"
  t, err := time.Parse(layout, datetime + " AST")

  if err != nil {
    fmt.Println("unable to use timestamp", datetime)
    os.Exit(1)
  }

  return t
}

func inInterval(timestamp time.Time, from time.Time, to time.Time) bool {
  uTimeStamp := timestamp.Unix()
  uFrom := from.Unix()
  uTo := to.Unix()

  return uTimeStamp >= uFrom && uTimeStamp <= uTo
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
