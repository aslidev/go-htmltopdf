package image

/*
#include <wkhtmltox/image.h>
#include <stdio.h>

extern void progress_changed(wkhtmltoimage_converter*, int);
extern void phase_changed(wkhtmltoimage_converter*, int);
extern void set_error(wkhtmltoimage_converter*, int);
extern void set_warning(wkhtmltoimage_converter*, int);

void progress_changed_cgo(wkhtmltoimage_converter *converter, int p) {
  progress_changed(converter, p);
}

void phase_changed_cgo(wkhtmltoimage_converter *converter) {
  phase_changed(converter);
}

void set_error_cgo(wkhtmltoimage_converter *converter, const char *msg) {
  set_error(converter, msg);
}

void set_warning_cgo(wkhtmltoimage_converter *converter, const char *msg) {
  set_warning(converter, msg);
}

*/
import "C"
