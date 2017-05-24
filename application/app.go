package application

const (
	STATUS_READY   string = "Listo"
	STATUS_RUNNING string = "Corriendo"
)

var (
	serverStatus string = STATUS_READY
)

func StartApp() {
	serverStatus = STATUS_RUNNING
}
