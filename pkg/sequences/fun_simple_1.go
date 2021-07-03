package sequences

import (
	. "github.com/ilackarms/yogayaml/pkg/model"
)

func StartsWithFeet() Node {
	return Sequence{
		SpeakAndPause("Begin seated.", 4),
		SpeakAndPause("Banda awareness. Breathing from the bottom up. Shoulders relax.", 4),
		SpeakAndPause("", 4),
	}
}

func FunSimple() Node {
	return Sequence{
		SpeakAndPause("Make your way into tadasana.", 6),
		SpeakAndPause("Make your way into tadasana.", 6),
		SpeakAndPause("Begin by setting an intention for your practice. Breathe and ground here.", 4),
		TadasanaToFold(),
		TabletopStretch(),
		ThreadTheNeedle(Left),
		ThreadTheNeedle(Right),
		SpeakAndPause("Inale in tabletop, activate your core.", 4),
		SpeakAndPause("Press up to hovering tabletop. Breathe here. Stretch out the toes.", 8),
		SpeakAndPause("Knees to the floor. Inhale here.", 4),
		DownDogStretch(),
		ChairFromDownDog(),
		SpeakAndPause("Inhale here, then exhale walk your hand up to your feet. Forward fold.", 6),
		InhaleExhale(1, 4, 4),
		SpeakAndPause("Inhale again.", 2),
		TadasanaFromFold(),
		SpeakAndPause("Recall your intention.", 6),
		SunSalutation(),
		WarriorWarmupFromDownDog(Right),
		WarriorWarmupFromDownDog(Left),
		SpeakAndPause("Nurture your intention.", 6),
		WarriorStandingFromDownDog(Right),
		WarriorStandingFromDownDog(Left),
		SunSalutation(),
		SpeakAndPause("Integrate your intention here. Nice work. Take your yoga with you off your mat.", 6),
	}
}
