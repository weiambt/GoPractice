package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> command args
// go run main.go run command args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

// 主进程(root权限)
func run() {

	cmd := exec.Command(
		"/proc/self/exe",
		append([]string{"child"}, //执行self的child() function
			os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//UTS PID namespace隔离
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
		//syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("主进程cmd run: %+v \n", cmd)
	_, _ = cmd.Process.Wait()

}

// 子进程，也就是容器进程
func child() {
	fmt.Printf("容器内部进程: running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	must(syscall.Chroot("/home/rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	syscall.Unmount("/proc", 0)

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
