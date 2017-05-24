package application

import (
	"testing"
)

func TestApplicationInitialStatus(t *testing.T) {
	exptd := "Listo"
	if serverStatus != exptd {
		t.Errorf("El estado debería ser '%s'", exptd)
	}
}

func TestApplicationRunningStatus(t *testing.T) {
	exptd := "Listo"
	if serverStatus != exptd {
		t.Errorf("El estado debería ser '%s'", exptd)
	}

	StartApp()

	exptd = "Corriendo"
	if serverStatus != exptd {
		t.Errorf("La aplicación corriendo debería tener estado '%s'", exptd)
	}
}
