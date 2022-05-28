package myConcat

import (
	"fmt"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	kafkaAusschnitt := "Als Gregor Samsa ..."
	sliceOfStringOfKafkaAusschnitt := strings.Split(kafkaAusschnitt, " ")
	kafkaAusschnitt = Concat(sliceOfStringOfKafkaAusschnitt)
	if kafkaAusschnitt != "Als Gregor Samsa ..." {
		t.Error("Expected:", "Als Gregor Samsa ...", "got:", kafkaAusschnitt)
	}
}

func ExampleConcat() {
	sliceOfStringOfKafkaAusschnitt := []string{"Als", "Gregor", "Samsa", "..."}
	fmt.Println(Concat(sliceOfStringOfKafkaAusschnitt))
	// Output:
	// Als Gregor Samsa ...
}

const kafka = "Als Gregor Samsa eines Morgens aus unruhigen Träumen erwachte, fand er sich in seinem Bett zu einem ungeheueren Ungeziefer verwandelt. Und es war ihnen wie eine Bestätigung ihrer neuen Träume und guten Absichten, als am Ziele ihrer Fahrt die Tochter als erste sich erhob und ihren jungen Körper dehnte. »Es ist ein eigentümlicher Apparat«, sagte der Offizier zu dem Forschungsreisenden und überblickte mit einem gewissermaßen bewundernden Blick den ihm doch wohlbekannten Apparat. Sie hätten noch ins Boot springen können, aber der Reisende hob ein schweres, geknotetes Tau vom Boden, drohte ihnen damit und hielt sie dadurch von dem Sprunge ab. In den letzten Jahrzehnten ist das Interesse an Hungerkünstlern sehr zurückgegangen. Aber sie überwanden sich, umdrängten den Käfig und wollten sich gar nicht fortrühren.Jemand musste Josef K. verleumdet haben, denn ohne dass er etwas Böses getan hätte, wurde er eines Morgens verhaftet. »Wie ein Hund!« sagte er, es war, als sollte die Scham ihn überleben."

var sliceStringTemp []string

func BenchmarkConcat(b *testing.B) {
	sliceStringTemp = strings.Split(kafka, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(sliceStringTemp)
	}
}
