package jail

import (
	"fmt"
	"log"
	"os"
)

// WriteConfig generates a FreeBSD jail.conf representing
// the passed jail struct
func WriteConfig(j Jail) error {

	file, err := os.Create(fmt.Sprintf("%s.conf", j.Name))
	if err != nil {
		log.Fatal("Cannot create config file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s {\n", j.Name)
	fmt.Fprintf(file, "	host.hostname = %s;\n", j.Hostname)
	fmt.Fprintf(file, "	ip4.addr = %s;\n", j.IP)
	fmt.Fprintf(file, `	path = "%s";%s`, j.Path, "\n")
	fmt.Fprintf(file, "	mount.devfs;\n")
	fmt.Fprintf(file, `	exec.start = "%s";%s`, j.ExecStart, "\n")
	fmt.Fprintf(file, `	exec.stop = "%s";%s`, j.ExecStop, "\n")
	fmt.Fprintf(file, "}\n")

	return nil
}
