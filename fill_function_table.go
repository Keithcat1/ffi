package ffi
import (
	"reflect"
	"fmt"
)
func (self *Library) FillFunctionTable(ptr interface{}) {

	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("FillFunctionTable: type %T is not a pointer", ptr))
	}
	indirect := reflect.Indirect(v)
	if indirect.Kind() != reflect.Struct {
		panic(fmt.Errorf("FillFunctionTable: %T is not a pointer to a struct", indirect.Interface()))
	}
	// getting a reflect.Type of interface type FFIValue is kind of hard. We need to create a FFIValue value, create a reflect.Value with a pointer to it and then indirect
	var temp FFIValue
	ffiValueType := reflect.Indirect(reflect.ValueOf(&temp)).Type()
	var tempResult FFIReturnable
	ffiReturnableType := reflect.Indirect(reflect.ValueOf(&tempResult)).Type()
	var tempError error
	errorType := reflect.Indirect(reflect.ValueOf(&tempError)).Type()
	kind := indirect.Type()
	for i := 0; i < kind.NumField(); i++ {
		field := kind.Field(i)
		fieldType := field.Type
		if fieldType.Kind() != reflect.Func || !field.IsExported() {
			continue
		}
		if fieldType.NumOut() > 2 {
		panic(fmt.Errorf("FillFunctionTable: field name %v, type %v cannot be made into a FFI wrapper function because it has more than two return parameters", field.Name, fieldType))
		}
			if fieldType.NumOut() > 0 && !reflect.PtrTo(fieldType.Out(0)).Implements(ffiReturnableType) {
				panic(fmt.Errorf("FillFunctionTable: Field name: %v, type: %v cannot be made into an FFI wrapper function because return parameter 0, type %v does not implement FFIReturnable", field.Name, fieldType, fieldType.Out(0)))
			}
			if fieldType.NumOut() > 1 && !fieldType.Out(1).Implements(errorType) {
				panic(fmt.Errorf("FillFunctionTable: Field name: %v, type: %v cannot be made into an FFI wrapper function because return parameter 1, type %v does not implement error", field.Name, fieldType, fieldType.Out(0)))
			}

		for j := 0; j<fieldType.NumIn(); j++ {
			if !fieldType.In(j).Implements(ffiValueType) {
				panic(fmt.Errorf("FillFunctionTable: Field name: %v, type: %v cannot be made into an FFI wrapper function because input parameter %v, type %v does not implement FFIValue", field.Name, fieldType, j, fieldType.In(j)))
			}
		}
		procName, exists := field.Tag.Lookup("proc")
		if !exists {
			procName = field.Name
		}
		proc := self.MustGetProcAddress(procName)
//		fmt.Printf("Loaded %v, proc %v\n", field.Name, proc)
		returnArgs := fieldType.NumOut()
		returnType := fieldType.Out(0)
		function := reflect.MakeFunc(fieldType, func(args []reflect.Value) []reflect.Value {
			proc := proc
//			fmt.Printf("%v (%v) just got called with proc %v, returning %v arguments\n", field.Name, fieldType, proc, returnArgs)
			var result reflect.Value
			if returnArgs >= 1 {
				result = reflect.New(returnType)
			}
			var err error
			switch len(args) {
				case 0: err = call0(result.Interface().(FFIReturnable), proc)
				case 1: err = call1(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue))
				case 2: err = call2(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue), args[1].Interface().(FFIValue))
				case 3: err = call3(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue), args[1].Interface().(FFIValue), args[2].Interface().(FFIValue))
				case 4: err = call4(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue), args[1].Interface().(FFIValue), args[2].Interface().(FFIValue), args[3].Interface().(FFIValue))
				case 5: err = call5(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue), args[1].Interface().(FFIValue), args[2].Interface().(FFIValue), args[3].Interface().(FFIValue), args[4].Interface().(FFIValue))
				case 6: err = call6(result.Interface().(FFIReturnable), proc, args[0].Interface().(FFIValue), args[1].Interface().(FFIValue), args[2].Interface().(FFIValue), args[3].Interface().(FFIValue), args[4].Interface().(FFIValue), args[5].Interface().(FFIValue))
			}
			if returnArgs == 1 {
				return []reflect.Value{reflect.Indirect(result)}
			} else if returnArgs == 2 {
				println("Returning two")
				return []reflect.Value{reflect.Indirect(result), reflect.ValueOf(err)}
			} else {
				println("Returning nothing!")
				return nil
			}
		})
		indirect.Field(i).Set(function)
//		fmt.Printf("Set field %v, type %v to value %v\n", field.Name, fieldType, function)
	}
}