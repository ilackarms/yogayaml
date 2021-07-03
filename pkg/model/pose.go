package model

import (
	"fmt"
	"github.com/ilackarms/yogayaml/pkg/action"
	"strings"
)

// TODO transitions

type Pose struct {
	Name         string
	Instructions []string
	Duration     int64
}

var _ Node = Pose{}

func NewPose(name string, duration int64, instructions ...string) Pose {
	return Pose{Name: name, Instructions: instructions, Duration: duration}
}

func (p Pose) Action() action.Action {
	return SpeakAndPause(p.Name+". "+strings.Join(p.Instructions, ". "), p.Duration).Action()
}

func (p Pose) Skip() bool {
	return false
}

/*
Instructions
*/
const Instruction_ShouldersBackAndBreathe = "breath expands the whole ribcage here." // "Chin tucked in slightly. Relax jaw and brows. Activate your bandas. Ground down while feeling your spine elongate. Breathe deep."

/*
Poses
*/
func Tadasana(duration int64) Pose {
	return NewPose(
		"Tadasana",
		duration,
		"Ground into your feet, arches flexed. Calves stabilize ankles. Glutes active. Shoulders melt down back.",
		Instruction_ShouldersBackAndBreathe,
	)
}

func Seated(duration int64) Pose {
	return NewPose(
		"Seated",
		duration,
		Instruction_ShouldersBackAndBreathe,
		"Legs crossed, feet flexed. Rest your hands comfortably. Spine straight. Hip flexors and abs stabilize back.  stabilize ankles. Activate your bandas. Ground down while feeling your spine elongate. Shoulders melt down back.",
	)
}

func SeatedFold(duration int64) Pose {
	return NewPose(
		"Seated Fold",
		duration,
		"From Staff, raise arms up inhale. Exhale, pull in your belly as you lean over your feet. Play with keeping upper back straight and folding from the hips. Play with the bend in the knees and the hamstrings.",
	)
}

func Staff(duration int64) Pose {
	return NewPose(
		"Staff",
		duration,
		"Legs straight, feet flexed. Spine straight. Handds push down into floor. Hip flexors and abs stabilize back. Activate your bandas. Ground down while feeling your spine elongate.",
		Instruction_ShouldersBackAndBreathe,
	)
}

func ReversePlank(duration int64) Pose {
	return NewPose(
		"Reverse Plank",
		duration,
		"Start from staff pose",
		"Move your hands back slightly",
		"Push your hips up while pointing your toes down",
		"Head tilts back",
		"Breathe",
		Instruction_ShouldersBackAndBreathe,
	)
}

func Bridge(duration int64) Pose {
	return NewPose(
		"Bridge",
		duration,
		"Start from lying down on your back",
		"Feet come up a few inches below the glutes",
		"Shuffle your shoulderblades under your back",
		"Inhale and press your hips up in the air",
		"",
		Instruction_ShouldersBackAndBreathe,
	)
}

func HappyBaby(duration int64) Pose {
	return NewPose(
		"Happy Baby",
		duration,
	)
}

func MorningStretch(duration int64) Pose {
	return NewPose(
		"Morning Stretch.",
		duration,
	)
}


func Shavasana(duration int64) Pose {
	return NewPose(
		"Shavasana.",
		duration,
	)
}

func ForwardFold(duration int64) Pose {
	return NewPose(
		"Forward Fold",
		duration,
		"Breathe and release the spine. Feet and calves strong.",
		Instruction_ShouldersBackAndBreathe,
	)
}

func FoldToTadasana(stepDuration int64) Node {
	return Sequence{
		SpeakAndPause("Inhale flat back.", stepDuration),
		SpeakAndPause("Exhale fold.", stepDuration),
		SpeakAndPause("Inhale urdva hastasana.", stepDuration),
		SpeakAndPause("Exhale tadasana.", stepDuration),
	}
}

func Plank(longVersion bool, stepDuration int64) Node {
	instr := []string{
	}
	if longVersion {
		instr = append(instr,
		"Inhale in Plank.",
		"Glutes active. Abs active. Press into your toes and stretch legs.",
			"Exhale here and remain firm.",
			"Stretch lightly into the toes and calves.",
			"Keep chest expanded, neck as an extension of the spine.",
		)
	}
	return NewPose(
		"Plank",
		stepDuration,
		instr...,
	)
}

func LowLunge(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		fmt.Sprintf("Exhale to low lunge, %s foot forward. Knee and weight goes straight down over foot. Ankle calf and arch engaged in %s foot.", side, side),
		"Back leg strong",
	}
	if longVersion {
		instr = append(instr,
			"Inhale here and remain firm.",
			"Stretch lightly into the hips, pulling core in.",
			"Keep chest expanded, neck as an extension of the spine.",
			"Breathe",
		)
	}
	return NewPose(
		"Low Lunge",
		stepDuration,
		instr...,
	)
}

func Crescent(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		fmt.Sprintf("%s knee down to the ground. Inhale and engage core. Slowly roll your spine up  while, %s foot forward. Knee and weight goes straight down over foot. Ankle calf and arch engaged in %s foot.", side.Opposite(), side, side),
		"Back leg presses down into the ground",
	}
	if longVersion {
		instr = append(instr,
			"Inhale here and remain firm.",
			"Stretch lightly into the hip flexors, pulling core in.",
			"Keep chest expanded, neck as an extension of the spine.",
			"Breathe",
		)
	}
	return NewPose(
		"Crescent "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func HighLunge(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		"Exhale and engage core. Back leg strong. Toes curled under.",
		fmt.Sprintf("With %s foot forward and core firm, slowly roll your spine up. Knee and weight goes straight down over foot. Ankle calf and arch engaged in %s foot.", side, side),
		"Arms raised overhead or in half angel.",
		"Back leg strong.",
		Instruction_ShouldersBackAndBreathe,
	}
	if longVersion {
		instr = append(instr,
			"Inhale here and remain firm.",
			"Stretch lightly into the hip flexors, pulling core in.",
			"Keep chest expanded, neck as an extension of the spine.",
			"Arms and chest strong",
			"Breathe",
		)
	}
	return NewPose(
		"High Lunge "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func Warrior2(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		"Exhale to warrior 2. Keep core engaged. Back foot lined up with front arch.",
		fmt.Sprintf("%s Knee and weight goes straight down over foot. Ankle calf and arch engaged in %s foot.", side, side),
		"Back leg strong.",
		"Arms strech out in both directions. Soften shoulders and keep chest opening.",
		Instruction_ShouldersBackAndBreathe,
	}
	if longVersion {
		instr = append(instr,
			"Inhale here and remain firm.",
			"Stretch lightly into the hip flexors, pulling core in.",
			"Keep chest expanded, neck as an extension of the spine.",
			"Arms and chest strong",
			"Breathe",
		)
	}
	return NewPose(
		"Warrior 2 "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func PeacefulWarrior(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		"Inhale peaceful warrior. Keep core engaged. Back foot lined up with front arch.",
		Instruction_ShouldersBackAndBreathe,
	}
	if longVersion {
		instr = append(instr,
			side.String() + " arm lifts, " + side.Opposite().String() + " arm gently drapes down back leg.",
			side.String() + " ribcage expands.",
			"Back leg strong.",
			"Inhale here and remain firm.",
			"Stretch lightly into the side body.",
			"Arms and chest strong",
			"Breathe",
		)
	}
	return NewPose(
		"Peaceful Warrior "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func ExtendedSideAngle(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{
		side.String() + " arm comes to the top of " + side.String() + " knee.",
		side.Opposite().String() + " arm lifts as " + side.Opposite().String() + " ribcage expands. Core strong here.",
		"Back leg strong.",
		Instruction_ShouldersBackAndBreathe,
	}
	if longVersion {
		instr = append(instr,
			"Inhale here and remain firm.",
			"Stretch into the side body.",
			"Arms and chest strong",
			"Breathe",
		)
	}
	return NewPose(
		"Extended Side Angle "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func Freestyle(nextAction string, duration int64) Pose {
	return NewPose(
		"move freely",
		duration,
		fmt.Sprintf("next pose will be %s. %v seconds here", nextAction, duration),
	)
}

func Cobra(longVersion bool, duration int64) Pose {
		instr := []string{}
	if longVersion {
		instr = append(instr,
		"Inhale to cobra. Active the whole posterior chain. Keep belly activated.",
		)
	}
	return NewPose(
		"Cobra Pose",
		duration,
		instr...,
	)
}

func DownDog(duration int64, longVersion bool) Pose {
	 instructions := []string{}
	 if longVersion {
		 instructions = append(instructions, "Keep your arches strong, upper back flat. Neck hangs loose. Press your torso back towards your knees. Slight bend in knees.")
	 }
	return NewPose(
		"Down Dog",
		duration,
		instructions...,
	)
}

func Triangle(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Triangle "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func HalfMoon(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Half Moon "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func Warrior3(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Warrior 3 "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func Pyramid(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Pyramid "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func RevolvedChair(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"RevolvedChair "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func RevolvedLunge(side Side, longVersion bool, stepDuration int64, kneeDown bool) Node {
	instr := []string{}
	if kneeDown {
		instr = append(instr,
			"back knee down.",
		)
	}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Revolved Lunge "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func SidePlank(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Side Plank "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func StandingSplit(side Side, longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Standing Split "+side.String()+" side.",
		stepDuration,
		instr...,
	)
}

func ForearmPlank(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Forearm Plank.",
		stepDuration,
		instr...,
	)
}

func Chaturanga(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Chaturanga.",
		stepDuration,
		instr...,
	)
}

func Boat(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Boat.",
		stepDuration,
		instr...,
	)
}


func Tabletop(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Tabletop.",
		stepDuration,
		instr...,
	)
}

func ChildsPose(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Childs Pose.",
		stepDuration,
		instr...,
	)
}

func Camel(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Camel.",
		stepDuration,
		instr...,
	)
}

func HoveringTabletop(longVersion bool, stepDuration int64) Node {
	instr := []string{}
	if longVersion {
		instr = append(instr,
			"",
		)
	}
	return NewPose(
		"Hovering Tabletop.",
		stepDuration,
		instr...,
	)
}

// Flows
func Vinyasa(longVersion bool, stepDuration int64) Node {
	return Sequence{
		SpeakAndPause("Vinyasa", 1),
		DownDog(stepDuration, longVersion),
		Plank(longVersion, stepDuration),
		Chaturanga(longVersion, stepDuration),
		SpeakAndPause("lower slowly to ground", stepDuration),
		Cobra(longVersion, stepDuration),
		SpeakAndPause("lower slowly to ground", stepDuration),
		DownDog(stepDuration, longVersion),
	}
}

func Warmup1()  Node {
	return Sequence{
		SpeakAndPause("Nourish the spine. Support the ankle. Breathe deep into the lungs and expand.", 3),
		SpeakAndPause("starting in forward fold.", 3),
		ForwardFold(10),
		SpeakAndPause("halfway lift.", 3),
		SpeakAndPause("exhale fold.", 10),
		Seated(10),
		Staff(10),
		SeatedFold(10),
		Seated(5),
		ReversePlank(20),
		Tabletop(false, 5),
		HoveringTabletop(false, 5),
		Freestyle("down dog", 30),
		DownDog(20, true),
		Vinyasa(true, 5),
		Plank(true, 10),
		DownDog(20, false),
	}
}

func StandingSideLevel1(side Side) Node {
	return Sequence{
		Crescent(side, false, 20),
		RevolvedLunge(side, false, 10, true),
		LowLunge(side, false, 5),
	}
}

func StandingSideLevel2(side Side) Node {
	return Sequence{
		Warrior2(side, false, 8),
		PeacefulWarrior(side, false, 5),
		ExtendedSideAngle(side, false, 6),
		PeacefulWarrior(side, false, 2),
		StandingSplit(side, false, 10),
		LowLunge(side, false, 5),
		Plank(false, 3),
	}
}

func StandingSideLevel3(side Side) Node {
	return Sequence{
		HighLunge(side, false, 8),
		Warrior2(side, false, 8),
		PeacefulWarrior(side, false, 5),
		RevolvedLunge(side, false, 5, false),
		HighLunge(side, false, 2),
		Warrior2(side, false, 2),
		HalfMoon(side, false, 10),
		PeacefulWarrior(side, false, 6),
		ExtendedSideAngle(side, false, 6),
		SidePlank(side, false, 3),
		Plank(false, 3),
	}
}

func Cooldown1() Node {
	return Sequence{
		DownDog(20, false),
		Tabletop(false, 20),
		ChildsPose(false, 20),
		Camel(false, 20),
		Seated(10),
		SeatedFold(10),
		HappyBaby(10),
		MorningStretch(10),
	}
}

