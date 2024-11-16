package godotgo

type EngineCallbacks struct {
	OnEngineStart   func()
	OnEngineUpdate  func(delta float32)
	OnEngineDestroy func()
}

func LinkEngine(info EngineCallbacks) {
}
