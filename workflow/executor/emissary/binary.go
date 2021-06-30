package emissary

import (
	"io"
	"os"
	"os/exec"
)

func copyBinary() error {
	name, err := exec.LookPath("argoexec")
	if err != nil {
		return err
	}
	in, err := os.Open(name)
	if err != nil {
		return err
	}
	defer func() { _ = in.Close() }()
	// argoexec needs to be executable from non-root user in the main container.
	// Therefore we set permission 0o555 == r-xr-xr-x.
	out, err := os.OpenFile("/var/run/argo/argoexec", os.O_RDWR|os.O_CREATE, 0o555)
	if err != nil {
		return err
	}
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
