package glfw

//#cgo LDFLAGS: -lglfw
//#include <GL/glfw3.h>
import "C"

const (
	//This is incremented when the API is changed in non-compatible ways.
	VersionMajor = C.GLFW_VERSION_MAJOR
	//This is incremented when features are added to the API but it
	//remains backward-compatible.
	VersionMinor = C.GLFW_VERSION_MINOR
	//This is incremented when a bug fix release is made that does not contain
	//any API changes.
	VersionRevision = C.GLFW_VERSION_REVISION
)

//Init initializes the GLFW library. Before most GLFW functions can be used,
//GLFW must be initialized, and before a program terminates GLFW should be
//terminated in order to free any resources allocated during or after
//initialization.
//
//If this function fails, it calls Terminate before returning. If it succeeds,
//you should call Terminate before the program exits.
//
//Additional calls to this function after successful initialization but before
//termination will succeed but will do nothing.
//
//This function may take several seconds to complete on some systems, while on
//other systems it may take only a fraction of a second to complete.
//
//NOTE: On Mac OS X, this function will change the current directory of the
//application to the Contents/Resources subdirectory of the application's
//bundle, if present.
func Init() bool {
	r := C.glfwInit()
	if r == C.GL_TRUE {
		return true
	}
	return false
}

//Terminate destroys all remaining windows, frees any allocated resources and
//sets the library to an uninitialized state. Once this is called, you must
//again call Init successfully before you will be able to use most GLFW
//functions.
//
//If GLFW has been successfully initialized, this function should be called
//before the program exits. If initialization fails, there is no need to call
//this function, as it is called by Init before it returns failure.
func Terminate() {
	C.glfwTerminate()
}

//GetVersion retrieves the major, minor and revision numbers of the GLFW
//library. It is intended for when you are using GLFW as a shared library and
//want to ensure that you are using the minimum required version.
//
//This function may be called before Init.
func GetVersion() (int, int, int) {
	var (
		major C.int
		minor C.int
		rev   C.int
	)
	C.glfwGetVersion(&major, &minor, &rev)
	return int(major), int(minor), int(rev)
}

//GetVersionString returns a static string generated at compile-time according
//to which configuration macros were defined. This is intended for use when
//submitting bug reports, to allow developers to see which code paths are
//enabled in a binary.
func GetVersionString() string {
	return C.GoString(C.glfwGetVersionString())
}
