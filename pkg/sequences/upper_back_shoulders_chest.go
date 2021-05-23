package sequences

import (
	"github.com/ghodss/yaml"
	"github.com/ilackarms/yogayaml/pkg/model"
	"log"
)

func init() {
	return
	sequenceYaml, err := yaml.Marshal(UpperBackShoulders())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("for your needs: YAML: \n\n---\n%s\n---\n", sequenceYaml)
}

func UpperBackShoulders() model.Sequence {
	return model.Sequence{
		model.SpeakAndPause("Upper Back & Shoulders", 5),
		model.SpeakAndPause("Stand in Tadasana and begin to settle in.", 5),
		model.SpeakAndPause("Ground into your arches. Activate your glutes. Activate your bandas.", 2),
		model.InhaleExhale(3, 4, 6),
		model.SpeakAndPause("Inhale, Urdva Hastasana.", 2),
		//model.SkipPrevious{},
		model.SpeakAndPause("Forward fold.", 1),
		model.SpeakAndPause("Inhale Flat back.", 2),
		model.SpeakAndPause("Exhale to fold.", 2),
		model.SpeakAndPause("Inhale to plank.", 2),
		model.SpeakAndPause("Exhale to belly.", 1),
		model.SpeakAndPause("Inhale to cobra. Active the whole posterior chain. Keep belly activated.", 2),
		model.SpeakAndPause("Exhale here.", 3),
		model.SpeakAndPause("Inhale here.", 2),
		model.SpeakAndPause("Exhale down.", 2),
		model.SpeakAndPause("Inhale cobra second round.", 2),
		model.SpeakAndPause("Exhale here.", 3),
		model.SpeakAndPause("Inhale, pick hands up off the floor.", 3),
		model.SpeakAndPause("Exhale down.", 1),
		model.SpeakAndPause("Inhale again.", 1),
		model.SpeakAndPause("Exhale to childs pose.", 2),
		model.SpeakAndPause("Lift up to to tabletop.", 2),
		model.SpeakAndPause("Freestyle cat cow.", 2),
		model.InhaleExhale(2, 4, 6),
		model.ThreadTheNeedle(model.Left),
		model.ThreadTheNeedle(model.Right),
		model.SpeakAndPause("Exhale to down dog.", 4),
		model.SpeakAndPause("walk your feet to your hands. forward fold", 2),
		model.TadasanaFromFold(),
		model.SpeakAndPause("Pick up your strap and bring it in front of you with your arms spread apart.", 2),
		model.SpeakAndPause("Bring the strap up over your head ````````````````````````````````````````````````````````````	and as far back as you can go. Spread your arms farther to get further back.", 2),
		model.SpeakAndPause("Breathe as you move.", 2),
		model.SpeakAndPause("Keep your feet grounded slightly wider than hip width, arches and glutes engaged. Back straight bandas active.", 4),
		model.SpeakAndPause("Slowly, two more.", 4),
		model.SpeakAndPause("As you exhale come into forward fold.", 4),
		model.TadasanaFromFold(),


	}
}

