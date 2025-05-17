package git_commands

type CloneCommands struct {
	*GitCommon
}

func NewCloneCommands(gitCommon *GitCommon) *CloneCommands {
	return &CloneCommands{GitCommon: gitCommon}
}

func (self *CloneCommands) SetRemoteRepository(url string, name string) error {
	cmdArgs := NewGitCmd("clone").Arg(url).Arg(name).ToArgv()

	return self.cmd.New(cmdArgs).Run()
}
