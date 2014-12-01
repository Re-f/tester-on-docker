// +build !inner

package docker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func init() {
	err := loadConfig()
	if nil != err {
		panic(err.Error())
	}
}
func RunTestCase(t *testing.T, tc func(t *testing.T)) {
	runTestCase(t, "", "", 1, false)
}

func RunTestCaseWithPrepare(t *testing.T, tagName string, tc func(t *testing.T)) {
	runTestCase(t, "", repository+":"+tagName, 1, false)
}

func getFuncInfo(skip int) (pkname, funcName string, err error) {
	funcPc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", "", fmt.Errorf("get func name and package name error")
	}
	// funcDesc: pkName.funcName
	funcDesc := runtime.FuncForPC(funcPc).Name()
	poitPos := strings.LastIndex(funcDesc, ".") + 1
	return funcDesc[0 : poitPos-1], funcDesc[poitPos:], nil
}

func runTestCase(t *testing.T, funcName, imageName string, skip int, isPrepare bool) string {
	// image ：default | given
	im := getImage()
	if "" != imageName {
		im.name = imageName
	}
	// funcName : defautl | given
	pkName, tmpFuncName, err := getFuncInfo(skip + 1)
	if nil != err {
		t.Fatalf(err.Error())
	}
	if !isPrepare {
		funcName = tmpFuncName
	}

	// compile
	err = compileInnerTestCase(pkName)
	if nil != err {
		t.Fatalf("complie tc error: " + err.Error())
	}

	// run tc
	cid, output, err := runContainer(funcName, filepath.Base(pkName), im, testing.Verbose(), !isDebug(), isPrepare)
	if nil != err {
		t.Fatalf("run container error: %v", err.Error())
	}
	fmt.Println(output)

	return cid
}

func Prepare(t *testing.T, funcName string, forceNew bool) {
	cid := ""
	isExist := isImageExist(repository, funcName)
	debugLog("[Info]prepare image exist: %v, force to build new image: %v", isExist, forceNew)
	if !isExist {
		forceNew = true
	}
	if forceNew {
		// @todo remove all containers base on this image
		if err := removeImage(repository + ":" + funcName); nil != err {
			debugLog("[Warning] remove old imgae error: %v", err.Error())
		}

		cid = runTestCase(t, funcName, "", 1, true)
		if err := buildImage(cid, funcName); nil != err {
			t.Fatalf("Prepare error: %v", err.Error())
		}
	}

	if err := removeContainer(cid); nil != err {
		debugLog("[Warning]Prepare: remove prepare container error: %v", err.Error())
	}

}
