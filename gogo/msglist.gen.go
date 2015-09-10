package fuzztests

var NewFuncs = []func() proto.Message{
	func() proto.Message { return &NinOptNative{} },
	func() proto.Message { return &NinRepNative{} },
	func() proto.Message { return &NinRepPackedNative{} },
	func() proto.Message { return &NinOptStruct{} },
	func() proto.Message { return &NinRepStruct{} },
	func() proto.Message { return &NinNestedStruct{} },
	func() proto.Message { return &Nil{} },
	func() proto.Message { return &NinOptEnum{} },
	func() proto.Message { return &NinRepEnum{} },
	func() proto.Message { return &NinOptEnumDefault{} },
	func() proto.Message { return &MyExtendable{} },
	func() proto.Message { return &OtherExtenable{} },
	func() proto.Message { return &NestedDefinition{} },
	func() proto.Message { return &NestedScope{} },
	func() proto.Message { return &NinOptNativeDefault{} },
	func() proto.Message { return &NinOptNative3{} },
	func() proto.Message { return &NinRepNative3{} },
	func() proto.Message { return &NinRepPackedNative3{} },
	func() proto.Message { return &NinOptStruct3{} },
	func() proto.Message { return &NinRepStruct3{} },
	func() proto.Message { return &NinNestedStruct3{} },
	func() proto.Message { return &Nil3{} },
	func() proto.Message { return &NinOptEnum3{} },
	func() proto.Message { return &NinRepEnum3{} },
	func() proto.Message { return &NestedDefinition3{} },
	func() proto.Message { return &NestedScope3{} },
}