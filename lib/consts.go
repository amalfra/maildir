package maildir

// the seperator between unique name and info
const colon = ':'

// default info, to which flags are appended
const info = "2,"

// subdirectories that are required in maildir
var subdirs = []string{"tmp", "new", "cur"}
