def getCommitHash() {
    def hash = sh 'git rev-parse --short origin/main'
    return hash
}

return this