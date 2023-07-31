package code

import (
	"net/http"

	"github.com/marmotedu/errors"
	"github.com/novalagung/gubrak"
)

// ErrCode implements `github.com/marmotedu/errors`.Coder interface
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

/**
这是 Go 语言中的一个技巧，用于在编译时检查某个类型是否实现了某个接口。

在这个语句中，_ 是一个匿名变量，errors.Coder 是一个接口，&ErrCode{} 是一个类型为 *ErrCode 的值。这个语句的意思是尝试将 &ErrCode{} 赋值给一个 errors.Coder 类型的变量。

如果 ErrCode 类型实现了 errors.Coder 接口，那么这个语句就可以编译通过。如果 ErrCode 类型没有实现 errors.Coder 接口，那么这个语句就会在编译时失败，编译器会报错，说 ErrCode 类型没有实现 errors.Coder 接口。

这个技巧通常用于确保某个类型实现了某个接口，而不需要等到运行时才发现问题。这可以帮助我们在编写代码时就发现问题，而不是在运行程序时才发现问题。
*/
var _ errors.Coder = &ErrCode{}

func (e ErrCode) HTTPStatus() int {
	if e.HTTP == 0 {
		return http.StatusInternalServerError
	}
	return e.HTTP
}

func (e ErrCode) String() string {
	return e.Ext
}

func (e ErrCode) Reference() string {
	return e.Ref
}

func (e ErrCode) Code() int {
	return e.C
}

/**
// nolint: unparam 是一个特殊的注释，用于告诉 linter（代码静态分析工具）忽略特定的警告或错误。在 Go 语言中，这种注释通常与一些 linter 工具一起使用，如 golangci-lint。

在这个特定的例子中，unparam 是一个 linter，它检查函数中未使用的参数。// nolint: unparam 注释告诉 linter 忽略这个函数中未使用的参数的警告。

这种注释通常用于以下情况：

当你知道某个函数的某个参数在未来可能会被使用，但现在还没有被使用时。
当你实现一个接口，但不需要使用所有的参数时。
*/
// nolint: unparam
func register(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}
	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}
	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}
	errors.MustRegister(coder)
}
