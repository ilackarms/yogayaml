package sequences

import (
	. "github.com/ilackarms/yogayaml/pkg/model"
)

func LyingOnBackWarmupToDownDog() Sequence {
	return Sequence{
		SpeakAndPause("Lie down on your back and breathe", 8),
		SpeakAndPause("Morning stretch", 4),
		SpeakAndPause("Breathe here.", 8),
		SpeakAndPause("Knees to your chest or happy baby", 4),
		SpeakAndPause("Breathe here.", 8),
		SpeakAndPause("Exhale, Let your knees drop to the left", 6),
		SpeakAndPause("Inhale back to center", 2),
		SpeakAndPause("Let your knees drop to the right", 6),
		SpeakAndPause("Inhale back to center", 2),
		SpeakAndPause("rock gently along your spine", 3),
		SpeakAndPause("roll up to a sitting position and press back to down dog", 8),
	}
}

func LowBackHipsGlutes() Sequence {
	return Sequence{
		LyingOnBackWarmupToDownDog(),
		SpeakAndPause("Low Back Hips Glutes. Starting from Tadasnaa", 5),
		ChairFromDownDog(),
		TadasanaFromFold(),
		GoddessPoseFromTadasana(),
		TadasanaToFold(),
		DownDogFromFold(),
		ClamshellGlutesFromDownDog(),
	}
}

/*
Inhale urdvahastasana  {{pause}}
Exhale and fold  {{pause}}
Inhale flat back  {{pause}}
Exhale to fold  {{pause}}
Inhale to plank  {{pause}}  {{pause}}
Exhale to belly  {{pause}}  {{pause}}
Inhale cobra  {{pause}}  {{pause}}
Exhale to belly  {{pause}}
Inhale to cobra, slowly.  {{pause}}
Exhale here  {{pause}}
Inhale again, hands lift up off ground  {{pause}}
Feel whole posterior chain activate  {{pause}}
Exhale here  {{pause}}
Inhale again  {{pause}}  {{pause}}
Exhale to childs pose.
Slowly rock hips back and draw chest and ribcage out and forward.
Stretching hips active arches  {{pause}}
Inhale  {{pause}}  {{pause}}
Exhale   {{pause}}  {{pause}}  {{pause}}  {{pause}}
Inhale to tabletop  {{pause}}  {{pause}}
Cat / Cow  {{pause}}  {{pause}}  {{pause}}  {{pause}}
Freestyle cat cow  {{pause}}  {{pause}}  {{pause}}  breathe. {{pause}}  {{pause}}  {{pause}}
Next inhale raise left arm to sky  {{pause}}
Exhale thread the needle  {{pause}}  {{pause}}  {{pause}} {{pause}}  {{pause}}
Breathe here, feel stretch in shoulder blade to fingers {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}} keep breathing  {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}}
Walk hand back and push back up to tabletop  {{pause}}  {{pause}}
Short Child’s pose  {{pause}}  {{pause}}
Back up to tabletop  {{pause}}
Next inhale raise right arm to sky  {{pause}}
Exhale thread the needle  {{pause}}  {{pause}}  {{pause}} {{pause}}  {{pause}}
Breathe here, feel stretch in shoulder blade to fingers {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}} keep breathing  {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}}  {{pause}}
Walk hand back and push back up to tabletop  {{pause}}  {{pause}}
Inhale here  {{pause}}  {{pause}}
Exhale to down dog  {{pause}}  {{pause}}  {{pause}}.

*/
