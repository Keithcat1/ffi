package main
import (
	"github.com/keithcat1/ffi"
	"fmt"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
)
type dll struct {
	SyzInitialize func() ffi.Int `proc:"syz_initialize"`
	SyzShutdown func() ffi.Int `proc:"syz_shutdown"`
	SyzCreateContext func(handle ffi.Uintptr, a, b ffi.Uintptr) ffi.Int `proc:"syz_createContext"`
	CreateStreamingGeneratorFromFile func(out ffi.Uintptr, ctx ffi.Uint64, path ffi.UTF8Pointer, a, b ffi.Uintptr) ffi.Int `proc:"syz_createStreamingGeneratorFromFile"`
CreateStreamingGeneratorFromEncodedData func(out ffi.Uintptr, ctx ffi.Uint64, data ffi.BytesPointer, a, b ffi.Uintptr) ffi.Int `proc:"syz_createStreamingGeneratorFromEncodedData"`
	CreateDirectSource func(out ffi.Uintptr, ctx ffi.Uint64, a, b ffi.Uintptr) ffi.Int `proc:"syz_createStreamingGeneratorFromFile"`
	SourceAddGenerator func(source, gen ffi.Uint64) ffi.Int `proc:"syz_sourceAddGenerator"`
CreateBufferFromEncodedData func(ffi.Uintptr, length ffi.Uint64, data ffi.BytesPointer, a, b ffi.Uintptr) ffi.Int `proc:"syz_createBufferFromEncodedData"`
}
func main() {
	var table dll
	lib := ffi.MustLoadLibrary("synthizer.dll")
	defer lib.MustUnload()
	lib.FillFunctionTable(&table)
	fmt.Printf("%#v\n", table)
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil { panic(err) }
	println(table.SyzInitialize())
	var ctx, gen, src ffi.Uint64
	println(table.SyzCreateContext(ffi.Uintptr(unsafe.Pointer(&ctx)), 0, 0))
	println(ctx)
println(table.CreateStreamingGeneratorFromEncodedData(ffi.Uintptr(unsafe.Pointer(&gen)), ctx, ffi.PointerToBytes(data), 0, 0))
	println(gen)
println(table.CreateDirectSource(ffi.Uintptr(unsafe.Pointer(&src)), ctx, 0, 0))
println(table.SourceAddGenerator(src, gen))
	time.Sleep(time.Second * 5)
	println(table.SyzShutdown())
}