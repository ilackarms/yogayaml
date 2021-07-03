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

func Warriors() Sequence {

	return Sequence{
		//LyingOnBackWarmupToDownDog(),
		SpeakAndPause("Warriors.", 3),
		TabletopStretch(),
		SpeakAndPause("starting standing at the top of the mat.", 3),
		Tadasana(4),
		TadasanaToFold(),
		Seated(4),
		SeatedFold(4),
		Staff(2),
		ReversePlank(4),
		SeatedFold(4),
		Seated(1),
		SpeakAndPause("Roll up to tabletop.", 5),
		Freestyle("down dog", 16),
		DownDogStretch(),
		SunSalutation(),
		TadasanaToFold(),
		Freestyle("plank", 4),
		Plank(true, 6),
		SpeakAndPause("Exhale slowly down to belly.", 5),
		Cobra(true, 6),
		SpeakAndPause("Exhale slowly down to belly.", 5),
		Cobra(false, 12),
		SpeakAndPause("Press back to tabletop.", 0),
		DownDog(4, false),

		LowLunge(Left, false, 2),
		Crescent(Left, false, 2),
		DownDog(4, false),

		LowLunge(Right, false, 2),
		Crescent(Right, false, 2),
		DownDog(4, false),

		LowLunge(Left, false, 2),
		HighLunge(Left, false, 2),
		Warrior2(Left, false, 2),
		PeacefulWarrior(Left, true, 2),
		ExtendedSideAngle(Left, true, 2),
		PeacefulWarrior(Left, false, 2),
		LowLunge(Left, false, 2),
		DownDog(4, false),

		LowLunge(Right, false, 2),
		HighLunge(Right, false, 2),
		Warrior2(Right, false, 2),
		PeacefulWarrior(Right, true, 2),
		ExtendedSideAngle(Right, false, 2),
		PeacefulWarrior(Right, true, 2),
		LowLunge(Right, false, 2),
		DownDog(4, false),

		ForwardFold(10),
		TadasanaFromFold(),
	}
}

func StaffGoddessChair() Sequence {
	return Sequence{
		//LyingOnBackWarmupToDownDog(),
		SpeakAndPause("Goddess Staff Chair.", 3),
		SpeakAndPause("starting on the ground in seated.", 3),
		Seated(4),
		Staff(4),
		SeatedFold(3),
		Staff(2),
		ReversePlank(4),
		SeatedFold(4),
		HappyBaby(4),
		Bridge(3),
		SpeakAndPause("lower back down.", 3),
		Bridge(4),
		HappyBaby(8),
		SpeakAndPause("Roll up to a seated position.", 3),
		SpeakAndPause("Come into tabletop.", 5),
		DownDogStretch(),
		ChairFromDownDog(),
		SpeakAndPause("Fordward fold.", 5),
		TadasanaFromFold(),
		GoddessPoseFromTadasana(),
		TadasanaToFold(),
		DownDogFromFold(),

		ForwardFold(4),
		FoldToTadasana(3),
		SunSalutation(),
		SpeakAndPause("Inhale urdva hastasana.", 2),
		ForwardFold(3),
		Plank(true, 6),
		SpeakAndPause("Exhale down to belly.", 2),
		SpeakAndPause("Cobra. Work the back", 4),
		SpeakAndPause("Exhale down.", 2),
		SpeakAndPause("Press back to child's pose.", 4),
		SpeakAndPause("Breathe here.", 4),
		SpeakAndPause("Inhale up to tabletop.", 2),

		ClamshellGlutesFromDownDog(),

		Staff(2),
		ReversePlank(4),
		SpeakAndPause("lower down to your back.", 3),
		HappyBaby(4),
		SpeakAndPause("you did it. well done. namaste.", 3),
	}
}

func July3_2021() Sequence {
	return Sequence{
		// warmup
		Warmup1(),

		StandingSideLevel1(Left),
		StandingSideLevel1(Right),
		Vinyasa(false, 2),


		StandingSideLevel2(Left),
		Vinyasa(false, 2),

		StandingSideLevel2(Right),
		Vinyasa(false, 2),


		StandingSideLevel3(Left),
		Vinyasa(false, 2),

		StandingSideLevel3(Right),
		Vinyasa(false, 2),

		Cooldown1(),

		SpeakAndPause("make your way into shavasana", 10),
		SpeakAndPause("continue to rest here or not. bring your yoga with  you always. beginning 2 minute shavasana timer", 2),

		Shavasana(120),

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
