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

func TestMean4(t *testing.T) {
	const SCRIPT = `
	var a = [32]
	Math.mean(a)
	`
	testScript1(SCRIPT, floatToValue(32), t)
}

func TestMean5(t *testing.T) {
	const SCRIPT = `
	var a = [32, 23]
	Math.mean(...a)
	`
	testScript1(SCRIPT, valueFloat(27.5), t)
}

func TestMedian(t *testing.T) {
	const SCRIPT = `
	var a = [32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34]
	Math.median(...a)
	`
	testScript1(SCRIPT, valueFloat(43.34), t)
}

func TestVar(t *testing.T) {
	const SCRIPT = `
	var a = [32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34]
	Math.var(...a)
	`
	testScript1(SCRIPT, valueFloat(253.6061111111111), t)
}

func TestStd(t *testing.T) {
	const SCRIPT = `
	var a = [32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34]
	Math.std(...a)
	`
	testScript1(SCRIPT, valueFloat(15.925015262507948), t)
}

func TestStd1(t *testing.T) {
	const SCRIPT = `
	var a = [32.32]
	Math.std(...a)
	`
	testScript1(SCRIPT, floatToValue(0), t)
}

func TestStd2(t *testing.T) {
	const SCRIPT = `
	Math.std()
	`
	testScript1(SCRIPT, _NaN, t)
}
