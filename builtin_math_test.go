package goja

import "testing"

func TestMean1(t *testing.T) {
	const SCRIPT = `
	var a = [32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34]
	Math.mean(...a)
	`
	testScript1(SCRIPT, valueFloat(35.96333333333333), t)
}

func TestMean2(t *testing.T) {
	const SCRIPT = `
	var a = [32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34]
	Math.mean(a)
	`
	testScript1(SCRIPT, valueFloat(35.96333333333333), t)
}

func TestMean3(t *testing.T) {
	const SCRIPT = `
	var a = [32.32]
	Math.mean(a)
	`
	testScript1(SCRIPT, valueFloat(32.32), t)
}
