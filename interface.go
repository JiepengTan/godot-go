package godotgo

type EngineCallbacks struct {
	OnEngineStart       func()
	OnEngineUpdate      func(float32)
	OnEngineFixedUpdate func(float32)
	OnEngineDestroy     func()

	OnKeyPressed  func(int64)
	OnKeyReleased func(int64)
}

type CallbackInfo struct {
	EngineCallbacks
}
