package storage

type StorageManager interface {
	initConfig()
	saveApp(App)
	getApps() []App
}
