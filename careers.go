package main

type qualification struct {
	trait string
	value int
}

type career struct {
	name          string
	qualification qualification
}

func (c career) String() string {
	return c.name
}

func professions() []career {
	return []career{
		{"aerospace system defense", qualification{"en", 5}},
		{"agent", qualification{"ss", 5}},
		{"athlete", qualification{"en", 8}},
		{"barbarian", qualification{"en", 5}},
		{"belter", qualification{"in", 4}},
		{"bureaucrat", qualification{"ss", 6}},
		{"colonist", qualification{"en", 5}},
		{"diplomat", qualification{"ss", 6}},
		{"drifter", qualification{"de", 5}},
		{"entertainer", qualification{"ss", 8}},
		{"hunter", qualification{"en", 5}},
		{"marine", qualification{"in", 6}},
		{"maritime system defense", qualification{"en", 5}},
		{"mercenary", qualification{"in", 4}},
		{"merchant", qualification{"in", 4}},
		{"navy", qualification{"in", 6}},
		{"noble", qualification{"ss", 8}},
		{"physician", qualification{"ed", 6}},
		{"pirate", qualification{"de", 5}},
		{"rogue", qualification{"de", 5}},
		{"scientist", qualification{"ed", 6}},
		{"scout", qualification{"in", 6}},
		{"surface system defense", qualification{"en", 5}},
		{"technician", qualification{"ed", 6}},
	}
}
