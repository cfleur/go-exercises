package main
import "os/exec"
func main() {
    exec.Command("open", "-a", "firefox", "https://courses.packtpub.com/enrollments").Run()
}