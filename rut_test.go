package ruttools

import "testing"

func TestGetDigit(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"6641891", "K"},
		{"47277337", "2"},
		{"21701460", "3"},
		{"12488882", "4"},
		{"23051153", "5"},
		{"32968060", "6"},
		{"53254916", "7"},
		{"65445455", "8"},
		{"64805350", "9"},
		{"84937308", "0"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			answer := GetDigit(tt.input)
			if answer != tt.want {
				t.Errorf("Rut %s expected %s, got %s", tt.input, tt.want, answer)
			}
		})
	}

}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"12679966-1", false},
		{"21858623-3", false},
		{"43527620-5", false},
		{"68562886-7", false},
		{"81114287-9", false},
		{"60469244-K", true},
		{"42028320-2", true},
		{"29843707-4", true},
		{"46335669-6", true},
		{"16133363-8", true},
		{"22700107-0", true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			answer := IsValid(tt.input)
			if answer != tt.want {
				t.Errorf("Rut %s expected %v, got %v", tt.input, tt.want, answer)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"99-030-307-K", "99030307-K"},
		{"94.228.360-1", "94228360-1"},
		{"47/277/337/2", "47277337-2"},
		{"21*701*460-3", "21701460-3"},
		{"12+488+882+4", "12488882-4"},
		{"     23 051    153    5     ", "23051153-5"},
		{"32_968_060_6", "32968060-6"},
		{"532549167", "53254916-7"},
		{"65,445,455-8", "65445455-8"},
		{"64+805*350/9", "64805350-9"},
		{"00000000849373080", "84937308-0"},
		{"        00000033*343+244_7", "33343244-7"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			answer := Format(tt.input)
			if answer != tt.want {
				t.Errorf("Rut %s expected %v, got %v", tt.input, tt.want, answer)
			}
		})
	}
}
