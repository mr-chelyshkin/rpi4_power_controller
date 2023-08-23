package utils

import (
	"bytes"
	"fmt"
	"os"
	"syscall"
)

func FileRead(path string) (bytes.Buffer, error) {
	var content bytes.Buffer

	exist, readable, _ := FileExist(path)
	if !exist {
		return content, fmt.Errorf("'%s' file doesn't exist", path)
	}
	if !readable {
		return content, fmt.Errorf("'%s' permission dinied (read)", path)
	}

	fd, err := syscall.Open(path, syscall.O_RDONLY, 0)
	if err != nil {
		return content, err
	}
	defer syscall.Close(fd)

	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(fd, buf)
		if n > 0 {
			content.Write(buf[:n])
		}
		if err == nil {
			if err == syscall.EAGAIN || err == syscall.EINTR {
				continue
			}
			break
		}
	}
	return content, nil
}

func FileWrite(path string, b []byte) error {
	exist, _, writable := FileExist(path)
	if !exist {
		return fmt.Errorf("'%s' file doesn't exist", path)
	}
	if !writable {
		return fmt.Errorf("'%s' permission dinied (write)", path)
	}

	fd, err := syscall.Open(path, syscall.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	_, err = syscall.Write(fd, b)
	if err != nil {
		return err
	}
	return nil
}

func FileParseEnv(b bytes.Buffer) map[string]string {
	data := make(map[string]string)

	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		p := bytes.SplitN(line, []byte("="), 2)
		if len(p) == 2 {
			key := string(bytes.TrimSpace(p[0]))
			value := string(bytes.Trim(bytes.TrimSpace(p[1]), `"`))
			data[key] = value
		}
	}

	return data
}

func FileExist(path string) (exist, readable, writable bool) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, false, false
	}

	fd, err := syscall.Open(path, os.O_RDONLY, 0644)
	if err == nil {
		readable = true
		syscall.Close(fd)
	}

	fd, err = syscall.Open(path, os.O_WRONLY, 0644)
	if err == nil {
		writable = true
		syscall.Close(fd)
	}

	return info != nil, readable, writable
}
