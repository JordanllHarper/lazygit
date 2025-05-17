package git_commands

type CloneCommands struct {
	*GitCommon
}

func NewCloneCommands(gitCommon *GitCommon) *CloneCommands {
	return &CloneCommands{GitCommon: gitCommon}
}

func (self *CloneCommands) SetRemoteRepository(url string) error {
	cmdArgs := NewGitCmd("clone").Arg(url).ToArgv()

	return self.cmd.New(cmdArgs).Run()
}
