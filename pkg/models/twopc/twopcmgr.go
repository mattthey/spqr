package twopc

type twopcmgr interface {
	Create2PhaseCommit(txid string, msg string) error
	//GetAll2PhaseCommits() (map[string][]string, error)
	//Delete2PhaseCommit(txid string, msg string) error
}
