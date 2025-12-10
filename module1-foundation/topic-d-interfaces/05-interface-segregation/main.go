// Assignment 5: The Interface Segregation
//
// Goal: Take a massive GodInterface. Break it into Reader, Writer, Closer.
//       Demonstrate combining them interface{ Reader; Writer }.
//
// Instructions:
// 1. Start with a "god interface" that does too much
// 2. Break it into small, focused interfaces
// 3. Compose interfaces when you need multiple capabilities
// 4. Functions should accept the smallest interface they need
//
// Key insight: Small interfaces are more flexible and testable

package main

import (
	"fmt"
	"io"
)

// BAD: God interface - does too much!
// type GodInterface interface {
//     Read(p []byte) (n int, err error)
//     Write(p []byte) (n int, err error)
//     Close() error
//     Seek(offset int64, whence int) (int64, error)
//     Flush() error
//     Reset()
// }

// GOOD: Small, focused interfaces (these already exist in io package)
// type Reader interface { Read(p []byte) (n int, err error) }
// type Writer interface { Write(p []byte) (n int, err error) }
// type Closer interface { Close() error }

// Composed interfaces - combine what you need
type ReadWriter interface {
	io.Reader
	io.Writer
}

type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

// Our implementation
type Buffer struct {
	data   []byte
	closed bool
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.closed {
		return 0, fmt.Errorf("buffer is closed")
	}
	n = copy(p, b.data)
	b.data = b.data[n:]
	return n, nil
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	if b.closed {
		return 0, fmt.Errorf("buffer is closed")
	}
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *Buffer) Close() error {
	b.closed = true
	fmt.Println("Buffer closed")
	return nil
}

// Functions accept the SMALLEST interface they need

// OnlyReads only needs to read
func OnlyReads(r io.Reader) {
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

// OnlyWrites only needs to write
func OnlyWrites(w io.Writer) {
	w.Write([]byte("Hello from OnlyWrites!"))
}

// NeedsReadWrite needs both
func NeedsReadWrite(rw ReadWriter) {
	rw.Write([]byte("Wrote this!"))
	buf := make([]byte, 1024)
	n, _ := rw.Read(buf)
	fmt.Printf("Then read: %s\n", buf[:n])
}

// NeedsEverything needs read, write, and close
func NeedsEverything(rwc ReadWriteCloser) {
	rwc.Write([]byte("Full access!"))
	buf := make([]byte, 1024)
	n, _ := rwc.Read(buf)
	fmt.Printf("Read: %s\n", buf[:n])
	rwc.Close()
}

func main() {
	// Buffer satisfies all our interfaces!
	buf := &Buffer{}

	fmt.Println("=== Interface Segregation Demo ===")

	// Pass to function that only needs Reader
	buf.Write([]byte("Some data to read"))
	OnlyReads(buf)

	// Pass to function that only needs Writer
	OnlyWrites(buf)

	// Pass to function that needs ReadWriter
	buf2 := &Buffer{}
	NeedsReadWrite(buf2)

	// Pass to function that needs ReadWriteCloser
	buf3 := &Buffer{}
	NeedsEverything(buf3)

	// KEY BENEFIT: Each function only depends on what it needs
	// This makes testing easier and code more flexible!
}

