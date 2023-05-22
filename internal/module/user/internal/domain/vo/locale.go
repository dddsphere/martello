package vo

type (
	Locale struct {
		Language  string
		Territory string
		Codeset   string
		Modifier  string
	}
)

func (l Locale) String() string {
	posixString := l.Language

	if l.Territory != "" {
		posixString += "_" + l.Territory
	}

	if l.Codeset != "" {
		posixString += "." + l.Codeset
	}

	if l.Modifier != "" {
		posixString += "@" + l.Modifier
	}

	return posixString
}
