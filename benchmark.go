package main

import (
  "os"
  "log"
  "strconv"
  "time"
  b "./bbpformula"
)

func main() {
  if len(os.Args) != 5 {
    log.Fatal("Please specify an output file, a range begin, a range end, and a range step size.")
  }

  filename := os.Args[1]
  rangeStart, err := strconv.Atoi(os.Args[2])
  if err != nil {
      log.Fatal(err)
  }

  rangeEnd, err := strconv.Atoi(os.Args[3])
  if err != nil {
      log.Fatal(err)
  }

  rangeStep, err := strconv.Atoi(os.Args[4])
  if err != nil {
      log.Fatal(err)
  }

  // If the file doesn't exist, create it, or append to the file
  f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
  }

  for n := rangeStart; n < rangeEnd; n+=rangeStep {
    computeAndWrite(f, n)
  }

  err = f.Close()
  if err != nil {
      log.Fatal(err)
  }
}

func computeAndWrite(f *os.File, n int) {
  timeStart := time.Now()
  _, err := b.ToStringBase(b.Log2(n), 2, 8)
  if err != nil {
      log.Fatal(err)
  }

  elapsed := float64(time.Since(timeStart))/float64(time.Millisecond)
  _, err = f.Write([]byte(strconv.Itoa(n)+","+strconv.FormatFloat(elapsed, 'f', -1, 64)+"\n"))
  if err != nil {
      log.Fatal(err)
  }
}
