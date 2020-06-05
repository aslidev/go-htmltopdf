package image

/*
#cgo CFLAGS: -I../wkhtmltopdf/include
#cgo LDFLAGS: -lwkhtmltox -L../wkhtmltopdf/lib
#include <wkhtmltox/image.h>

void progress_changed_cgo(wkhtmltoimage_converter *converter, int p);
void phase_changed_cgo(wkhtmltoimage_converter *converter);
void set_error_cgo(wkhtmltoimage_converter *converter, const char *msg);
void set_warning_cgo(wkhtmltoimage_converter *converter, const char *msg);

*/
import "C"
import (
	"errors"
	"unsafe"
)

var inputData string

type intFuncCallback func(int)
type stringFuncCallback func(string)

type Html2image struct {
	gs        *C.wkhtmltoimage_global_settings
	converter *C.wkhtmltoimage_converter
}

var progressChangedCallback intFuncCallback
var phaseChangedCallback, errorCallback, warningCallback stringFuncCallback

//export progress_changed
func progress_changed(converter *C.wkhtmltoimage_converter, p C.int) {
	if progressChangedCallback != nil {
		progressChangedCallback(int(p))
	}
}

//export phase_changed
func phase_changed(converter *C.wkhtmltoimage_converter) {
	var phase C.int = C.wkhtmltoimage_current_phase(converter)
	var phaseDescription = C.GoString(C.wkhtmltoimage_phase_description(converter, phase))
	if phaseChangedCallback != nil {
		phaseChangedCallback(phaseDescription)
	}
}

//export set_error
func set_error(converter *C.wkhtmltoimage_converter, msg *C.char) {
	if errorCallback != nil {
		errorCallback(C.GoString(msg))
	}
}

//export set_warning
func set_warning(converter *C.wkhtmltoimage_converter, msg *C.char) {
	if warningCallback != nil {
		warningCallback(C.GoString(msg))
	}
}

func (h2p *Html2image) initCallback() {
	C.wkhtmltoimage_set_progress_changed_callback(h2p.converter, C.wkhtmltoimage_int_callback(unsafe.Pointer(C.progress_changed_cgo)))
	C.wkhtmltoimage_set_phase_changed_callback(h2p.converter, C.wkhtmltoimage_void_callback(unsafe.Pointer(C.phase_changed_cgo)))
	C.wkhtmltoimage_set_error_callback(h2p.converter, C.wkhtmltoimage_str_callback(unsafe.Pointer(C.set_error_cgo)))
	C.wkhtmltoimage_set_warning_callback(h2p.converter, C.wkhtmltoimage_str_callback(unsafe.Pointer(C.set_warning_cgo)))
}

func (h2p *Html2image) OnProgressChanged(callback intFuncCallback) {
	progressChangedCallback = callback
}

func (h2p *Html2image) OnPhaseChanged(callback stringFuncCallback) {
	phaseChangedCallback = callback
}

func (h2p *Html2image) OnError(callback stringFuncCallback) {
	errorCallback = callback
}

func (h2p *Html2image) OnWarning(callback stringFuncCallback) {
	warningCallback = callback
}

func New() *Html2image {
	gs := C.wkhtmltoimage_create_global_settings()
	return &Html2image{gs: gs}
}

func (h2p *Html2image) SetGlobalSettings(global_settings [][2]string) {
	for _, row := range global_settings {
		C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString(row[0]), C.CString(row[1]))
	}
}

func (h2p *Html2image) SetWidth(width string) *Html2image {
	C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString("screenWidth"), C.CString(width))
	return h2p
}

func (h2p *Html2image) SetHeight(height string) *Html2image {
	C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString("screenHeight"), C.CString(height))
	return h2p
}

func (h2p *Html2image) SetURL(url string) {
	C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString("page"), C.CString(url))
}

func (h2p *Html2image) SetOutputFileName(name string) {
	C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString("out"), C.CString(name))
}

func (h2p *Html2image) SetBufferedOutput() {
	C.wkhtmltoimage_set_global_setting(h2p.gs, C.CString("out"), C.CString(""))
}

func (h2p *Html2image) SetData(data string) {
	inputData = data
}

func (h2p *Html2image) CreateImage(data string) (error, []byte) {
	converter := C.wkhtmltoimage_create_converter(h2p.gs, C.CString(data))
	h2p.converter = converter
	h2p.initCallback()

	res := C.wkhtmltoimage_convert(h2p.converter)

	if res != 1 {
		return errors.New("Conversion failed"), nil
	}

	var ptr *C.uchar
	length := C.wkhtmltoimage_get_output(h2p.converter, &ptr)

	// GoBytes accepts C.int as length, so the length needs to be
	// type casted from long to int
	outData := C.GoBytes(unsafe.Pointer(ptr), C.int(length))
	C.wkhtmltoimage_destroy_converter(h2p.converter)
	return nil, outData
}

func init() {
	C.wkhtmltoimage_init(1)
}

func (h2p *Html2image) Destroy() {
	C.wkhtmltoimage_deinit()
}
