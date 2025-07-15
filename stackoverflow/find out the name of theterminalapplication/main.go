package main

import (
    "log"
    "os"
    "syscall"
    "unsafe"
)

func main() {
    log.SetFlags(0)

    myPid := os.Getpid()

    h, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
    if err != nil {
        log.Fatal(err)
    }

    defer syscall.CloseHandle(h)

    found, parentExe, err := findMyParentsExePath(h, myPid)
    if err != nil {
        log.Fatal(err)
    }

    if !found {
        log.Fatal("weird: no parent!?")
    }

    log.Println(parentExe)
}

func findMyParentsExePath(hSnap syscall.Handle, myPid int) (bool, string, error) {
    found, parentPid, err := findMyParentsPid(hSnap, myPid)
    if err != nil || !found {
        return found, "", err
    }

    var parentExe string

    found = false
    err = enumerateProcSnap(hSnap, func(
        pe *syscall.ProcessEntry32) bool {
        if pe.ProcessID != uint32(parentPid) {
            return true
        }

        parentExe = syscall.UTF16ToString(pe.ExeFile[:])
        found = true

        return false
    })

    if err != nil || !found {
        return found, "", err
    }

    return true, parentExe, nil
}

func findMyParentsPid(hSnap syscall.Handle, myPid int) (bool, int, error) {
    var (
        found     bool
        parentPid uint32
    )

    err := enumerateProcSnap(hSnap, func(pe *syscall.ProcessEntry32) bool {
        if pe.ProcessID != uint32(myPid) {
            return true
        }

        parentPid = pe.ParentProcessID
        found = true

        return false
    })

    if err != nil || !found {
        return found, 0, err
    }

    return true, int(parentPid), nil
}

func enumerateProcSnap(
    hSnap syscall.Handle,
    cb func(pe *syscall.ProcessEntry32) bool,
) error {
    var pe syscall.ProcessEntry32
    pe.Size = uint32(unsafe.Sizeof(pe))

    var err error

    err = syscall.Process32First(hSnap, &pe)

    for {
        if err != nil {
            if err == syscall.ERROR_NO_MORE_FILES {
                return nil
            }

            return err
        }

        proceed := cb(&pe)
        if !proceed {
            break
        }

        err = syscall.Process32Next(hSnap, &pe)
    }

    return nil
}