// Set the file and directory name
msiFile := `package.msi`
targetDir := `C:\test folder`

// Set to the misexec application, but don't pass command line arguments
cmd := exec.Command("msiexec")

// Manually set the command line arguments so they are not escaped
cmd.SysProcAttr = &syscall.SysProcAttr{
    HideWindow:    false,
    CmdLine:       fmt.Sprintf(` /a "%v" TARGETDIR="%v"`, msiFile, targetDir), // Leave a space at the beginning
    CreationFlags: 0,
}

// Run the install
err := cmd.Run()
if err != nil {
    log.Fatalln(err)
}