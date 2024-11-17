/*------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
package api

import "unsafe"

var (
	FFI GDExtensionInterface
)

type GDExtensionInterface struct {
	GetProcAddress                                func(p_function_name unsafe.Pointer)
	GetGodotVersion                               func(r_godot_version unsafe.Pointer)
	MemAlloc                                      func(p_bytes int64)
	MemRealloc                                    func(p_ptr unsafe.Pointer, p_bytes int64)
	MemFree                                       func(p_ptr unsafe.Pointer)
	PrintError                                    func(p_description unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	PrintErrorWithMessage                         func(p_description unsafe.Pointer, p_message unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	PrintWarning                                  func(p_description unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	PrintWarningWithMessage                       func(p_description unsafe.Pointer, p_message unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	PrintScriptError                              func(p_description unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	PrintScriptErrorWithMessage                   func(p_description unsafe.Pointer, p_message unsafe.Pointer, p_function unsafe.Pointer, p_file unsafe.Pointer, p_line int32, p_editor_notify bool)
	GetNativeStructSize                           func(p_name unsafe.Pointer)
	VariantNewCopy                                func(r_dest unsafe.Pointer, p_src unsafe.Pointer)
	VariantNewNil                                 func(r_dest unsafe.Pointer)
	VariantDestroy                                func(p_self unsafe.Pointer)
	VariantCall                                   func(p_self unsafe.Pointer, p_method unsafe.Pointer, p_args unsafe.Pointer, p_argument_count int64, r_return unsafe.Pointer, r_error unsafe.Pointer)
	VariantCallStatic                             func(p_type any /*VariantType*/, p_method unsafe.Pointer, p_args unsafe.Pointer, p_argument_count int64, r_return unsafe.Pointer, r_error unsafe.Pointer)
	VariantEvaluate                               func(p_op any /*VariantOperator*/, p_a unsafe.Pointer, p_b unsafe.Pointer, r_return unsafe.Pointer, r_valid unsafe.Pointer)
	VariantSet                                    func(p_self unsafe.Pointer, p_key unsafe.Pointer, p_value unsafe.Pointer, r_valid unsafe.Pointer)
	VariantSetNamed                               func(p_self unsafe.Pointer, p_key unsafe.Pointer, p_value unsafe.Pointer, r_valid unsafe.Pointer)
	VariantSetKeyed                               func(p_self unsafe.Pointer, p_key unsafe.Pointer, p_value unsafe.Pointer, r_valid unsafe.Pointer)
	VariantSetIndexed                             func(p_self unsafe.Pointer, p_index int64, p_value unsafe.Pointer, r_valid unsafe.Pointer, r_oob unsafe.Pointer)
	VariantGet                                    func(p_self unsafe.Pointer, p_key unsafe.Pointer, r_ret unsafe.Pointer, r_valid unsafe.Pointer)
	VariantGetNamed                               func(p_self unsafe.Pointer, p_key unsafe.Pointer, r_ret unsafe.Pointer, r_valid unsafe.Pointer)
	VariantGetKeyed                               func(p_self unsafe.Pointer, p_key unsafe.Pointer, r_ret unsafe.Pointer, r_valid unsafe.Pointer)
	VariantGetIndexed                             func(p_self unsafe.Pointer, p_index int64, r_ret unsafe.Pointer, r_valid unsafe.Pointer, r_oob unsafe.Pointer)
	VariantIterInit                               func(p_self unsafe.Pointer, r_iter unsafe.Pointer, r_valid unsafe.Pointer)
	VariantIterNext                               func(p_self unsafe.Pointer, r_iter unsafe.Pointer, r_valid unsafe.Pointer)
	VariantIterGet                                func(p_self unsafe.Pointer, r_iter unsafe.Pointer, r_ret unsafe.Pointer, r_valid unsafe.Pointer)
	VariantHash                                   func(p_self unsafe.Pointer)
	VariantRecursiveHash                          func(p_self unsafe.Pointer, p_recursion_count int64)
	VariantHashCompare                            func(p_self unsafe.Pointer, p_other unsafe.Pointer)
	VariantBooleanize                             func(p_self unsafe.Pointer)
	VariantDuplicate                              func(p_self unsafe.Pointer, r_ret unsafe.Pointer, p_deep bool)
	VariantStringify                              func(p_self unsafe.Pointer, r_ret unsafe.Pointer)
	VariantGetType                                func(p_self unsafe.Pointer)
	VariantHasMethod                              func(p_self unsafe.Pointer, p_method unsafe.Pointer)
	VariantHasMember                              func(p_type any /*VariantType*/, p_member unsafe.Pointer)
	VariantHasKey                                 func(p_self unsafe.Pointer, p_key unsafe.Pointer, r_valid unsafe.Pointer)
	VariantGetTypeName                            func(p_type any /*VariantType*/, r_name unsafe.Pointer)
	VariantCanConvert                             func(p_from any /*VariantType*/, p_to any /*VariantType*/)
	VariantCanConvertStrict                       func(p_from any /*VariantType*/, p_to any /*VariantType*/)
	GetVariantFromTypeConstructor                 func(p_type any /*VariantType*/)
	GetVariantToTypeConstructor                   func(p_type any /*VariantType*/)
	VariantGetPtrOperatorEvaluator                func(p_operator any /*VariantOperator*/, p_type_a any /*VariantType*/, p_type_b any /*VariantType*/)
	VariantGetPtrBuiltinMethod                    func(p_type any /*VariantType*/, p_method unsafe.Pointer, p_hash int64)
	VariantGetPtrConstructor                      func(p_type any /*VariantType*/, p_constructor int32)
	VariantGetPtrDestructor                       func(p_type any /*VariantType*/)
	VariantConstruct                              func(p_type any /*VariantType*/, r_base unsafe.Pointer, p_args unsafe.Pointer, p_argument_count int32, r_error unsafe.Pointer)
	VariantGetPtrSetter                           func(p_type any /*VariantType*/, p_member unsafe.Pointer)
	VariantGetPtrGetter                           func(p_type any /*VariantType*/, p_member unsafe.Pointer)
	VariantGetPtrIndexedSetter                    func(p_type any /*VariantType*/)
	VariantGetPtrIndexedGetter                    func(p_type any /*VariantType*/)
	VariantGetPtrKeyedSetter                      func(p_type any /*VariantType*/)
	VariantGetPtrKeyedGetter                      func(p_type any /*VariantType*/)
	VariantGetPtrKeyedChecker                     func(p_type any /*VariantType*/)
	VariantGetConstantValue                       func(p_type any /*VariantType*/, p_constant unsafe.Pointer, r_ret unsafe.Pointer)
	VariantGetPtrUtilityFunction                  func(p_function unsafe.Pointer, p_hash int64)
	StringNewWithLatin1Chars                      func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNewWithUtf8Chars                        func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNewWithUtf16Chars                       func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNewWithUtf32Chars                       func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNewWithWideChars                        func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNewWithLatin1CharsAndLen                func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_size int64)
	StringNewWithUtf8CharsAndLen                  func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_size int64)
	StringNewWithUtf8CharsAndLen2                 func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_size int64)
	StringNewWithUtf16CharsAndLen                 func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_char_count int64)
	StringNewWithUtf16CharsAndLen2                func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_char_count int64, p_default_little_endian bool)
	StringNewWithUtf32CharsAndLen                 func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_char_count int64)
	StringNewWithWideCharsAndLen                  func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_char_count int64)
	StringToLatin1Chars                           func(p_self unsafe.Pointer, r_text unsafe.Pointer, p_max_write_length int64)
	StringToUtf8Chars                             func(p_self unsafe.Pointer, r_text unsafe.Pointer, p_max_write_length int64)
	StringToUtf16Chars                            func(p_self unsafe.Pointer, r_text unsafe.Pointer, p_max_write_length int64)
	StringToUtf32Chars                            func(p_self unsafe.Pointer, r_text unsafe.Pointer, p_max_write_length int64)
	StringToWideChars                             func(p_self unsafe.Pointer, r_text unsafe.Pointer, p_max_write_length int64)
	StringOperatorIndex                           func(p_self unsafe.Pointer, p_index int64)
	StringOperatorIndexConst                      func(p_self unsafe.Pointer, p_index int64)
	StringOperatorPlusEqString                    func(p_self unsafe.Pointer, p_b unsafe.Pointer)
	StringOperatorPlusEqChar                      func(p_self unsafe.Pointer, p_b rune)
	StringOperatorPlusEqCstr                      func(p_self unsafe.Pointer, p_b unsafe.Pointer)
	StringOperatorPlusEqWcstr                     func(p_self unsafe.Pointer, p_b unsafe.Pointer)
	StringOperatorPlusEqC32str                    func(p_self unsafe.Pointer, p_b unsafe.Pointer)
	StringResize                                  func(p_self unsafe.Pointer, p_resize int64)
	StringNameNewWithLatin1Chars                  func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_is_static bool)
	StringNameNewWithUtf8Chars                    func(r_dest unsafe.Pointer, p_contents unsafe.Pointer)
	StringNameNewWithUtf8CharsAndLen              func(r_dest unsafe.Pointer, p_contents unsafe.Pointer, p_size int64)
	XmlParserOpenBuffer                           func(p_instance unsafe.Pointer, p_buffer unsafe.Pointer, p_size int64)
	FileAccessStoreBuffer                         func(p_instance unsafe.Pointer, p_src unsafe.Pointer, p_length uint64)
	FileAccessGetBuffer                           func(p_instance unsafe.Pointer, p_dst unsafe.Pointer, p_length uint64)
	ImagePtrw                                     func(p_instance unsafe.Pointer)
	ImagePtr                                      func(p_instance unsafe.Pointer)
	WorkerThreadPoolAddNativeGroupTask            func(p_instance unsafe.Pointer, _funcPtr1 unsafe.Pointer /*void(*p_func)(void * ,uint32_t)*/, p_userdata unsafe.Pointer, p_elements int, p_tasks int, p_high_priority bool, p_description unsafe.Pointer)
	WorkerThreadPoolAddNativeTask                 func(p_instance unsafe.Pointer, _funcPtr1 unsafe.Pointer /*void(*p_func)(void * )*/, p_userdata unsafe.Pointer, p_high_priority bool, p_description unsafe.Pointer)
	PackedByteArrayOperatorIndex                  func(p_self unsafe.Pointer, p_index int64)
	PackedByteArrayOperatorIndexConst             func(p_self unsafe.Pointer, p_index int64)
	PackedFloat32ArrayOperatorIndex               func(p_self unsafe.Pointer, p_index int64)
	PackedFloat32ArrayOperatorIndexConst          func(p_self unsafe.Pointer, p_index int64)
	PackedFloat64ArrayOperatorIndex               func(p_self unsafe.Pointer, p_index int64)
	PackedFloat64ArrayOperatorIndexConst          func(p_self unsafe.Pointer, p_index int64)
	PackedInt32ArrayOperatorIndex                 func(p_self unsafe.Pointer, p_index int64)
	PackedInt32ArrayOperatorIndexConst            func(p_self unsafe.Pointer, p_index int64)
	PackedInt64ArrayOperatorIndex                 func(p_self unsafe.Pointer, p_index int64)
	PackedInt64ArrayOperatorIndexConst            func(p_self unsafe.Pointer, p_index int64)
	PackedStringArrayOperatorIndex                func(p_self unsafe.Pointer, p_index int64)
	PackedStringArrayOperatorIndexConst           func(p_self unsafe.Pointer, p_index int64)
	PackedVector2ArrayOperatorIndex               func(p_self unsafe.Pointer, p_index int64)
	PackedVector2ArrayOperatorIndexConst          func(p_self unsafe.Pointer, p_index int64)
	PackedVector3ArrayOperatorIndex               func(p_self unsafe.Pointer, p_index int64)
	PackedVector3ArrayOperatorIndexConst          func(p_self unsafe.Pointer, p_index int64)
	PackedVector4ArrayOperatorIndex               func(p_self unsafe.Pointer, p_index int64)
	PackedVector4ArrayOperatorIndexConst          func(p_self unsafe.Pointer, p_index int64)
	PackedColorArrayOperatorIndex                 func(p_self unsafe.Pointer, p_index int64)
	PackedColorArrayOperatorIndexConst            func(p_self unsafe.Pointer, p_index int64)
	ArrayOperatorIndex                            func(p_self unsafe.Pointer, p_index int64)
	ArrayOperatorIndexConst                       func(p_self unsafe.Pointer, p_index int64)
	ArrayRef                                      func(p_self unsafe.Pointer, p_from unsafe.Pointer)
	ArraySetTyped                                 func(p_self unsafe.Pointer, p_type any /*VariantType*/, p_class_name unsafe.Pointer, p_script unsafe.Pointer)
	DictionaryOperatorIndex                       func(p_self unsafe.Pointer, p_key unsafe.Pointer)
	DictionaryOperatorIndexConst                  func(p_self unsafe.Pointer, p_key unsafe.Pointer)
	ObjectMethodBindCall                          func(p_method_bind unsafe.Pointer, p_instance unsafe.Pointer, p_args unsafe.Pointer, p_arg_count int64, r_ret unsafe.Pointer, r_error unsafe.Pointer)
	ObjectMethodBindPtrcall                       func(p_method_bind unsafe.Pointer, p_instance unsafe.Pointer, p_args unsafe.Pointer, r_ret unsafe.Pointer)
	ObjectDestroy                                 func(p_o unsafe.Pointer)
	GlobalGetSingleton                            func(p_name unsafe.Pointer)
	ObjectGetInstanceBinding                      func(p_o unsafe.Pointer, p_token unsafe.Pointer, p_callbacks unsafe.Pointer)
	ObjectSetInstanceBinding                      func(p_o unsafe.Pointer, p_token unsafe.Pointer, p_binding unsafe.Pointer, p_callbacks unsafe.Pointer)
	ObjectFreeInstanceBinding                     func(p_o unsafe.Pointer, p_token unsafe.Pointer)
	ObjectSetInstance                             func(p_o unsafe.Pointer, p_classname unsafe.Pointer, p_instance unsafe.Pointer)
	ObjectGetClassName                            func(p_object unsafe.Pointer, p_library unsafe.Pointer, r_class_name unsafe.Pointer)
	ObjectCastTo                                  func(p_object unsafe.Pointer, p_class_tag unsafe.Pointer)
	ObjectGetInstanceFromId                       func(p_instance_id int64)
	ObjectGetInstanceId                           func(p_object unsafe.Pointer)
	ObjectHasScriptMethod                         func(p_object unsafe.Pointer, p_method unsafe.Pointer)
	ObjectCallScriptMethod                        func(p_object unsafe.Pointer, p_method unsafe.Pointer, p_args unsafe.Pointer, p_argument_count int64, r_return unsafe.Pointer, r_error unsafe.Pointer)
	RefGetObject                                  func(p_ref unsafe.Pointer)
	RefSetObject                                  func(p_ref unsafe.Pointer, p_object unsafe.Pointer)
	ScriptInstanceCreate                          func(p_info unsafe.Pointer, p_instance_data unsafe.Pointer)
	ScriptInstanceCreate2                         func(p_info unsafe.Pointer, p_instance_data unsafe.Pointer)
	ScriptInstanceCreate3                         func(p_info unsafe.Pointer, p_instance_data unsafe.Pointer)
	PlaceHolderScriptInstanceCreate               func(p_language unsafe.Pointer, p_script unsafe.Pointer, p_owner unsafe.Pointer)
	PlaceHolderScriptInstanceUpdate               func(p_placeholder unsafe.Pointer, p_properties unsafe.Pointer, p_values unsafe.Pointer)
	ObjectGetScriptInstance                       func(p_object unsafe.Pointer, p_language unsafe.Pointer)
	CallableCustomCreate                          func(r_callable unsafe.Pointer, p_callable_custom_info unsafe.Pointer)
	CallableCustomCreate2                         func(r_callable unsafe.Pointer, p_callable_custom_info unsafe.Pointer)
	CallableCustomGetUserData                     func(p_callable unsafe.Pointer, p_token unsafe.Pointer)
	ClassdbConstructObject                        func(p_classname unsafe.Pointer)
	ClassdbGetMethodBind                          func(p_classname unsafe.Pointer, p_methodname unsafe.Pointer, p_hash int64)
	ClassdbGetClassTag                            func(p_classname unsafe.Pointer)
	ClassdbRegisterExtensionClass                 func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_parent_class_name unsafe.Pointer, p_extension_funcs unsafe.Pointer)
	ClassdbRegisterExtensionClass2                func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_parent_class_name unsafe.Pointer, p_extension_funcs unsafe.Pointer)
	ClassdbRegisterExtensionClass3                func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_parent_class_name unsafe.Pointer, p_extension_funcs unsafe.Pointer)
	ClassdbRegisterExtensionClassMethod           func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_method_info unsafe.Pointer)
	ClassdbRegisterExtensionClassVirtualMethod    func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_method_info unsafe.Pointer)
	ClassdbRegisterExtensionClassIntegerConstant  func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_enum_name unsafe.Pointer, p_constant_name unsafe.Pointer, p_constant_value int64, p_is_bitfield bool)
	ClassdbRegisterExtensionClassProperty         func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_info unsafe.Pointer, p_setter unsafe.Pointer, p_getter unsafe.Pointer)
	ClassdbRegisterExtensionClassPropertyIndexed  func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_info unsafe.Pointer, p_setter unsafe.Pointer, p_getter unsafe.Pointer, p_index int64)
	ClassdbRegisterExtensionClassPropertyGroup    func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_group_name unsafe.Pointer, p_prefix unsafe.Pointer)
	ClassdbRegisterExtensionClassPropertySubgroup func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_subgroup_name unsafe.Pointer, p_prefix unsafe.Pointer)
	ClassdbRegisterExtensionClassSignal           func(p_library unsafe.Pointer, p_class_name unsafe.Pointer, p_signal_name unsafe.Pointer, p_argument_info unsafe.Pointer, p_argument_count int64)
	ClassdbUnregisterExtensionClass               func(p_library unsafe.Pointer, p_class_name unsafe.Pointer)
	GetLibraryPath                                func(p_library unsafe.Pointer, r_path unsafe.Pointer)
	EditorAddPlugin                               func(p_class_name unsafe.Pointer)
	EditorRemovePlugin                            func(p_class_name unsafe.Pointer)
}
