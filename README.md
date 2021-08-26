# GOCICADA
## 功能
便捷的处理错误码
## 使用
### 定义错误码
```go
type UDMException interface {
	// UeIdNotEqual @CicadaError(code = 201001, message = "参数中各 ueId 不相等")
	UeIdNotEqual()

	// PlamIdNotEqual @CicadaError(code = 201002, message = "参数中各 plmnId 不相等")
	PlamIdNotEqual()
}
```
必须按照`@CicadaError(code = xxx, message = "xxx")`格式注释在函数上，框架才会正确获取到`函数名`、`code`、`message`信息。
### 扫描错误码
调用`CicadaScan(path string)`函数。

**path**：错误码文件位置（基于项目的相对路径），默认到`/src/exceptions`中扫描。
```go
CicadaScan("/src/xxx/exceptions")
```