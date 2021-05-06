package model

import (
	"fmt"
	"time"
)

type Side string

func (s Side) Opposite() Side {
	switch s {
	case Left:
		return Right
	default:
		return Left
	}
}

func (s Side) String() string {
	return string(s)
}

const (
	Left Side = "left"
	Right Side = "right"
)

func SpeakAndPause(speech string, pauseSec int64) Node {
	return SpeechNode{
		Template: speech,
		Next: PauseNode{
			Duration: time.Second * time.Duration(pauseSec),
		},
	}
}

func InhaleExhale(times int, inSec, exSec int64) Node {
	return LoopNode{
		Node: Sequence{
			SpeakAndPause("inhale", inSec),
			SpeakAndPause("exhale", exSec),
		},
		Times: times,
	}
}
func ThreadTheNeedle(side string) Node {
	return Sequence{
		SpeakAndPause("Slowly return to tabletop.", 2),
		SpeakAndPause(fmt.Sprintf("Inhale %v arm up.", side), 2),
		SpeakAndPause("Exhale thread the needle.", 2),
		InhaleExhale(3, 4, 6),
	}
}

func TadasanaFromFold() Node {
	return Sequence{
		SpeakAndPause("Exhale to fold.", 2),
		SpeakAndPause("Inhate Flat back.", 2),
		SpeakAndPause("Exhale to fold.", 2),
		SpeakAndPause("Inhale, Urdva Hastasana.", 2),
		SpeakAndPause("Arms down, Tadasana.", 2),
	}
}

func TadasanaToFold() Node {
	return Sequence{
		SpeakAndPause("Stand in Tadasana and begin to settle in.", 2),
		SpeakAndPause("Ground into your arches. Activate your glutes. Activate your bandas.", 5),
		InhaleExhale(2, 4, 6),
		SpeakAndPause("Inhale, Urdva Hastasana.", 3),
		SpeakAndPause("Forward fold.", 3),
	}
}

func DownDogFromFold() Node {
	return Sequence{
		SpeakAndPause("Inhate Flat back.", 3),
		SpeakAndPause("Exhale to fold.", 3),
		SpeakAndPause("Inhale to plank.", 3),
		SpeakAndPause("Exhale to belly.", 3),
		SpeakAndPause("Inhale to cobra. Active the whole posterior chain. Keep belly activated.", 4),
		SpeakAndPause("Exhale here.", 3),
		SpeakAndPause("Inhale here.", 2),
		SpeakAndPause("Exhale down.", 2),
		SpeakAndPause("Inhale legs and chest engaged.", 1),
		SpeakAndPause("Exhale back and up to down dog.", 4),
	}
}

func ChairFromDownDog() Node {
	return Sequence{
		SpeakAndPause("Walk your feet to your hands.", 2),
		SpeakAndPause("Exhale forward fold.", 2),
		SpeakAndPause("Inhale Chair pose.", 2),
		SpeakAndPause("Tailbone tucked under, breathing into belly, belly and bandas pulled in.", 6),
		SpeakAndPause("Exhale and fold.", 4),
		SpeakAndPause("Option to add a strap here.", 3),
		SpeakAndPause("Inhale Chair pose.", 2),
		InhaleExhale(3, 4, 6),
		SpeakAndPause("Forward fold.", 3),
	}
}

func GoddessPoseFromTadasana() Node {
	return Sequence{
		SpeakAndPause("Walk your feet out", 5),
		SpeakAndPause("Goddess Pose", 5),
		SpeakAndPause("Bandas engaged. Ujayii Breath", 5),
		SpeakAndPause("Exhale and forward fold", 5),
		SpeakAndPause("Keep legs strong. push into outer foot while lifting or relaxing toes", 5),
		SpeakAndPause("Slowly roll it up on inhale. Bandas engaged.", 5),
		SpeakAndPause("Walk feet back together. Tadasana.", 5),
	}
}

func ClamshellGlutesFromDownDog() Node {
	return Sequence{
		SpeakAndPause("Bend your knees and make your way onto your left side. Make sure you have your strap.", 5),
		SpeakAndPause("Reclining pose on left side. Put strap around thighs.", 2),
		SpeakAndPause("Raise your right thigh to external rotation position.", 2),
		SpeakAndPause("Exhale and rotate leg keeping thigh at same height.", 2),
		SpeakAndPause("Inhale and rotate back.", 2),
		SpeakAndPause("Continue for 10 seconds.", 5),
		SpeakAndPause("3 more.", 5),
		SpeakAndPause("Switch to your other side.", 6),
		SpeakAndPause("Raise your right thigh to external rotation position.", 2),
		SpeakAndPause("Exhale and rotate leg keeping thigh at same height.", 2),
		SpeakAndPause("Inhale and rotate back.", 2),
		SpeakAndPause("Continue for 10 seconds.", 5),
		SpeakAndPause("3 more.", 5),
		SpeakAndPause("Roll back to your stomach and press up to table top", 5),
		SpeakAndPause("Child's pose", 5),
		SpeakAndPause("Press back to down dog", 5),
	}
}

func WarriorSequenceFromDownDog() Node {
	return Sequence{
		WarriorWarmupFromDownDog("right"),
		WarriorWarmupFromDownDog("left"),
		WarriorStandingFromDownDog("right"),
		WarriorStandingFromDownDog("left"),
	}
}

func WarriorWarmupFromDownDog(side Side) Node {
	return Sequence{
		SpeakAndPause("Exhale in down dog", 5),
		SpeakAndPause("Inhale and raise your "+side.String()+" leg.", 5),
		SpeakAndPause("Exhale and bring your weight forward, bringing "+side.String()+" leg to your chest.", 4),
		SpeakAndPause("Inhale and raise your "+side.String()+" leg.", 4),
		SpeakAndPause("Exhale and bring your weight forward, bringing "+side.String()+" foot to the floor directly under your knee.", 5),
		SpeakAndPause("Back knee down. Inhale up slowly into crescent lunge. Stretch your arms back into angel, opening chest.", 5),
		SpeakAndPause("Fold back to low lunge.", 5),
		SpeakAndPause("Inhale here.", 5),
		SpeakAndPause("Step back to down dog.", 10),
	}
}

func WarriorStandingFromDownDog(side Side) Node {
	return Sequence{
		SpeakAndPause("Exhale in down dog", 5),
		SpeakAndPause("Inhale and raise your "+side.String()+" leg.", 5),
		SpeakAndPause("Exhale and bring your weight forward, bringing "+side.String()+" foot to the floor directly under your knee.", 5),
		SpeakAndPause("Front and back legs are strong, core engaged. Inhale up slowly into high lunge. Stretch your arms back into angel, opening chest.", 8),
		SpeakAndPause("Breathe here.", 5),
		SpeakAndPause("On your next exhale open up to warrior 2. Be mindful of your arches and glutes. Tummy pulls in.", 10),

		SpeakAndPause("Inhale here.", 5),
		SpeakAndPause("Exhale into peaceful warrior. Expand the "+side.String()+" side of your ribcage. Back hand slowly traces down back leg. Keep arches engaged here.", 10),
		SpeakAndPause("Breathe here.", 10),
		SpeakAndPause("On your next inhale tilt forward into extended side angle.", 10),
		SpeakAndPause("Breathe here.", 10),
		SpeakAndPause("Inhale again back up to warrior 2. Breathe here.", 10),
		SpeakAndPause("Inhale, one more peaeful warrior. Arches strong", 6),
		SpeakAndPause("Exhale back to low lunge", 4),
		SpeakAndPause("Inhale here", 4),
		SpeakAndPause("Back to down dog", 6),
	}
}
