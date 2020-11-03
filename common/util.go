package common

import (
	"net"
	"os"
	"text/template"
)

func GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
		}
	}
	return "", nil
}

func WriteFile(filename, content string) (int, error) {
	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}

	return file.WriteString(content)
}

func RenderTemplateFile(text string, content interface{}, outputFile string) error {
	tmp, err := template.New(outputFile).Parse(text)
	if err != nil {
		return err
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	return tmp.Execute(file, content)
}
