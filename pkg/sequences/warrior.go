package sequences

import (
	. "github.com/ilackarms/yogayaml/pkg/model"
)

func WarriorSequence() Sequence {
	return Sequence{
		WarriorSequenceFromDownDog(),
	}
}
